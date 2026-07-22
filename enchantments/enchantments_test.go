package enchantments_test

import (
	"math"
	"testing"

	"github.com/go-mclib/data/enchantments"
)

func TestGeneratedLootingEquipmentDropEffect(t *testing.T) {
	looting := enchantments.Get("minecraft:looting")
	if looting == nil {
		t.Fatal("looting definition is absent")
	}
	if !looting.MatchesSlot("mainhand") || looting.MatchesSlot("offhand") {
		t.Fatalf("looting slots = %v", looting.Slots)
	}
	if len(looting.EquipmentDropEffects) != 1 {
		t.Fatalf("looting equipment drop effects = %+v", looting.EquipmentDropEffects)
	}
	effect := looting.EquipmentDropEffects[0]
	if effect.Enchanted != "attacker" || effect.Affected != "victim" || !effect.RequiresAttackerPlayer {
		t.Fatalf("looting equipment drop effect = %+v", effect)
	}
	if got := effect.Apply(3, 0.085); math.Abs(float64(got-0.115)) > 1e-6 {
		t.Fatalf("looting III drop chance = %v, want 0.115", got)
	}
}
