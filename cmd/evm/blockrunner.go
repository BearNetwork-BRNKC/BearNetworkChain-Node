

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/tests"
	"github.com/urfave/cli/v2"
)

var RunFlag = &cli.StringFlag{
	Name:  "run",
	Value: ".*",
	Usage: "Run only those tests matching the regular expression.",
}

var blockTestCommand = &cli.Command{
	Action:    blockTestCmd,
	Name:      "blocktest",
	Usage:     "Executes the given blockchain tests",
	ArgsUsage: "<file>",
	Flags:     []cli.Flag{RunFlag},
}

func blockTestCmd(ctx *cli.Context) error {
	if len(ctx.Args().First()) == 0 {
		return errors.New("path-to-test argument required")
	}

	var tracer vm.EVMLogger
	// Configure the EVM logger
	if ctx.Bool(MachineFlag.Name) {
		tracer = logger.NewJSONLogger(&logger.Config{
			EnableMemory:     !ctx.Bool(DisableMemoryFlag.Name),
			DisableStack:     ctx.Bool(DisableStackFlag.Name),
			DisableStorage:   ctx.Bool(DisableStorageFlag.Name),
			EnableReturnData: !ctx.Bool(DisableReturnDataFlag.Name),
		}, os.Stderr)
	}
	// Load the test content from the input file
	src, err := os.ReadFile(ctx.Args().First())
	if err != nil {
		return err
	}
	var tests map[string]tests.BlockTest
	if err = json.Unmarshal(src, &tests); err != nil {
		return err
	}
	re, err := regexp.Compile(ctx.String(RunFlag.Name))
	if err != nil {
		return fmt.Errorf("invalid regex -%s: %v", RunFlag.Name, err)
	}

	// Run them in order
	var keys []string
	for key := range tests {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, name := range keys {
		if !re.MatchString(name) {
			continue
		}
		test := tests[name]
		if err := test.Run(false, rawdb.HashScheme, tracer, func(res error, chain *core.BlockChain) {
			if ctx.Bool(DumpFlag.Name) {
				if state, _ := chain.State(); state != nil {
					fmt.Println(string(state.Dump(nil)))
				}
			}
		}); err != nil {
			return fmt.Errorf("test %v: %w", name, err)
		}
	}
	return nil
}
