

package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/cmd/devp2p/internal/ethtest"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
)

var (
	rlpxCommand = &cli.Command{
		Name:  "rlpx",
		Usage: "RLPx Commands",
		Subcommands: []*cli.Command{
			rlpxPingCommand,
			rlpxEthTestCommand,
			rlpxSnapTestCommand,
		},
	}
	rlpxPingCommand = &cli.Command{
		Name:   "ping",
		Usage:  "ping <node>",
		Action: rlpxPing,
	}
	rlpxEthTestCommand = &cli.Command{
		Name:      "eth-test",
		Usage:     "Runs eth protocol tests against a node",
		ArgsUsage: "<node>",
		Action:    rlpxEthTest,
		Flags: []cli.Flag{
			testPatternFlag,
			testTAPFlag,
			testChainDirFlag,
			testNodeFlag,
			testNodeJWTFlag,
			testNodeEngineFlag,
		},
	}
	rlpxSnapTestCommand = &cli.Command{
		Name:      "snap-test",
		Usage:     "Runs snap protocol tests against a node",
		ArgsUsage: "",
		Action:    rlpxSnapTest,
		Flags: []cli.Flag{
			testPatternFlag,
			testTAPFlag,
			testChainDirFlag,
			testNodeFlag,
			testNodeJWTFlag,
			testNodeEngineFlag,
		},
	}
)

func rlpxPing(ctx *cli.Context) error {
	n := getNodeArg(ctx)
	fd, err := net.Dial("tcp", fmt.Sprintf("%v:%d", n.IP(), n.TCP()))
	if err != nil {
		return err
	}
	conn := rlpx.NewConn(fd, n.Pubkey())
	ourKey, _ := crypto.GenerateKey()
	_, err = conn.Handshake(ourKey)
	if err != nil {
		return err
	}
	code, data, _, err := conn.Read()
	if err != nil {
		return err
	}
	switch code {
	case 0:
		var h ethtest.Hello
		if err := rlp.DecodeBytes(data, &h); err != nil {
			return fmt.Errorf("invalid handshake: %v", err)
		}
		fmt.Printf("%+v\n", h)
	case 1:
		var msg []p2p.DiscReason
		if rlp.DecodeBytes(data, &msg); len(msg) == 0 {
			return errors.New("invalid disconnect message")
		}
		return fmt.Errorf("received disconnect message: %v", msg[0])
	default:
		return fmt.Errorf("invalid message code %d, expected handshake (code zero)", code)
	}
	return nil
}

// rlpxEthTest runs the eth protocol test suite.
func rlpxEthTest(ctx *cli.Context) error {
	p := cliTestParams(ctx)
	suite, err := ethtest.NewSuite(p.node, p.chainDir, p.engineAPI, p.jwt)
	if err != nil {
		exit(err)
	}
	return runTests(ctx, suite.EthTests())
}

// rlpxSnapTest runs the snap protocol test suite.
func rlpxSnapTest(ctx *cli.Context) error {
	p := cliTestParams(ctx)
	suite, err := ethtest.NewSuite(p.node, p.chainDir, p.engineAPI, p.jwt)
	if err != nil {
		exit(err)
	}
	return runTests(ctx, suite.SnapTests())
}

type testParams struct {
	node      *enode.Node
	engineAPI string
	jwt       string
	chainDir  string
}

func cliTestParams(ctx *cli.Context) *testParams {
	nodeStr := ctx.String(testNodeFlag.Name)
	if nodeStr == "" {
		exit(fmt.Errorf("missing -%s", testNodeFlag.Name))
	}
	node, err := parseNode(nodeStr)
	if err != nil {
		exit(err)
	}
	p := testParams{
		node:      node,
		engineAPI: ctx.String(testNodeEngineFlag.Name),
		jwt:       ctx.String(testNodeJWTFlag.Name),
		chainDir:  ctx.String(testChainDirFlag.Name),
	}
	if p.engineAPI == "" {
		exit(fmt.Errorf("missing -%s", testNodeEngineFlag.Name))
	}
	if p.jwt == "" {
		exit(fmt.Errorf("missing -%s", testNodeJWTFlag.Name))
	}
	if p.chainDir == "" {
		exit(fmt.Errorf("missing -%s", testChainDirFlag.Name))
	}
	return &p
}
