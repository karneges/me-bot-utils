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
	"log"
	"strings"
)

func GetMagicEdenActionTransaction(transaction *rpc.GetTransactionResult) (globalTypes.MarketAction, error) {

	var transactionDecoded solana.Transaction
	jsonBytes, err := transaction.Transaction.MarshalJSON()
	if err != nil {
		return globalTypes.MarketAction{}, err
	}
	err = json.Unmarshal(jsonBytes, &transactionDecoded)
	if err != nil {
		return globalTypes.MarketAction{}, err
	}
	marketActionInstruction, isFound := lo.Find(transactionDecoded.Message.Instructions, func(tr solana.CompiledInstruction) bool {
		if transactionDecoded.Message.AccountKeys[tr.ProgramIDIndex].Equals(constance.MAGIC_EDEN_V2_PROGRAM_ID) {
			hexData := hex.EncodeToString(tr.Data)
			return strings.Contains(hexData, constance.SaleMatcher) || strings.Contains(hexData, constance.ListingMatcher) || strings.Contains(hexData, constance.CancelListingMatcher)

		}
		return false
	})
	if !isFound {
		return globalTypes.MarketAction{
			ActionType: constance.StatusMap[constance.UnCategorisedMatcher],
			TimeStamp:  transaction.BlockTime.Time(),
		}, nil
	}
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
			ActionType:     constance.StatusMap[constance.ListingMatcher],
			Price:          lo.ToPtr(price),
			EscrowAccount:  lo.ToPtr(transactionDecoded.Message.AccountKeys[marketActionInstruction.Accounts[8]].String()),
			ListingAccount: lo.ToPtr(transactionDecoded.Message.AccountKeys[marketActionInstruction.Accounts[2]]),
			Mint:           mint,
			Owner:          owner,
			TimeStamp:      transaction.BlockTime.Time(),
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
	Signatures chan []*rpc.TransactionSignature
	Err        chan error
}

func GetSignatureCh(client *rpc.Client) SignatureOutputCh {
	var currentLastSignature *rpc.TransactionSignature
	signatures := SignatureOutputCh{
		Signatures: make(chan []*rpc.TransactionSignature),
		Err:        make(chan error),
	}
	firstSignature, err := client.GetSignaturesForAddressWithOpts(context.TODO(), constance.MAGIC_EDEN_V2_PROGRAM_ID, &rpc.GetSignaturesForAddressOpts{
		Limit:      lo.ToPtr(1),
		Commitment: rpc.CommitmentProcessed,
	})
	if err != nil {
		signatures.Err <- err
		return signatures
	}
	currentLastSignature = firstSignature[0]
	println(currentLastSignature.Signature.String())
	go func() {
		for {
			tailSignature, err := client.GetSignaturesForAddressWithOpts(context.TODO(), constance.MAGIC_EDEN_V2_PROGRAM_ID, &rpc.GetSignaturesForAddressOpts{
				Until:      currentLastSignature.Signature,
				Commitment: rpc.CommitmentProcessed,
				Limit:      lo.ToPtr(50),
			})

			if err != nil {
				log.Printf("get confirmed Signatures error: %v", err.Error())
				//signatures.Err <- err
				//break
			}
			if len(tailSignature) > 0 {
				//log.Printf(
				//	"Current signature %v vs next signature %v",
				//	currentLastSignature.Signature.String(),
				//	tailSignature[0].Signature.String(),
				//	)
				currentLastSignature = tailSignature[0]
			} else {
				continue
			}
			signaturesWithoutErrors := lo.Filter(tailSignature, func(tr *rpc.TransactionSignature, i int) bool {
				return tr.Err == nil
			})
			signatures.Signatures <- signaturesWithoutErrors
		}
	}()
	return signatures
}

func getSigner(transaction solana.Transaction) solana.PublicKey {
	return transaction.Message.Signers()[0]
}
