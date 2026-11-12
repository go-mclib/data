package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync"
)

const maxOpenRegions = 32

var errStorageClosed = errors.New("region storage closed")

// RegionStorage manages a directory of region files with an LRU cache.
type RegionStorage struct {
	dir    string
	mu     sync.Mutex
	closed bool

	// LRU cache of open region files
	cache    map[int64]*regionEntry
	lruOrder []int64 // oldest first
}

// regionEntry is a cached region file and the number of operations currently
// using it; entries in use are never evicted, so a region is only closed when
// no caller still holds it.
type regionEntry struct {
	region *regionFile
	refs   int
}

// NewRegionStorage creates a storage backed by .mca files in dir.
func NewRegionStorage(dir string) (*RegionStorage, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &RegionStorage{
		dir:   dir,
		cache: make(map[int64]*regionEntry),
	}, nil
}

func regionKey(rx, rz int32) int64 {
	return int64(rx)<<32 | int64(uint32(rz))
}

func regionCoords(chunkX, chunkZ int32) (int32, int32) {
	return chunkX >> 5, chunkZ >> 5
}

// acquire returns the (cached) region file for a chunk with its reference
// count raised; the caller must release it. Without create, a missing file
// yields (nil, nil) instead of creating an empty region on disk.
func (rs *RegionStorage) acquire(chunkX, chunkZ int32, create bool) (*regionEntry, error) {
	rx, rz := regionCoords(chunkX, chunkZ)
	key := regionKey(rx, rz)

	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.closed {
		return nil, errStorageClosed
	}
	if entry, ok := rs.cache[key]; ok {
		rs.touchLRU(key)
		entry.refs++
		return entry, nil
	}

	path := filepath.Join(rs.dir, fmt.Sprintf("r.%d.%d.mca", rx, rz))
	if !create {
		if _, err := os.Stat(path); err != nil {
			return nil, nil
		}
	}

	rs.evictIdle()

	region, err := openRegion(path)
	if err != nil {
		return nil, err
	}

	entry := &regionEntry{region: region, refs: 1}
	rs.cache[key] = entry
	rs.lruOrder = append(rs.lruOrder, key)
	return entry, nil
}

func (rs *RegionStorage) release(entry *regionEntry) {
	rs.mu.Lock()
	entry.refs--
	rs.mu.Unlock()
}

func (rs *RegionStorage) touchLRU(key int64) {
	if i := slices.Index(rs.lruOrder, key); i >= 0 {
		rs.lruOrder = append(slices.Delete(rs.lruOrder, i, i+1), key)
	}
}

// evictIdle closes least recently used regions that no operation holds until
// there is room for one more. When every cached region is in use, the cache
// temporarily grows past maxOpenRegions.
func (rs *RegionStorage) evictIdle() {
	for i := 0; len(rs.cache) >= maxOpenRegions && i < len(rs.lruOrder); {
		key := rs.lruOrder[i]
		entry := rs.cache[key]
		if entry.refs > 0 {
			i++
			continue
		}
		entry.region.Close()
		delete(rs.cache, key)
		rs.lruOrder = slices.Delete(rs.lruOrder, i, i+1)
	}
}

// ReadChunk reads a chunk from the appropriate region file. Returns nil, nil if absent.
func (rs *RegionStorage) ReadChunk(chunkX, chunkZ int32) ([]byte, error) {
	entry, err := rs.acquire(chunkX, chunkZ, false)
	if err != nil || entry == nil {
		return nil, err
	}
	defer rs.release(entry)
	return entry.region.ReadChunk(chunkX, chunkZ)
}

// WriteChunk writes a chunk to the appropriate region file, creating it if needed.
func (rs *RegionStorage) WriteChunk(chunkX, chunkZ int32, data []byte) error {
	entry, err := rs.acquire(chunkX, chunkZ, true)
	if err != nil {
		return err
	}
	defer rs.release(entry)
	return entry.region.WriteChunk(chunkX, chunkZ, data)
}

// HasChunk checks if a chunk exists in the appropriate region file.
func (rs *RegionStorage) HasChunk(chunkX, chunkZ int32) bool {
	entry, err := rs.acquire(chunkX, chunkZ, false)
	if err != nil || entry == nil {
		return false
	}
	defer rs.release(entry)
	return entry.region.HasChunk(chunkX, chunkZ)
}

// Close closes all open region files and makes later calls fail. A region
// operation already in progress finishes first (region files serialize access
// behind their own lock).
func (rs *RegionStorage) Close() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.closed = true
	var errs []error
	for _, entry := range rs.cache {
		errs = append(errs, entry.region.Close())
	}
	rs.cache = nil
	rs.lruOrder = nil
	return errors.Join(errs...)
}
