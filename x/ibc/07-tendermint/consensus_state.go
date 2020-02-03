package tendermint

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	commitment "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
	tmtypes "github.com/tendermint/tendermint/types"
)

// ConsensusState defines a Tendermint consensus state
type ConsensusState struct {
	Timestamp    time.Time             `json:"timestamp" yaml:"timestamp"`
	Root         commitment.RootI      `json:"root" yaml:"root"`
	ValidatorSet *tmtypes.ValidatorSet `json:"validator_set" yaml:"validator_set"`
}

// ClientType returns Tendermint
func (ConsensusState) ClientType() clientexported.ClientType {
	return clientexported.Tendermint
}

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() commitment.RootI {
	return cs.Root
}

// GetTimestamp returns block time at which the consensus state was stored
func (cs ConsensusState) GetTimestamp() time.Time {
	return cs.Timestamp
}

// ValidateBasic defines a basic validation for the tendermint consensus state.
func (cs ConsensusState) ValidateBasic() error {
	if cs.Root == nil || cs.Root.IsEmpty() {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "root cannot be empty")
	}
	if cs.ValidatorSet == nil {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "validator set cannot be nil")
	}
	if cs.Timestamp.IsZero() {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "timestamp cannot be zero Unix time")
	}
	return nil
}
