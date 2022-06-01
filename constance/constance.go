package constance

import (
	"github.com/gagliardetto/solana-go"
)

var MAGIC_EDEN_V2_PROGRAM_ID = solana.MustPublicKeyFromBase58("M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K")

const (
	SaleMatcher          = "254ad99d4f31"
	ListingMatcher       = "33e685a4017f83"
	CancelListingMatcher = "c6c682cba35f"
	UnCategorisedMatcher = "UnCategorisedMatcher"
	InitialStateMatcher  = "InitialStateMatcher"
)

var StatusMap = map[string]string{SaleMatcher: "SaleMatcher", ListingMatcher: "ListingMatcher", CancelListingMatcher: "CancelListingMatcher", UnCategorisedMatcher: "UnCategorisedMatcher", InitialStateMatcher: "InitialStateMatcher"}

func IsMarketAction(action string) bool {
	switch action {
	case StatusMap[SaleMatcher], StatusMap[ListingMatcher], StatusMap[CancelListingMatcher]:
		return true
	default:
		return false

	}
}
