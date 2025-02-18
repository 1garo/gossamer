// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package babe

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/crypto/sr25519"
)

type handleSlotFunc = func(epoch uint64, slot Slot, authorityIndex uint32,
	preRuntimeDigest *types.PreRuntimeDigest) error

var (
	errEpochPast = errors.New("cannot run epoch that has already passed")
)

type epochHandler struct {
	slotHandler slotHandler
	epochNumber uint64
	firstSlot   uint64

	constants constants
	epochData *epochData

	slotToPreRuntimeDigest map[uint64]*types.PreRuntimeDigest

	handleSlot handleSlotFunc
}

func newEpochHandler(epochNumber, firstSlot uint64, epochData *epochData, constants constants,
	handleSlot handleSlotFunc, keypair *sr25519.Keypair) (*epochHandler, error) {
	// determine which slots we'll be authoring in by pre-calculating VRF output
	slotToPreRuntimeDigest := make(map[uint64]*types.PreRuntimeDigest, constants.epochLength)
	for i := firstSlot; i < firstSlot+constants.epochLength; i++ {
		preRuntimeDigest, err := claimSlot(epochNumber, i, epochData, keypair)
		if err == nil {
			slotToPreRuntimeDigest[i] = preRuntimeDigest
			continue
		}

		if errors.Is(err, errNotOurTurnToPropose) {
			continue
		}

		return nil, fmt.Errorf("failed to create new epoch handler: %w", err)
	}

	return &epochHandler{
		slotHandler:            newSlotHandler(constants.slotDuration),
		epochNumber:            epochNumber,
		firstSlot:              firstSlot,
		constants:              constants,
		epochData:              epochData,
		handleSlot:             handleSlot,
		slotToPreRuntimeDigest: slotToPreRuntimeDigest,
	}, nil
}

// run executes the block production for each available successfully claimed slot
// it is important to note that any error will be transmitted through errCh
func (h *epochHandler) run(ctx context.Context, errCh chan<- error) {
	defer close(errCh)
	currSlot := getCurrentSlot(h.constants.slotDuration)

	// if currSlot < h.firstSlot, it means we're at genesis and waiting for the first slot to arrive.
	// we have to check it here to prevent int overflow.
	if currSlot >= h.firstSlot && currSlot-h.firstSlot > h.constants.epochLength {
		logger.Warnf("attempted to start epoch that has passed: current slot=%d, start slot of epoch=%d",
			currSlot, h.firstSlot,
		)
		errCh <- errEpochPast
		return
	}

	logger.Debugf("authoring in %d slots in epoch %d", len(h.slotToPreRuntimeDigest), h.epochNumber)

	for {
		currentSlot, err := h.slotHandler.waitForNextSlot(ctx)
		if err != nil {
			errCh <- err
			return
		}

		// check if the slot is an authoring slot otherwise wait for the next slot
		preRuntimeDigest, has := h.slotToPreRuntimeDigest[currentSlot.number]
		if !has {
			continue
		}

		err = h.handleSlot(h.epochNumber, currentSlot, h.epochData.authorityIndex, preRuntimeDigest)
		if err != nil {
			logger.Warnf("failed to handle slot %d: %s", currentSlot.number, err)
		}
	}
}
