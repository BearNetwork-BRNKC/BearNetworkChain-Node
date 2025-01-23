

package flags

import "github.com/urfave/cli/v2"

const (
	EthCategory        = "ETHEREUM"
	LightCategory      = "LIGHT CLIENT"
	DevCategory        = "DEVELOPER CHAIN"
	StateCategory      = "STATE HISTORY MANAGEMENT"
	TxPoolCategory     = "TRANSACTION POOL (EVM)"
	BlobPoolCategory   = "TRANSACTION POOL (BLOB)"
	PerfCategory       = "PERFORMANCE TUNING"
	AccountCategory    = "ACCOUNT"
	APICategory        = "API AND CONSOLE"
	NetworkingCategory = "NETWORKING"
	MinerCategory      = "MINER"
	GasPriceCategory   = "GAS PRICE ORACLE"
	VMCategory         = "VIRTUAL MACHINE"
	LoggingCategory    = "LOGGING AND DEBUGGING"
	MetricsCategory    = "METRICS AND STATS"
	MiscCategory       = "MISC"
	TestingCategory    = "TESTING"
	DeprecatedCategory = "ALIASED (deprecated)"
)

func init() {
	cli.HelpFlag.(*cli.BoolFlag).Category = MiscCategory
	cli.VersionFlag.(*cli.BoolFlag).Category = MiscCategory
}
