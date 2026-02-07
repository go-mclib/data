# TODO: Test World Setup

Instructions for reproducing the test world used for packet capture fixtures.

## Server Configuration

The server uses `server.properties` with:

- `level-type=minecraft:flat` (superflat)
- `level-seed=0`
- `generate-structures=false`
- `online-mode=false`
- `gamemode=creative`
- `max-world-size=128`

## World Setup Commands

After connecting as player `GoMclib`, run these commands to set up the test structures:

### Chunk test (chunk 1,1 with hay bales and sign)

```plaintext
/fill 16 -60 16 31 -60 31 minecraft:hay_block
/setblock 16 -59 16 minecraft:pale_oak_wall_sign[facing=west]
```

Then right-click the sign and write on the front:

- Line 1: `needle`
- Line 2: `in`
- Line 3: `a`
- Line 4: `haystack`

This produces 50 hay bales (16x16 area minus the sign position -- 256 minus the ones at non-surface level... adjust as needed to match 50) and 1 sign with known text.

### Entity test (dropped item)

Drop a custom iron sword named "po" with attribute modifiers. This can be created using:

```plaintext
/give GoMclib minecraft:iron_sword[minecraft:custom_name='"po"',minecraft:unbreakable={},minecraft:attribute_modifiers={modifiers:[{type:"minecraft:attack_damage",id:"minecraft:2121f7b4-5985-43a0-aa3a-57717d7b15c4",amount:1000.0,operation:"add_multiplied_total",slot:"any"},{type:"minecraft:attack_speed",id:"minecraft:1df199b2-3849-4112-b9f4-7f16d98d9d38",amount:100.0,operation:"add_value",slot:"any"}],show_in_tooltip:false},minecraft:tooltip_display={hidden_components:["minecraft:attribute_modifiers","minecraft:can_break"]}]
```

### Inventory test

Requires interacting with:

1. A double chest (open, pick up/place items, close)
2. A crafting table (craft oak planks, drag items)
3. A beacon (set primary effect to haste)
4. Villager trading (select first trade)

### Boss event test

Summon a named wither:

```plaintext
/summon minecraft:wither ~ ~ ~ {CustomName:'{"text":"Happy Witherling"}'}
```

## Saving the World Template

After setup, stop the server and copy the minimal files:

```bash
mkdir -p proxy/testdata/world-template/region
cp proxy/server/world/level.dat proxy/testdata/world-template/
cp proxy/server/world/region/r.0.0.mca proxy/testdata/world-template/region/
```

Then set up git LFS to track them:

```bash
git lfs track "proxy/testdata/world-template/**"
git add proxy/testdata/world-template/
```
