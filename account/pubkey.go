package account

import (
	"encoding/hex"
	"strings"

	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

type PKAccount struct {
	publicKey ethsecp256k1.PubKey
}

func NewPKAccount(pubkey string) *PKAccount {
	pubkey = strings.Split(strings.Split(pubkey, "{")[1], "}")[0]
	key, err := hex.DecodeString(pubkey)
	if err != nil {
		panic(err)
	}
	return &PKAccount{
		publicKey: ethsecp256k1.PubKey{
			Key: key,
		},
	}
}

func (pka *PKAccount) PublicKey() cryptoTypes.PubKey {
	return &pka.publicKey
}

func (pka *PKAccount) AccAddress() types.AccAddress {
	pub := pka.PublicKey()
	addr := types.AccAddress(pub.Address())

	return addr
}
