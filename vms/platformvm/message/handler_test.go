// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
)

type CounterHandler struct {
	Tx int
}

func (h *CounterHandler) HandleTx(ids.NodeID, uint32, *Tx) error {
	h.Tx++
	return nil
}

func TestHandleTx(t *testing.T) {
	assert := assert.New(t)

	handler := CounterHandler{}
	msg := Tx{}

	err := msg.Handle(&handler, ids.EmptyNodeID, 0)
	assert.NoError(err)
	assert.Equal(1, handler.Tx)
}

func TestNoopHandler(t *testing.T) {
	assert := assert.New(t)

	handler := NoopHandler{
		Log: logging.NoLog{},
	}

	err := handler.HandleTx(ids.EmptyNodeID, 0, nil)
	assert.NoError(err)
}
