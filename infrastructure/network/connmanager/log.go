package connmanager

import (
	"github.com/quantumx-coin/qtmd/infrastructure/logger"
	"github.com/quantumx-coin/qtmd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
