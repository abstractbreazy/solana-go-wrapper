package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

func TestClient_Balance(t *testing.T) {
	var (
		ctx         = context.Background()
		rpcEndpoint = rpc.DevNet_RPC
		wsEndpoint  = rpc.DevNet_WS
		err         error
		account = "8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"
	)

	client, err := New(
		ctx, 
		rpcEndpoint, 
		wsEndpoint,
	)
	require.NoError(t, err)

	out, err := client.GetBalance(
		ctx, 
		solana.MustPublicKeyFromBase58(account),
	)
	require.NoError(t, err)

	t.Logf("Balance of the account: %v", out)
}