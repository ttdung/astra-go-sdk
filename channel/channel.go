package channel

import (
	"context"
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/ttdung/astra-go-sdk/account"
	"github.com/ttdung/astra-go-sdk/common"
	channelTypes "github.com/ttdung/channel_v0.46/x/channel/types"
)

type Channel struct {
	rpcClient client.Context
}

type SignMsgRequest struct {
	Msg      types.Msg
	GasLimit uint64
	GasPrice string
}

func NewChannel(rpcClient client.Context) *Channel {
	return &Channel{rpcClient}
}

func (cn *Channel) SignMultisigMsg(req SignMsgRequest,
	account *account.PrivateKeySerialized,
	multiSigPubkey cryptoTypes.PubKey) (string, error) {

	err := req.Msg.ValidateBasic()
	if err != nil {
		return "", err
	}

	newTx := common.NewTxMulSign(cn.rpcClient, account, req.GasLimit, req.GasPrice, 0, 2)
	txBuilder, err := newTx.BuildUnsignedTx(req.Msg)
	if err != nil {
		return "", err
	}

	err = newTx.SignTxWithSignerAddress(txBuilder, multiSigPubkey)
	if err != nil {
		return "", errors.Wrap(err, "SignTx")
	}

	sign, err := common.TxBuilderSignatureJsonEncoder(cn.rpcClient.TxConfig, txBuilder)
	if err != nil {
		return "", err
	}

	return sign, nil
}

func (cn *Channel) SignCommitmentMultisigMsg(req SignMsgRequest,
	account *account.PrivateKeySerialized,
	multiSigPubkey cryptoTypes.PubKey) (string, error) {

	err := req.Msg.ValidateBasic()
	if err != nil {
		return "", err
	}

	accNum, accSeq, err := cn.rpcClient.AccountRetriever.GetAccountNumberSequence(
		cn.rpcClient,
		types.AccAddress(multiSigPubkey.Address()))
	if err != nil {
		return "", errors.Wrap(err, "GetAccountNumberSequence")
	}

	log.Printf("SignCommitmentMultisigMsg: req %v \n", req)
	log.Printf("SignCommitmentMultisigMsg: sign with acc %v \n", account.AccAddress().String())
	log.Printf("SignCommitmentMultisigMsg: multiSigPubkey %v \n", multiSigPubkey.Address())
	log.Printf("SignCommitmentMultisigMsg: accNum %v, accSeq+1 %v, \n", accNum, accSeq+1)
	newTx := common.NewTxMulSign(cn.rpcClient, account, req.GasLimit, req.GasPrice, accSeq+1, accNum)
	txBuilder, err := newTx.BuildUnsignedTx(req.Msg)
	if err != nil {
		return "", err
	}

	err = newTx.SignTxWithSignerAddress(txBuilder, multiSigPubkey)
	if err != nil {
		return "", errors.Wrap(err, "SignTx")
	}

	sign, err := common.TxBuilderSignatureJsonEncoder(cn.rpcClient.TxConfig, txBuilder)
	if err != nil {
		return "", err
	}

	return sign, nil
}

func (cn *Channel) ListChannel() (*channelTypes.QueryAllChannelResponse, error) {
	channelClient := channelTypes.NewQueryClient(cn.rpcClient)
	return channelClient.ChannelAll(context.Background(), &channelTypes.QueryAllChannelRequest{})
}
