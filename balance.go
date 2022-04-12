package solego

import (
	"context"
	"fmt"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// GetBalance returns the balance of the account of provided publicKey.
func (cl *Client) GetBalance(
	ctx context.Context,
	account solana.PublicKey,
) (string, error) {

	out, err := cl.rpc.GetBalance(
		ctx,
		account,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return "", err
	}

	return cl.LamportsToSOL(out.Value), nil
}

// LamportsToSOL converts lamports to sol.
func (cl *Client) LamportsToSOL(
	lamports uint64,
) string {

	var (
		ls   = new(big.Float).SetUint64(lamports)
		sols = new(big.Float).Quo(ls, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	)

	return fmt.Sprintf("%s %s", sols.String(), "SOL")
}
