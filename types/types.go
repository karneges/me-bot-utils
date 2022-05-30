package types

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/karneges/me-bot-utils/mpl_token_metadata"
	"time"
)

type NftData struct {
	Name                 string         `json:"name"`
	Symbol               string         `json:"symbol"`
	Description          string         `json:"description"`
	SellerFeeBasisPoints int            `json:"seller_fee_basis_points"`
	Image                string         `json:"image"`
	ExternalUrl          string         `json:"external_url"`
	Attributes           []NftAttribute `json:"attributes"`
	Collection           struct {
		Name   string `json:"name"`
		Family string `json:"family"`
	} `json:"collection"`
	Properties struct {
		Files []struct {
			Uri  string `json:"uri"`
			Type string `json:"type"`
		} `json:"files"`
		Category string `json:"category"`
		Creators []struct {
			Address string `json:"address"`
			Share   int    `json:"share"`
		} `json:"creators"`
	} `json:"properties"`
}

type NftAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NftDataWithMetadata struct {
	MetaData mpl_token_metadata.Metadata
	Data     NftData
}

type NftEntity struct {
	Mint     solana.PublicKey
	Rarity   float64
	Creators []solana.PublicKey
}

type NftEntityWithState struct {
	NftEntity
	Owner         solana.PublicKey
	State         string
	Price         *uint64
	EscrowAccount *string
}
type MarketAction struct {
	ActionType    string
	EscrowAccount *string
	Price         *uint64
	Mint          solana.PublicKey
	Owner         solana.PublicKey
	TimeStamp     time.Time
}
type MarketActionWithSignature struct {
	MarketAction
	Signature solana.Signature
}

type TransactionWithSignatureData struct {
	*rpc.TransactionWithMeta
	*rpc.TransactionSignature
}
