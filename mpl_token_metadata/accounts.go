// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type UseAuthorityRecord struct {
	Key         Key
	AllowedUses uint64
	Bump        uint8
}

var UseAuthorityRecordDiscriminator = [8]byte{227, 200, 230, 197, 244, 198, 172, 50}

func (obj UseAuthorityRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(UseAuthorityRecordDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `AllowedUses` param:
	err = encoder.Encode(obj.AllowedUses)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UseAuthorityRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(UseAuthorityRecordDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[227 200 230 197 244 198 172 50]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `AllowedUses`:
	err = decoder.Decode(&obj.AllowedUses)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

type CollectionAuthorityRecord struct {
	Key  Key
	Bump uint8
}

var CollectionAuthorityRecordDiscriminator = [8]byte{156, 48, 108, 31, 212, 219, 100, 168}

func (obj CollectionAuthorityRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CollectionAuthorityRecordDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CollectionAuthorityRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CollectionAuthorityRecordDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[156 48 108 31 212 219 100 168]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

type Metadata struct {
	Key                 Key
	UpdateAuthority     ag_solanago.PublicKey
	Mint                ag_solanago.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8         `bin:"optional"`
	TokenStandard       *TokenStandard `bin:"optional"`
	Collection          *Collection    `bin:"optional"`
	Uses                *Uses          `bin:"optional"`
}

var MetadataDiscriminator = [8]byte{72, 11, 121, 26, 111, 181, 85, 93}

func (obj Metadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `UpdateAuthority` param:
	err = encoder.Encode(obj.UpdateAuthority)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `PrimarySaleHappened` param:
	err = encoder.Encode(obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `EditionNonce` param (optional):
	{
		if obj.EditionNonce == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TokenStandard` param (optional):
	{
		if obj.TokenStandard == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uses` param (optional):
	{
		if obj.Uses == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *Metadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	//{
	//	discriminator, err := decoder.ReadTypeID()
	//	if err != nil {
	//		return err
	//	}
	//	if !discriminator.Equal(MetadataDiscriminator[:]) {
	//		return fmt.Errorf(
	//			"wrong discriminator: wanted %s, got %s",
	//			"[72 11 121 26 111 181 85 93]",
	//			fmt.Sprint(discriminator[:]))
	//	}
	//}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `UpdateAuthority`:
	err = decoder.Decode(&obj.UpdateAuthority)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `PrimarySaleHappened`:
	err = decoder.Decode(&obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `EditionNonce` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TokenStandard` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MasterEditionV2 struct {
	Key    Key
	Supply uint64
	MaxSupply *uint64 `bin:"optional"`
}

var MasterEditionV2Discriminator = [8]byte{101, 59, 163, 207, 238, 16, 170, 159}

func (obj MasterEditionV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MasterEditionV2Discriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	// Serialize `MaxSupply` param (optional):
	{
		if obj.MaxSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *MasterEditionV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MasterEditionV2Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[101 59 163 207 238 16 170 159]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MasterEditionV1 struct {
	Key    Key
	Supply uint64
	MaxSupply                        *uint64 `bin:"optional"`
	PrintingMint                     ag_solanago.PublicKey
	OneTimePrintingAuthorizationMint ag_solanago.PublicKey
}

var MasterEditionV1Discriminator = [8]byte{79, 165, 41, 167, 180, 191, 141, 185}

func (obj MasterEditionV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MasterEditionV1Discriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	// Serialize `MaxSupply` param (optional):
	{
		if obj.MaxSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PrintingMint` param:
	err = encoder.Encode(obj.PrintingMint)
	if err != nil {
		return err
	}
	// Serialize `OneTimePrintingAuthorizationMint` param:
	err = encoder.Encode(obj.OneTimePrintingAuthorizationMint)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MasterEditionV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MasterEditionV1Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[79 165 41 167 180 191 141 185]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PrintingMint`:
	err = decoder.Decode(&obj.PrintingMint)
	if err != nil {
		return err
	}
	// Deserialize `OneTimePrintingAuthorizationMint`:
	err = decoder.Decode(&obj.OneTimePrintingAuthorizationMint)
	if err != nil {
		return err
	}
	return nil
}

type Edition struct {
	Key    Key
	Parent ag_solanago.PublicKey
	Edition uint64
}

var EditionDiscriminator = [8]byte{234, 117, 249, 74, 7, 99, 235, 167}

func (obj Edition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(EditionDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Parent` param:
	err = encoder.Encode(obj.Parent)
	if err != nil {
		return err
	}
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Edition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(EditionDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[234 117 249 74 7 99 235 167]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Parent`:
	err = decoder.Decode(&obj.Parent)
	if err != nil {
		return err
	}
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

type ReservationListV2 struct {
	Key                   Key
	MasterEdition         ag_solanago.PublicKey
	SupplySnapshot        *uint64 `bin:"optional"`
	Reservations          []Reservation
	TotalReservationSpots uint64
	CurrentReservationSpots uint64
}

var ReservationListV2Discriminator = [8]byte{193, 233, 97, 55, 245, 135, 103, 218}

func (obj ReservationListV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ReservationListV2Discriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `MasterEdition` param:
	err = encoder.Encode(obj.MasterEdition)
	if err != nil {
		return err
	}
	// Serialize `SupplySnapshot` param (optional):
	{
		if obj.SupplySnapshot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Reservations` param:
	err = encoder.Encode(obj.Reservations)
	if err != nil {
		return err
	}
	// Serialize `TotalReservationSpots` param:
	err = encoder.Encode(obj.TotalReservationSpots)
	if err != nil {
		return err
	}
	// Serialize `CurrentReservationSpots` param:
	err = encoder.Encode(obj.CurrentReservationSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationListV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ReservationListV2Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[193 233 97 55 245 135 103 218]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `MasterEdition`:
	err = decoder.Decode(&obj.MasterEdition)
	if err != nil {
		return err
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Reservations`:
	err = decoder.Decode(&obj.Reservations)
	if err != nil {
		return err
	}
	// Deserialize `TotalReservationSpots`:
	err = decoder.Decode(&obj.TotalReservationSpots)
	if err != nil {
		return err
	}
	// Deserialize `CurrentReservationSpots`:
	err = decoder.Decode(&obj.CurrentReservationSpots)
	if err != nil {
		return err
	}
	return nil
}

type ReservationListV1 struct {
	Key            Key
	MasterEdition  ag_solanago.PublicKey
	SupplySnapshot *uint64 `bin:"optional"`
	Reservations   []ReservationV1
}

var ReservationListV1Discriminator = [8]byte{239, 79, 12, 206, 116, 153, 1, 140}

func (obj ReservationListV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ReservationListV1Discriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `MasterEdition` param:
	err = encoder.Encode(obj.MasterEdition)
	if err != nil {
		return err
	}
	// Serialize `SupplySnapshot` param (optional):
	{
		if obj.SupplySnapshot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Reservations` param:
	err = encoder.Encode(obj.Reservations)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationListV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ReservationListV1Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[239 79 12 206 116 153 1 140]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `MasterEdition`:
	err = decoder.Decode(&obj.MasterEdition)
	if err != nil {
		return err
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupplySnapshot)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Reservations`:
	err = decoder.Decode(&obj.Reservations)
	if err != nil {
		return err
	}
	return nil
}

type EditionMarker struct {
	Key    Key
	Ledger [31]uint8
}

var EditionMarkerDiscriminator = [8]byte{233, 10, 18, 230, 129, 172, 37, 234}

func (obj EditionMarker) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(EditionMarkerDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Ledger` param:
	err = encoder.Encode(obj.Ledger)
	if err != nil {
		return err
	}
	return nil
}

func (obj *EditionMarker) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(EditionMarkerDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[233 10 18 230 129 172 37 234]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Ledger`:
	err = decoder.Decode(&obj.Ledger)
	if err != nil {
		return err
	}
	return nil
}
