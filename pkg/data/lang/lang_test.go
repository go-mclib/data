package lang_test

import (
	"testing"

	"github.com/go-mclib/data/pkg/data/lang"
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{"item.minecraft.iron_sword", "Iron Sword"},
		{"item.minecraft.diamond_sword", "Diamond Sword"},
		{"item.minecraft.apple", "Apple"},
		{"item.minecraft.golden_apple", "Golden Apple"},
		{"item.minecraft.stick", "Stick"},
		{"item.minecraft.diamond", "Diamond"},
		{"block.minecraft.stone", "Stone"},
		{"block.minecraft.dirt", "Dirt"},
		{"block.minecraft.diamond_block", "Block of Diamond"},
		{"block.minecraft.iron_block", "Block of Iron"},
		{"block.minecraft.oak_planks", "Oak Planks"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := lang.Translate(tt.key); got != tt.want {
				t.Errorf("Translate(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestTranslateNotFound(t *testing.T) {
	if got := lang.Translate("nonexistent.translation.key"); got != "" {
		t.Errorf("Translate for nonexistent key = %q, want empty string", got)
	}
}

func TestTranslateUI(t *testing.T) {
	// test some UI translations that are unlikely to change
	tests := []struct {
		key  string
		want string
	}{
		{"menu.singleplayer", "Singleplayer"},
		{"menu.multiplayer", "Multiplayer"},
		{"menu.options", "Options..."},
		{"menu.quit", "Quit Game"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := lang.Translate(tt.key); got != tt.want {
				t.Errorf("Translate(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}
