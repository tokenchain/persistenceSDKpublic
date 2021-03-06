package acl

import (
	"github.com/commitHub/commitBlockchain/types"
	wire "github.com/commitHub/commitBlockchain/wire"
)

//RegisterWire : Most users shouldn't use this, but this comes handy for tests.
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(DefineACLBody{}, "commit-blockchain/DefineACLBody", nil)
	cdc.RegisterConcrete(DefineOrganizationBody{}, "commit-blockchain/DefineOrganizationBody", nil)
	cdc.RegisterConcrete(DefineZoneBody{}, "commit-blockchain/DefineZoneBody", nil)
}

//RegisterACLAccount :  register acl account type and interface
func RegisterACLAccount(cdc *wire.Codec) {
	cdc.RegisterInterface((*types.ACLAccount)(nil), nil)
	cdc.RegisterConcrete(&types.BaseACLAccount{}, "commit-blockchain/AclAccount", nil)
}

var msgCdc = wire.NewCodec()

func init() {
	RegisterWire(msgCdc)
	RegisterACLAccount(msgCdc)
}
