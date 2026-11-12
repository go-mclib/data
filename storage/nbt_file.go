package storage

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-mclib/nbt"
)

// SaveCompressedNBT writes a root compound to a gzip-compressed NBT file,
// creating parent directories. The file is replaced atomically, so a crash
// mid-write leaves the previous contents intact.
func SaveCompressedNBT(path string, root nbt.Compound) error {
	if dir := filepath.Dir(path); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	data, err := nbt.EncodeFile(root, "")
	if err != nil {
		return fmt.Errorf("encode NBT: %w", err)
	}
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(data); err != nil {
		return fmt.Errorf("gzip write: %w", err)
	}
	if err := gw.Close(); err != nil {
		return fmt.Errorf("gzip close: %w", err)
	}
	return WriteFileAtomic(path, buf.Bytes())
}

// WriteFileAtomic replaces path with data via a synced temporary file in the
// same directory and a rename, like vanilla's Util.safeReplaceFile.
func WriteFileAtomic(path string, data []byte) error {
	tmp, err := writeTempFile(filepath.Dir(path), data)
	if err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}

// MoveAside renames an unreadable save file to "<path>.corrupt-<unix>" so a
// later save cannot overwrite it, and returns the new path.
func MoveAside(path string) (string, error) {
	corrupt := fmt.Sprintf("%s.corrupt-%d", path, time.Now().Unix())
	return corrupt, os.Rename(path, corrupt)
}

// writeTempFile writes data to a new synced temporary file in dir and returns
// its path. The caller renames it into place or removes it.
func writeTempFile(dir string, data []byte) (string, error) {
	f, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return "", err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = os.Chmod(f.Name(), 0644)
	}
	if err != nil {
		os.Remove(f.Name())
		return "", err
	}
	return f.Name(), nil
}

// LoadCompressedNBT reads a gzip-compressed NBT file's root compound.
// Returns nil, nil when the file does not exist.
func LoadCompressedNBT(path string) (nbt.Compound, error) {
	compressed, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	gr, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer gr.Close()
	data, err := io.ReadAll(gr)
	if err != nil {
		return nil, fmt.Errorf("gzip read: %w", err)
	}
	tag, _, err := nbt.DecodeFile(data, nbt.WithMaxBytes(0))
	if err != nil {
		return nil, fmt.Errorf("decode NBT: %w", err)
	}
	root, ok := tag.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("root is not a compound")
	}
	return root, nil
}
