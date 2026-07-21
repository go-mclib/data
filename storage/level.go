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

// LevelData is the subset of vanilla level.dat the server persists.
type LevelData struct {
	SpawnX, SpawnY, SpawnZ int32
	Time                   int64 // world age, never frozen
	DayTime                int64 // visual day clock
	Seed                   int64 // world generation seed
	BorderSet              bool
	BorderCenterX          float64
	BorderCenterZ          float64
	BorderSize             float64
	BorderSizeLerpTime     int64
	BorderSizeLerpTarget   float64
	BorderSafeZone         float64
	BorderDamagePerBlock   float64
	BorderWarningBlocks    float64
	BorderWarningTime      float64
}

// SaveLevel writes <dir>/level.dat (gzip-compressed NBT, vanilla layout).
func SaveLevel(dir string, ld *LevelData) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	border := *ld
	if !border.BorderSet {
		border.BorderSize = 59999968
		border.BorderSizeLerpTarget = 59999968
		border.BorderSafeZone = 5
		border.BorderDamagePerBlock = 0.2
		border.BorderWarningBlocks = 5
		border.BorderWarningTime = 15
	}
	root := nbt.Compound{
		"Data": nbt.Compound{
			"DataVersion": nbt.Int(DataVersion),
			"SpawnX":      nbt.Int(ld.SpawnX),
			"SpawnY":      nbt.Int(ld.SpawnY),
			"SpawnZ":      nbt.Int(ld.SpawnZ),
			"Time":        nbt.Long(ld.Time),
			"DayTime":     nbt.Long(ld.DayTime),
			"WorldGenSettings": nbt.Compound{
				"seed": nbt.Long(ld.Seed),
			},
			"BorderCenterX":        nbt.Double(border.BorderCenterX),
			"BorderCenterZ":        nbt.Double(border.BorderCenterZ),
			"BorderSize":           nbt.Double(border.BorderSize),
			"BorderSizeLerpTime":   nbt.Long(border.BorderSizeLerpTime),
			"BorderSizeLerpTarget": nbt.Double(border.BorderSizeLerpTarget),
			"BorderSafeZone":       nbt.Double(border.BorderSafeZone),
			"BorderDamagePerBlock": nbt.Double(border.BorderDamagePerBlock),
			"BorderWarningBlocks":  nbt.Double(border.BorderWarningBlocks),
			"BorderWarningTime":    nbt.Double(border.BorderWarningTime),
		},
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
	return os.WriteFile(filepath.Join(dir, "level.dat"), buf.Bytes(), 0644)
}

// LoadLevel reads <dir>/level.dat. Returns nil, nil if the file doesn't exist.
func LoadLevel(dir string) (*LevelData, error) {
	compressed, err := os.ReadFile(filepath.Join(dir, "level.dat"))
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
		return nil, fmt.Errorf("root is not compound")
	}
	d := root.GetCompound("Data")
	ld := &LevelData{
		SpawnX:  d.GetInt("SpawnX"),
		SpawnY:  d.GetInt("SpawnY"),
		SpawnZ:  d.GetInt("SpawnZ"),
		Time:    d.GetLong("Time"),
		DayTime: d.GetLong("DayTime"),
		Seed:    d.GetCompound("WorldGenSettings").GetLong("seed"),
	}
	if d.Get("BorderSize") != nil {
		ld.BorderSet = true
		ld.BorderCenterX = d.GetDouble("BorderCenterX")
		ld.BorderCenterZ = d.GetDouble("BorderCenterZ")
		ld.BorderSize = d.GetDouble("BorderSize")
		ld.BorderSizeLerpTime = d.GetLong("BorderSizeLerpTime")
		ld.BorderSizeLerpTarget = d.GetDouble("BorderSizeLerpTarget")
		ld.BorderSafeZone = d.GetDouble("BorderSafeZone")
		ld.BorderDamagePerBlock = d.GetDouble("BorderDamagePerBlock")
		ld.BorderWarningBlocks = d.GetDouble("BorderWarningBlocks")
		ld.BorderWarningTime = d.GetDouble("BorderWarningTime")
	}
	return ld, nil
}
