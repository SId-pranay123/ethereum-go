package consensus

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestPoSValidation(t *testing.T) {
	validators := make(map[common.Address]*big.Int)
	validators[common.HexToAddress("0x0000000000000000000000000000000000000001")] = big.NewInt(5000) // Validator with sufficient stake

	pos := NewPoS(validators)

	header := &types.Header{
		Coinbase: common.HexToAddress("0x0000000000000000000000000000000000000000"),
	}

	if !pos.ValidateBlock(header) {
		t.Errorf("Block should be valid")
	}
}
