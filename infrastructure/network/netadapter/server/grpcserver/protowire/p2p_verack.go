package protowire

import (
	"github.com/quantumx-coin/qtmd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *HoosatdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "HoosatdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *HoosatdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
