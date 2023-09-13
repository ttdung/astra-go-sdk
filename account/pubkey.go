package account

import (
	"encoding/hex"
	"github.com/evmos/evmos/v12/crypto/ethsecp256k1"
	"strings"

	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	//"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

type PKAccount struct {
	publicKey ethsecp256k1.PubKey //evm
	//publicKey secp256k1.PubKey
}

func NewPKAccount(pubkey string) (*PKAccount, error) {
	var key []byte
	if strings.Contains(pubkey, "{") {
		pubkey = strings.Split(strings.Split(pubkey, "{")[1], "}")[0]
		k, err := hex.DecodeString(pubkey)
		if err != nil {
			return nil, err
		}
		key = k
	}

	key = []byte(pubkey)

	return &PKAccount{
		//publicKey: secp256k1.PubKey{
		publicKey: ethsecp256k1.PubKey{
			Key: key,
		},
	}, nil
}

func (pka *PKAccount) PublicKey() cryptoTypes.PubKey {
	return &pka.publicKey
}

func (pka *PKAccount) AccAddress() types.AccAddress {
	pub := pka.PublicKey()
	addr := types.AccAddress(pub.Address())

	return addr
}
