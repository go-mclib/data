import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.mojang.serialization.JsonOps;
import net.minecraft.SharedConstants;
import net.minecraft.core.BlockPos;
import net.minecraft.core.Direction;
import net.minecraft.core.Registry;
import net.minecraft.core.RegistryAccess;
import net.minecraft.core.registries.BuiltInRegistries;
import net.minecraft.resources.Identifier;
import net.minecraft.server.Bootstrap;
import net.minecraft.server.packs.PackType;
import net.minecraft.server.packs.repository.ServerPacksSource;
import net.minecraft.server.packs.resources.MultiPackResourceManager;
import net.minecraft.tags.BlockTags;
import net.minecraft.tags.TagLoader;
import net.minecraft.world.effect.MobEffectInstance;
import net.minecraft.world.entity.EntityDimensions;
import net.minecraft.world.entity.EntityType;
import net.minecraft.world.item.Item;
import net.minecraft.world.level.EmptyBlockGetter;
import net.minecraft.world.level.block.Block;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.world.level.block.FireBlock;
import net.minecraft.world.level.block.state.BlockState;
import net.minecraft.world.level.block.state.properties.Property;
import net.minecraft.world.level.material.FlowingFluid;
import net.minecraft.world.level.material.Fluids;
import net.minecraft.world.phys.AABB;
import net.minecraft.world.phys.shapes.VoxelShape;

import java.io.Writer;
import java.lang.reflect.Method;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

// standalone dumper: boots Minecraft's registries straight from the server jar
// (real Mojang names — no Fabric, no remapping) and writes the game's own data
// as JSON. The actual game code is the single source of truth.
//
// covers the data that previously needed PrismarineJS / source scraping or
// hand-kept Go tables (collision shapes, entity hitboxes, hardness, mob
// categories, friction, light emission/dampening, effect colors, crafting
// remainders) plus the built-in registry id maps. Vanilla tags are bound before querying state
// predicates; dynamic item components and biomes come from server datagen.
public final class Dump {
	private static final Gson GSON = new GsonBuilder().setPrettyPrinting().disableHtmlEscaping().create();
	private static final Method FIRE_IGNITE_ODDS = fireOddsMethod("getIgniteOdds");
	private static final Method FIRE_BURN_ODDS = fireOddsMethod("getBurnOdds");

	public static void main(final String[] args) throws Exception {
		final Path out = Path.of(args.length > 0 ? args[0] : "go-mclib-dump");
		Files.createDirectories(out);

		SharedConstants.tryDetectVersion();
		Bootstrap.bootStrap();
		final var packs = ServerPacksSource.createVanillaTrustedRepository();
		packs.reload();
		packs.setSelected(List.of("vanilla"));
		try (var resources = new MultiPackResourceManager(PackType.SERVER_DATA, packs.openAllSelected())) {
			TagLoader.loadTagsForExistingRegistries(resources, RegistryAccess.fromRegistryOfRegistries(BuiltInRegistries.REGISTRY))
				.forEach(Registry.PendingTags::apply);
		}

		write(out, "blocks.json", blocks());
		write(out, "entities.json", entities());
		write(out, "registries.json", registries());
		write(out, "potions.json", potions());
		write(out, "effects.json", effects());
		write(out, "items.json", items());

		System.out.println("[mcdump] dump written to " + out.toAbsolutePath());
	}

	private static JsonObject potions() {
		final JsonObject result = new JsonObject();
		for (final var potion : sortedByKey(BuiltInRegistries.POTION)) {
			final JsonArray effects = new JsonArray();
			for (final var effect : potion.getEffects()) {
				effects.add(MobEffectInstance.CODEC.encodeStart(JsonOps.INSTANCE, effect).getOrThrow());
			}
			result.add(BuiltInRegistries.POTION.getKey(potion).toString(), effects);
		}
		return result;
	}

	// effects: every mob effect's particle/potion color.
	private static JsonObject effects() {
		final JsonObject result = new JsonObject();
		for (final var effect : sortedByKey(BuiltInRegistries.MOB_EFFECT)) {
			final JsonObject obj = new JsonObject();
			obj.addProperty("color", effect.getColor());
			result.add(BuiltInRegistries.MOB_EFFECT.getKey(effect).toString(), obj);
		}
		return result;
	}

	// items: the crafting remainder of every item that leaves one behind.
	private static JsonObject items() {
		final JsonObject result = new JsonObject();
		for (final Item item : sortedByKey(BuiltInRegistries.ITEM)) {
			final var remainder = item.getCraftingRemainder();
			if (remainder == null) {
				continue;
			}
			final JsonObject obj = new JsonObject();
			obj.addProperty("craftingRemainder", BuiltInRegistries.ITEM.getKey(remainder.item().value()).toString());
			result.add(BuiltInRegistries.ITEM.getKey(item).toString(), obj);
		}
		return result;
	}

	// blocks: every state's palette id, properties, collision shape (deduped into
	// a shapes table), replaceability, flammability, top-face sturdiness, hardness, and tool requirement.
	private static JsonObject blocks() {
		final JsonObject root = new JsonObject();
		final JsonObject blocks = new JsonObject();
		final List<JsonArray> shapeTable = new ArrayList<>();
		final List<JsonArray> coordinateTable = new ArrayList<>();
		final Map<String, Integer> shapeIndex = new HashMap<>();

		for (final Block block : sortedByKey(BuiltInRegistries.BLOCK)) {
			final Identifier key = BuiltInRegistries.BLOCK.getKey(block);
			final BlockState def = block.defaultBlockState();

			final JsonObject bobj = new JsonObject();
			bobj.addProperty("protocolId", BuiltInRegistries.BLOCK.getId(block));
			bobj.addProperty("hardness", round(def.getDestroySpeed(EmptyBlockGetter.INSTANCE, BlockPos.ZERO)));
			bobj.addProperty("requiresTool", def.requiresCorrectToolForDrops());
			bobj.addProperty("friction", block.getFriction());

			final JsonArray states = new JsonArray();
			for (final BlockState state : block.getStateDefinition().getPossibleStates()) {
				final JsonObject sobj = new JsonObject();
				sobj.addProperty("id", Block.getId(state)); // global block-state (palette) id
				final JsonObject props = new JsonObject();
				for (final Property<?> prop : state.getProperties()) {
					props.addProperty(prop.getName(), valueName(state, prop));
				}
				sobj.add("properties", props);
				final VoxelShape shape = state.getCollisionShape(EmptyBlockGetter.INSTANCE, BlockPos.ZERO);
				sobj.addProperty("shape", shapeIndexOf(shape, shapeTable, coordinateTable, shapeIndex));
				sobj.addProperty("replaceable", state.canBeReplaced());
				sobj.addProperty("igniteOdds", fireOdds(FIRE_IGNITE_ODDS, state));
				sobj.addProperty("burnOdds", fireOdds(FIRE_BURN_ODDS, state));
				sobj.addProperty("sturdyTop", state.isFaceSturdy(EmptyBlockGetter.INSTANCE, BlockPos.ZERO, Direction.UP));
				sobj.addProperty("blocksMotion", state.is(BlockTags.BLOCKS_MOTION));
				sobj.addProperty("blocksFluidFlow", state.is(BlockTags.BLOCKS_FLUID_FLOW));
				sobj.addProperty("largeCollisionShape", state.hasLargeCollisionShape());
				sobj.addProperty("suffocating", state.isSuffocating(EmptyBlockGetter.INSTANCE, BlockPos.ZERO));
				int sturdyFaces = 0;
				for (final Direction direction : Direction.values()) {
					if (state.isFaceSturdy(EmptyBlockGetter.INSTANCE, BlockPos.ZERO, direction)) {
						sturdyFaces |= 1 << direction.get3DDataValue();
					}
				}
				sobj.addProperty("sturdyFaces", sturdyFaces);
				final var fluid = state.getFluidState();
				// Both source and flowing variants share the same contained-fluid kind.
				sobj.addProperty("fluidKind", fluid.getType().isSame(Fluids.WATER) ? 1 : fluid.getType().isSame(Fluids.LAVA) ? 2 : 0);
				sobj.addProperty("fluidAmount", fluid.getAmount());
				sobj.addProperty("fluidFalling", !fluid.isEmpty() && fluid.hasProperty(FlowingFluid.FALLING) && fluid.getValue(FlowingFluid.FALLING));
				sobj.addProperty("lightEmission", state.getLightEmission());
				sobj.addProperty("lightDampening", state.getLightDampening());
				sobj.addProperty("mapColor", Byte.toUnsignedInt(state.getMapColor(EmptyBlockGetter.INSTANCE, BlockPos.ZERO).getPackedId(net.minecraft.world.level.material.MapColor.Brightness.NORMAL)) >> 2);
				states.add(sobj);
			}
			bobj.add("states", states);
			blocks.add(key.toString(), bobj);
		}

		root.add("blocks", blocks);
		final JsonArray shapes = new JsonArray();
		for (final JsonArray s : shapeTable) {
			shapes.add(s);
		}
		root.add("shapes", shapes);
		final JsonArray coordinates = new JsonArray();
		for (final JsonArray coords : coordinateTable) {
			coordinates.add(coords);
		}
		root.add("shapeCoordinates", coordinates);
		return root;
	}

	private static int shapeIndexOf(final VoxelShape shape, final List<JsonArray> table, final List<JsonArray> coordinateTable, final Map<String, Integer> index) {
		final JsonArray boxes = new JsonArray();
		final StringBuilder key = new StringBuilder();
		for (final AABB b : shape.toAabbs()) {
			final JsonArray box = new JsonArray();
			for (final double d : new double[]{b.minX, b.minY, b.minZ, b.maxX, b.maxY, b.maxZ}) {
				box.add(d);
			}
			boxes.add(box);
			key.append(box).append('|');
		}
		final JsonArray coordinates = new JsonArray();
		for (final Direction.Axis axis : Direction.Axis.values()) {
			final JsonArray coords = new JsonArray();
			for (final double coord : shape.getCoords(axis)) {
				coords.add(coord);
			}
			coordinates.add(coords);
		}
		key.append(coordinates);
		final Integer existing = index.get(key.toString());
		if (existing != null) {
			return existing;
		}
		final int id = table.size();
		table.add(boxes);
		coordinateTable.add(coordinates);
		index.put(key.toString(), id);
		return id;
	}

	@SuppressWarnings({"unchecked", "rawtypes"})
	private static String valueName(final BlockState state, final Property prop) {
		return prop.getName(state.getValue(prop));
	}

	// entities: hitbox width/height, eye height, mob category, client tracking range.
	private static JsonObject entities() {
		final JsonObject root = new JsonObject();
		for (final EntityType<?> type : sortedByKey(BuiltInRegistries.ENTITY_TYPE)) {
			final EntityDimensions dim = type.getDimensions();
			final JsonObject obj = new JsonObject();
			obj.addProperty("protocolId", BuiltInRegistries.ENTITY_TYPE.getId(type));
			obj.addProperty("width", round(dim.width()));
			obj.addProperty("height", round(dim.height()));
			obj.addProperty("eyeHeight", round(dim.eyeHeight()));
			obj.addProperty("category", type.getCategory().getName());
			obj.addProperty("trackingRange", type.clientTrackingRange());
			root.add(BuiltInRegistries.ENTITY_TYPE.getKey(type).toString(), obj);
		}
		return root;
	}

	// registries: every built-in registry -> { entry name: protocol id }. (datapack
	// registries such as biomes need a loaded world and still come from the datagen.)
	private static JsonObject registries() {
		final JsonObject root = new JsonObject();
		final TreeMap<String, JsonObject> byName = new TreeMap<>();
		for (final Registry<?> reg : BuiltInRegistries.REGISTRY) {
			byName.put(reg.key().identifier().toString(), registryEntries(reg));
		}
		byName.forEach(root::add);
		return root;
	}

	@SuppressWarnings({"unchecked", "rawtypes"})
	private static JsonObject registryEntries(final Registry reg) {
		final TreeMap<Integer, String> byId = new TreeMap<>();
		for (final Object value : reg) {
			final Identifier key = reg.getKey(value);
			if (key != null) {
				byId.put(reg.getId(value), key.toString());
			}
		}
		final JsonObject entries = new JsonObject();
		byId.forEach((id, name) -> entries.addProperty(name, id));
		return entries;
	}

	// --- helpers ---

	private static Method fireOddsMethod(final String name) {
		try {
			final Method method = FireBlock.class.getDeclaredMethod(name, BlockState.class);
			method.setAccessible(true);
			return method;
		} catch (final ReflectiveOperationException exception) {
			throw new ExceptionInInitializerError(exception);
		}
	}

	private static int fireOdds(final Method method, final BlockState state) {
		try {
			return (int) method.invoke(Blocks.FIRE, state);
		} catch (final ReflectiveOperationException exception) {
			throw new IllegalStateException("cannot inspect fire odds", exception);
		}
	}

	private static <T> List<T> sortedByKey(final Registry<T> reg) {
		final List<T> list = new ArrayList<>();
		reg.forEach(list::add);
		list.sort(Comparator.comparing(v -> reg.getKey(v).toString()));
		return list;
	}

	private static double round(final double d) {
		return Math.round(d * 1.0e6) / 1.0e6;
	}

	private static void write(final Path dir, final String name, final JsonElement json) throws Exception {
		try (final Writer w = Files.newBufferedWriter(dir.resolve(name), StandardCharsets.UTF_8)) {
			GSON.toJson(json, w);
		}
	}
}
