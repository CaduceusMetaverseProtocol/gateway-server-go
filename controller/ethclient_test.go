package controller

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)
var helloContractstring = `[{"constant":false,"inputs":[{"name":"_greeting","type":"string"}],"name":"setGreeting","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"getGreeting","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_greeting","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`

func TestGetBlance(t *testing.T) {
	ec :=NewEmController()
	balance,err := ec.GetBalance("0x9321A2F1DC2B417B7165870A27873D2A7813A616",nil)
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(balance)
}

func TestEmControler_Call(t *testing.T) {
	ec :=NewEmController()
	from := common.HexToAddress("0x9321A2F1DC2B417B7165870A27873D2A7813A616")
	to := common.HexToAddress("0x3bfC61cbA24a835F99a5855e613b239c7110B619")
	//param := "hello ethermint,this is setGreeting"
	res,err := ec.doCall(from,to,"getGreeting",helloContractstring,false)
	if err!=nil{
		t.Errorf(err.Error())
	}
	// unpack("getGreeting",res)
	fmt.Println(string(res))
}
//func unpack(message string,data []byte ){
//	abi, err := ethabi.JSON(strings.NewReader(helloContractstring))
//	if err != nil {
//		return
//	}
//	method, _ := abi.Methods[message]
//	res,err := method.Outputs.Unpack(data)
//	fmt.Println(res)
//}

func TestEmControler_EstimateGas(t *testing.T) {
	ec :=NewEmController()
	from := common.HexToAddress("0x9321A2F1DC2B417B7165870A27873D2A7813A616")
	to := common.HexToAddress("0x3bfC61cbA24a835F99a5855e613b239c7110B619")
	//param := "hello ethermint,this is setGreeting"
	res,err := ec.doEstimate(from,to,"getGreeting",helloContractstring,)
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(res)
}

func TestEmControler_SendTransaction(t *testing.T) {
	ec :=NewEmController()
	from := common.HexToAddress("0x9321A2F1DC2B417B7165870A27873D2A7813A616")
	to := common.HexToAddress("0x3bfC61cbA24a835F99a5855e613b239c7110B619")
	param := "nihao ethermint,this is setGreeting"
	value := hexutil.Big(*(big.NewInt(0)))
	txId,err := ec.doSendTransaction(from,&to,&value,"setGreeting",helloContractstring,param)
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(txId) //0x9929146efdab638ceae5a75ec1fb7c8480a2388f75b79326b8e9acecf4badd03
}

func TestEmControler_GetBlockByNumber(t *testing.T) {
	ec :=NewEmController()
	ec.GetBlockByNumber(nil)
}

func TestEmControler_GetBlockNumber(t *testing.T) {
	ec :=NewEmController()
	number,err := ec.GetBlockNumber()
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(number)
}

func TestEmControler_GetBlockByHash(t *testing.T) {
	ec :=NewEmController()
	//hash := common.HexToHash("")
	ec.GetBlockByHash("hash")
}

func TestEmControler_GetTransactionByHash(t *testing.T) {
	ec :=NewEmController()
	//hash := common.HexToHash("0x8e114fe45af750cb4cedfc9284168d8cd6c71506dbe2b5310e2bc541fed7651a")
	ec.GetTransactionByHash("hash")
}

func TestEmControler_GetTransactionReceipt(t *testing.T) {
	ec :=NewEmController()
	//hash := common.HexToHash("0x2da293ea5329fcfe17caaf593addef02f0bba6e4e0d137694ae3ff6381e2491e")
	ec.GetTransactionReceipt("hash")
}

func TestEmControler_GetCode(t *testing.T) {
	ec :=NewEmController()
	//to := common.HexToAddress("0x3bfC61cbA24a835F99a5855e613b239c7110B619")
	code,err := ec.GetCode("",nil)
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(string(code))
}

func TestEmControler_GetNonceAt(t *testing.T) {
	ec :=NewEmController()
	//from := common.HexToAddress("0x9321A2F1DC2B417B7165870A27873D2A7813A616")
	nonce,err := ec.GetNonceAt("from",nil)
	if err!=nil{
		t.Errorf(err.Error())
	}
	fmt.Println(nonce)
}

//func TestRawTransaction(t *testing.T) {
//	ec :=NewEmControler()
//	from := common.HexToAddress("0x9321A2F1DC2B417B7165870A27873D2A7813A616")
//	to := common.HexToAddress("0x3bfC61cbA24a835F99a5855e613b239c7110B619")
//	param := "nihao ethermint,this is setGreeting"
//	nonce := hexutil.Uint64(2825)
//	txHash,err := RawTransaction(from,&to,&nonce,"setGreeting",helloContractstring,param)
//	if err!=nil{
//		t.Errorf(err.Error())
//	}
//
//	privstr := "0x59ADEC6D2A86154A082654A9FF0F30523BE8099BAD5442EBA6928E8FE2D25B18"
//	privbyte,err := hexutil.Decode(privstr)
//	if err!=nil{
//		t.Errorf(err.Error())
//	}
//	priv,err := ethcrypto.ToECDSA(privbyte)
//	if err!=nil{
//		t.Errorf(err.Error())
//	}
//	//privKey := ethsecp256k1.PrivKey(ethcrypto.FromECDSA(priv))
//	sign := Sign(txHash,priv)
//	txId,err := ec.doSendRawTransaction(from,&to,&nonce,"setGreeting",helloContractstring,sign,param)
//	if err!=nil{
//		t.Errorf(err.Error())
//	}
//	fmt.Println(txId)//0x2da293ea5329fcfe17caaf593addef02f0bba6e4e0d137694ae3ff6381e2491e
//}
//0xC9D057F7EE6EEAF68A52158B1EFE126D91923271
//0x39DBFE8AD79D698C26F0FB2B2782E1C7529836589CEBA99F421628B748B18F53
func TestRawTransaction(t *testing.T) {
	rlpStr := "6aa559c9f139da67d7307778fe7d3b716ce976634d29fdd609dbd6f9bb593661"
	hash,err := hex.DecodeString(rlpStr)
	if err!=nil{
		t.Errorf(err.Error())
	}
		//privstr := "0x59ADEC6D2A86154A082654A9FF0F30523BE8099BAD5442EBA6928E8FE2D25B18"  //mykey
	privstr := "0x39DBFE8AD79D698C26F0FB2B2782E1C7529836589CEBA99F421628B748B18F53" //mykey1
		privbyte,err := hexutil.Decode(privstr)
		if err!=nil{
			t.Errorf(err.Error())
		}
		priv,err := ethcrypto.ToECDSA(privbyte)
		if err!=nil{
			t.Errorf(err.Error())
		}
	sign := Sign(hash,priv)
	signstr := hex.EncodeToString(sign)
	fmt.Println(signstr)
}

func Sign(txHash []byte, priv *ecdsa.PrivateKey) []byte {
	//txHash := msg.RLPSignBytes(chainID)
	sig, err := ethcrypto.Sign(txHash[:], priv)
	if err != nil {
		panic(err.Error())
	}
	return sig
}
