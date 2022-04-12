package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
)

func TestTransaction_New(t *testing.T) {
	var (
		ctx              = context.Background()
		rpcEndpoint      = rpc.DevNet_RPC
		wsEndpoint       = rpc.DevNet_WS
		err              error
		client           *Client
		lamports         = solana.LAMPORTS_PER_SOL / 2
		fundingAccount   = "5anuiBqm3sgdAYivBptyb8dGvZmQKH1cCHwRLGb4Fa6TZ5bEUJzkc5wNcPjexSKcazKcERP4cPS5YAkuzQttYgsC"
		recipientAccount = "8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"
		tx               *solana.Transaction
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

	require.Equal(t,
		[]solana.PublicKey{
			solana.MustPublicKeyFromBase58("24ccQHKPJKyhJUBSkweHMtR7cufDGasrQajswNsUmqdC"),
			solana.MustPublicKeyFromBase58("8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"),
			solana.MustPublicKeyFromBase58("11111111111111111111111111111111"),
		},
		tx.Message.AccountKeys,
	)

	require.Equal(t,
		solana.MessageHeader{
			NumRequiredSignatures:       1,
			NumReadonlySignedAccounts:   0,
			NumReadonlyUnsignedAccounts: 1,
		},
		tx.Message.Header,
	)

	require.Equal(t,
		solana.MustHashFromBase58(recent.Value.Blockhash.String()),
		tx.Message.RecentBlockhash,
	)

	decodedData, err := base58.Decode(tx.Message.Instructions[0].Data.String())
	require.NoError(t, err)

	require.Equal(t,
		[]solana.CompiledInstruction{
			{
				ProgramIDIndex: 2,
				Accounts: []uint16{
					0,
					1,
				},
				Data: solana.Base58(decodedData),
			},
		},
		tx.Message.Instructions,
	)

	require.NoError(t, tx.VerifySignatures())
	require.Equal(t, len(tx.Signatures), len(tx.Message.Signers()))

	t.Logf("Transaction: %v", tx)
}

func TestTransaction_Send(t *testing.T) {

	var (
		ctx              = context.Background()
		rpcEndpoint      = rpc.DevNet_RPC
		wsEndpoint       = rpc.DevNet_WS
		err              error
		client           *Client
		lamports         = solana.LAMPORTS_PER_SOL / 2
		fundingAccount   = "7xK9a35KstNYHruT1yhqr93ffZ33EipSTN7Cq2ECzNZz"
		recipientAccount = "8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j"
	)

	client, err = New(
		ctx,
		rpcEndpoint,
		wsEndpoint,
	)
	require.NoError(t, err)
	client.Accounts = Accounts

	_, err = client.requestAirdrop(
		ctx,
		fundingAccount,
		solana.LAMPORTS_PER_SOL*1,
	)
	require.NoError(t, err)

	sig, err := client.SendTransaction(
		ctx,
		lamports,
		fundingAccount,
		recipientAccount,
	)
	require.NoError(t, err)

	t.Logf("tx hash: %v", sig)
}

func TestTransaction_RequestAirdrop(t *testing.T) {

	var (
		ctx         = context.Background()
		rpcEndpoint = rpc.DevNet_RPC
		wsEndpoint  = rpc.DevNet_WS
		err         error
		client      *Client
		lamports    = solana.LAMPORTS_PER_SOL
		account     = "7xK9a35KstNYHruT1yhqr93ffZ33EipSTN7Cq2ECzNZz"
	)

	client, err = New(
		ctx,
		rpcEndpoint,
		wsEndpoint,
	)
	require.NoError(t, err)

	sig, err := client.requestAirdrop(
		ctx,
		account,
		lamports,
	)
	require.NoError(t, err)

	t.Logf("tx hash: %v", sig)
}
