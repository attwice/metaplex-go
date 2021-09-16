// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Validates that a given safety deposit box has in it contents that match the expected WinningConfig in the auction manager.
// A stateful call, this will error out if you call it a second time after validation has occurred.
type DeprecatedValidateSafetyDepositBoxV1 struct {

	// [0] = [WRITE] uninitializedSafetyDepositValidationTicket
	// ··········· Uninitialized Safety deposit validation ticket, pda of seed ['metaplex', program id, auction manager key, safety deposit key]
	//
	// [1] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [2] = [WRITE] metadataAccount
	// ··········· Metadata account
	//
	// [3] = [WRITE] originalAuthorityLookup
	// ··········· Original authority lookup - unallocated uninitialized pda account with seed ['metaplex', auction key, metadata key]
	// ··········· We will store original authority here to return it later.
	//
	// [4] = [] whitelistedCreatorEntry
	// ··········· A whitelisted creator entry for the store of this auction manager pda of ['metaplex', store key, creator key]
	// ··········· where creator key comes from creator list of metadata, any will do
	//
	// [5] = [] auctionManagerStoreKey
	// ··········· The auction manager's store key
	//
	// [6] = [] safetyDepositBox
	// ··········· Safety deposit box account
	//
	// [7] = [] safetyDepositBoxStorage
	// ··········· Safety deposit box storage account where the actual nft token is stored
	//
	// [8] = [] mintAccount
	// ··········· Mint account of the token in the safety deposit box
	//
	// [9] = [] editionOrMasterEditionRecordKey
	// ··········· Edition OR MasterEdition record key
	// ··········· Remember this does not need to be an existing account (may not be depending on token), just is a pda with seed
	// ··········· of ['metadata', program id, Printing mint id, 'edition']. - remember PDA is relative to token metadata program.
	//
	// [10] = [] vaultAccount
	// ··········· Vault account
	//
	// [11] = [SIGNER] authority
	// ··········· Authority
	//
	// [12] = [SIGNER] metadataAuthority
	// ··········· [optional] Metadata Authority - Signer only required if doing a full ownership txfer
	//
	// [13] = [SIGNER] payer
	// ··········· Payer
	//
	// [14] = [] tokenMetadataProgram
	// ··········· Token metadata program
	//
	// [15] = [] system
	// ··········· System
	//
	// [16] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [17] = [WRITE] limitedEditionPrintingMint
	// ··········· Limited edition Printing mint account (optional - only if using sending Limited Edition)
	//
	// [18] = [SIGNER] limitedEditionPrintingMintAuthority
	// ··········· Limited edition Printing mint Authority account, this will TEMPORARILY TRANSFER MINTING AUTHORITY to the auction manager
	// ··········· until all limited editions have been redeemed for authority tokens.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewDeprecatedValidateSafetyDepositBoxV1InstructionBuilder creates a new `DeprecatedValidateSafetyDepositBoxV1` instruction builder.
func NewDeprecatedValidateSafetyDepositBoxV1InstructionBuilder() *DeprecatedValidateSafetyDepositBoxV1 {
	nd := &DeprecatedValidateSafetyDepositBoxV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 19),
	}
	return nd
}

// SetUninitializedSafetyDepositValidationTicketAccount sets the "uninitializedSafetyDepositValidationTicket" account.
// Uninitialized Safety deposit validation ticket, pda of seed ['metaplex', program id, auction manager key, safety deposit key]
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetUninitializedSafetyDepositValidationTicketAccount(uninitializedSafetyDepositValidationTicket ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(uninitializedSafetyDepositValidationTicket).WRITE()
	return inst
}

// GetUninitializedSafetyDepositValidationTicketAccount gets the "uninitializedSafetyDepositValidationTicket" account.
// Uninitialized Safety deposit validation ticket, pda of seed ['metaplex', program id, auction manager key, safety deposit key]
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetUninitializedSafetyDepositValidationTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetMetadataAccount sets the "metadataAccount" account.
// Metadata account
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetMetadataAccount(metadataAccount ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadataAccount).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadataAccount" account.
// Metadata account
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetOriginalAuthorityLookupAccount sets the "originalAuthorityLookup" account.
// Original authority lookup - unallocated uninitialized pda account with seed ['metaplex', auction key, metadata key]
// We will store original authority here to return it later.
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetOriginalAuthorityLookupAccount(originalAuthorityLookup ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(originalAuthorityLookup).WRITE()
	return inst
}

// GetOriginalAuthorityLookupAccount gets the "originalAuthorityLookup" account.
// Original authority lookup - unallocated uninitialized pda account with seed ['metaplex', auction key, metadata key]
// We will store original authority here to return it later.
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetOriginalAuthorityLookupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetWhitelistedCreatorEntryAccount sets the "whitelistedCreatorEntry" account.
// A whitelisted creator entry for the store of this auction manager pda of ['metaplex', store key, creator key]
// where creator key comes from creator list of metadata, any will do
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetWhitelistedCreatorEntryAccount(whitelistedCreatorEntry ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(whitelistedCreatorEntry)
	return inst
}

// GetWhitelistedCreatorEntryAccount gets the "whitelistedCreatorEntry" account.
// A whitelisted creator entry for the store of this auction manager pda of ['metaplex', store key, creator key]
// where creator key comes from creator list of metadata, any will do
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetWhitelistedCreatorEntryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuctionManagerStoreKeyAccount sets the "auctionManagerStoreKey" account.
// The auction manager's store key
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetAuctionManagerStoreKeyAccount(auctionManagerStoreKey ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionManagerStoreKey)
	return inst
}

// GetAuctionManagerStoreKeyAccount gets the "auctionManagerStoreKey" account.
// The auction manager's store key
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetAuctionManagerStoreKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(safetyDepositBox)
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetSafetyDepositBoxStorageAccount sets the "safetyDepositBoxStorage" account.
// Safety deposit box storage account where the actual nft token is stored
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetSafetyDepositBoxStorageAccount(safetyDepositBoxStorage ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(safetyDepositBoxStorage)
	return inst
}

// GetSafetyDepositBoxStorageAccount gets the "safetyDepositBoxStorage" account.
// Safety deposit box storage account where the actual nft token is stored
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetSafetyDepositBoxStorageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetMintAccount sets the "mintAccount" account.
// Mint account of the token in the safety deposit box
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetMintAccount(mintAccount ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mintAccount)
	return inst
}

// GetMintAccount gets the "mintAccount" account.
// Mint account of the token in the safety deposit box
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetEditionOrMasterEditionRecordKeyAccount sets the "editionOrMasterEditionRecordKey" account.
// Edition OR MasterEdition record key
// Remember this does not need to be an existing account (may not be depending on token), just is a pda with seed
// of ['metadata', program id, Printing mint id, 'edition']. - remember PDA is relative to token metadata program.
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetEditionOrMasterEditionRecordKeyAccount(editionOrMasterEditionRecordKey ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(editionOrMasterEditionRecordKey)
	return inst
}

// GetEditionOrMasterEditionRecordKeyAccount gets the "editionOrMasterEditionRecordKey" account.
// Edition OR MasterEdition record key
// Remember this does not need to be an existing account (may not be depending on token), just is a pda with seed
// of ['metadata', program id, Printing mint id, 'edition']. - remember PDA is relative to token metadata program.
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetEditionOrMasterEditionRecordKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetVaultAccount sets the "vaultAccount" account.
// Vault account
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetVaultAccount(vaultAccount ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(vaultAccount)
	return inst
}

// GetVaultAccount gets the "vaultAccount" account.
// Vault account
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetAuthorityAccount sets the "authority" account.
// Authority
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetAuthorityAccount(authority ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Authority
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetMetadataAuthorityAccount sets the "metadataAuthority" account.
// [optional] Metadata Authority - Signer only required if doing a full ownership txfer
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetMetadataAuthorityAccount(metadataAuthority ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(metadataAuthority).SIGNER()
	return inst
}

// GetMetadataAuthorityAccount gets the "metadataAuthority" account.
// [optional] Metadata Authority - Signer only required if doing a full ownership txfer
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetMetadataAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetPayerAccount(payer ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13]
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14]
}

// SetSystemAccount sets the "system" account.
// System
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetSystemAccount(system ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16]
}

// SetLimitedEditionPrintingMintAccount sets the "limitedEditionPrintingMint" account.
// Limited edition Printing mint account (optional - only if using sending Limited Edition)
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetLimitedEditionPrintingMintAccount(limitedEditionPrintingMint ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(limitedEditionPrintingMint).WRITE()
	return inst
}

// GetLimitedEditionPrintingMintAccount gets the "limitedEditionPrintingMint" account.
// Limited edition Printing mint account (optional - only if using sending Limited Edition)
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetLimitedEditionPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[17]
}

// SetLimitedEditionPrintingMintAuthorityAccount sets the "limitedEditionPrintingMintAuthority" account.
// Limited edition Printing mint Authority account, this will TEMPORARILY TRANSFER MINTING AUTHORITY to the auction manager
// until all limited editions have been redeemed for authority tokens.
func (inst *DeprecatedValidateSafetyDepositBoxV1) SetLimitedEditionPrintingMintAuthorityAccount(limitedEditionPrintingMintAuthority ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(limitedEditionPrintingMintAuthority).SIGNER()
	return inst
}

// GetLimitedEditionPrintingMintAuthorityAccount gets the "limitedEditionPrintingMintAuthority" account.
// Limited edition Printing mint Authority account, this will TEMPORARILY TRANSFER MINTING AUTHORITY to the auction manager
// until all limited editions have been redeemed for authority tokens.
func (inst *DeprecatedValidateSafetyDepositBoxV1) GetLimitedEditionPrintingMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18]
}

func (inst DeprecatedValidateSafetyDepositBoxV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedValidateSafetyDepositBoxV1),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedValidateSafetyDepositBoxV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedValidateSafetyDepositBoxV1) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UninitializedSafetyDepositValidationTicket is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OriginalAuthorityLookup is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.WhitelistedCreatorEntry is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionManagerStoreKey is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SafetyDepositBoxStorage is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MintAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.EditionOrMasterEditionRecordKey is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.VaultAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Authority is not set")
		}

		// [12] = MetadataAuthority is optional

		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.LimitedEditionPrintingMint is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.LimitedEditionPrintingMintAuthority is not set")
		}
	}
	return nil
}

func (inst *DeprecatedValidateSafetyDepositBoxV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedValidateSafetyDepositBoxV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=19]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("uninitializedSafetyDepositValidationTicket", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                            auctionManager", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("                                  metadata", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("                   originalAuthorityLookup", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                   whitelistedCreatorEntry", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("                    auctionManagerStoreKey", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("                          safetyDepositBox", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                   safetyDepositBoxStorage", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("                                      mint", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("           editionOrMasterEditionRecordKey", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("                                     vault", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("                                 authority", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("                         metadataAuthority", inst.AccountMetaSlice[12]))
						accountsBranch.Child(ag_format.Meta("                                     payer", inst.AccountMetaSlice[13]))
						accountsBranch.Child(ag_format.Meta("                      tokenMetadataProgram", inst.AccountMetaSlice[14]))
						accountsBranch.Child(ag_format.Meta("                                    system", inst.AccountMetaSlice[15]))
						accountsBranch.Child(ag_format.Meta("                                rentSysvar", inst.AccountMetaSlice[16]))
						accountsBranch.Child(ag_format.Meta("                limitedEditionPrintingMint", inst.AccountMetaSlice[17]))
						accountsBranch.Child(ag_format.Meta("       limitedEditionPrintingMintAuthority", inst.AccountMetaSlice[18]))
					})
				})
		})
}

func (obj DeprecatedValidateSafetyDepositBoxV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeprecatedValidateSafetyDepositBoxV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeprecatedValidateSafetyDepositBoxV1Instruction declares a new DeprecatedValidateSafetyDepositBoxV1 instruction with the provided parameters and accounts.
func NewDeprecatedValidateSafetyDepositBoxV1Instruction(
	// Accounts:
	uninitializedSafetyDepositValidationTicket ag_solanago.PublicKey,
	auctionManager ag_solanago.PublicKey,
	metadataAccount ag_solanago.PublicKey,
	originalAuthorityLookup ag_solanago.PublicKey,
	whitelistedCreatorEntry ag_solanago.PublicKey,
	auctionManagerStoreKey ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	safetyDepositBoxStorage ag_solanago.PublicKey,
	mintAccount ag_solanago.PublicKey,
	editionOrMasterEditionRecordKey ag_solanago.PublicKey,
	vaultAccount ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	metadataAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	limitedEditionPrintingMint ag_solanago.PublicKey,
	limitedEditionPrintingMintAuthority ag_solanago.PublicKey) *DeprecatedValidateSafetyDepositBoxV1 {
	return NewDeprecatedValidateSafetyDepositBoxV1InstructionBuilder().
		SetUninitializedSafetyDepositValidationTicketAccount(uninitializedSafetyDepositValidationTicket).
		SetAuctionManagerAccount(auctionManager).
		SetMetadataAccount(metadataAccount).
		SetOriginalAuthorityLookupAccount(originalAuthorityLookup).
		SetWhitelistedCreatorEntryAccount(whitelistedCreatorEntry).
		SetAuctionManagerStoreKeyAccount(auctionManagerStoreKey).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetSafetyDepositBoxStorageAccount(safetyDepositBoxStorage).
		SetMintAccount(mintAccount).
		SetEditionOrMasterEditionRecordKeyAccount(editionOrMasterEditionRecordKey).
		SetVaultAccount(vaultAccount).
		SetAuthorityAccount(authority).
		SetMetadataAuthorityAccount(metadataAuthority).
		SetPayerAccount(payer).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetSystemAccount(system).
		SetRentSysvarAccount(rentSysvar).
		SetLimitedEditionPrintingMintAccount(limitedEditionPrintingMint).
		SetLimitedEditionPrintingMintAuthorityAccount(limitedEditionPrintingMintAuthority)
}
