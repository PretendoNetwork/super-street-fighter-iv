package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/encryption"
	"github.com/PretendoNetwork/super-street-fighter-iv/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.SecureEndpoint.DefaultStreamSettings.EncryptionAlgorithm = encryption.NewQuazalRC4Encryption()
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(0, 7, 8))
	globals.SecureServer.SetFragmentSize(900)
	globals.SecureServer.SessionKeyLength = 16
	globals.SecureServer.PRUDPV0Settings.EncryptedConnect = true
	globals.SecureServer.PRUDPV0Settings.LegacyConnectionSignature = true
	globals.SecureServer.AccessKey = "7edcd5b4"

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== Super Street Fighter IV - Secure ===")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("==================================")
	})

	globals.SecureEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Error(err.Error())
	})

	registerCommonSecureServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SSFIV_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
