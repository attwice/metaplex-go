// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// If the auction manager is in a Disbursing or Finished state, then this means Auction must be in Ended state.
// Then this end point can be used as a signed proxy to use auction manager's authority over the auction to claim bid funds
// into the accept payment account on the auction manager for a given bid. Auction has no opinions on how bids are redeemed,
// only that they exist, have been paid, and have a winning place. It is up to the implementer of the auction to determine redemption,
// and auction manager does this via bid redemption tickets and the vault contract which ensure the user always
// can get their NFT once they have paid. Therefore, once they have paid, and the auction is over, the artist can claim
// funds at any time without any danger to the user of losing out on their NFT, because the AM will honor their bid with an NFT
// at ANY time.
type ClaimBid struct {

	// [0] = [WRITE] acceptPayment
	// ··········· The accept payment account on the auction manager
	//
	// [1] = [WRITE] bidderPotToken
	// ··········· The bidder pot token account
	//
	// [2] = [WRITE] bidderPotPDA
	// ··········· The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key] -
	// ··········· relative to the auction program, not auction manager
	//
	// [3] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [4] = [] auction
	// ··········· The auction
	//
	// [5] = [] bidderWallet
	// ··········· The bidder wallet
	//
	// [6] = [] tokenMint
	// ··········· Token mint of the auction
	//
	// [7] = [] vault
	// ··········· Vault
	//
	// [8] = [] store
	// ··········· Store
	//
	// [9] = [] auctionProgram
	// ··········· Auction program
	//
	// [10] = [] clockSysvar
	// ··········· Clock sysvar
	//
	// [11] = [] tokenProgram
	// ··········· Token program
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewClaimBidInstructionBuilder creates a new `ClaimBid` instruction builder.
func NewClaimBidInstructionBuilder() *ClaimBid {
	nd := &ClaimBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetAcceptPaymentAccount sets the "acceptPayment" account.
// The accept payment account on the auction manager
func (inst *ClaimBid) SetAcceptPaymentAccount(acceptPayment ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(acceptPayment).WRITE()
	return inst
}

// GetAcceptPaymentAccount gets the "acceptPayment" account.
// The accept payment account on the auction manager
func (inst *ClaimBid) GetAcceptPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBidderPotTokenAccount sets the "bidderPotToken" account.
// The bidder pot token account
func (inst *ClaimBid) SetBidderPotTokenAccount(bidderPotToken ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bidderPotToken).WRITE()
	return inst
}

// GetBidderPotTokenAccount gets the "bidderPotToken" account.
// The bidder pot token account
func (inst *ClaimBid) GetBidderPotTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetBidderPotPDAAccount sets the "bidderPotPDA" account.
// The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key] -
// relative to the auction program, not auction manager
func (inst *ClaimBid) SetBidderPotPDAAccount(bidderPotPDA ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bidderPotPDA).WRITE()
	return inst
}

// GetBidderPotPDAAccount gets the "bidderPotPDA" account.
// The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key] -
// relative to the auction program, not auction manager
func (inst *ClaimBid) GetBidderPotPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *ClaimBid) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *ClaimBid) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuctionAccount sets the "auction" account.
// The auction
func (inst *ClaimBid) SetAuctionAccount(auction ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// The auction
func (inst *ClaimBid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetBidderWalletAccount sets the "bidderWallet" account.
// The bidder wallet
func (inst *ClaimBid) SetBidderWalletAccount(bidderWallet ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(bidderWallet)
	return inst
}

// GetBidderWalletAccount gets the "bidderWallet" account.
// The bidder wallet
func (inst *ClaimBid) GetBidderWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenMintAccount sets the "tokenMint" account.
// Token mint of the auction
func (inst *ClaimBid) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
// Token mint of the auction
func (inst *ClaimBid) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetVaultAccount sets the "vault" account.
// Vault
func (inst *ClaimBid) SetVaultAccount(vault ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
// Vault
func (inst *ClaimBid) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetStoreAccount sets the "store" account.
// Store
func (inst *ClaimBid) SetStoreAccount(store ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store
func (inst *ClaimBid) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetAuctionProgramAccount sets the "auctionProgram" account.
// Auction program
func (inst *ClaimBid) SetAuctionProgramAccount(auctionProgram ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(auctionProgram)
	return inst
}

// GetAuctionProgramAccount gets the "auctionProgram" account.
// Auction program
func (inst *ClaimBid) GetAuctionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetClockSysvarAccount sets the "clockSysvar" account.
// Clock sysvar
func (inst *ClaimBid) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clockSysvar" account.
// Clock sysvar
func (inst *ClaimBid) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *ClaimBid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *ClaimBid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

func (inst ClaimBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ClaimBid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimBid) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AcceptPayment is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BidderPotToken is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BidderPotPDA is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BidderWallet is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AuctionProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" acceptPayment", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("bidderPotToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("  bidderPotPDA", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("auctionManager", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("       auction", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("  bidderWallet", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("     tokenMint", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("         vault", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("         store", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("auctionProgram", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("   clockSysvar", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice[11]))
					})
				})
		})
}

func (obj ClaimBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClaimBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClaimBidInstruction declares a new ClaimBid instruction with the provided parameters and accounts.
func NewClaimBidInstruction(
	// Accounts:
	acceptPayment ag_solanago.PublicKey,
	bidderPotToken ag_solanago.PublicKey,
	bidderPotPDA ag_solanago.PublicKey,
	auctionManager ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	bidderWallet ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	auctionProgram ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ClaimBid {
	return NewClaimBidInstructionBuilder().
		SetAcceptPaymentAccount(acceptPayment).
		SetBidderPotTokenAccount(bidderPotToken).
		SetBidderPotPDAAccount(bidderPotPDA).
		SetAuctionManagerAccount(auctionManager).
		SetAuctionAccount(auction).
		SetBidderWalletAccount(bidderWallet).
		SetTokenMintAccount(tokenMint).
		SetVaultAccount(vault).
		SetStoreAccount(store).
		SetAuctionProgramAccount(auctionProgram).
		SetClockSysvarAccount(clockSysvar).
		SetTokenProgramAccount(tokenProgram)
}
