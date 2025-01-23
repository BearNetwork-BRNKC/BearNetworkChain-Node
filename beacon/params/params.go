

package params

const (
	EpochLength      = 32
	SyncPeriodLength = 8192

	BLSSignatureSize = 96
	BLSPubkeySize    = 48

	SyncCommitteeSize          = 512
	SyncCommitteeBitmaskSize   = SyncCommitteeSize / 8
	SyncCommitteeSupermajority = (SyncCommitteeSize*2 + 2) / 3
)

const (
	StateIndexGenesisTime       = 32
	StateIndexGenesisValidators = 33
	StateIndexForkVersion       = 141
	StateIndexLatestHeader      = 36
	StateIndexBlockRoots        = 37
	StateIndexStateRoots        = 38
	StateIndexHistoricRoots     = 39
	StateIndexFinalBlock        = 105
	StateIndexSyncCommittee     = 54
	StateIndexNextSyncCommittee = 55
	StateIndexExecPayload       = 56
	StateIndexExecHead          = 908
)
