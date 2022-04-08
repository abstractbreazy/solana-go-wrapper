package solego

import (
	"context"
	"fmt"
	"log"

	_ "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)
type SolClient struct {
	*rpc.Client
}

func New() (c *SolClient, err error) {
	client := rpc.New(rpc.TestNet_RPC)

	var ver *rpc.GetVersionResult
	if ver, err = client.GetVersion(context.Background()); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	log.Printf("current solana version %v", ver.SolanaCore)

	return &SolClient{client}, nil
}