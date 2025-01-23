

package t8ntool

import (
	"encoding/json"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/log"
)

// traceWriter is an vm.EVMLogger which also holds an inner logger/tracer.
// When the TxEnd event happens, the inner tracer result is written to the file, and
// the file is closed.
type traceWriter struct {
	inner vm.EVMLogger
	f     io.WriteCloser
}

// Compile-time interface check
var _ = vm.EVMLogger((*traceWriter)(nil))

func (t *traceWriter) CaptureTxEnd(restGas uint64) {
	t.inner.CaptureTxEnd(restGas)
	defer t.f.Close()

	if tracer, ok := t.inner.(tracers.Tracer); ok {
		result, err := tracer.GetResult()
		if err != nil {
			log.Warn("Error in tracer", "err", err)
			return
		}
		err = json.NewEncoder(t.f).Encode(result)
		if err != nil {
			log.Warn("Error writing tracer output", "err", err)
			return
		}
	}
}

func (t *traceWriter) CaptureTxStart(gasLimit uint64) { t.inner.CaptureTxStart(gasLimit) }
func (t *traceWriter) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {
	t.inner.CaptureStart(env, from, to, create, input, gas, value)
}

func (t *traceWriter) CaptureEnd(output []byte, gasUsed uint64, err error) {
	t.inner.CaptureEnd(output, gasUsed, err)
}

func (t *traceWriter) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	t.inner.CaptureEnter(typ, from, to, input, gas, value)
}

func (t *traceWriter) CaptureExit(output []byte, gasUsed uint64, err error) {
	t.inner.CaptureExit(output, gasUsed, err)
}

func (t *traceWriter) CaptureState(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
	t.inner.CaptureState(pc, op, gas, cost, scope, rData, depth, err)
}
func (t *traceWriter) CaptureFault(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {
	t.inner.CaptureFault(pc, op, gas, cost, scope, depth, err)
}
