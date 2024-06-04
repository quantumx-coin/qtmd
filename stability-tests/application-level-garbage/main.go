package main

import (
	"fmt"
	"os"

	"github.com/Hoosat-Oy/HTND/infrastructure/config"
	"github.com/Hoosat-Oy/HTND/infrastructure/network/netadapter/standalone"
	"github.com/Hoosat-Oy/HTND/stability-tests/common"
	"github.com/Hoosat-Oy/HTND/util/panics"
	"github.com/Hoosat-Oy/HTND/util/profiling"
)

func main() {
	defer panics.HandlePanic(log, "applicationLevelGarbage-main", nil)
	err := parseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %+v", err)
		os.Exit(1)
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())
	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	htndConfig := config.DefaultConfig()
	htndConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(htndConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating minimalNetAdapter: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	blocksChan, err := readBlocks()
	if err != nil {
		log.Errorf("Error reading blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	err = sendBlocks(cfg.NodeP2PAddress, minimalNetAdapter, blocksChan)
	if err != nil {
		log.Errorf("Error sending blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}
}
