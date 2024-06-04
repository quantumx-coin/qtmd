package main

import (
	"path/filepath"

	"github.com/quantumx-coin/qtmd/infrastructure/config"
	"github.com/quantumx-coin/qtmd/stability-tests/common"
	"github.com/jessevdk/go-flags"
)

const (
	defaultLogFilename    = "many_tips.log"
	defaultErrLogFilename = "many_tips_err.log"
)

var (
	// Default configuration options
	defaultLogFile    = filepath.Join(common.DefaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(common.DefaultAppDir, defaultErrLogFilename)
)

type configFlags struct {
	LogLevel       string `short:"d" long:"loglevel" description:"Set log level {trace, debug, info, warn, error, critical}"`
	NumberOfBlocks uint64 `short:"n" long:"numblocks" description:"Number of blocks to mine" required:"false"`
	Profile        string `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	config.NetworkFlags
}

var cfg *configFlags

func activeConfig() *configFlags {
	return cfg
}

func parseConfig() error {
	cfg = &configFlags{}
	parser := flags.NewParser(cfg, flags.PrintErrors|flags.HelpFlag)
	_, err := parser.Parse()
	if err != nil {
		return err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return err
	}

	initLog(defaultLogFile, defaultErrLogFile)

	return nil
}
