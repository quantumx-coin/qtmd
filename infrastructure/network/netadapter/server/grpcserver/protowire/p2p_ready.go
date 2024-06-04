package protowire

import (
	"github.com/quantumx-coin/qtmd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *HoosatdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "HoosatdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *HoosatdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
