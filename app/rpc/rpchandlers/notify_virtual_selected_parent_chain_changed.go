package rpchandlers

import (
	"github.com/Hoosat-Oy/HTND/app/appmessage"
	"github.com/Hoosat-Oy/HTND/app/rpc/rpccontext"
	"github.com/Hoosat-Oy/HTND/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualSelectedParentChainChanged handles the respectively named RPC command
func HandleNotifyVirtualSelectedParentChainChanged(context *rpccontext.Context, router *router.Router,
	request appmessage.Message) (appmessage.Message, error) {

	notifyVirtualSelectedParentChainChangedRequest := request.(*appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage)

	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualSelectedParentChainChangedNotifications(
		notifyVirtualSelectedParentChainChangedRequest.IncludeAcceptedTransactionIDs)

	response := appmessage.NewNotifyVirtualSelectedParentChainChangedResponseMessage()
	return response, nil
}
