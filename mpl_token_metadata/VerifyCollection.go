// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// VerifyCollection is the `VerifyCollection` instruction.
type VerifyCollection struct {

	// [0] = [WRITE] metadata
	//
	// [1] = [WRITE, SIGNER] collectionAuthority
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] collectionMint
	//
	// [4] = [] collection
	//
	// [5] = [] collectionMasterEditionAccount
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVerifyCollectionInstructionBuilder creates a new `VerifyCollection` instruction builder.
func NewVerifyCollectionInstructionBuilder() *VerifyCollection {
	nd := &VerifyCollection{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetMetadataAccount sets the "metadata" account.
func (inst *VerifyCollection) SetMetadataAccount(metadata ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *VerifyCollection) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAuthorityAccount sets the "collectionAuthority" account.
func (inst *VerifyCollection) SetCollectionAuthorityAccount(collectionAuthority ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collectionAuthority).WRITE().SIGNER()
	return inst
}

// GetCollectionAuthorityAccount gets the "collectionAuthority" account.
func (inst *VerifyCollection) GetCollectionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *VerifyCollection) SetPayerAccount(payer ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *VerifyCollection) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *VerifyCollection) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *VerifyCollection) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionAccount sets the "collection" account.
func (inst *VerifyCollection) SetCollectionAccount(collection ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collection)
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *VerifyCollection) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCollectionMasterEditionAccountAccount sets the "collectionMasterEditionAccount" account.
func (inst *VerifyCollection) SetCollectionMasterEditionAccountAccount(collectionMasterEditionAccount ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(collectionMasterEditionAccount)
	return inst
}

// GetCollectionMasterEditionAccountAccount gets the "collectionMasterEditionAccount" account.
func (inst *VerifyCollection) GetCollectionMasterEditionAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst VerifyCollection) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_VerifyCollection,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst VerifyCollection) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *VerifyCollection) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CollectionAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.collection is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CollectionMasterEditionAccount is not set")
		}
	}
	return nil
}

func (inst *VerifyCollection) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("VerifyCollection")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    collectionAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         collectionMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             collection", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("collectionMasterEdition", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj VerifyCollection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *VerifyCollection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewVerifyCollectionInstruction declares a new VerifyCollection instruction with the provided parameters and accounts.
func NewVerifyCollectionInstruction(
	// Accounts:
	metadata ag_solanago.PublicKey,
	collectionAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	collectionMasterEditionAccount ag_solanago.PublicKey) *VerifyCollection {
	return NewVerifyCollectionInstructionBuilder().
		SetMetadataAccount(metadata).
		SetCollectionAuthorityAccount(collectionAuthority).
		SetPayerAccount(payer).
		SetCollectionMintAccount(collectionMint).
		SetCollectionAccount(collection).
		SetCollectionMasterEditionAccountAccount(collectionMasterEditionAccount)
}
