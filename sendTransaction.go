package solego

import (
	"context"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
)

func (c *Client) sendTransaction(
	ctx context.Context, 
	lamports uint64,
	fundingAccount string,
	recipientAccount string,
) (sig *solana.Signature, err error) {
	
	pk, err := solana.PrivateKeyFromBase58((c.Accounts)[fundingAccount])
	if err != nil {
		return nil, err
	}

	log.Printf("funding account pk: %v\n", pk)

	recent, err := c.rpc.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	log.Printf("account pk: %v\n", solana.MustPublicKeyFromBase58(recipientAccount))

	tx, err := c.newTransaction(
		lamports,
		pk,
		solana.MustPublicKeyFromBase58(recipientAccount),
		recent.Value.Blockhash,
	)
	if err != nil {
		return nil, err
	}

	sign, err := confirm.SendAndConfirmTransaction(
		ctx,
		c.rpc,
		c.wss,
		tx,
	)

	if err != nil {
		return nil, err
	}

	return &sign, err
}