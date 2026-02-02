// Reads a captured packet JSON file and prints decoded packet structures
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-mclib/data/pkg/data/items"
	"github.com/go-mclib/data/pkg/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type CapturedPacket struct {
	Direction string `json:"direction"`
	State     string `json:"state"`
	PacketID  string `json:"packet_id"` // hex format "0x00"
	RawData   string `json:"raw_data"`
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

func getPacketName(p jp.Packet) string {
	t := reflect.TypeOf(p)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func formatItemStack(stack *items.ItemStack, indent string) string {
	if stack.IsEmpty() {
		return "Empty"
	}

	var sb strings.Builder
	sb.WriteString("ItemStack {\n")
	sb.WriteString(indent)
	sb.WriteString(fmt.Sprintf("  ID: %d\n", stack.ID))
	sb.WriteString(indent)
	sb.WriteString(fmt.Sprintf("  Count: %d\n", stack.Count))
	sb.WriteString(indent)
	sb.WriteString("  Components: ")

	if stack.Components == nil {
		sb.WriteString("nil\n")
	} else {
		// use reflection to display non-zero component fields
		cv := reflect.ValueOf(stack.Components).Elem()
		ct := cv.Type()

		sb.WriteString("{\n")
		hasAny := false
		for i := 0; i < cv.NumField(); i++ {
			field := ct.Field(i)
			if !field.IsExported() {
				continue
			}
			fv := cv.Field(i)

			// skip zero values
			if fv.IsZero() {
				continue
			}

			hasAny = true
			sb.WriteString(indent)
			sb.WriteString("    ")
			sb.WriteString(field.Name)
			sb.WriteString(": ")
			sb.WriteString(formatValue(fv, indent+"    "))
			sb.WriteString("\n")
		}
		if !hasAny {
			sb.WriteString(indent)
			sb.WriteString("    (all defaults)\n")
		}
		sb.WriteString(indent)
		sb.WriteString("  }\n")
	}

	sb.WriteString(indent)
	sb.WriteString("}")
	return sb.String()
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
				for i := range 16 {
					bytes[i] = byte(v.Index(i).Uint())
				}
				return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:16])
			}
		case "Slot":
			// convert ns.Slot to items.ItemStack for better display
			if slot, ok := v.Interface().(ns.Slot); ok {
				if slot.IsEmpty() {
					return "Empty"
				}
				stack, err := items.FromSlot(slot)
				if err != nil {
					return fmt.Sprintf("Slot{decode error: %v}", err)
				}
				return formatItemStack(stack, indent)
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
			for i := range 16 {
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
		registry, ok := packets.PacketRegistries[registryKey]
		if !ok {
			fmt.Printf("// [%d] WARNING: unknown state/direction %s (packet %s)\n\n", i, registryKey, cap.PacketID)
			continue
		}

		factory, ok := registry[packetID]
		if !ok {
			fmt.Printf("// [%d] WARNING: unknown packet %s in %s\n\n", i, cap.PacketID, registryKey)
			continue
		}

		rawData, err := hex.DecodeString(cap.RawData)
		if err != nil {
			fmt.Printf("// [%d] WARNING: invalid hex data: %v\n\n", i, err)
			continue
		}

		p := factory()
		packetName := getPacketName(p)
		buf := ns.NewReader(rawData)
		if err := p.Read(buf); err != nil {
			fmt.Printf("// [%d] WARNING: failed to decode %s: %v\n", i, packetName, err)
			fmt.Printf("//     raw data: %s\n\n", cap.RawData)
			continue
		}

		fmt.Printf("// [%d] %s %s\n", i, cap.Direction, cap.PacketID)
		fmt.Printf("%s {\n", packetName)
		fmt.Print(formatPacket(p))
		fmt.Printf("}\n\n")
	}
}
