dagconfig
========

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/quantumx-coin/qtmd/dagconfig)

Package dagconfig defines DAG configuration parameters for the standard
Hoosatd networks and provides the ability for callers to define their own custom
Hoosatd networks.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/quantumx-coin/qtmd/util"
	"github.com/quantumx-coin/qtmd/domain/dagconfig"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Hoosat network")

// By default (without --testnet), use mainnet.
var dagParams = &dagconfig.MainnetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		dagParams = &dagconfig.TestnetParams
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKey := make([]byte, 32)
	addr, err := util.NewAddressPubKey(pubKey, dagParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```
