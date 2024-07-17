package consensus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type PoS struct {
	Validators map[common.Address]*big.Int // Mapping of validator addresses to their stakes
}

func NewPoS(validators map[common.Address]*big.Int) *PoS {
	return &PoS{Validators: validators}
}

func (pos *PoS) ValidateBlock(header *types.Header) bool {
	// Simulate checking if the signer of the block is a validator and has enough stake
	stake, isValidator := pos.Validators[header.Coinbase]
	return isValidator && stake.Cmp(big.NewInt(1000)) > 0 // Check if stake is above some threshold
}
