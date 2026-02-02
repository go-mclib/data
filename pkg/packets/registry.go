package packets

import (
	packets_data "github.com/go-mclib/data/pkg/data/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
)

// PacketFactory creates a new packet instance.
type PacketFactory func() jp.Packet

// PacketRegistries maps phase+bound to a map of packet ID -> factory function.
// Useful for packet-capturing tools like proxy, or tests.
// Keys are in format "phase_bound" like "configuration_c2s" or "play_s2c".
var PacketRegistries = map[string]map[int]PacketFactory{
	"handshake_c2s": {
		int(packets_data.C2SIntentionID): func() jp.Packet { return &C2SIntention{} },
	},
	"status_c2s": {
		int(packets_data.C2SStatusRequestID):     func() jp.Packet { return &C2SStatusRequest{} },
		int(packets_data.C2SPingRequestStatusID): func() jp.Packet { return &C2SPingRequestStatus{} },
	},
	"status_s2c": {
		int(packets_data.S2CStatusResponseID):     func() jp.Packet { return &S2CStatusResponse{} },
		int(packets_data.S2CPongResponseStatusID): func() jp.Packet { return &S2CPongResponseStatus{} },
	},
	"login_c2s": {
		int(packets_data.C2SHelloID):               func() jp.Packet { return &C2SHello{} },
		int(packets_data.C2SKeyID):                 func() jp.Packet { return &C2SKey{} },
		int(packets_data.C2SCustomQueryAnswerID):   func() jp.Packet { return &C2SCustomQueryAnswer{} },
		int(packets_data.C2SLoginAcknowledgedID):   func() jp.Packet { return &C2SLoginAcknowledged{} },
		int(packets_data.C2SCookieResponseLoginID): func() jp.Packet { return &C2SCookieResponseLogin{} },
	},
	"login_s2c": {
		int(packets_data.S2CLoginDisconnectID):    func() jp.Packet { return &S2CLoginDisconnectLogin{} },
		int(packets_data.S2CHelloID):              func() jp.Packet { return &S2CHello{} },
		int(packets_data.S2CLoginFinishedID):      func() jp.Packet { return &S2CLoginFinished{} },
		int(packets_data.S2CLoginCompressionID):   func() jp.Packet { return &S2CLoginCompression{} },
		int(packets_data.S2CCustomQueryID):        func() jp.Packet { return &S2CCustomQuery{} },
		int(packets_data.S2CCookieRequestLoginID): func() jp.Packet { return &S2CCookieRequestLogin{} },
	},
	"configuration_c2s": {
		int(packets_data.C2SClientInformationConfigurationID): func() jp.Packet { return &C2SClientInformationConfiguration{} },
		int(packets_data.C2SCookieResponseConfigurationID):    func() jp.Packet { return &C2SCookieResponseConfiguration{} },
		int(packets_data.C2SCustomPayloadConfigurationID):     func() jp.Packet { return &C2SCustomPayloadConfiguration{} },
		int(packets_data.C2SFinishConfigurationID):            func() jp.Packet { return &C2SFinishConfiguration{} },
		int(packets_data.C2SKeepAliveConfigurationID):         func() jp.Packet { return &C2SKeepAliveConfiguration{} },
		int(packets_data.C2SPongConfigurationID):              func() jp.Packet { return &C2SPongConfiguration{} },
		int(packets_data.C2SResourcePackConfigurationID):      func() jp.Packet { return &C2SResourcePackConfiguration{} },
		int(packets_data.C2SSelectKnownPacksID):               func() jp.Packet { return &C2SSelectKnownPacks{} },
		int(packets_data.C2SCustomClickActionConfigurationID): func() jp.Packet { return &C2SCustomClickActionConfiguration{} },
	},
	"configuration_s2c": {
		int(packets_data.S2CCookieRequestConfigurationID):       func() jp.Packet { return &S2CCookieRequestConfiguration{} },
		int(packets_data.S2CCustomPayloadConfigurationID):       func() jp.Packet { return &S2CCustomPayloadConfiguration{} },
		int(packets_data.S2CDisconnectConfigurationID):          func() jp.Packet { return &S2CDisconnectConfiguration{} },
		int(packets_data.S2CFinishConfigurationID):              func() jp.Packet { return &S2CFinishConfiguration{} },
		int(packets_data.S2CKeepAliveConfigurationID):           func() jp.Packet { return &S2CKeepAliveConfiguration{} },
		int(packets_data.S2CPingConfigurationID):                func() jp.Packet { return &S2CPingConfiguration{} },
		int(packets_data.S2CResetChatID):                        func() jp.Packet { return &S2CResetChat{} },
		int(packets_data.S2CRegistryDataID):                     func() jp.Packet { return &S2CRegistryData{} },
		int(packets_data.S2CResourcePackPopConfigurationID):     func() jp.Packet { return &S2CResourcePackPopConfiguration{} },
		int(packets_data.S2CResourcePackPushConfigurationID):    func() jp.Packet { return &S2CResourcePackPushConfiguration{} },
		int(packets_data.S2CStoreCookieConfigurationID):         func() jp.Packet { return &S2CStoreCookieConfiguration{} },
		int(packets_data.S2CTransferConfigurationID):            func() jp.Packet { return &S2CTransferConfiguration{} },
		int(packets_data.S2CUpdateEnabledFeaturesID):            func() jp.Packet { return &S2CUpdateEnabledFeatures{} },
		int(packets_data.S2CUpdateTagsConfigurationID):          func() jp.Packet { return &S2CUpdateTagsConfiguration{} },
		int(packets_data.S2CSelectKnownPacksID):                 func() jp.Packet { return &S2CSelectKnownPacks{} },
		int(packets_data.S2CCustomReportDetailsConfigurationID): func() jp.Packet { return &S2CCustomReportDetailsConfiguration{} },
		int(packets_data.S2CServerLinksConfigurationID):         func() jp.Packet { return &S2CServerLinksConfiguration{} },
		int(packets_data.S2CClearDialogConfigurationID):         func() jp.Packet { return &S2CClearDialogConfiguration{} },
		int(packets_data.S2CShowDialogConfigurationID):          func() jp.Packet { return &S2CShowDialogConfiguration{} },
		int(packets_data.S2CCodeOfConductID):                    func() jp.Packet { return &S2CCodeOfConduct{} },
	},
}
