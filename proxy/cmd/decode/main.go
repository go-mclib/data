// decode reads a captured packet JSON file and prints decoded packet structures
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-mclib/data/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type CapturedPacket struct {
	Direction string `json:"direction"`
	State     string `json:"state"`
	PacketID  string `json:"packet_id"` // hex format "0x00"
	RawData   string `json:"raw_data"`
}

type packetFactory func() jp.Packet

var packetRegistries = map[string]map[int]struct {
	name    string
	factory packetFactory
}{
	"handshake_c2s": {
		int(packets.C2SIntentionID): {"C2SIntention", func() jp.Packet { return &packets.C2SIntention{} }},
	},
	"status_c2s": {
		int(packets.C2SStatusRequestID):     {"C2SStatusRequest", func() jp.Packet { return &packets.C2SStatusRequest{} }},
		int(packets.C2SPingRequestStatusID): {"C2SPingRequestStatus", func() jp.Packet { return &packets.C2SPingRequestStatus{} }},
	},
	"status_s2c": {
		int(packets.S2CStatusResponseID):     {"S2CStatusResponse", func() jp.Packet { return &packets.S2CStatusResponse{} }},
		int(packets.S2CPongResponseStatusID): {"S2CPongResponseStatus", func() jp.Packet { return &packets.S2CPongResponseStatus{} }},
	},
	"login_c2s": {
		int(packets.C2SHelloID):               {"C2SHello", func() jp.Packet { return &packets.C2SHello{} }},
		int(packets.C2SKeyID):                 {"C2SKey", func() jp.Packet { return &packets.C2SKey{} }},
		int(packets.C2SCustomQueryAnswerID):   {"C2SCustomQueryAnswer", func() jp.Packet { return &packets.C2SCustomQueryAnswer{} }},
		int(packets.C2SLoginAcknowledgedID):   {"C2SLoginAcknowledged", func() jp.Packet { return &packets.C2SLoginAcknowledged{} }},
		int(packets.C2SCookieResponseLoginID): {"C2SCookieResponseLogin", func() jp.Packet { return &packets.C2SCookieResponseLogin{} }},
	},
	"login_s2c": {
		int(packets.S2CLoginDisconnectLoginID): {"S2CLoginDisconnectLogin", func() jp.Packet { return &packets.S2CLoginDisconnectLogin{} }},
		int(packets.S2CHelloID):                {"S2CHello", func() jp.Packet { return &packets.S2CHello{} }},
		int(packets.S2CLoginFinishedID):        {"S2CLoginFinished", func() jp.Packet { return &packets.S2CLoginFinished{} }},
		int(packets.S2CLoginCompressionID):     {"S2CLoginCompression", func() jp.Packet { return &packets.S2CLoginCompression{} }},
		int(packets.S2CCustomQueryID):          {"S2CCustomQuery", func() jp.Packet { return &packets.S2CCustomQuery{} }},
		int(packets.S2CCookieRequestLoginID):   {"S2CCookieRequestLogin", func() jp.Packet { return &packets.S2CCookieRequestLogin{} }},
	},
	"configuration_c2s": {
		int(packets.C2SClientInformationConfigurationID): {"C2SClientInformationConfiguration", func() jp.Packet { return &packets.C2SClientInformationConfiguration{} }},
		int(packets.C2SCookieResponseConfigurationID):    {"C2SCookieResponseConfiguration", func() jp.Packet { return &packets.C2SCookieResponseConfiguration{} }},
		int(packets.C2SCustomPayloadConfigurationID):     {"C2SCustomPayloadConfiguration", func() jp.Packet { return &packets.C2SCustomPayloadConfiguration{} }},
		int(packets.C2SFinishConfigurationID):            {"C2SFinishConfiguration", func() jp.Packet { return &packets.C2SFinishConfiguration{} }},
		int(packets.C2SKeepAliveConfigurationID):         {"C2SKeepAliveConfiguration", func() jp.Packet { return &packets.C2SKeepAliveConfiguration{} }},
		int(packets.C2SPongConfigurationID):              {"C2SPongConfiguration", func() jp.Packet { return &packets.C2SPongConfiguration{} }},
		int(packets.C2SResourcePackConfigurationID):      {"C2SResourcePackConfiguration", func() jp.Packet { return &packets.C2SResourcePackConfiguration{} }},
		int(packets.C2SSelectKnownPacksID):               {"C2SSelectKnownPacks", func() jp.Packet { return &packets.C2SSelectKnownPacks{} }},
	},
	"configuration_s2c": {
		int(packets.S2CCookieRequestConfigurationID):       {"S2CCookieRequestConfiguration", func() jp.Packet { return &packets.S2CCookieRequestConfiguration{} }},
		int(packets.S2CCustomPayloadConfigurationID):       {"S2CCustomPayloadConfiguration", func() jp.Packet { return &packets.S2CCustomPayloadConfiguration{} }},
		int(packets.S2CDisconnectConfigurationID):          {"S2CDisconnectConfiguration", func() jp.Packet { return &packets.S2CDisconnectConfiguration{} }},
		int(packets.S2CFinishConfigurationID):              {"S2CFinishConfiguration", func() jp.Packet { return &packets.S2CFinishConfiguration{} }},
		int(packets.S2CKeepAliveConfigurationID):           {"S2CKeepAliveConfiguration", func() jp.Packet { return &packets.S2CKeepAliveConfiguration{} }},
		int(packets.S2CPingConfigurationID):                {"S2CPingConfiguration", func() jp.Packet { return &packets.S2CPingConfiguration{} }},
		int(packets.S2CResetChatID):                        {"S2CResetChat", func() jp.Packet { return &packets.S2CResetChat{} }},
		int(packets.S2CResourcePackPopConfigurationID):     {"S2CResourcePackPopConfiguration", func() jp.Packet { return &packets.S2CResourcePackPopConfiguration{} }},
		int(packets.S2CResourcePackPushConfigurationID):    {"S2CResourcePackPushConfiguration", func() jp.Packet { return &packets.S2CResourcePackPushConfiguration{} }},
		int(packets.S2CStoreCookieConfigurationID):         {"S2CStoreCookieConfiguration", func() jp.Packet { return &packets.S2CStoreCookieConfiguration{} }},
		int(packets.S2CTransferConfigurationID):            {"S2CTransferConfiguration", func() jp.Packet { return &packets.S2CTransferConfiguration{} }},
		int(packets.S2CUpdateEnabledFeaturesID):            {"S2CUpdateEnabledFeatures", func() jp.Packet { return &packets.S2CUpdateEnabledFeatures{} }},
		int(packets.S2CUpdateTagsConfigurationID):          {"S2CUpdateTagsConfiguration", func() jp.Packet { return &packets.S2CUpdateTagsConfiguration{} }},
		int(packets.S2CSelectKnownPacksID):                 {"S2CSelectKnownPacks", func() jp.Packet { return &packets.S2CSelectKnownPacks{} }},
		int(packets.S2CCustomReportDetailsConfigurationID): {"S2CCustomReportDetailsConfiguration", func() jp.Packet { return &packets.S2CCustomReportDetailsConfiguration{} }},
		int(packets.S2CServerLinksConfigurationID):         {"S2CServerLinksConfiguration", func() jp.Packet { return &packets.S2CServerLinksConfiguration{} }},
		int(packets.S2CClearDialogConfigurationID):         {"S2CClearDialogConfiguration", func() jp.Packet { return &packets.S2CClearDialogConfiguration{} }},
		int(packets.S2CCodeOfConductID):                    {"S2CCodeOfConduct", func() jp.Packet { return &packets.S2CCodeOfConduct{} }},
	},
}

func parsePacketID(s string) (int, error) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		id, err := strconv.ParseInt(s[2:], 16, 32)
		return int(id), err
	}
	id, err := strconv.ParseInt(s, 10, 32)
	return int(id), err
}

func formatValue(v reflect.Value, indent string) string {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return formatValue(v.Elem(), indent)

	case reflect.Struct:
		t := v.Type()
		typeName := t.Name()

		// handle special types
		switch typeName {
		case "UUID":
			// format UUID as standard hex string
			if v.Len() == 16 {
				bytes := make([]byte, 16)
				for i := 0; i < 16; i++ {
					bytes[i] = byte(v.Index(i).Uint())
				}
				return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:16])
			}
		}

		// check for PrefixedOptional
		if strings.HasPrefix(typeName, "PrefixedOptional") {
			hasValue := v.FieldByName("HasValue")
			value := v.FieldByName("Value")
			if hasValue.IsValid() && hasValue.Bool() {
				return formatValue(value, indent)
			}
			return "None"
		}

		// regular struct
		var sb strings.Builder
		sb.WriteString("{\n")
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}
			fieldValue := v.Field(i)
			sb.WriteString(indent)
			sb.WriteString("  ")
			sb.WriteString(field.Name)
			sb.WriteString(" ")
			sb.WriteString(field.Type.Name())
			if field.Type.Name() == "" {
				sb.WriteString(field.Type.String())
			}
			sb.WriteString(" = ")
			sb.WriteString(formatValue(fieldValue, indent+"  "))
			sb.WriteString("\n")
		}
		sb.WriteString(indent)
		sb.WriteString("}")
		return sb.String()

	case reflect.Slice:
		if v.Len() == 0 {
			return "[]"
		}
		// for byte slices, show hex
		if v.Type().Elem().Kind() == reflect.Uint8 {
			bytes := make([]byte, v.Len())
			for i := 0; i < v.Len(); i++ {
				bytes[i] = byte(v.Index(i).Uint())
			}
			if len(bytes) > 64 {
				return fmt.Sprintf("0x%s... (%d bytes)", hex.EncodeToString(bytes[:64]), len(bytes))
			}
			return fmt.Sprintf("0x%s", hex.EncodeToString(bytes))
		}
		// for other slices
		var sb strings.Builder
		sb.WriteString("[\n")
		for i := 0; i < v.Len(); i++ {
			sb.WriteString(indent)
			sb.WriteString("  ")
			sb.WriteString(formatValue(v.Index(i), indent+"  "))
			if i < v.Len()-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
		sb.WriteString(indent)
		sb.WriteString("]")
		return sb.String()

	case reflect.Array:
		// for UUID (16-byte array)
		if v.Type().Elem().Kind() == reflect.Uint8 && v.Len() == 16 {
			bytes := make([]byte, 16)
			for i := 0; i < 16; i++ {
				bytes[i] = byte(v.Index(i).Uint())
			}
			return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:16])
		}
		// other arrays
		var sb strings.Builder
		sb.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(formatValue(v.Index(i), indent))
		}
		sb.WriteString("]")
		return sb.String()

	case reflect.String:
		s := v.String()
		if len(s) > 100 {
			return fmt.Sprintf("%q... (%d chars)", s[:100], len(s))
		}
		return fmt.Sprintf("%q", s)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", v.Float())

	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())

	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

func formatPacket(p jp.Packet) string {
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	var sb strings.Builder

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldValue := v.Field(i)
		sb.WriteString("  ")
		sb.WriteString(field.Name)
		sb.WriteString(" ")
		typeName := field.Type.Name()
		if typeName == "" {
			typeName = field.Type.String()
		}
		sb.WriteString(typeName)
		sb.WriteString(" = ")
		sb.WriteString(formatValue(fieldValue, "  "))
		sb.WriteString("\n")
	}

	return sb.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: decode <capture.json>\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var captured []CapturedPacket
	if err := json.Unmarshal(data, &captured); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Decoding %d packets from %s\n\n", len(captured), filename)

	for i, cap := range captured {
		packetID, err := parsePacketID(cap.PacketID)
		if err != nil {
			fmt.Printf("// [%d] WARNING: invalid packet ID %q\n\n", i, cap.PacketID)
			continue
		}

		registryKey := cap.State + "_" + cap.Direction
		registry, ok := packetRegistries[registryKey]
		if !ok {
			fmt.Printf("// [%d] WARNING: unknown state/direction %s (packet %s)\n\n", i, registryKey, cap.PacketID)
			continue
		}

		entry, ok := registry[packetID]
		if !ok {
			fmt.Printf("// [%d] WARNING: unknown packet %s in %s\n\n", i, cap.PacketID, registryKey)
			continue
		}

		rawData, err := hex.DecodeString(cap.RawData)
		if err != nil {
			fmt.Printf("// [%d] WARNING: invalid hex data for %s: %v\n\n", i, entry.name, err)
			continue
		}

		p := entry.factory()
		buf := ns.NewReader(rawData)
		if err := p.Read(buf); err != nil {
			fmt.Printf("// [%d] WARNING: failed to decode %s: %v\n", i, entry.name, err)
			fmt.Printf("//     raw data: %s\n\n", cap.RawData)
			continue
		}

		fmt.Printf("// [%d] %s %s\n", i, cap.Direction, cap.PacketID)
		fmt.Printf("%s {\n", entry.name)
		fmt.Print(formatPacket(p))
		fmt.Printf("}\n\n")
	}
}
