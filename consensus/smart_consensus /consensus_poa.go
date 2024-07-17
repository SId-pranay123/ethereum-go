package smartconsensus

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// SmartPoA implements a simple Proof of Authority consensus mechanism.
type SmartPoA struct {
	authorities map[common.Address]bool // Set of nodes that are allowed to produce new blocks
}

// NewSmartPoA initializes and returns a new SmartPoA with a list of authorities.
func NewSmartPoA(authorityAddresses []common.Address) *SmartPoA {
	auths := make(map[common.Address]bool)
	for _, addr := range authorityAddresses {
		auths[addr] = true
	}
	return &SmartPoA{authorities: auths}
}

// Author retrieves the Ethereum address of the block's creator.
func (poa *SmartPoA) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyHeader checks whether the header conforms to the set consensus rules.
func (poa *SmartPoA) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	if _, authorized := poa.authorities[header.Coinbase]; !authorized {
		return errors.New("block's coinbase is not an authorized authority")
	}
	return nil
}

// Prepare initializes the consensus fields of a block header.
func (poa *SmartPoA) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	header.Difficulty = big.NewInt(1) // Fixed difficulty, as PoA does not involve complex computation
	return nil
}

// Finalize applies any state changes at block finalization.
func (poa *SmartPoA) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction) error {
	// PoA does not typically involve rewards, so this can be left empty unless specific incentives are designed.
	return nil
}

// APIs returns the RPC APIs this consensus engine provides.
func (poa *SmartPoA) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return nil
}

// Close terminates any background threads maintained by the consensus engine.
func (poa *SmartPoA) Close() error {
	return nil
}
