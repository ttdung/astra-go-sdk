package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sync"
	"testing"

	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	signingTypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/ttdung/astra-go-sdk/bank"
	"github.com/ttdung/astra-go-sdk/channel"
	"github.com/ttdung/astra-go-sdk/common"
	"github.com/ttdung/astra-go-sdk/config"
)

type AstraSdkTestSuite struct {
	suite.Suite
	Client *Client
}

func (suite *AstraSdkTestSuite) SetupTest() {
	err := godotenv.Load("./../dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := &config.Config{
		ChainId:       os.Getenv("CHAIN_ID"),
		Endpoint:      os.Getenv("END_POINT"),
		CoinType:      60,
		PrefixAddress: "astra",
		TokenSymbol:   "aastra",
	}

	client := NewClient(cfg)
	suite.Client = client
}

func TestAstraSdkTestSuite(t *testing.T) {
	suite.Run(t, new(AstraSdkTestSuite))
}

func (suite *AstraSdkTestSuite) TestInitBank() {
	bankClient := suite.Client.NewBankClient()
	balance, err := bankClient.Balance("astra1hlue55l54erxk3asqzvkfsm5yl0du50twljtlp")
	if err != nil {
		panic(err)
	}

	fmt.Println(balance.String())
}

func (suite *AstraSdkTestSuite) TestInitChannel() {
	channelClient := suite.Client.NewChannelClient()
	channels, err := channelClient.ListChannel()
	if err != nil {
		panic(err)
	}
	fmt.Println(channels)
}

func (suite *AstraSdkTestSuite) TestGenAccount() {
	accClient := suite.Client.NewAccountClient()
	acc, err := accClient.CreateAccount()
	if err != nil {
		panic(err)
	}

	data, _ := acc.String()

	fmt.Println(data)
}

func (suite *AstraSdkTestSuite) TestGenMulSignAccount() {
	accClient := suite.Client.NewAccountClient()
	acc, addr, pubKey, err := accClient.CreateMulSignAccount(3, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println("addr", addr)
	fmt.Println("pucKey", pubKey)
	fmt.Println("list key")
	for i, item := range acc {
		fmt.Println("index", i)
		fmt.Println(item.String())
	}
}

func (suite *AstraSdkTestSuite) TestTransfer() {
	bankClient := suite.Client.NewBankClient()

	amount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).SetUint64(uint64(math.Pow10(18))))
	fmt.Println("amount", amount.String())

	request := &bank.TransferRequest{
		PrivateKey: "valve season sauce knife burden benefit zone field ask carpet fury vital action donate trade street ability artwork ball uniform garbage sugar warm differ",
		Receiver:   "astra1p6sscujfpygmrrxqlwqeqqw6r5lxk2x9gz9glh",
		Amount:     amount,
		GasLimit:   200000,
		GasPrice:   "0.001aastra",
	}

	txBuilder, err := bankClient.TransferRawData(request)
	if err != nil {
		panic(err)
	}

	txJson, err := common.TxBuilderJsonEncoder(suite.Client.rpcClient.TxConfig, txBuilder)
	if err != nil {
		panic(err)
	}

	fmt.Println("rawData", string(txJson))

	txByte, err := common.TxBuilderJsonDecoder(suite.Client.rpcClient.TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	fmt.Println(ethCommon.BytesToHash(txByte).String())

	res, err := suite.Client.rpcClient.BroadcastTxCommit(txByte)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func (suite *AstraSdkTestSuite) TestTransferWithPrivateKey() {
	bankClient := suite.Client.NewBankClient()
	amount := big.NewInt(0).Mul(big.NewInt(20), big.NewInt(0).SetUint64(uint64(math.Pow10(18))))
	fmt.Println("amount", amount.String())

	request := &bank.TransferRequest{
		PrivateKey: "69e2ece17baa00b1112217f530661a8b9d0ecabc8fe122fc1f403761c86a1ccc",
		Receiver:   "astra1p6sscujfpygmrrxqlwqeqqw6r5lxk2x9gz9glh",
		Amount:     amount,
		GasLimit:   200000,
		GasPrice:   "0.001aastra",
	}

	txBuilder, err := bankClient.TransferRawDataWithPrivateKey(request)
	if err != nil {
		panic(err)
	}

	txJson, err := common.TxBuilderJsonEncoder(suite.Client.rpcClient.TxConfig, txBuilder)
	if err != nil {
		panic(err)
	}

	fmt.Println("rawData", string(txJson))

	txByte, err := common.TxBuilderJsonDecoder(suite.Client.rpcClient.TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	fmt.Println(ethCommon.BytesToHash(txByte).String())

	res, err := suite.Client.rpcClient.BroadcastTxCommit(txByte)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func (suite *AstraSdkTestSuite) TestTransferMultiSign() {
	//main address
	/*
		addr astra1ha0vgh05zzlwdeejxq9aq7gqr6jzs7stdhlfra
		pucKey {"@type":"/cosmos.crypto.multisig.LegacyAminoPubKey","threshold":2,"public_keys":[{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A0ATAOfWQM6XXCA5po9DBsKVGmWudnIN55arHhDYhR89"},{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A0ks8ww7AVKYQRsKgZSQi9wTfoQzKNt30gLOMpOJNSPn"},{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A9Q4nSS73SG+Tclghh1JEtfng5vd41dgmG7HJrYW4/Ml"}]}
	*/

	//child address
	/*
		index 0
		{
		 "address": "astra1dmdsy082730stdletm7z6zulfxuez4lsx3tztx",
		 "hexAddress": "0x6Edb023ceAF45F05b7f95efC2d0B9f49B99157F0",
		 "mnemonic": "ignore risk morning strike school street radar silk recipe health december system inflict gold foster item end twenty magic shine oppose island loop impact",
		 "privateKey": "7f1d3df4044f09b1edfab34c7e3fee92396ea23861e96a8ac7429efcf158d794",
		 "publicKey": "{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A0ATAOfWQM6XXCA5po9DBsKVGmWudnIN55arHhDYhR89\"}",
		 "type": "eth_secp256k1",
		 "validatorKey": "astravaloper1dmdsy082730stdletm7z6zulfxuez4lsrg2nsg"
		} <nil>
		index 1
		{
		 "address": "astra1fd39nlc4hsl7ma9knpjwlhcrnunz66dnvf5agx",
		 "hexAddress": "0x4b6259ff15Bc3FEdf4B69864EfdF039F262d69B3",
		 "mnemonic": "seven mean snap illness couch excite item topic tobacco erosion tourist blue van possible wolf gadget combine excess brush goddess glory subway few mind",
		 "privateKey": "8dca20a27b0bfdcf1dacc9b2f71d4b7e7d269a4b87949707c12ef2ba328fd0e9",
		 "publicKey": "{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A0ks8ww7AVKYQRsKgZSQi9wTfoQzKNt30gLOMpOJNSPn\"}",
		 "type": "eth_secp256k1",
		 "validatorKey": "astravaloper1fd39nlc4hsl7ma9knpjwlhcrnunz66dnfs4vng"
		} <nil>
		index 2
		{
		 "address": "astra1gc0v03kjrg9uv7duvzqsndv3nhkhehvkwuhkdr",
		 "hexAddress": "0x461EC7C6D21a0BC679bC608109b5919DEd7Cdd96",
		 "mnemonic": "swap exhaust letter left light trust diet piano pride rifle trust orbit clip suggest achieve unaware please guess lawsuit doctor use bargain jealous weekend",
		 "privateKey": "e3f46776e933129611b3cb6418176dcd2a9badd8188fb4804d5b822548200bac",
		 "publicKey": "{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A9Q4nSS73SG+Tclghh1JEtfng5vd41dgmG7HJrYW4/Ml\"}",
		 "type": "eth_secp256k1",
		 "validatorKey": "astravaloper1gc0v03kjrg9uv7duvzqsndv3nhkhehvkt9k8kd"
		}
	*/

	pk, err := common.DecodePublicKey(
		suite.Client.rpcClient,
		"{\"@type\":\"/cosmos.crypto.multisig.LegacyAminoPubKey\",\"threshold\":2,\"public_keys\":[{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A0ATAOfWQM6XXCA5po9DBsKVGmWudnIN55arHhDYhR89\"},{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A0ks8ww7AVKYQRsKgZSQi9wTfoQzKNt30gLOMpOJNSPn\"},{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A9Q4nSS73SG+Tclghh1JEtfng5vd41dgmG7HJrYW4/Ml\"}]}",
	)
	if err != nil {
		panic(err)
	}

	from := types.AccAddress(pk.Address())
	fmt.Println("from", from.String())

	bankClient := suite.Client.NewBankClient()

	listPrivate := []string{
		"ignore risk morning strike school street radar silk recipe health december system inflict gold foster item end twenty magic shine oppose island loop impact",
		"seven mean snap illness couch excite item topic tobacco erosion tourist blue van possible wolf gadget combine excess brush goddess glory subway few mind",
	}

	thread := 2
	listRawdata := make([][]byte, 0)

	for i := 0; i < thread; i++ {
		amount := big.NewInt(0).Mul(big.NewInt(10+int64(i)), big.NewInt(0).SetUint64(uint64(math.Pow10(18))))
		fmt.Println("amount", amount.String())

		fmt.Println("start signer")
		signList := make([][]signingTypes.SignatureV2, 0)
		for i, s := range listPrivate {
			fmt.Println("index", i)
			request := &bank.SignTxWithSignerAddressRequest{
				SignerPrivateKey: s,
				SignerPublicKey:  pk,
				Receiver:         "astra156dh69y8j39eynue4jahrezg32rgl8eck5rhsl",
				Amount:           amount,
				GasLimit:         200000,
				GasPrice:         "0.001aastra",
			}

			txBuilder, err := bankClient.SignTxWithSignerAddress(request)
			if err != nil {
				panic(err)
			}

			sign, err := common.TxBuilderSignatureJsonEncoder(suite.Client.rpcClient.TxConfig, txBuilder)
			if err != nil {
				panic(err)
			}

			fmt.Println("sign-data", string(sign))

			signByte, err := common.TxBuilderSignatureJsonDecoder(suite.Client.rpcClient.TxConfig, sign)
			if err != nil {
				panic(err)
			}

			signList = append(signList, signByte)
		}

		fmt.Println("start transfer")
		//200
		request := &bank.TransferMultiSignRequest{
			MulSignAccPublicKey: pk,
			Receiver:            "astra156dh69y8j39eynue4jahrezg32rgl8eck5rhsl",
			Amount:              amount,
			GasLimit:            200000,
			GasPrice:            "0.001aastra",
			Sigs:                signList,
		}

		txBuilder, err := bankClient.TransferMultiSignRawData(request)
		if err != nil {
			panic(err)
		}

		txJson, err := common.TxBuilderJsonEncoder(suite.Client.rpcClient.TxConfig, txBuilder)
		if err != nil {
			panic(err)
		}

		fmt.Println("rawData", string(txJson))

		txByte, err := common.TxBuilderJsonDecoder(suite.Client.rpcClient.TxConfig, txJson)
		if err != nil {
			panic(err)
		}

		txHash := common.TxHash(txByte)
		fmt.Println("txHash", txHash)

		listRawdata = append(listRawdata, txByte)
	}

	var wg sync.WaitGroup
	wg.Add(thread)

	go func(item []byte, client client.Context) {
		_, err := client.BroadcastTxCommit(item)
		if err != nil {
			panic(err)
		}

		//fmt.Println("BroadcastTxCommit", res)
		defer wg.Done()
	}(listRawdata[0], suite.Client.rpcClient)

	go func(item []byte, client client.Context) {
		_, err := client.BroadcastTxCommit(item)
		if err != nil {
			panic(err)
		}

		//fmt.Println("BroadcastTxCommit", res)
		defer wg.Done()
	}(listRawdata[1], suite.Client.rpcClient)

	wg.Wait()
}

func (suite *AstraSdkTestSuite) TestAddressValid() {
	addressCheck := "astra1hann2zj3sx3ympd40ptxdmpd4nd4eypm45zhhr"
	addressCheck = "astra19a3mu6k0y326mcny60m3x70qfxtkms20sn5j8p"
	receiver, err := types.AccAddressFromBech32(addressCheck)
	if err != nil {
		panic(err)
	}

	fmt.Println(receiver.String())
	assert.Equal(suite.T(), addressCheck, receiver.String(), "they should be equal")

	rs, _ := common.IsAddressValid(addressCheck)
	assert.Equal(suite.T(), rs, true)
}

func (suite *AstraSdkTestSuite) TestConvertHexToCosmosAddress() {
	eth := "0x9cc92bd19df168539ba7c73b450db998b0e79761"
	cosmos := "astra1nnyjh5va79598xa8cua52rdenzcw09mpwfekts"

	rs, _ := common.EthAddressToCosmosAddress(eth)
	fmt.Println(rs)
	assert.Equal(suite.T(), cosmos, rs)

	rs1, _ := common.CosmosAddressToEthAddress(cosmos)
	fmt.Println(rs1)
	assert.Equal(suite.T(), eth, rs1)
}

func (suite *AstraSdkTestSuite) TestCheckTx() {
	bankClient := suite.Client.NewBankClient()
	//rs, err := bankClient.CheckTx("646F944DCDB201F674C109E6EF9A594ADBCC33B8F0FA054D7B3F4ABE4CCA2AEB")
	rs, err := bankClient.CheckTx("25D2704C3ABDFE3DBBC1A8202A15D43A2E86D7F1F24AD2E704A7F50FCB75FB94")
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.Code)
	if rs != nil && common.BlockedStatus(rs.Code) {
		fmt.Println("blocked")
	}
}

func (suite *AstraSdkTestSuite) TestImportAccountViaHdPath() {
	accClient := suite.Client.NewAccountClient()

	_, err := common.VerifyHdPath("m/44'/60'/0'/0/0")
	if err != nil {
		panic(err)
	}

	nmemonic := "secret immense amount trial polar security mother scare useful hen squeeze confirm right size best trash team clock matter grow copy quiz capital ill"

	for i := 100083357; i <= (100083357 + 20); i++ {
		s := fmt.Sprintf("m/44'/60'/%v'/1/0", i)
		wallet, err := accClient.ImportHdPath(
			nmemonic,
			s,
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("index ", i, s)
		fmt.Println(wallet.String())
	}
}

func (suite *AstraSdkTestSuite) TestImportByNmemonic() {
	accClient := suite.Client.NewAccountClient()
	key, err := accClient.ImportAccount("effort behave trash gaze youth food north brain poverty drive armed split kind script fox frog breeze cliff bright raise napkin question payment upset")
	if err != nil {
		panic(err)
	}

	fmt.Println(key.String())
}

func (suite *AstraSdkTestSuite) TestImportByPrivatekey() {
	accClient := suite.Client.NewAccountClient()
	key, err := accClient.ImportPrivateKey("b8f7f2e5bab9c0b08df50cb5aa93ca8d1f5fe4aa11677ebf05232930d28349a9")
	if err != nil {
		panic(err)
	}

	fmt.Println(key.String())
}

func (suite *AstraSdkTestSuite) TestScanner() {
	bankClient := suite.Client.NewBankClient()
	c := suite.Client.NewScanner(bankClient)
	//listTx, err := c.ScanByBlockHeight(2040457) //cosmos
	//listTx, err := c.ScanByBlockHeight(1871260) //erc20
	listTx, err := c.ScanByBlockHeight(2030365) //erc20
	if err != nil {
		panic(err)
	}

	rs, _ := json.MarshalIndent(listTx, " ", " ")
	fmt.Println(string(rs))
}

func (suite *AstraSdkTestSuite) TestGetTxDetail() {
	bankClient := suite.Client.NewBankClient()
	rs, err := bankClient.TxDetail("2A01E4B7FC44FE90387241AA6A067D420838940CA4B3605E5CB4AB39BFDD0320")

	if err != nil {
		panic(err)
	}

	rsMarshal, _ := json.MarshalIndent(rs, " ", " ")

	fmt.Println(string(rsMarshal))

}

func (suite *AstraSdkTestSuite) TestSequenceNumberFromPk() {
	mulSignAccPubKey := "{\"@type\":\"/cosmos.crypto.multisig.LegacyAminoPubKey\",\"threshold\":2,\"public_keys\":[{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A0UjEVXxXA7JY2oou5HPH7FuPSyJ2hAfDMc4XThXiopM\"},{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A6DFr74kQmk/k88fCTPCxmf9kyFJMhFUF21IPFY7XoV2\"},{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"AgPQELGzKmlAaSb01OKbmuL1f17MHJshkh9s9xAWxMa3\"}]}"

	walletMultiPub, err := common.DecodePublicKey(suite.Client.RpcClient(), mulSignAccPubKey)
	if err != nil {
		panic(err)
	}

	masterHexAddr := ethCommon.BytesToAddress(walletMultiPub.Address().Bytes())
	fmt.Println(masterHexAddr)

	bankClient := suite.Client.NewBankClient()
	accNum, accSeq, err := bankClient.AccountRetriever(masterHexAddr.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(accNum)
	fmt.Println(accSeq)
}

func (suite *AstraSdkTestSuite) TestConvertToDecimal() {
	amount, err := common.ConvertToDecimal("740000000000", 18)
	fmt.Println(err)
	fmt.Println(amount)
}

func (suite *AstraSdkTestSuite) TestOpenChannel() {
	channelClient := suite.Client.NewChannelClient()
	acc := suite.Client.NewAccountClient()
	account1, err := acc.ImportAccount("gadget final blue appear hero retire wild account message social health hobby decade neglect common egg cruel certain phrase myself alert enlist brother sure")
	if err != nil {
		panic(err)
	}
	account2, err := acc.ImportAccount("salute debate real reject wreck topple derive night height job range enrich juice develop crush install method always vacant napkin blush beyond hedgehog tortoise")
	if err != nil {
		panic(err)
	}

	bankClient := suite.Client.NewBankClient()
	balance1Origin, err := bankClient.Balance(account1.AccAddress().String())
	balance2Origin, err := bankClient.Balance(account2.AccAddress().String())
	if err != nil {
		panic(err)
	}
	fmt.Println("Balance 1 - origin: ", balance1Origin)

	multisigAddr, multiSigPubkey, err := acc.CreateMulSignAccountFromTwoAccount(account1.PublicKey(), account2.PublicKey(), 2)
	if err != nil {
		panic(err)
	}

	fmt.Println("multisigAddr", multisigAddr)
	fmt.Println("Deposit init multisignAddr")
	amount := big.NewInt(1)

	fmt.Println("deposit amount", amount.String())
	request1 := &bank.TransferRequest{
		PrivateKey: "gadget final blue appear hero retire wild account message social health hobby decade neglect common egg cruel certain phrase myself alert enlist brother sure",
		Receiver:   multisigAddr,
		Amount:     amount,
		GasLimit:   200000,
		GasPrice:   "0.001aastra",
	}

	txResult1, err := bankClient.TransferRawDataAndBroadcast(request1)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx transfer result code", txResult1.Code)

	openChannelRequest := channel.SignMsgRequest{
		Msg: &channelTypes.MsgOpenChannel{
			Creator: multisigAddr,
			PartA:   account1.AccAddress().String(),
			PartB:   account2.AccAddress().String(),
			CoinA: &types.Coin{
				Denom:  "astra",
				Amount: types.NewInt(1),
			},
			CoinB: &types.Coin{
				Denom:  "astra",
				Amount: types.NewInt(1),
			},
			MultisigAddr: multisigAddr,
			Sequence:     "8",
		},
		GasLimit: 200000,
		GasPrice: "0.001aastra",
	}

	signList := make([][]signingTypes.SignatureV2, 0)
	strSig1, err := channelClient.SignMultisigMsg(openChannelRequest, account1, multiSigPubkey)
	if err != nil {
		panic(err)
	}
	signByte1, err := common.TxBuilderSignatureJsonDecoder(suite.Client.rpcClient.TxConfig, strSig1)
	if err != nil {
		panic(err)
	}

	signList = append(signList, signByte1)

	strSig2, err := channelClient.SignMultisigMsg(openChannelRequest, account2, multiSigPubkey)
	if err != nil {
		panic(err)
	}
	signByte2, err := common.TxBuilderSignatureJsonDecoder(suite.Client.rpcClient.TxConfig, strSig2)
	if err != nil {
		panic(err)
	}

	signList = append(signList, signByte2)

	fmt.Println("new tx multisign")

	newTx := common.NewTxMulSign(suite.Client.rpcClient,
		nil,
		openChannelRequest.GasLimit,
		openChannelRequest.GasPrice,
		0,
		2)

	txBuilderMultiSign, err := newTx.BuildUnsignedTx(openChannelRequest.Msg)
	if err != nil {
		panic(err)
	}

	err = newTx.CreateTxMulSign(txBuilderMultiSign, multiSigPubkey, suite.Client.coinType, signList)
	if err != nil {
		panic(err)
	}

	txJson, err := common.TxBuilderJsonEncoder(suite.Client.rpcClient.TxConfig, txBuilderMultiSign)
	if err != nil {
		panic(err)
	}
	fmt.Println("rawData", string(txJson))

	txByte, err := common.TxBuilderJsonDecoder(suite.Client.rpcClient.TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	fmt.Println(ethCommon.BytesToHash(txByte).String())

	txResult2, err := suite.Client.rpcClient.BroadcastTxCommit(txByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx openchannel result code", txResult2.Code)

	balance1After, err := bankClient.Balance(account1.AccAddress().String())
	balance2After, err := bankClient.Balance(account2.AccAddress().String())
	if err != nil {
		panic(err)
	}
	fmt.Println("balance 1 - after", balance1After)
	fmt.Println("balance 2 - after", balance1After)

	diff := balance1Origin.Sub(balance1Origin, balance1After)
	diff2 := balance1Origin.Sub(balance2Origin, balance2After)
	fmt.Println("Account1 Balance decrease", diff)
	fmt.Println("Account2 Balance decrease", diff2)
}
