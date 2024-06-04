package rpchandlers

import (
	"github.com/quantumx-coin/qtmd/app/appmessage"
	"github.com/quantumx-coin/qtmd/app/rpc/rpccontext"
	"github.com/quantumx-coin/qtmd/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
