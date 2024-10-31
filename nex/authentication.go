package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/encryption"
	"github.com/PretendoNetwork/super-street-fighter-iv/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.AuthenticationEndpoint.DefaultStreamSettings.EncryptionAlgorithm = encryption.NewQuazalRC4Encryption()
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(0, 7, 8))
	globals.AuthenticationServer.SetFragmentSize(900)
	globals.AuthenticationServer.SessionKeyLength = 16
	globals.AuthenticationServer.PRUDPV0Settings.EncryptedConnect = true
	globals.AuthenticationServer.PRUDPV0Settings.LegacyConnectionSignature = true
	globals.AuthenticationServer.AccessKey = "7edcd5b4"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== Super Street Fighter IV - Auth ===")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("================================")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SSFIV_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
