

package blobpool

import (
	"github.com/ethereum/go-ethereum/log"
)

// Config are the configuration parameters of the blob transaction pool.
type Config struct {
	Datadir   string // Data directory containing the currently executable blobs
	Datacap   uint64 // Soft-cap of database storage (hard cap is larger due to overhead)
	PriceBump uint64 // Minimum price bump percentage to replace an already existing nonce
}

// DefaultConfig contains the default configurations for the transaction pool.
var DefaultConfig = Config{
	Datadir:   "blobpool",
	Datacap:   10 * 1024 * 1024 * 1024 / 4, // TODO(karalabe): /4 handicap for rollout, gradually bump back up to 10GB
	PriceBump: 100,                         // either have patience or be aggressive, no mushy ground
}

// sanitize checks the provided user configurations and changes anything that's
// unreasonable or unworkable.
func (config *Config) sanitize() Config {
	conf := *config
	if conf.Datacap < 1 {
		log.Warn("Sanitizing invalid blobpool storage cap", "provided", conf.Datacap, "updated", DefaultConfig.Datacap)
		conf.Datacap = DefaultConfig.Datacap
	}
	if conf.PriceBump < 1 {
		log.Warn("Sanitizing invalid blobpool price bump", "provided", conf.PriceBump, "updated", DefaultConfig.PriceBump)
		conf.PriceBump = DefaultConfig.PriceBump
	}
	return conf
}
