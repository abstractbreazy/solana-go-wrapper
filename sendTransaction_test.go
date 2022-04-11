package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

func TestSendTransaction(t *testing.T) {

	var (
		ctx = context.Background()
		rpcEndpoint 	 = rpc.DevNet_RPC
		wsEndpoint 		 = rpc.DevNet_WS
		err 			 error 
		client           *Client
		lamports    	 = solana.LAMPORTS_PER_SOL / 2
		fundingAccount 	 = "7xK9a35KstNYHruT1yhqr93ffZ33EipSTN7Cq2ECzNZz"
		recipientAccount = "8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"
	)

	client, err = New(
		ctx, 
		rpcEndpoint, 
		wsEndpoint,
	)
	require.NoError(t, err)
	client.Accounts = Accounts

	sig, err := client.sendTransaction(
		ctx,
		lamports,
		fundingAccount,
		recipientAccount,
	)
	require.NoError(t, err)


	t.Logf("tx hash: %v", sig.String())
}