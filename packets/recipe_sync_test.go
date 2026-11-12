package packets_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestUpdateRecipesWireLayout(t *testing.T) {
	packet := &packets.S2CUpdateRecipes{
		PropertySets: []packets.RecipePropertySet{{
			Key: "minecraft:furnace_input", Items: []ns.VarInt{1, 300},
		}},
		StonecutterRecipes: []packets.StonecutterRecipeEntry{{
			Ingredient: packets.RecipeIngredient{Items: []ns.VarInt{4}},
			ResultItem: 6, ResultCount: 2,
		}},
	}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	want := append([]byte{0x01, 0x17}, []byte("minecraft:furnace_input")...)
	want = append(want,
		0x02, 0x01, 0xac, 0x02,
		0x01,
		0x02, 0x04,
		0x05, 0x06, 0x02, 0x00, 0x00,
	)
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("wire = %x, want %x", buf.Bytes(), want)
	}

	var decoded packets.S2CUpdateRecipes
	reader := ns.NewReader(want)
	if err := decoded.Read(reader); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(decoded.PropertySets) != 1 || len(decoded.PropertySets[0].Items) != 2 || decoded.PropertySets[0].Items[1] != 300 ||
		len(decoded.StonecutterRecipes) != 1 || decoded.StonecutterRecipes[0].ResultCount != 2 {
		t.Fatalf("decoded = %+v", decoded)
	}
	if _, err := reader.ReadByte(); err != io.EOF {
		t.Fatalf("trailing data: %v", err)
	}
}

func TestEmptyInitialRecipeBookWire(t *testing.T) {
	add := &packets.S2CRecipeBookAdd{Replace: true}
	buf := ns.NewWriter()
	if err := add.Write(buf); err != nil {
		t.Fatalf("add Write: %v", err)
	}
	if want := []byte{0, 1}; !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("add wire = %x, want %x", buf.Bytes(), want)
	}
	var decoded packets.S2CRecipeBookAdd
	if err := decoded.Read(ns.NewReader(buf.Bytes())); err != nil || !bool(decoded.Replace) {
		t.Fatalf("add decode = %+v, %v", decoded, err)
	}

	remove := &packets.S2CRecipeBookRemove{Recipes: []ns.VarInt{1, 300}}
	buf = ns.NewWriter()
	if err := remove.Write(buf); err != nil {
		t.Fatalf("remove Write: %v", err)
	}
	if want := []byte{2, 1, 0xac, 0x02}; !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("remove wire = %x, want %x", buf.Bytes(), want)
	}
}

func TestRecipeBookEntryWire(t *testing.T) {
	packet := &packets.S2CRecipeBookAdd{
		Entries: []packets.RecipeBookEntry{{
			DisplayID: 7,
			Display: packets.RecipeDisplay{
				Type: 1, Width: 1, Height: 1,
				Ingredients: []packets.RecipeSlotDisplay{{Type: 10, Children: []packets.RecipeSlotDisplay{{Type: 4, ItemID: 1}}}},
				Result:      packets.RecipeSlotDisplay{Type: 5, ItemID: 2, Count: 4},
				Station:     packets.RecipeSlotDisplay{Type: 4, ItemID: 3},
			},
			Group: -1, Category: 2,
			HasRequirements: true,
			Requirements:    []packets.RecipeIngredient{{Items: []ns.VarInt{1}}},
			Flags:           3,
		}},
		Replace: true,
	}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	want := []byte{1, 7, 1, 1, 1, 1, 10, 1, 4, 1, 5, 2, 4, 0, 0, 4, 3, 0, 2, 1, 1, 2, 1, 3, 1}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("wire = %x, want %x", buf.Bytes(), want)
	}
	var decoded packets.S2CRecipeBookAdd
	if err := decoded.Read(ns.NewReader(buf.Bytes())); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(decoded.Entries) != 1 || decoded.Entries[0].Display.Width != 1 || decoded.Entries[0].Requirements[0].Items[0] != 1 || decoded.Entries[0].Flags != 3 {
		t.Fatalf("decoded = %+v", decoded)
	}
}

func TestPotionSlotDisplayCarriesChild(t *testing.T) {
	result := ns.NewSlot(6, 1)
	result.Components.Add = []ns.RawSlotComponent{{ID: items.ComponentDamage, Data: []byte{3}}}
	packet := &packets.S2CRecipeBookAdd{Entries: []packets.RecipeBookEntry{{
		Display: packets.RecipeDisplay{
			Type:        0,
			Ingredients: []packets.RecipeSlotDisplay{{Type: 2, Children: []packets.RecipeSlotDisplay{{Type: 4, ItemID: 5}}}},
			Result:      packets.RecipeSlotDisplay{Type: 5, Item: result}, Station: packets.RecipeSlotDisplay{Type: 4, ItemID: 7},
		},
		Group: -1,
	}}}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	var decoded packets.S2CRecipeBookAdd
	if err := decoded.Read(ns.NewReader(buf.Bytes())); err != nil {
		t.Fatalf("Read: %v", err)
	}
	child := decoded.Entries[0].Display.Ingredients[0].Children
	if len(child) != 1 || child[0].ItemID != 5 {
		t.Fatalf("potion display child = %+v", child)
	}
	if !items.SameStack(decoded.Entries[0].Display.Result.Item, result) {
		t.Fatalf("recipe result template = %+v", decoded.Entries[0].Display.Result.Item)
	}
}

func TestRecipeIngredientTagWire(t *testing.T) {
	packet := &packets.S2CUpdateRecipes{StonecutterRecipes: []packets.StonecutterRecipeEntry{{
		Ingredient: packets.RecipeIngredient{Tag: "minecraft:stone_crafting_materials"},
		ResultItem: 1, ResultCount: 1,
	}}}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	reader := ns.NewReader(buf.Bytes())
	var decoded packets.S2CUpdateRecipes
	if err := decoded.Read(reader); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if got := decoded.StonecutterRecipes[0].Ingredient.Tag; got != "minecraft:stone_crafting_materials" {
		t.Fatalf("tag = %q", got)
	}
}

func TestRecipeTagDisplayHolderSetWire(t *testing.T) {
	packet := &packets.S2CRecipeBookAdd{Entries: []packets.RecipeBookEntry{{
		DisplayID: 7,
		Display: packets.RecipeDisplay{
			Type: 0,
			Ingredients: []packets.RecipeSlotDisplay{
				{Type: 6, Holders: packets.RecipeIngredient{Tag: "minecraft:planks"}},
				{Type: 6, Holders: packets.RecipeIngredient{Items: []ns.VarInt{1, 300}}},
			},
			Result:  packets.RecipeSlotDisplay{Type: 4, ItemID: 2},
			Station: packets.RecipeSlotDisplay{Type: 4, ItemID: 3},
		},
		Group: -1, Category: 2, Flags: 3,
	}}, Replace: true}
	want := append([]byte{1, 7, 0, 2, 6, 0, 16}, []byte("minecraft:planks")...)
	want = append(want, 6, 3, 1, 0xac, 2, 4, 2, 4, 3, 0, 2, 0, 3, 1)
	writer := ns.NewWriter()
	if err := packet.Write(writer); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(writer.Bytes(), want) {
		t.Fatalf("holder-set display wire = %x, want %x", writer.Bytes(), want)
	}
	var decoded packets.S2CRecipeBookAdd
	reader := ns.NewReader(want)
	if err := decoded.Read(reader); err != nil {
		t.Fatal(err)
	}
	ingredients := decoded.Entries[0].Display.Ingredients
	if ingredients[0].Holders.Tag != "minecraft:planks" ||
		len(ingredients[1].Holders.Items) != 2 || ingredients[1].Holders.Items[0] != 1 || ingredients[1].Holders.Items[1] != 300 {
		t.Fatalf("decoded holder sets = %+v", ingredients)
	}
	if decoded.Entries[0].Flags != 3 || !bool(decoded.Replace) {
		t.Fatal("holder-set display desynchronized the following fields")
	}
	if _, err := reader.ReadByte(); err != io.EOF {
		t.Fatalf("trailing data: %v", err)
	}
}
