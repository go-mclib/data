// Code generated for Minecraft 1.21.11 (Protocol 774); DO NOT EDIT.

package entities

// Entity metadata serializer type IDs.
const (
	SerializerBYTE                             = 0
	SerializerINT                              = 1
	SerializerLONG                             = 2
	SerializerFLOAT                            = 3
	SerializerSTRING                           = 4
	SerializerCOMPONENT                        = 5
	SerializerOPTIONAL_COMPONENT               = 6
	SerializerITEM_STACK                       = 7
	SerializerBOOLEAN                          = 8
	SerializerROTATIONS                        = 9
	SerializerBLOCK_POS                        = 10
	SerializerOPTIONAL_BLOCK_POS               = 11
	SerializerDIRECTION                        = 12
	SerializerOPTIONAL_LIVING_ENTITY_REFERENCE = 13
	SerializerBLOCK_STATE                      = 14
	SerializerOPTIONAL_BLOCK_STATE             = 15
	SerializerPARTICLE                         = 16
	SerializerPARTICLES                        = 17
	SerializerVILLAGER_DATA                    = 18
	SerializerOPTIONAL_UNSIGNED_INT            = 19
	SerializerPOSE                             = 20
	SerializerCAT_VARIANT                      = 21
	SerializerCOW_VARIANT                      = 22
	SerializerWOLF_VARIANT                     = 23
	SerializerWOLF_SOUND_VARIANT               = 24
	SerializerFROG_VARIANT                     = 25
	SerializerPIG_VARIANT                      = 26
	SerializerCHICKEN_VARIANT                  = 27
	SerializerZOMBIE_NAUTILUS_VARIANT          = 28
	SerializerOPTIONAL_GLOBAL_POS              = 29
	SerializerPAINTING_VARIANT                 = 30
	SerializerSNIFFER_STATE                    = 31
	SerializerARMADILLO_STATE                  = 32
	SerializerCOPPER_GOLEM_STATE               = 33
	SerializerWEATHERING_COPPER_STATE          = 34
	SerializerVECTOR3                          = 35
	SerializerQUATERNION                       = 36
	SerializerRESOLVABLE_PROFILE               = 37
	SerializerHUMANOID_ARM                     = 38
)

// serializerNames maps serializer IDs to names.
var serializerNames = map[int32]string{
	0:  "BYTE",
	1:  "INT",
	2:  "LONG",
	3:  "FLOAT",
	4:  "STRING",
	5:  "COMPONENT",
	6:  "OPTIONAL_COMPONENT",
	7:  "ITEM_STACK",
	8:  "BOOLEAN",
	9:  "ROTATIONS",
	10: "BLOCK_POS",
	11: "OPTIONAL_BLOCK_POS",
	12: "DIRECTION",
	13: "OPTIONAL_LIVING_ENTITY_REFERENCE",
	14: "BLOCK_STATE",
	15: "OPTIONAL_BLOCK_STATE",
	16: "PARTICLE",
	17: "PARTICLES",
	18: "VILLAGER_DATA",
	19: "OPTIONAL_UNSIGNED_INT",
	20: "POSE",
	21: "CAT_VARIANT",
	22: "COW_VARIANT",
	23: "WOLF_VARIANT",
	24: "WOLF_SOUND_VARIANT",
	25: "FROG_VARIANT",
	26: "PIG_VARIANT",
	27: "CHICKEN_VARIANT",
	28: "ZOMBIE_NAUTILUS_VARIANT",
	29: "OPTIONAL_GLOBAL_POS",
	30: "PAINTING_VARIANT",
	31: "SNIFFER_STATE",
	32: "ARMADILLO_STATE",
	33: "COPPER_GOLEM_STATE",
	34: "WEATHERING_COPPER_STATE",
	35: "VECTOR3",
	36: "QUATERNION",
	37: "RESOLVABLE_PROFILE",
	38: "HUMANOID_ARM",
}

// serializerWireTypes maps serializer IDs to wire types.
var serializerWireTypes = map[int32]string{
	0:  "byte",
	1:  "varint",
	2:  "varlong",
	3:  "float32",
	4:  "string",
	5:  "nbt",
	6:  "optional_nbt",
	7:  "slot",
	8:  "bool",
	9:  "rotations",
	10: "position",
	11: "optional_position",
	12: "varint",
	13: "optional_uuid",
	14: "varint",
	15: "varint",
	16: "particle",
	17: "particle_list",
	18: "villager_data",
	19: "optional_varint",
	20: "varint",
	21: "varint",
	22: "varint",
	23: "varint",
	24: "varint",
	25: "varint",
	26: "varint",
	27: "varint",
	28: "varint",
	29: "optional_global_pos",
	30: "id_or_inline",
	31: "varint",
	32: "varint",
	33: "varint",
	34: "varint",
	35: "vector3f",
	36: "quaternionf",
	37: "resolvable_profile",
	38: "varint",
}
