package solego

import (
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/text"
)

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