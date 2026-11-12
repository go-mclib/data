package packets_test

import (
	"encoding/json"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestStatusResponseDescriptionShapes(t *testing.T) {
	for name, raw := range map[string]string{
		"string": `{"description":"§aHello","players":{"max":20,"online":1},"version":{"name":"26.3","protocol":777}}`,
		"object": `{"description":{"text":"§aHello"},"players":{"max":20,"online":1},"version":{"name":"26.3","protocol":777}}`,
	} {
		t.Run(name, func(t *testing.T) {
			var status packets.StatusResponse
			if err := json.Unmarshal([]byte(raw), &status); err != nil {
				t.Fatal(err)
			}
			if status.Description.Text != "§aHello" || status.Players.Max != 20 || status.Version.Protocol != 777 {
				t.Fatalf("decoded %+v", status)
			}
		})
	}
}

func TestStatusResponseMarshalsComponentDescription(t *testing.T) {
	bold := true
	status := packets.StatusResponse{
		Description: ns.TextComponent{Text: "hi", Extra: []ns.TextComponent{{Text: "!", Bold: &bold}}},
		Players:     packets.StatusPlayers{Max: 1, Sample: []packets.StatusPlayer{{ID: "00000000-0000-0000-0000-000000000000", Name: "A"}}},
		Version:     packets.StatusVersion{Name: "26.3", Protocol: 777},
	}
	raw, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	if desc := fields["description"]; len(desc) == 0 || desc[0] != '{' {
		t.Fatalf("description = %s, want component object", desc)
	}
	var back packets.StatusResponse
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatal(err)
	}
	if back.Description.String() != "hi!" || len(back.Players.Sample) != 1 {
		t.Fatalf("round trip %+v", back)
	}
}
