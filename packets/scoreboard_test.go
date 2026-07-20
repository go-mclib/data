package packets_test

import (
	"testing"

	"github.com/go-mclib/data/packets"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestScoreboardPacketsRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   *packets.S2CSetObjective
	}{
		{"add", &packets.S2CSetObjective{
			ObjectiveName: "kills", Mode: 0,
			DisplayName: ns.TextComponent{Text: "Kills"}, RenderType: 0,
		}},
		{"change fixed format", &packets.S2CSetObjective{
			ObjectiveName: "kills", Mode: 2,
			DisplayName: ns.TextComponent{Text: "Eliminations", Color: "red"}, RenderType: 1,
			NumberFormat: ns.Some(ns.NumberFormat{Type: ns.NumberFormatFixed, Fixed: ns.TextComponent{Text: "!"}}),
		}},
		{"remove", &packets.S2CSetObjective{ObjectiveName: "kills", Mode: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { roundTrip(t, tt.in, &packets.S2CSetObjective{}) })
	}

	roundTrip(t, &packets.S2CSetScore{
		EntityName: "Steve", ObjectiveName: "kills", Value: 7,
		DisplayName:  ns.Some(ns.TextComponent{Text: "Steve"}),
		NumberFormat: ns.Some(ns.NumberFormat{Type: ns.NumberFormatStyled, Style: nbt.Compound{"color": nbt.String("gold")}}),
	}, &packets.S2CSetScore{})

	color := ns.VarInt(12)
	roundTrip(t, &packets.S2CSetPlayerTeam{
		TeamName: "red", Method: 0,
		Parameters: &packets.PlayerTeamParameters{
			DisplayName:       ns.TextComponent{Text: "Red Team"},
			MemberNamePrefix:  ns.TextComponent{Text: "[R] "},
			NameTagVisibility: 2, CollisionRule: 3,
			Color: ns.Some(color), Options: 3,
		},
		Players: ns.PrefixedArray[ns.String]{"Steve", "Alex"},
	}, &packets.S2CSetPlayerTeam{})

	roundTrip(t, &packets.S2CSetPlayerTeam{
		TeamName: "red", Method: 3,
		Players: ns.PrefixedArray[ns.String]{"Herobrine"},
	}, &packets.S2CSetPlayerTeam{})
}
