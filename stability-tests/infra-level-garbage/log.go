package main

import (
	"github.com/quantumx-coin/qtmd/infrastructure/logger"
	"github.com/quantumx-coin/qtmd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
