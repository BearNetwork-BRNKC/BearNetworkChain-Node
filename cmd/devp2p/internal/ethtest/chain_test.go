

package ethtest

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/stretchr/testify/assert"
)

// TestEthProtocolNegotiation tests whether the test suite
// can negotiate the highest eth protocol in a status message exchange
func TestEthProtocolNegotiation(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		conn     *Conn
		caps     []p2p.Cap
		expected uint32
	}{
		{
			conn: &Conn{
				ourHighestProtoVersion: 65,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "eth", Version: 65},
			},
			expected: uint32(65),
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 65,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "eth", Version: 65},
			},
			expected: uint32(65),
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 65,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "eth", Version: 65},
			},
			expected: uint32(65),
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 64,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "eth", Version: 65},
			},
			expected: 64,
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 65,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 0},
				{Name: "eth", Version: 89},
				{Name: "eth", Version: 65},
			},
			expected: uint32(65),
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 64,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "wrongProto", Version: 65},
			},
			expected: uint32(64),
		},
		{
			conn: &Conn{
				ourHighestProtoVersion: 65,
			},
			caps: []p2p.Cap{
				{Name: "eth", Version: 63},
				{Name: "eth", Version: 64},
				{Name: "wrongProto", Version: 65},
			},
			expected: uint32(64),
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tt.conn.negotiateEthProtocol(tt.caps)
			assert.Equal(t, tt.expected, uint32(tt.conn.negotiatedProtoVersion))
		})
	}
}

// TestChainGetHeaders tests whether the test suite can correctly
// respond to a GetBlockHeaders request from a node.
func TestChainGetHeaders(t *testing.T) {
	t.Parallel()

	dir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	chain, err := NewChain(dir)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		req      eth.GetBlockHeadersPacket
		expected []*types.Header
	}{
		{
			req: eth.GetBlockHeadersPacket{
				GetBlockHeadersRequest: &eth.GetBlockHeadersRequest{
					Origin:  eth.HashOrNumber{Number: uint64(2)},
					Amount:  uint64(5),
					Skip:    1,
					Reverse: false,
				},
			},
			expected: []*types.Header{
				chain.blocks[2].Header(),
				chain.blocks[4].Header(),
				chain.blocks[6].Header(),
				chain.blocks[8].Header(),
				chain.blocks[10].Header(),
			},
		},
		{
			req: eth.GetBlockHeadersPacket{
				GetBlockHeadersRequest: &eth.GetBlockHeadersRequest{
					Origin:  eth.HashOrNumber{Number: uint64(chain.Len() - 1)},
					Amount:  uint64(3),
					Skip:    0,
					Reverse: true,
				},
			},
			expected: []*types.Header{
				chain.blocks[chain.Len()-1].Header(),
				chain.blocks[chain.Len()-2].Header(),
				chain.blocks[chain.Len()-3].Header(),
			},
		},
		{
			req: eth.GetBlockHeadersPacket{
				GetBlockHeadersRequest: &eth.GetBlockHeadersRequest{
					Origin:  eth.HashOrNumber{Hash: chain.Head().Hash()},
					Amount:  uint64(1),
					Skip:    0,
					Reverse: false,
				},
			},
			expected: []*types.Header{
				chain.Head().Header(),
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			headers, err := chain.GetHeaders(&tt.req)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, headers, tt.expected)
		})
	}
}
