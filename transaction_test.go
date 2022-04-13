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

func TestTransaction_VerifySignatures(t *testing.T) {
	var (
		ctx         = context.Background()
		rpcEndpoint = rpc.DevNet_RPC
		wsEndpoint  = rpc.DevNet_WS
		err         error
		client      *Client
	)

	type testCase struct {
		Transaction string
	}

	testCases := []testCase{
		{
			Transaction: "AVBFwRrn4wroV9+NVQfgg/GbjFtQFodLnNI5oTpDMQiQ4HfZNyFzcFamHSSFW4p5wc3efeEKvykbmk8jzf2LCQwBAAIGjYddInd/DSl2KJCP18GhEDlaJyPKVrgBGGsr3TF6jSYPgr3AdITNKr2UQVQ5I+Wh5StQv/a5XdLr6VN4Y21My1M/Y1FNK5wQLKJa1LYfN/HAudufFVtc0fRPR6AMUJ9UrkRI7sjY/PnpcXLF7A7SBvJrWu+o8+7QIaD8sL9aXkGFDy1uAqR6+CTQmradxC1wyyjL+iSft+5XudJWwSdi7wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAi+i1vCST+HNO0DEchpEJImMHhZ1BReuf7poRqmXpeA8CBAUBAgMCAgcAAwAAAAEABQIAAAwCAAAA6w0AAAAAAAA=",
		},
		{
			Transaction: "AWwhMTxKhl9yZOlidY0u3gYmy+J/6V3kFSXU7GgK5zwN+SwljR2dOlHgKtUDRX8uee2HtfeyL3t4lB3n749L4QQBAAIEFg+6wTr33dgF0xcKPeDGvZcSah4CwNJZ0Khu+CHW5cehpkZfTC6/JEwx2AvJXCc0WjQk5CjC3vM+ztnpDT9wGwan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAA3OXr4eScO58RTLVUTFCpnsDWktY/Vnla4Cmsg9nqi+Jr/+AAgahV8wmBK4mnz9WwJSryq8x2Ic0asytADGhLZAEDAwABAigCAAAABwAAAAEAAAAAAAAAz+dyuQIAAAAIn18BAAAAAPsVKAcAAAAA",
		},
		{
			Transaction: "ARZsk8+AvvT9onUT8FU1VRaiC8Sp+FKveOwhdPoigWHA+MGNcIOqbow6mwSILEYvvyOB/fi3UQ/xKQCjEtxBRgIBAAIFKIX92BRrkgEfrLEXAvXtw7OgPPhHU+62C8DB5QPoMgNSbKXgdub0sr7Yp3Nvdrsp6SDoJ4gdoyRad2AV+Japj0dRtYW4OxE78FvRZTeqHFy2My/m12/afGIPS8iUnMGlBqfVFxjHdMkoVmOYaR1etoteuKObS21cc1VbIQAAAAC/jt8clGtWu0PSX5i4e2vlERcwCmEmGvn5+U7telqAiK4hdAN78GteFjqtJrxLXxpVNKsu1lfdcFPXa/Kcg4e5AQQEAQADAicmMiQQAiGujz0xoTQSQCgAMPOroDk5F0hQ/BgzEkBBvVKWIY41EkA=",
		},
		{
			Transaction: "Ad7TPpYTvSpO//KNA5YTZVojVwz4NlH4gH9ktl+rTObJcgo8QkqmHK4t6DQr9dD58B/A/5/N7v9K+0j6y1TVCAsBAAMFA9maY4S727Z/lOSb08nHehVFsC32kTKMMPjPJp111bKM0Fl1Dg04vV2x9nL2TCqSHmjT8xg6wUAzjZa1+6YCBQan1RcZLwqvxvJl4/t3zHragsUp0L47E24tAFUgAAAABqfVFxjHdMkoVmOYaR1etoteuKObS21cc1VbIQAAAAAHYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAAJDQfslK1yQFkGqDXWu6cthRNuYGlajYMOmtoSJB6hmPAQQEAQIDAE0CAAAAAwAAAAAAAAD5FSgHAAAAAPoVKAcAAAAA+xUoBwAAAADECMJOPX7e7fOF5Hrq9xhdch2Uqhg8vQOYyZM/6V983gHQ0gNiAAAAAA==",
		},
		{
			Transaction: "Ak8jvC3ch5hq3lhOHPkACoFepIUON2zEN4KRcw4lDS6GBsQfnSdzNGPETm/yi0hPKk75/i2VXFj0FLUWnGR64ADyUbqnirFjFtaSNgcGi02+Tm7siT4CPpcaTq0jxfYQK/h9FdxXXPnLry74J+RE8yji/BtJ/Cjxbx+TIHigeIYJAgEBBByE1Y6EqCJKsr7iEupU6lsBHtBdtI4SK3yWMCFA0iEKeFPgnGmtp+1SIX1Ak+sN65iBaR7v4Iim5m1OEuFQTgi9N57UnhNpCNuUePaTt7HJaFBmyeZB3deXeKWVudpY3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWVECK/n3a7QR6OKWYR4DuAVjS6FXgZj82W0dJpSIPnEBAwQAAgEDDAIAAABAQg8AAAAAAA==",
		},
	}
	
	client, err = New(
		ctx,
		rpcEndpoint,
		wsEndpoint,
	)
	require.NoError(t, err)

	for _, tc := range testCases {
		tx, err := client.VerifySignatures(tc.Transaction)
		require.NoError(t, err)

		require.Equal(t, len(tx.Signatures), len(tx.Message.Signers()))
	}

}


