package storage

import (
	"crypto/rand"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-mclib/nbt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegionFileRoundTrip(t *testing.T) {
	r, err := openRegion(filepath.Join(t.TempDir(), "r.0.0.mca"))
	require.NoError(t, err)
	defer r.Close()

	data := []byte("hello chunk")
	require.NoError(t, r.WriteChunk(3, 5, data))

	assert.True(t, r.HasChunk(3, 5))
	assert.False(t, r.HasChunk(0, 0))

	got, err := r.ReadChunk(3, 5)
	require.NoError(t, err)
	assert.Equal(t, data, got)

	got, err = r.ReadChunk(0, 0)
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestRegionFileOverwriteKeepsOldSectorsUntilCommitted(t *testing.T) {
	r, err := openRegion(filepath.Join(t.TempDir(), "r.0.0.mca"))
	require.NoError(t, err)
	defer r.Close()

	require.NoError(t, r.WriteChunk(1, 1, []byte("first")))
	oldOffset := r.locations[chunkIndex(1, 1)] >> 8
	require.NoError(t, r.WriteChunk(1, 1, []byte("second")))

	// the new copy must not overwrite the old one in place
	assert.NotEqual(t, oldOffset, r.locations[chunkIndex(1, 1)]>>8)
	assert.False(t, r.sectors[oldOffset], "old sectors are freed after the header update")

	got, err := r.ReadChunk(1, 1)
	require.NoError(t, err)
	assert.Equal(t, []byte("second"), got)
}

func TestRegionFilePersistence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "r.0.0.mca")

	r, err := openRegion(path)
	require.NoError(t, err)
	require.NoError(t, r.WriteChunk(10, 20, []byte("persistent")))
	r.Close()

	r2, err := openRegion(path)
	require.NoError(t, err)
	defer r2.Close()

	got, err := r2.ReadChunk(10, 20)
	require.NoError(t, err)
	assert.Equal(t, []byte("persistent"), got)
}

func TestRegionFileClearsInvalidLocations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "r.0.0.mca")
	header := make([]byte, headerBytes+sectorSize)
	binary.BigEndian.PutUint32(header[chunkIndex(0, 0)*4:], 2<<8|0) // zero sectors
	binary.BigEndian.PutUint32(header[chunkIndex(1, 0)*4:], 1<<8|1) // overlaps header
	binary.BigEndian.PutUint32(header[chunkIndex(2, 0)*4:], 9<<8|1) // past end of file
	require.NoError(t, os.WriteFile(path, header, 0644))

	r, err := openRegion(path)
	require.NoError(t, err)
	defer r.Close()

	for x := range int32(3) {
		got, err := r.ReadChunk(x, 0)
		require.NoError(t, err)
		assert.Nil(t, got)
		assert.False(t, r.HasChunk(x, 0))
	}
}

func TestRegionFileOversizedChunkIsExternal(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.0.0.mca")
	r, err := openRegion(path)
	require.NoError(t, err)

	big := make([]byte, (maxChunkSectors+8)*sectorSize) // incompressible
	_, err = rand.Read(big)
	require.NoError(t, err)
	require.NoError(t, r.WriteChunk(4, 7, big))

	loc := r.locations[chunkIndex(4, 7)]
	assert.Equal(t, uint32(1), loc&0xFF, "region keeps a one-sector stub")
	external := filepath.Join(dir, "c.4.7.mcc")
	assert.FileExists(t, external)
	require.NoError(t, r.Close())

	r, err = openRegion(path)
	require.NoError(t, err)
	defer r.Close()
	got, err := r.ReadChunk(4, 7)
	require.NoError(t, err)
	assert.Equal(t, big, got)

	require.NoError(t, r.WriteChunk(4, 7, []byte("small")))
	assert.NoFileExists(t, external)
	got, err = r.ReadChunk(4, 7)
	require.NoError(t, err)
	assert.Equal(t, []byte("small"), got)
}

func TestRegionStorageNeverEvictsRegionsInUse(t *testing.T) {
	rs, err := NewRegionStorage(t.TempDir())
	require.NoError(t, err)
	defer rs.Close()

	require.NoError(t, rs.WriteChunk(0, 0, []byte("held")))
	held, err := rs.acquire(0, 0, false)
	require.NoError(t, err)

	// open enough other regions to force eviction
	for i := range int32(maxOpenRegions + 4) {
		require.NoError(t, rs.WriteChunk((i+1)*32, 0, []byte("other")))
	}

	got, err := held.region.ReadChunk(0, 0)
	require.NoError(t, err, "held region must stay open")
	assert.Equal(t, []byte("held"), got)
	rs.release(held)

	require.NoError(t, rs.WriteChunk(64*32, 0, []byte("next")))
	assert.LessOrEqual(t, len(rs.cache), maxOpenRegions)
}

func TestSaveCompressedNBTReplacesAtomically(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.dat")
	require.NoError(t, SaveCompressedNBT(path, nbt.Compound{"v": nbt.Int(1)}))
	require.NoError(t, SaveCompressedNBT(path, nbt.Compound{"v": nbt.Int(2)}))

	root, err := LoadCompressedNBT(path)
	require.NoError(t, err)
	assert.Equal(t, int32(2), root.GetInt("v"))

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, entries, 1, "no temporary files left behind")
}
