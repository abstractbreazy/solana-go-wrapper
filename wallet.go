package solego

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// NewWallet creates a new wallet account.
func (c *Client) NewWallet(
	ctx context.Context,
	lamports uint64,
	) (*solana.PrivateKey, *solana.Signature, error) {

	pk := solana.NewWallet()

	out, err := c.rpc.RequestAirdrop(
		ctx, 
		pk.PublicKey(),
		lamports,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return nil, nil, err
	}

	return &pk.PrivateKey, &out, nil
}