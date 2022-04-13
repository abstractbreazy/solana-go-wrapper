package solego

import (
	"context"
	"encoding/base64"
	"os"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/text"
)

// SendTransaction transfer tokens from one account to another, and returns transaction signature.
func (c *Client) SendTransaction(
	ctx context.Context,
	lamports uint64,
	fundingAccount string,
	recipientAccount string,
) (sig *solana.Signature, err error) {

	pk, err := solana.PrivateKeyFromBase58((c.Accounts)[fundingAccount])
	if err != nil {
		return nil, err
	}

	recent, err := c.rpc.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

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

// NewTransaction creates a new Solana signed transaction.
func (cl *Client) newTransaction(
	lamports uint64,
	fundingAccount solana.PrivateKey,
	recipientAccount solana.PublicKey,
	recentBlockHash solana.Hash,
) (*solana.Transaction, error) {

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				lamports,
				fundingAccount.PublicKey(),
				recipientAccount,
			).Build(),
		},
		recentBlockHash,
		solana.TransactionPayer(fundingAccount.PublicKey()),
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if fundingAccount.PublicKey().Equals(key) {
				return &fundingAccount
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	tx.EncodeToTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	return tx, nil
}

func (c *Client) requestAirdrop(
	ctx context.Context,
	account string,
	lamports uint64,
) (*solana.Signature, error) {

	out, err := c.rpc.RequestAirdrop(
		ctx,
		solana.MustPublicKeyFromBase58(account),
		lamports,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// VerifySignatures wraps solana.VerifySignatures and verifies all the signatures in the transaction.
func (c *Client) VerifySignatures(sig string) (*solana.Transaction, error) {

	txBin, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(txBin))
	if err != nil{
		return nil, err
	}
	err = tx.VerifySignatures()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
