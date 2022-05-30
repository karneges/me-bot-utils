// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMetadataAccount is the `CreateMetadataAccount` instruction.
type CreateMetadataAccount struct {
	CreateMetadataAccountArgs *CreateMetadataAccountArgs

	// [0] = [WRITE] metadata
	//
	// [1] = [] mint
	//
	// [2] = [SIGNER] mintAuthority
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] updateAuthority
	//
	// [5] = [] systemProgram
	//
	// [6] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMetadataAccountInstructionBuilder creates a new `CreateMetadataAccount` instruction builder.
func NewCreateMetadataAccountInstructionBuilder() *CreateMetadataAccount {
	nd := &CreateMetadataAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetCreateMetadataAccountArgs sets the "createMetadataAccountArgs" parameter.
func (inst *CreateMetadataAccount) SetCreateMetadataAccountArgs(createMetadataAccountArgs CreateMetadataAccountArgs) *CreateMetadataAccount {
	inst.CreateMetadataAccountArgs = &createMetadataAccountArgs
	return inst
}

// SetMetadataAccount sets the "metadata" account.
func (inst *CreateMetadataAccount) SetMetadataAccount(metadata ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *CreateMetadataAccount) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *CreateMetadataAccount) SetMintAccount(mint ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CreateMetadataAccount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *CreateMetadataAccount) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *CreateMetadataAccount) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateMetadataAccount) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateMetadataAccount) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *CreateMetadataAccount) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(updateAuthority)
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *CreateMetadataAccount) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMetadataAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMetadataAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateMetadataAccount) SetRentAccount(rent ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateMetadataAccount) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CreateMetadataAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMetadataAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMetadataAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMetadataAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreateMetadataAccountArgs == nil {
			return errors.New("CreateMetadataAccountArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMetadataAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMetadataAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreateMetadataAccountArgs", *inst.CreateMetadataAccountArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  mintAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CreateMetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreateMetadataAccountArgs` param:
	err = encoder.Encode(obj.CreateMetadataAccountArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreateMetadataAccountArgs`:
	err = decoder.Decode(&obj.CreateMetadataAccountArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMetadataAccountInstruction declares a new CreateMetadataAccount instruction with the provided parameters and accounts.
func NewCreateMetadataAccountInstruction(
	// Parameters:
	createMetadataAccountArgs CreateMetadataAccountArgs,
	// Accounts:
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMetadataAccount {
	return NewCreateMetadataAccountInstructionBuilder().
		SetCreateMetadataAccountArgs(createMetadataAccountArgs).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetUpdateAuthorityAccount(updateAuthority).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}