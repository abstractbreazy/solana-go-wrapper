package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {

	var (
		ctx 			 = context.Background()
		rpcEndpoint 	 = rpc.TestNet_RPC
		wsEndpoint 		 = rpc.TestNet_WS
		err 			 error
		client      	 *Client
		lamports    	 = solana.LAMPORTS_PER_SOL / 2
		fundingAccount 	 = "5anuiBqm3sgdAYivBptyb8dGvZmQKH1cCHwRLGb4Fa6TZ5bEUJzkc5wNcPjexSKcazKcERP4cPS5YAkuzQttYgsC"
		recipientAccount = "8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"
		tx 				 *solana.Transaction
	)

	client, err = New(
		ctx,
		rpcEndpoint,
		wsEndpoint,
	)
	require.NoError(t, err)

	recent, err := client.rpc.GetRecentBlockhash(
		ctx, 
		rpc.CommitmentFinalized,
	)
	require.NoError(t, err)

	tx, err = client.newTransaction(
		lamports, 
		solana.MustPrivateKeyFromBase58(fundingAccount),
		solana.MustPublicKeyFromBase58(recipientAccount),
		recent.Value.Blockhash,
	)
	require.NoError(t, err)
	require.NotEmpty(t, tx)

	t.Logf("Transaction: %v", tx)
}