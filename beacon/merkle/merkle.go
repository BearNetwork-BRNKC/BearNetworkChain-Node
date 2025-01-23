

// Package merkle implements proof verifications in binary merkle trees.
package merkle

import (
	"crypto/sha256"
	"errors"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Value represents either a 32 byte leaf value or hash node in a binary merkle tree/partial proof.
type Value [32]byte

// Values represent a series of merkle tree leaves/nodes.
type Values []Value

var valueT = reflect.TypeOf(Value{})

// UnmarshalJSON parses a merkle value in hex syntax.
func (m *Value) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(valueT, input, m[:])
}

// VerifyProof verifies a Merkle proof branch for a single value in a
// binary Merkle tree (index is a generalized tree index).
func VerifyProof(root common.Hash, index uint64, branch Values, value Value) error {
	hasher := sha256.New()
	for _, sibling := range branch {
		hasher.Reset()
		if index&1 == 0 {
			hasher.Write(value[:])
			hasher.Write(sibling[:])
		} else {
			hasher.Write(sibling[:])
			hasher.Write(value[:])
		}
		hasher.Sum(value[:0])
		if index >>= 1; index == 0 {
			return errors.New("branch has extra items")
		}
	}
	if index != 1 {
		return errors.New("branch is missing items")
	}
	if common.Hash(value) != root {
		return errors.New("root mismatch")
	}
	return nil
}
