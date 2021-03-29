package main

import (
	gen "github.com/whyrusleeping/cbor-gen"

	"github.com/chenjianmei111/go-state-types/abi"
)

func main() {
	// Common types
	if err := gen.WriteTupleEncodersToFile("./abi/cbor_gen.go", "abi",
		abi.PieceInfo{},
		abi.SectorID{},
	); err != nil {
		panic(err)
	}
}
