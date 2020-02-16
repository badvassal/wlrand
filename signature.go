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

func CreateSignature(cfg randomizeCfg) (*msq.Body, error) {
	sig := Signature{
		Description: "Randomized by wlrand",
		Version:     version.VersionStr(),
		Cfg:         cfg,
	}

	j, err := json.MarshalIndent(sig, "    ", "")
	if err != nil {
		return nil, wlerr.Errorf("failed to marshal signature to JSON")
	}

	return &msq.Body{
		SecSection:   j,
		PlainSection: nil,
	}, nil
}

func FindSignature(bodies []msq.Body) (int, *Signature) {
	sig := &Signature{}
	for i, b := range bodies {
		if err := json.Unmarshal(b.SecSection, sig); err == nil {
			return i, sig
		}
	}

	return -1, nil
}
