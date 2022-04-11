package solego

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go/rpc"

	"github.com/stretchr/testify/require"
)

var (
	// it's a mapping of account pubkey to account private key.
	Accounts = map[string]string{
		"315t2p6VxYNVP1sfn1344vR6TEXrMvfQh7sC5rtGoQ9E": "2axx2Ew8AjTSm8tPpTehcCycdg4hh4DCrsXbwRrkpHESbjz7XNfYTgWmCkE29hnoF59KLrHJeapTP523aFuGDJja",
		"24ccQHKPJKyhJUBSkweHMtR7cufDGasrQajswNsUmqdC": "5anuiBqm3sgdAYivBptyb8dGvZmQKH1cCHwRLGb4Fa6TZ5bEUJzkc5wNcPjexSKcazKcERP4cPS5YAkuzQttYgsC",
		"8PxBC1rmGYPNGQr23fF8xZbpUrieqGMMfpitJTiFyi8j": "5Spz5GjAAafAwKVT5h42Rtz4wHEESeADUwqPS5v5vxMCK5K8UoqdUFppJwoxhf9mkxmW6ZYKjNseqEfrJZRTMUy9",
		"7xK9a35KstNYHruT1yhqr93ffZ33EipSTN7Cq2ECzNZz": "5y4tQbhbwERwGvQ47W4ghMiDKBUvXDHhZjC6dnXUHJW2XH1BudRMAf7t5FBATm2waLEwmUCNyTWYPxTdBMrkvgUt",
		"2TG9onyPXfe6hB5af1nLzeYBaJW6t9KVTnoCZ3kUG9Uj": "4R4L5Q7NZiPANiHGLAcPngmoLi5qmuyquKbugV1JKkDHycLV4Av6cdGfBa2ezJUw7qYpFWrJLN9p9crjL61FhiF1",
		"EhMqHyHqV8Ja5xmLvzHp2tCi3g8pttQ2Q8a7VDZUtBDi": "2hUyzXZrcpueRjgtSXqvnUnQ4fuUgZuUNn5rzGvZgSaLUj6s5t3hhS3MB6t3gToNdQqZDmuRHDH5Z7rdnMRFoBTN",
		"AdMisi7rXcZwYR9BAUq2gHSE7ZdnnHByfx9AT3x1XgWp": "TKsAexzmUxbMrzokAmYMFwx5dKJkWUpFry1dGxU3v2S4rP4i18WSJKX3vtwG8qHERx3v5YvbFD3k85chJEpsCYU",
		"ChDesJJLJb9kwcVc3KZmXLNvXXZbKzL2ji2xcFxYscL1": "2kYGPxsYuqHsdfdPpqGtUfzNdKxTef24ERsqqotYQP3syWRM661tvrq6S2gXdV8Nz4QyZNHgMuGFmTohNe32woiR",
		"5iZtTQjqKpNTBZCtSpKd9h6cNpdxbahHnH9PBssUrbZr": "5qyfAwaTeC7V4U9cvfZfjV1zVbmhQUf3VTuNanMPW6qtpARCnR6VQSng4nZFMw2pibRNdQYZR2r5u2zn5gAkuhVi",
		"HSEFoD3n3DYYhHkvSPcDvk2vwPrJkULu32Vms2M7o9wP": "KAPh7idad5RQAxgYSQ2nTNsTjM4esNjiozbrh7cbpSpMxwFDhuUahKyGjU7eVZ7hZfA9AG2EtWgt2iZCNDFEQiM",
	}

)

func TestClient_New(t *testing.T) {
	var (
		ctx 		= context.Background()
		rpcEndpoint = rpc.TestNet_RPC
		wsEndpoint 	= rpc.TestNet_WS
		err 		error
	)
	
	client, err := New(
		ctx, 
		rpcEndpoint, 
		wsEndpoint,
	)
	require.NoError(t, err)
	
	_, err = client.rpc.GetVersion(ctx)
	require.NoError(t, err)
}