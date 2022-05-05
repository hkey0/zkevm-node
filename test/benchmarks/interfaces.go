package benchmarks

import "context"

// Consumer interfaces required by the package.

// stateInterface gathers the methods required to interact with the state.
type stateInterface interface {
	GetLastBatchNumber(ctx context.Context, txBundleID string) (uint64, error)
}