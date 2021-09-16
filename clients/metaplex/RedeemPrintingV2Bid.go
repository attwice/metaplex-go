// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Note: This requires that auction manager be in a Running state and that be of the V1 type.
//
// If an auction is complete, you can redeem your printing v2 bid for a specific item here. If you are the first to do this,
// The auction manager will switch from Running state to Disbursing state. If you are the last, this may change
// the auction manager state to Finished provided that no authorities remain to be delegated for Master Edition tokens.
//
// NOTE: Please note that it is totally possible to redeem a bid 2x - once for a prize you won and once at the RedeemParticipationBid point for an open edition
// that comes as a 'token of appreciation' for bidding. They are not mutually exclusive unless explicitly set to be that way.
type RedeemPrintingV2Bid struct {
	Args *RedeemPrintingV2BidArgs

	// [0] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [1] = [WRITE] safetyDepositTokenStorage
	// ··········· Safety deposit token storage account
	//
	// [2] = [WRITE] singleItemAccount
	// ··········· Account containing 1 token of your new mint type.
	// ··········· MUST be an associated token account of pda [wallet, token program, mint] relative to ata program.
	//
	// [3] = [WRITE] bidRedemptionKey
	// ··········· Bid redemption key -
	// ··········· Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
	//
	// [4] = [WRITE] safetyDepositBox
	// ··········· Safety deposit box account
	//
	// [5] = [WRITE] vaultAccount
	// ··········· Vault account
	//
	// [6] = [] safetyDepositConfig
	// ··········· Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
	// ··········· This account will only get used in the event this is an AuctionManagerV2
	//
	// [7] = [] auction
	// ··········· Auction
	//
	// [8] = [] bidderMetadata
	// ··········· Your BidderMetadata account
	//
	// [9] = [] bidder
	// ··········· Your Bidder account - Only needs to be signer if payer does not own
	//
	// [10] = [SIGNER] payer
	// ··········· Payer
	//
	// [11] = [] tokenProgram
	// ··········· Token program
	//
	// [12] = [] tokenVaultProgram
	// ··········· Token Vault program
	//
	// [13] = [] tokenMetadataProgram
	// ··········· Token metadata program
	//
	// [14] = [] store
	// ··········· Store
	//
	// [15] = [] system
	// ··········· System
	//
	// [16] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [17] = [WRITE] prizeTrackingTicket
	// ··········· Prize tracking ticket (pda of ['metaplex', program id, auction manager key, metadata mint id])
	//
	// [18] = [WRITE] newMetadataKey
	// ··········· New Metadata key (pda of ['metadata', program id, mint id])
	//
	// [19] = [WRITE] newEditionPDA
	// ··········· New Edition (pda of ['metadata', program id, mint id, 'edition'])
	//
	// [20] = [WRITE] masterEdition
	// ··········· Master Edition of token in vault V2 (pda of ['metadata', program id, master metadata mint id, 'edition']) PDA is relative to token metadata.
	//
	// [21] = [WRITE] newTokenMint
	// ··········· Mint of new token
	//
	// [22] = [WRITE] editionPDA
	// ··········· Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number])
	// ··········· where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE). PDA is relative to token metadata.
	//
	// [23] = [SIGNER] mintAuthority
	// ··········· Mint authority of new mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [24] = [] metadataAccount
	// ··········· Metadata account of token in vault
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRedeemPrintingV2BidInstructionBuilder creates a new `RedeemPrintingV2Bid` instruction builder.
func NewRedeemPrintingV2BidInstructionBuilder() *RedeemPrintingV2Bid {
	nd := &RedeemPrintingV2Bid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 25),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *RedeemPrintingV2Bid) SetArgs(args RedeemPrintingV2BidArgs) *RedeemPrintingV2Bid {
	inst.Args = &args
	return inst
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *RedeemPrintingV2Bid) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *RedeemPrintingV2Bid) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetSafetyDepositTokenStorageAccount sets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *RedeemPrintingV2Bid) SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(safetyDepositTokenStorage).WRITE()
	return inst
}

// GetSafetyDepositTokenStorageAccount gets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *RedeemPrintingV2Bid) GetSafetyDepositTokenStorageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSingleItemAccount sets the "singleItemAccount" account.
// Account containing 1 token of your new mint type.
// MUST be an associated token account of pda [wallet, token program, mint] relative to ata program.
func (inst *RedeemPrintingV2Bid) SetSingleItemAccount(singleItemAccount ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(singleItemAccount).WRITE()
	return inst
}

// GetSingleItemAccount gets the "singleItemAccount" account.
// Account containing 1 token of your new mint type.
// MUST be an associated token account of pda [wallet, token program, mint] relative to ata program.
func (inst *RedeemPrintingV2Bid) GetSingleItemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetBidRedemptionKeyAccount sets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *RedeemPrintingV2Bid) SetBidRedemptionKeyAccount(bidRedemptionKey ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bidRedemptionKey).WRITE()
	return inst
}

// GetBidRedemptionKeyAccount gets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *RedeemPrintingV2Bid) GetBidRedemptionKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *RedeemPrintingV2Bid) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(safetyDepositBox).WRITE()
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *RedeemPrintingV2Bid) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetVaultAccount sets the "vaultAccount" account.
// Vault account
func (inst *RedeemPrintingV2Bid) SetVaultAccount(vaultAccount ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultAccount).WRITE()
	return inst
}

// GetVaultAccount gets the "vaultAccount" account.
// Vault account
func (inst *RedeemPrintingV2Bid) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSafetyDepositConfigAccount sets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used in the event this is an AuctionManagerV2
func (inst *RedeemPrintingV2Bid) SetSafetyDepositConfigAccount(safetyDepositConfig ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(safetyDepositConfig)
	return inst
}

// GetSafetyDepositConfigAccount gets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used in the event this is an AuctionManagerV2
func (inst *RedeemPrintingV2Bid) GetSafetyDepositConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *RedeemPrintingV2Bid) SetAuctionAccount(auction ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *RedeemPrintingV2Bid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetBidderMetadataAccount sets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *RedeemPrintingV2Bid) SetBidderMetadataAccount(bidderMetadata ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(bidderMetadata)
	return inst
}

// GetBidderMetadataAccount gets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *RedeemPrintingV2Bid) GetBidderMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetBidderAccount sets the "bidder" account.
// Your Bidder account - Only needs to be signer if payer does not own
func (inst *RedeemPrintingV2Bid) SetBidderAccount(bidder ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(bidder)
	return inst
}

// GetBidderAccount gets the "bidder" account.
// Your Bidder account - Only needs to be signer if payer does not own
func (inst *RedeemPrintingV2Bid) GetBidderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *RedeemPrintingV2Bid) SetPayerAccount(payer ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *RedeemPrintingV2Bid) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *RedeemPrintingV2Bid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *RedeemPrintingV2Bid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetTokenVaultProgramAccount sets the "tokenVaultProgram" account.
// Token Vault program
func (inst *RedeemPrintingV2Bid) SetTokenVaultProgramAccount(tokenVaultProgram ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenVaultProgram)
	return inst
}

// GetTokenVaultProgramAccount gets the "tokenVaultProgram" account.
// Token Vault program
func (inst *RedeemPrintingV2Bid) GetTokenVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *RedeemPrintingV2Bid) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *RedeemPrintingV2Bid) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13]
}

// SetStoreAccount sets the "store" account.
// Store
func (inst *RedeemPrintingV2Bid) SetStoreAccount(store ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store
func (inst *RedeemPrintingV2Bid) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14]
}

// SetSystemAccount sets the "system" account.
// System
func (inst *RedeemPrintingV2Bid) SetSystemAccount(system ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System
func (inst *RedeemPrintingV2Bid) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *RedeemPrintingV2Bid) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *RedeemPrintingV2Bid) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16]
}

// SetPrizeTrackingTicketAccount sets the "prizeTrackingTicket" account.
// Prize tracking ticket (pda of ['metaplex', program id, auction manager key, metadata mint id])
func (inst *RedeemPrintingV2Bid) SetPrizeTrackingTicketAccount(prizeTrackingTicket ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(prizeTrackingTicket).WRITE()
	return inst
}

// GetPrizeTrackingTicketAccount gets the "prizeTrackingTicket" account.
// Prize tracking ticket (pda of ['metaplex', program id, auction manager key, metadata mint id])
func (inst *RedeemPrintingV2Bid) GetPrizeTrackingTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[17]
}

// SetNewMetadataKeyAccount sets the "newMetadataKey" account.
// New Metadata key (pda of ['metadata', program id, mint id])
func (inst *RedeemPrintingV2Bid) SetNewMetadataKeyAccount(newMetadataKey ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(newMetadataKey).WRITE()
	return inst
}

// GetNewMetadataKeyAccount gets the "newMetadataKey" account.
// New Metadata key (pda of ['metadata', program id, mint id])
func (inst *RedeemPrintingV2Bid) GetNewMetadataKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18]
}

// SetNewEditionPDAAccount sets the "newEditionPDA" account.
// New Edition (pda of ['metadata', program id, mint id, 'edition'])
func (inst *RedeemPrintingV2Bid) SetNewEditionPDAAccount(newEditionPDA ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(newEditionPDA).WRITE()
	return inst
}

// GetNewEditionPDAAccount gets the "newEditionPDA" account.
// New Edition (pda of ['metadata', program id, mint id, 'edition'])
func (inst *RedeemPrintingV2Bid) GetNewEditionPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[19]
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master Edition of token in vault V2 (pda of ['metadata', program id, master metadata mint id, 'edition']) PDA is relative to token metadata.
func (inst *RedeemPrintingV2Bid) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
// Master Edition of token in vault V2 (pda of ['metadata', program id, master metadata mint id, 'edition']) PDA is relative to token metadata.
func (inst *RedeemPrintingV2Bid) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[20]
}

// SetNewTokenMintAccount sets the "newTokenMint" account.
// Mint of new token
func (inst *RedeemPrintingV2Bid) SetNewTokenMintAccount(newTokenMint ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(newTokenMint).WRITE()
	return inst
}

// GetNewTokenMintAccount gets the "newTokenMint" account.
// Mint of new token
func (inst *RedeemPrintingV2Bid) GetNewTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[21]
}

// SetEditionPDAAccount sets the "editionPDA" account.
// Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number])
// where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE). PDA is relative to token metadata.
func (inst *RedeemPrintingV2Bid) SetEditionPDAAccount(editionPDA ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[22] = ag_solanago.Meta(editionPDA).WRITE()
	return inst
}

// GetEditionPDAAccount gets the "editionPDA" account.
// Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number])
// where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE). PDA is relative to token metadata.
func (inst *RedeemPrintingV2Bid) GetEditionPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[22]
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority of new mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *RedeemPrintingV2Bid) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[23] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority of new mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *RedeemPrintingV2Bid) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[23]
}

// SetMetadataAccount sets the "metadataAccount" account.
// Metadata account of token in vault
func (inst *RedeemPrintingV2Bid) SetMetadataAccount(metadataAccount ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	inst.AccountMetaSlice[24] = ag_solanago.Meta(metadataAccount)
	return inst
}

// GetMetadataAccount gets the "metadataAccount" account.
// Metadata account of token in vault
func (inst *RedeemPrintingV2Bid) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[24]
}

func (inst RedeemPrintingV2Bid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RedeemPrintingV2Bid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RedeemPrintingV2Bid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RedeemPrintingV2Bid) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SafetyDepositTokenStorage is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SingleItemAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BidRedemptionKey is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SafetyDepositConfig is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BidderMetadata is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Bidder is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenVaultProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.PrizeTrackingTicket is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.NewMetadataKey is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.NewEditionPDA is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.NewTokenMint is not set")
		}
		if inst.AccountMetaSlice[22] == nil {
			return errors.New("accounts.EditionPDA is not set")
		}
		if inst.AccountMetaSlice[23] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[24] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
	}
	return nil
}

func (inst *RedeemPrintingV2Bid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RedeemPrintingV2Bid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=25]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           auctionManager", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("safetyDepositTokenStorage", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("               singleItem", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("         bidRedemptionKey", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         safetyDepositBox", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("                    vault", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("      safetyDepositConfig", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                  auction", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("           bidderMetadata", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("                   bidder", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("             tokenProgram", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("        tokenVaultProgram", inst.AccountMetaSlice[12]))
						accountsBranch.Child(ag_format.Meta("     tokenMetadataProgram", inst.AccountMetaSlice[13]))
						accountsBranch.Child(ag_format.Meta("                    store", inst.AccountMetaSlice[14]))
						accountsBranch.Child(ag_format.Meta("                   system", inst.AccountMetaSlice[15]))
						accountsBranch.Child(ag_format.Meta("               rentSysvar", inst.AccountMetaSlice[16]))
						accountsBranch.Child(ag_format.Meta("      prizeTrackingTicket", inst.AccountMetaSlice[17]))
						accountsBranch.Child(ag_format.Meta("           newMetadataKey", inst.AccountMetaSlice[18]))
						accountsBranch.Child(ag_format.Meta("            newEditionPDA", inst.AccountMetaSlice[19]))
						accountsBranch.Child(ag_format.Meta("            masterEdition", inst.AccountMetaSlice[20]))
						accountsBranch.Child(ag_format.Meta("             newTokenMint", inst.AccountMetaSlice[21]))
						accountsBranch.Child(ag_format.Meta("               editionPDA", inst.AccountMetaSlice[22]))
						accountsBranch.Child(ag_format.Meta("            mintAuthority", inst.AccountMetaSlice[23]))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice[24]))
					})
				})
		})
}

func (obj RedeemPrintingV2Bid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RedeemPrintingV2Bid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewRedeemPrintingV2BidInstruction declares a new RedeemPrintingV2Bid instruction with the provided parameters and accounts.
func NewRedeemPrintingV2BidInstruction(
	// Parameters:
	args RedeemPrintingV2BidArgs,
	// Accounts:
	auctionManager ag_solanago.PublicKey,
	safetyDepositTokenStorage ag_solanago.PublicKey,
	singleItemAccount ag_solanago.PublicKey,
	bidRedemptionKey ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	vaultAccount ag_solanago.PublicKey,
	safetyDepositConfig ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	bidderMetadata ag_solanago.PublicKey,
	bidder ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenVaultProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	prizeTrackingTicket ag_solanago.PublicKey,
	newMetadataKey ag_solanago.PublicKey,
	newEditionPDA ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	newTokenMint ag_solanago.PublicKey,
	editionPDA ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	metadataAccount ag_solanago.PublicKey) *RedeemPrintingV2Bid {
	return NewRedeemPrintingV2BidInstructionBuilder().
		SetArgs(args).
		SetAuctionManagerAccount(auctionManager).
		SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage).
		SetSingleItemAccount(singleItemAccount).
		SetBidRedemptionKeyAccount(bidRedemptionKey).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetVaultAccount(vaultAccount).
		SetSafetyDepositConfigAccount(safetyDepositConfig).
		SetAuctionAccount(auction).
		SetBidderMetadataAccount(bidderMetadata).
		SetBidderAccount(bidder).
		SetPayerAccount(payer).
		SetTokenProgramAccount(tokenProgram).
		SetTokenVaultProgramAccount(tokenVaultProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetStoreAccount(store).
		SetSystemAccount(system).
		SetRentSysvarAccount(rentSysvar).
		SetPrizeTrackingTicketAccount(prizeTrackingTicket).
		SetNewMetadataKeyAccount(newMetadataKey).
		SetNewEditionPDAAccount(newEditionPDA).
		SetMasterEditionAccount(masterEdition).
		SetNewTokenMintAccount(newTokenMint).
		SetEditionPDAAccount(editionPDA).
		SetMintAuthorityAccount(mintAuthority).
		SetMetadataAccount(metadataAccount)
}
