package mappers

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/karneges/me-bot-utils/constance"
	globalTypes "github.com/karneges/me-bot-utils/types"
	"github.com/samber/lo"
	"strings"
)

func GetMagicEdenActionTransaction(transaction globalTypes.TransactionWithSignatureData) (globalTypes.MarketAction, error) {
	var transactionDecoded solana.Transaction
	jsonBytes, err := transaction.Transaction.MarshalJSON()

	err = json.Unmarshal(jsonBytes, &transactionDecoded)
	if err != nil {
		return globalTypes.MarketAction{}, err
	}
	marketActionInstruction, _ := lo.Find(transactionDecoded.Message.Instructions, func(tr solana.CompiledInstruction) bool {
		hexData := hex.EncodeToString(tr.Data)
		return strings.Contains(hexData, constance.SaleMatcher) || strings.Contains(hexData, constance.ListingMatcher) || strings.Contains(hexData, constance.CancelListingMatcher)
	})
	hexData := hex.EncodeToString(marketActionInstruction.Data)
	var mint solana.PublicKey
	if len(transaction.Meta.PostTokenBalances) > 0 {
		mint = transaction.Meta.PostTokenBalances[0].Mint
	}
	owner := getSigner(transactionDecoded)
	if strings.Contains(hexData, constance.SaleMatcher) {
		return globalTypes.MarketAction{
			ActionType: constance.StatusMap[constance.SaleMatcher],
			Mint:       mint,
			Owner:      owner,
			TimeStamp:  transaction.BlockTime.Time(),
		}, nil
	}

	if strings.Contains(hexData, constance.ListingMatcher) {
		price, err := extractPriceFromLogs(transaction.Meta.LogMessages)
		if err != nil {
			return globalTypes.MarketAction{}, err
		}
		return globalTypes.MarketAction{
			ActionType:    constance.StatusMap[constance.ListingMatcher],
			Price:         lo.ToPtr(price),
			EscrowAccount: lo.ToPtr(transactionDecoded.Message.AccountKeys[marketActionInstruction.Accounts[8]].String()),
			Mint:          mint,
			Owner:         owner,
			TimeStamp:     transaction.BlockTime.Time(),
		}, nil
	}

	if strings.Contains(hexData, constance.CancelListingMatcher) {
		return globalTypes.MarketAction{
			ActionType: constance.StatusMap[constance.CancelListingMatcher],
			Mint:       mint,
			Owner:      owner,
			TimeStamp:  transaction.BlockTime.Time(),
		}, nil
	}
	return globalTypes.MarketAction{
		ActionType: constance.StatusMap[constance.UnCategorisedMatcher],
		Owner:      owner,
		TimeStamp:  transaction.BlockTime.Time(),
	}, nil
}

func extractPriceFromLogs(logs []string) (uint64, error) {
	priceRaw, _ := lo.Find(logs, func(log string) bool {
		return strings.Contains(log, `Program log: {"price"`)
	})
	jsonPart := strings.Split(priceRaw, "Program log: ")
	if len(jsonPart) < 2 {
		return 0, nil
	}
	type PriceLog struct {
		Price        uint64 `json:"Price"`
		SellerExpiry int64  `json:"seller_expiry"`
		BuyerExpiry  int64  `json:"buyer_expiry"`
	}
	var priceLog PriceLog
	err := json.Unmarshal([]byte(jsonPart[1]), &priceLog)
	if err != nil {
		return 0, err
	}
	return priceLog.Price, nil
}

type SignatureOutputCh struct {
	signatures chan *rpc.TransactionSignature
	err        chan error
}

func GetSignatureCh(client *rpc.Client) SignatureOutputCh {
	var currentLastSignature *rpc.TransactionSignature
	signatures := SignatureOutputCh{
		signatures: make(chan *rpc.TransactionSignature),
		err:        make(chan error),
	}
	firstSignature, err := client.GetSignaturesForAddressWithOpts(context.TODO(), constance.MAGIC_EDEN_V2_PROGRAM_ID, &rpc.GetSignaturesForAddressOpts{
		Limit: lo.ToPtr(1),
	})
	if err != nil {
		signatures.err <- err
		return signatures
	}
	currentLastSignature = firstSignature[0]
	go func() {
		for {
			tailSignature, err := client.GetSignaturesForAddressWithOpts(context.TODO(), constance.MAGIC_EDEN_V2_PROGRAM_ID, &rpc.GetSignaturesForAddressOpts{
				Until:      currentLastSignature.Signature,
				Commitment: rpc.CommitmentConfirmed,
			})
			if err != nil {
				signatures.err <- err
				break
			}
			if len(tailSignature) > 0 && tailSignature[0].BlockTime.Time().After(currentLastSignature.BlockTime.Time()) {
				currentLastSignature = tailSignature[0]
			} else {
				continue
			}
			signaturesWithoutErrors := lo.Filter(tailSignature, func(tr *rpc.TransactionSignature, i int) bool {
				return tr.Err == nil
			})
			for _, signature := range signaturesWithoutErrors {
				signatures.signatures <- signature
			}
		}
	}()
	return signatures
}

func getSigner(transaction solana.Transaction) solana.PublicKey {
	return transaction.Message.Signers()[0]
}
