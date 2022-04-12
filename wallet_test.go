package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {

	var (
		ctx         = context.Background()
		rpcEndpoint = rpc.DevNet_RPC
		wsEndpoint  = rpc.DevNet_WS
		err         error
		client      *Client
		lamports    = solana.LAMPORTS_PER_SOL*1
	)

	client, err = New(
		ctx,
		rpcEndpoint,
		wsEndpoint,
	)
	require.NoError(t, err)

	privateKey, _, err := client.NewWallet(
		ctx, 
		lamports,
	)
	require.NoError(t, err)
	public := privateKey.PublicKey()

	pk2, err := solana.WalletFromPrivateKeyBase58(privateKey.String())
	require.NoError(t, err)

	require.Equal(t, privateKey, &pk2.PrivateKey)
	require.Equal(t, public, pk2.PublicKey())
}