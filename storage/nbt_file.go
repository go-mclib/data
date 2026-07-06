package storage

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-mclib/nbt"
)

// SaveCompressedNBT writes a root compound to a gzip-compressed NBT file,
// creating parent directories.
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
	return os.WriteFile(path, buf.Bytes(), 0644)
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
