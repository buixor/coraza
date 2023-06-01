package experimental

import (
	"github.com/corazawaf/coraza/v3/types"
)

func ToFullInterface(waf types.Transaction) FullTransaction {
	x := waf.(FullTransaction)
	return x
}

type FullTransaction interface {
	types.Transaction
	RemoveRuleByID(int)
}
