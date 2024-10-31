package globals

import (
	"context"

	pb_account "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid *types.PID) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXPassword(ctx, &pb_account.GetNEXPasswordRequest{Pid: pid.LegacyValue()})
	if err != nil {
		globals.Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	return response.Password, 0
}
