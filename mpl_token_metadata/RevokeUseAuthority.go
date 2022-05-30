// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RevokeUseAuthority is the `RevokeUseAuthority` instruction.
type RevokeUseAuthority struct {

	// [0] = [WRITE] useAuthorityRecord
	//
	// [1] = [WRITE, SIGNER] owner
	//
	// [2] = [] user
	//
	// [3] = [WRITE] ownerTokenAccount
	//
	// [4] = [] mint
	//
	// [5] = [] metadata
	//
	// [6] = [] tokenProgram
	//
	// [7] = [] systemProgram
	//
	// [8] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRevokeUseAuthorityInstructionBuilder creates a new `RevokeUseAuthority` instruction builder.
func NewRevokeUseAuthorityInstructionBuilder() *RevokeUseAuthority {
	nd := &RevokeUseAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetUseAuthorityRecordAccount sets the "useAuthorityRecord" account.
func (inst *RevokeUseAuthority) SetUseAuthorityRecordAccount(useAuthorityRecord ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(useAuthorityRecord).WRITE()
	return inst
}

// GetUseAuthorityRecordAccount gets the "useAuthorityRecord" account.
func (inst *RevokeUseAuthority) GetUseAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *RevokeUseAuthority) SetOwnerAccount(owner ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *RevokeUseAuthority) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserAccount sets the "user" account.
func (inst *RevokeUseAuthority) SetUserAccount(user ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(user)
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *RevokeUseAuthority) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOwnerTokenAccountAccount sets the "ownerTokenAccount" account.
func (inst *RevokeUseAuthority) SetOwnerTokenAccountAccount(ownerTokenAccount ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ownerTokenAccount).WRITE()
	return inst
}

// GetOwnerTokenAccountAccount gets the "ownerTokenAccount" account.
func (inst *RevokeUseAuthority) GetOwnerTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintAccount sets the "mint" account.
func (inst *RevokeUseAuthority) SetMintAccount(mint ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *RevokeUseAuthority) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *RevokeUseAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *RevokeUseAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *RevokeUseAuthority) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *RevokeUseAuthority) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RevokeUseAuthority) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RevokeUseAuthority) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *RevokeUseAuthority) SetRentAccount(rent ag_solanago.PublicKey) *RevokeUseAuthority {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *RevokeUseAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst RevokeUseAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RevokeUseAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RevokeUseAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RevokeUseAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UseAuthorityRecord is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OwnerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *RevokeUseAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RevokeUseAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("useAuthorityRecord", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        ownerToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              rent", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj RevokeUseAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RevokeUseAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRevokeUseAuthorityInstruction declares a new RevokeUseAuthority instruction with the provided parameters and accounts.
func NewRevokeUseAuthorityInstruction(
	// Accounts:
	useAuthorityRecord ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	ownerTokenAccount ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *RevokeUseAuthority {
	return NewRevokeUseAuthorityInstructionBuilder().
		SetUseAuthorityRecordAccount(useAuthorityRecord).
		SetOwnerAccount(owner).
		SetUserAccount(user).
		SetOwnerTokenAccountAccount(ownerTokenAccount).
		SetMintAccount(mint).
		SetMetadataAccount(metadata).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}