package main

import (
	"encoding/json"

	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
	"github.com/badvassal/wlrand/version"
)

type Signature struct {
	Description string
	Version     string
	Cfg         randomizeCfg
}

func CreateSignatureMSQBlock(cfg randomizeCfg) (*msq.Block, error) {
	sig := Signature{
		Description: "Randomized by wlrand",
		Version:     version.VersionStr(),
		Cfg:         cfg,
	}

	j, err := json.MarshalIndent(sig, "    ", "")
	if err != nil {
		return nil, wlerr.Errorf("failed to marshal signature to JSON")
	}

	return &msq.Block{
		EncSection:   j,
		PlainSection: nil,
	}, nil
}

func FindSignatureMSQBlock(blocks []msq.Block) (int, *Signature) {
	sig := &Signature{}
	for i, b := range blocks {
		if err := json.Unmarshal(b.EncSection, sig); err == nil {
			return i, sig
		}
	}

	return -1, nil
}
