package solego

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// GetBalance returns the balance of the account of provided publicKey.
func (cl *Client) GetBalance(
	ctx context.Context,
	account solana.PublicKey,
) (uint64, error) {

	out, err := cl.rpc.GetBalance(
		ctx,
		account,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return 0, err
	}

	return out.Value, nil
}