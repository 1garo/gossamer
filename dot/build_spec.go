// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package dot

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ChainSafe/gossamer/dot/state"
	"github.com/ChainSafe/gossamer/dot/telemetry"
	"github.com/ChainSafe/gossamer/internal/log"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/genesis"
	"github.com/ChainSafe/gossamer/lib/utils"
)

// BuildSpec object for working with building genesis JSON files
type BuildSpec struct {
	genesis *genesis.Genesis
}

// ToJSON outputs genesis JSON in human-readable form
func (b *BuildSpec) ToJSON() ([]byte, error) {
	tmpGen := &genesis.Genesis{
		Name:       b.genesis.Name,
		ID:         b.genesis.ID,
		ChainType:  b.genesis.ChainType,
		Bootnodes:  b.genesis.Bootnodes,
		ProtocolID: b.genesis.ProtocolID,
		Properties: b.genesis.Properties,
		Genesis: genesis.Fields{
			Runtime: b.genesis.GenesisFields().Runtime,
		},
	}
	return json.MarshalIndent(tmpGen, "", "    ")
}

// ToJSONRaw outputs genesis JSON in raw form
func (b *BuildSpec) ToJSONRaw() ([]byte, error) {
	tmpGen := &genesis.Genesis{
		Name:       b.genesis.Name,
		ID:         b.genesis.ID,
		ChainType:  b.genesis.ChainType,
		Bootnodes:  b.genesis.Bootnodes,
		ProtocolID: b.genesis.ProtocolID,
		Properties: b.genesis.Properties,
		Genesis: genesis.Fields{
			Raw: b.genesis.GenesisFields().Raw,
		},
	}
	return json.MarshalIndent(tmpGen, "", "    ")
}

// BuildFromGenesis builds a BuildSpec based on the human-readable genesis file at path
func BuildFromGenesis(path string, authCount int) (*BuildSpec, error) {
	gen, err := genesis.NewGenesisFromJSON(path, authCount)
	if err != nil {
		return nil, err
	}
	bs := &BuildSpec{
		genesis: gen,
	}
	return bs, nil
}

// WriteGenesisSpecFile writes the build-spec in the output filepath
func WriteGenesisSpecFile(data []byte, fp string) error {
	// if file already exists then dont apply any written on it
	if utils.PathExists(fp) {
		return fmt.Errorf("file %s already exists, rename to avoid overwriting", fp)
	}

	if err := os.MkdirAll(filepath.Dir(fp), os.ModeDir|os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(fp, data, 0600)
}

// BuildFromDB builds a BuildSpec from the DB located at path
func BuildFromDB(path string) (*BuildSpec, error) {
	tmpGen := &genesis.Genesis{
		Name:       "",
		ID:         "",
		Bootnodes:  nil,
		ProtocolID: "",
		Genesis: genesis.Fields{
			Runtime: nil,
		},
	}
	tmpGen.Genesis.Raw = make(map[string]map[string]string)
	tmpGen.Genesis.Runtime = new(genesis.Runtime)

	config := state.Config{
		Path:      path,
		LogLevel:  log.Info,
		Telemetry: telemetry.NewNoopMailer(),
	}

	stateSrvc := state.NewService(config)

	err := stateSrvc.SetupBase()
	if err != nil {
		return nil, fmt.Errorf("cannot setup state database: %w", err)
	}

	// start state service (initialise state database)
	err = stateSrvc.Start()
	if err != nil {
		return nil, fmt.Errorf("cannot start state service: %w", err)
	}
	// set genesis fields data
	ent, err := stateSrvc.Storage.Entries(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get storage trie entries: %w", err)
	}
	err = genesis.BuildFromMap(ent, tmpGen)
	if err != nil {
		return nil, fmt.Errorf("failed to build from map: %w", err)
	}
	// set genesisData
	gd, err := stateSrvc.DB().Get(common.GenesisDataKey)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve genesis data: %w", err)
	}
	gData := &genesis.Data{}
	err = json.Unmarshal(gd, gData)
	if err != nil {
		return nil, err
	}
	tmpGen.Name = gData.Name
	tmpGen.ID = gData.ID
	tmpGen.Bootnodes = common.BytesToStringArray(gData.Bootnodes)
	tmpGen.ProtocolID = gData.ProtocolID

	bs := &BuildSpec{
		genesis: tmpGen,
	}
	return bs, nil
}
