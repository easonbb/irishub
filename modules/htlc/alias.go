// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/irisnet/irishub/modules/htlc/internal/types
// ALIASGEN: github.com/irisnet/irishub/modules/htlc/internal/keeper
package htlc

import (
	"github.com/irisnet/irishub/modules/htlc/internal/keeper"
	"github.com/irisnet/irishub/modules/htlc/internal/types"
)

const (
	EventTypeCreateHTLC                = types.EventTypeCreateHTLC
	EventTypeClaimHTLC                 = types.EventTypeClaimHTLC
	EventTypeRefundHTLC                = types.EventTypeRefundHTLC
	EventTypeExpiredHTLC               = types.EventTypeExpiredHTLC
	AttributeValueCategory             = types.AttributeValueCategory
	AttributeValueSender               = types.AttributeValueSender
	AttributeValueReceiver             = types.AttributeValueReceiver
	AttributeValueReceiverOnOtherChain = types.AttributeValueReceiverOnOtherChain
	AttributeValueAmount               = types.AttributeValueAmount
	AttributeValueHashLock             = types.AttributeValueHashLock
	AttributeValueSecret               = types.AttributeValueSecret
	OPEN                               = types.OPEN
	COMPLETED                          = types.COMPLETED
	EXPIRED                            = types.EXPIRED
	ModuleName                         = types.ModuleName
	RouterKey                          = types.RouterKey
	StoreKey                           = types.StoreKey
	QuerierRoute                       = types.QuerierRoute
	TypeMsgCreateHTLC                  = types.TypeMsgCreateHTLC
	TypeMsgClaimHTLC                   = types.TypeMsgClaimHTLC
	TypeMsgRefundHTLC                  = types.TypeMsgRefundHTLC
	SecretLength                       = types.SecretLength
	HashLockLength                     = types.HashLockLength
	MaxLengthForAddressOnOtherChain    = types.MaxLengthForAddressOnOtherChain
	MinTimeLock                        = types.MinTimeLock
	MaxTimeLock                        = types.MaxTimeLock
	QueryHTLC                          = types.QueryHTLC
)

var (
	// functions aliases
	RegisterCodec              = types.RegisterCodec
	ErrUnknownHTLC             = types.ErrUnknownHTLC
	ErrInvalidHashLock         = types.ErrInvalidHashLock
	ErrInvalidSecret           = types.ErrInvalidSecret
	ErrHashLockAlreadyExists   = types.ErrHashLockAlreadyExists
	ErrStateIsNotOpen          = types.ErrStateIsNotOpen
	ErrStateIsNotExpired       = types.ErrStateIsNotExpired
	NewGenesisState            = types.NewGenesisState
	DefaultGenesisState        = types.DefaultGenesisState
	ValidateGenesis            = types.ValidateGenesis
	NewHTLC                    = types.NewHTLC
	HTLCStateFromString        = types.HTLCStateFromString
	NewMsgCreateHTLC           = types.NewMsgCreateHTLC
	NewMsgClaimHTLC            = types.NewMsgClaimHTLC
	NewMsgRefundHTLC           = types.NewMsgRefundHTLC
	SHA256                     = types.SHA256
	GetHashLock                = types.GetHashLock
	NewKeeper                  = keeper.NewKeeper
	KeyHTLC                    = keeper.KeyHTLC
	KeyHTLCExpireQueue         = keeper.KeyHTLCExpireQueue
	KeyHTLCExpireQueueSubspace = keeper.KeyHTLCExpireQueueSubspace
	NewQuerier                 = keeper.NewQuerier

	// variable aliases
	ModuleCdc             = types.ModuleCdc
	HTLCStateToStringMap  = types.HTLCStateToStringMap
	StringToHTLCStateMap  = types.StringToHTLCStateMap
	KeyDelimiter          = keeper.KeyDelimiter
	PrefixHTLC            = keeper.PrefixHTLC
	PrefixHTLCExpireQueue = keeper.PrefixHTLCExpireQueue
)

type (
	SupplyKeeper    = types.SupplyKeeper
	GenesisState    = types.GenesisState
	HTLC            = types.HTLC
	HTLCState       = types.HTLCState
	HTLCSecret      = types.HTLCSecret
	HTLCHashLock    = types.HTLCHashLock
	MsgCreateHTLC   = types.MsgCreateHTLC
	MsgClaimHTLC    = types.MsgClaimHTLC
	MsgRefundHTLC   = types.MsgRefundHTLC
	QueryHTLCParams = types.QueryHTLCParams
	Keeper          = keeper.Keeper
)