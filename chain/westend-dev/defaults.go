// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package westenddev

import (
	cfg "github.com/ChainSafe/gossamer/config"
)

const (
	// defaultName is the default node name
	defaultName = "Westend"
	// defaultID is the default node ID
	defaultID = "westend_dev"
	// defaultBasePath is the default basepath for the westend dev node
	defaultBasePath = "~/.gossamer/westend-dev"
	// defaultChainSpec is the default chain spec for the westend dev node
	defaultChainSpec = "./chain/westend-dev/westend-dev-spec-raw.json"
)

// DefaultConfig returns a westend dev node configuration
func DefaultConfig() *cfg.Config {
	config := cfg.DefaultConfig()
	config.BasePath = defaultBasePath
	config.ID = defaultID
	config.Name = defaultName
	config.ChainSpec = defaultChainSpec
	config.RPC.RPCExternal = true
	config.RPC.UnsafeRPC = true
	config.RPC.WSExternal = true
	config.RPC.UnsafeWSExternal = true

	return config
}
