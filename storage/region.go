// Package storage implements Minecraft's Anvil region file format for
// persisting chunk and player data.
package storage

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	sectorSize      = 4096
	headerSectors   = 2 // location table + timestamp table
	headerBytes     = headerSectors * sectorSize
	chunksPerAxis   = 32
	chunksPerFile   = chunksPerAxis * chunksPerAxis
	compressionZlib = 2
	compressionNone = 3
	// externalFlag marks a chunk whose stream lives in a c.<x>.<z>.mcc file
	// next to the region (vanilla RegionFile.EXTERNAL_STREAM_FLAG).
	externalFlag = 0x80
	// maxChunkSectors is the most sectors a location entry can address
	// (8-bit count); larger chunks are stored externally.
	maxChunkSectors = 255
)

// regionFile handles reading and writing chunks in one .mca region file.
// Each region file holds up to 32×32 = 1024 chunks.
type regionFile struct {
	path       string
	file       *os.File
	mu         sync.Mutex
	locations  [chunksPerFile]uint32 // sector offset (24 bits) | sector count (8 bits)
	timestamps [chunksPerFile]uint32
	sectors    []bool // bitmap of used sectors
}

// openRegion opens or creates a region file at the given path.
func openRegion(path string) (*regionFile, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	r := &regionFile{path: path, file: f}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	if stat.Size() < headerBytes {
		// new file — write empty header
		err = r.initHeader()
	} else {
		err = r.readHeader(stat.Size())
	}
	if err != nil {
		f.Close()
		return nil, err
	}

	return r, nil
}

func (r *regionFile) initHeader() error {
	header := make([]byte, headerBytes)
	if _, err := r.file.WriteAt(header, 0); err != nil {
		return err
	}
	r.sectors = make([]bool, headerSectors)
	r.sectors[0] = true
	r.sectors[1] = true
	return nil
}

// readHeader loads the location and timestamp tables. Like vanilla
// (RegionFile.java), entries that overlap the header, have zero sectors or
// start past the end of the file are cleared instead of trusted.
func (r *regionFile) readHeader(size int64) error {
	var buf [headerBytes]byte
	if _, err := r.file.ReadAt(buf[:], 0); err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	totalSectors := max(int((size+sectorSize-1)/sectorSize), headerSectors)
	r.sectors = make([]bool, totalSectors)
	r.sectors[0] = true
	r.sectors[1] = true

	for i := range chunksPerFile {
		r.timestamps[i] = binary.BigEndian.Uint32(buf[sectorSize+i*4:])
		loc := binary.BigEndian.Uint32(buf[i*4:])
		if loc == 0 {
			continue
		}
		offset, count := int(loc>>8), int(loc&0xFF)
		if offset < headerSectors || count == 0 || int64(offset)*sectorSize > size {
			continue
		}
		r.locations[i] = loc
		r.markSectors(offset, count, true)
	}

	return nil
}

func chunkIndex(chunkX, chunkZ int32) int {
	return int(chunkX&31) + int(chunkZ&31)*chunksPerAxis
}

func (r *regionFile) externalChunkPath(chunkX, chunkZ int32) string {
	return filepath.Join(filepath.Dir(r.path), fmt.Sprintf("c.%d.%d.mcc", chunkX, chunkZ))
}

// HasChunk returns true if the region file contains data for the given chunk.
func (r *regionFile) HasChunk(chunkX, chunkZ int32) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.locations[chunkIndex(chunkX, chunkZ)] != 0
}

// ReadChunk reads and decompresses chunk data. Returns nil, nil if chunk is absent.
func (r *regionFile) ReadChunk(chunkX, chunkZ int32) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	loc := r.locations[chunkIndex(chunkX, chunkZ)]
	if loc == 0 {
		return nil, nil
	}

	// read raw sector data; the last sector may be truncated at end of file
	raw := make([]byte, int(loc&0xFF)*sectorSize)
	n, err := r.file.ReadAt(raw, int64(loc>>8)*sectorSize)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("read sectors: %w", err)
	}
	raw = raw[:n]
	if len(raw) < 5 {
		return nil, fmt.Errorf("chunk (%d,%d) header truncated: %d bytes", chunkX, chunkZ, len(raw))
	}

	// parse chunk header: 4-byte length + 1-byte compression
	length := int(binary.BigEndian.Uint32(raw[0:4]))
	if length < 1 || length > len(raw)-4 {
		return nil, fmt.Errorf("invalid chunk length %d", length)
	}
	compression := raw[4]
	compressed := raw[5 : 4+length]
	if compression&externalFlag != 0 {
		compression &^= externalFlag
		if compressed, err = os.ReadFile(r.externalChunkPath(chunkX, chunkZ)); err != nil {
			return nil, fmt.Errorf("read external chunk: %w", err)
		}
	}

	switch compression {
	case compressionZlib:
		zr, err := zlib.NewReader(bytes.NewReader(compressed))
		if err != nil {
			return nil, fmt.Errorf("zlib init: %w", err)
		}
		defer zr.Close()
		return io.ReadAll(zr)
	case compressionNone:
		return compressed, nil
	default:
		return nil, fmt.Errorf("unsupported compression type %d", compression)
	}
}

// WriteChunk compresses and writes chunk data to the region file. Like
// vanilla, the new sectors are written and the header updated before the old
// sectors are released, so a failed or interrupted write keeps the old chunk.
// Chunks needing more than 255 sectors go to an external .mcc file.
func (r *regionFile) WriteChunk(chunkX, chunkZ int32, data []byte) error {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	if _, err := zw.Write(data); err != nil {
		return fmt.Errorf("zlib write: %w", err)
	}
	if err := zw.Close(); err != nil {
		return fmt.Errorf("zlib close: %w", err)
	}
	compressed := buf.Bytes()

	r.mu.Lock()
	defer r.mu.Unlock()

	externalPath := r.externalChunkPath(chunkX, chunkZ)

	// chunk payload: 4-byte length + 1-byte compression + compressed data
	payloadLen := 1 + len(compressed)
	neededSectors := (4 + payloadLen + sectorSize - 1) / sectorSize
	var sectorData []byte
	externalTmp := ""
	if neededSectors > maxChunkSectors {
		tmp, err := writeTempFile(filepath.Dir(r.path), compressed)
		if err != nil {
			return fmt.Errorf("write external chunk: %w", err)
		}
		externalTmp = tmp
		neededSectors = 1
		sectorData = make([]byte, sectorSize)
		binary.BigEndian.PutUint32(sectorData[0:4], 1)
		sectorData[4] = compressionZlib | externalFlag
	} else {
		sectorData = make([]byte, neededSectors*sectorSize)
		binary.BigEndian.PutUint32(sectorData[0:4], uint32(payloadLen))
		sectorData[4] = compressionZlib
		copy(sectorData[5:], compressed)
	}

	newOffset := r.allocateSectors(neededSectors)
	fail := func(err error) error {
		r.markSectors(newOffset, neededSectors, false)
		if externalTmp != "" {
			os.Remove(externalTmp)
		}
		return err
	}
	if _, err := r.file.WriteAt(sectorData, int64(newOffset)*sectorSize); err != nil {
		return fail(fmt.Errorf("write sectors: %w", err))
	}

	idx := chunkIndex(chunkX, chunkZ)
	oldLoc, oldTimestamp := r.locations[idx], r.timestamps[idx]
	r.locations[idx] = uint32(newOffset)<<8 | uint32(neededSectors)
	r.timestamps[idx] = uint32(time.Now().Unix())
	if err := r.writeHeader(); err != nil {
		r.locations[idx], r.timestamps[idx] = oldLoc, oldTimestamp
		return fail(fmt.Errorf("write header: %w", err))
	}

	var commitErr error
	if externalTmp != "" {
		commitErr = os.Rename(externalTmp, externalPath)
	} else if err := os.Remove(externalPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		commitErr = err
	}
	if oldLoc != 0 {
		r.markSectors(int(oldLoc>>8), int(oldLoc&0xFF), false)
	}
	return commitErr
}

// markSectors sets the used state of sectors [start, start+count), growing
// the bitmap as needed. The header sectors always stay used.
func (r *regionFile) markSectors(start, count int, used bool) {
	if end := start + count; end > len(r.sectors) {
		r.sectors = append(r.sectors, make([]bool, end-len(r.sectors))...)
	}
	for s := max(start, headerSectors); s < start+count; s++ {
		r.sectors[s] = used
	}
}

// allocateSectors reserves the first run of count free sectors, appending at
// the end of the file when no gap is large enough.
func (r *regionFile) allocateSectors(count int) int {
	run := 0
	for s := headerSectors; s < len(r.sectors); s++ {
		if r.sectors[s] {
			run = 0
			continue
		}
		if run++; run == count {
			start := s - count + 1
			r.markSectors(start, count, true)
			return start
		}
	}
	start := len(r.sectors) - run
	r.markSectors(start, count, true)
	return start
}

func (r *regionFile) writeHeader() error {
	var buf [headerBytes]byte
	for i := range chunksPerFile {
		binary.BigEndian.PutUint32(buf[i*4:], r.locations[i])
		binary.BigEndian.PutUint32(buf[sectorSize+i*4:], r.timestamps[i])
	}
	_, err := r.file.WriteAt(buf[:], 0)
	return err
}

// Close closes the underlying file.
func (r *regionFile) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.file.Close()
}
