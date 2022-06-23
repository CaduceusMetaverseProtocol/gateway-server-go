package controller

import (
	"context"
	"errors"
	"fmt"
	rpctypes "github.com/cosmos/ethermint/rpc/types"
	"github.com/ethereum/go-ethereum"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	models "github.com/gateway-server-go/models"
	"math/big"
	"strings"
)

var (
	Gas = uint64(1000000)
	Gasprice = big.NewInt(100)
	ChainIDEpoch = big.NewInt(2)
)

func (ec *EmController)Call(fromName ,contractName,message string,isView bool,parameters ...interface{}  )(interface{}, error){
	//step1 get contractabi string and contract address based on the contractName
	var from common.Address
	if fromName != ""{
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return nil, err
	}
	from = common.HexToAddress(account.Address)
	}
	contract,err := models.GetContractByName(contractName)
	if err !=nil{
		return nil, err
	}
	to := common.HexToAddress(contract.Address)
	res,err := ec.doCall(from,to,message,contract.ABI,isView,parameters...)
	return unpack(message,contract.ABI,res)
}

func (ec *EmController)doCall(from ,to common.Address,message,abistr string,isView bool,parameters ...interface{} )([]byte, error){
	//step1 get contractabi string and contract address based on the contractName
	msg,err := createMsg(from,to,message,abistr,isView,parameters...)
	if err!= nil{
		return nil, err
	}
	return ec.emClient.CallContract(context.Background(),*msg,nil)
}

func (ec *EmController)GetBalance(accountName string, blockNumber *big.Int)(*big.Int, error) {
	account,err := models.GetAccountByName(accountName)
	if err !=nil{
		return nil, err
	}
	address := common.HexToAddress(account.Address)
	balance ,err := ec.emClient.BalanceAt(context.Background(),address , blockNumber)
	if err!= nil{
		return nil, err
	}
	return balance, nil
}

func (ec *EmController)GetCode(contractName string, blockNumber *big.Int)([]byte, error){
	account,err := models.GetContractByName(contractName)
	if err !=nil{
		return nil, err
	}
	address := common.HexToAddress(account.Address)
	return ec.emClient.CodeAt(context.Background(),address , blockNumber )
}
func (ec *EmController)SendBaseRawTransaction(tx string)(string,error) {
	return ec.emClient.EmSendBaseRawTransaction(context.Background(),tx)
}
func (ec *EmController)SendRawTransaction(fromName ,contractName,message string,value int64,sign []byte,parameters ...interface{} )(string,error){
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return "", err
	}
	from := common.HexToAddress(account.Address)
	var abi string
	var toAdd string
	if message != ""{
		contract,err := models.GetContractByName(contractName)
		if err !=nil{
			return "", err
		}
		abi = contract.ABI
		toAdd = contract.Address
	}else{
		acc,err := models.GetAccountByName(contractName)
		if err !=nil{
			return "", err
		}
		toAdd = acc.Address
	}
	to := common.HexToAddress(toAdd)
	vb := hexutil.Big(*(big.NewInt(value)))
	nonce,err := ec.getNonceAt(from,nil)
	if err != nil{
		return "", err
	}
	nonceN := hexutil.Uint64(nonce)
	return ec.doSendRawTransaction(from,&to,&nonceN,&vb,message,abi,sign,parameters...)
}

func (ec *EmController)doSendRawTransaction(from common.Address,to *common.Address,nonce *hexutil.Uint64,value *hexutil.Big,message,abistr string,sign []byte,parameters ...interface{} )(string,error){
	arg,err := createSendTxArgs(from,to,nonce,value,message,abistr,parameters...)
	if err!=nil{
		return "",err
	}
	msg,err := generateFromArgs(*arg)
	if err!=nil{
		return "",err
	}
	err = addSign(msg,ChainIDEpoch,sign)
	if err!=nil{
		return "",err
	}
	res,err := ec.emClient.EmSendRawTransaction(context.Background(),msg)
	return res,err
}


func (ec *EmController)SendTransaction(fromName ,contractName,message string,value int64,parameters ...interface{} )(string,error){
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return "", err
	}
	from := common.HexToAddress(account.Address)
	var abi string
	var toAdd string
	if message != ""{
		contract,err := models.GetContractByName(contractName)
		if err !=nil{
			return "", err
		}
		abi = contract.ABI
		toAdd = contract.Address
	}else{
		acc,err := models.GetAccountByName(contractName)
		if err !=nil{
			return "", err
		}
		toAdd = acc.Address
	}
	to := common.HexToAddress(toAdd)
	vb := hexutil.Big(*(big.NewInt(value)))
	return ec.doSendTransaction(from,&to,&vb,message,abi,parameters)
}

func (ec *EmController)doSendTransaction(from  common.Address,to *common.Address,value *hexutil.Big,message,abistr string,parameters ...interface{} )(string,error){
	arg,err := createSendTxArgs(from,to,nil,value,message,abistr,parameters...)
	if err!=nil{
		return "",err
	}
	res,err := ec.emClient.EmSendTransaction(context.Background(),arg)
	return res,err
}

func (ec *EmController)EstimateGas(fromName ,contractName,message string,parameters ...interface{})(uint64, error){
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return 0, err
	}
	from := common.HexToAddress(account.Address)
	contract,err := models.GetContractByName(contractName)
	if err !=nil{
		return 0, err
	}
	to := common.HexToAddress(contract.Address)
	return ec.doEstimate(from,to,message,contract.ABI,parameters...)
}

func (ec *EmController)doEstimate(from ,to common.Address,message,abistr string,parameters ...interface{} )(uint64, error){
	msg,err := createMsg(from,to,message,abistr,false,parameters...)
	if err!= nil{
		return 0, err
	}
	return ec.emClient.EstimateGas(context.Background(),*msg)
}

func (ec *EmController)GetBlockByNumber(blockNumber *big.Int)(map[string]interface{}, error){
	block,err := ec.emClient.EmBlockByNumber(context.Background(),blockNumber)
	return block,err
}

func (ec *EmController)GetBlockByHash(hexHash string)(map[string]interface{}, error){
	hash := common.HexToHash(hexHash)
	if len(hash) == 0{
		return nil,fmt.Errorf("hex hash is error，please check it! ")
	}
	block,err :=ec.emClient.EmBlockByHash(context.Background(),hash)
	return block,err
}

func (ec *EmController)GetBlockNumber()(uint64, error){
	return ec.emClient.BlockNumber(context.Background())

}

func (ec *EmController)GetTransactionReceipt(hexHash string)(interface{}, error){
	hash := common.HexToHash(hexHash)
	if len(hash) == 0{
		return nil,fmt.Errorf("hex hash is error，please check it! ")
	}
	receipt,err := ec.emClient.TransactionReceipt(context.Background(),hash)
	//fmt.Println(receipt,err)
	return receipt, err
}

func (ec *EmController)GetNonceAt(fromName string,blockNumber *big.Int)(uint64,error){
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return 0, err
	}
	address:= common.HexToAddress(account.Address)
	return ec.emClient.NonceAt(context.Background(),address,blockNumber)
}

func (ec *EmController)getNonceAt(account common.Address,blockNumber *big.Int)(uint64,error){
	return ec.emClient.NonceAt(context.Background(),account,blockNumber)
}

func (ec *EmController)GetTransactionByHash(hexHash string)(interface{},bool, error){
	hash := common.HexToHash(hexHash)
	if len(hash) == 0{
		return nil,false,fmt.Errorf("hex hash is error，please check it! ")
	}
	 tx,isPending,err := ec.emClient.TransactionByHash(context.Background(),hash)
	//fmt.Println(tx,ispending,err)
	return tx,isPending, err
}


func createMsg(from ,to common.Address,message,abistr string,isView bool,parames ...interface{} )(*ethereum.CallMsg,error) {
	//step1 get contractabi string and contract address based on the contractName
	abi, err := ethabi.JSON(strings.NewReader(abistr))
	if err != nil {
		return nil, err
	}
	if isView{
		method, exist := abi.Methods[message]
		if !exist {
			return nil, fmt.Errorf("method '%s' not found", message)
		}
		if method.StateMutability != "view"{
			return nil, errors.New(fmt.Sprintf("function %s cannot be read only",message))
		}
	}

	var ap []byte
	if len(parames) == 0 {
		ap, err = abi.Pack(message)
	} else {
		ap, err = abi.Pack(message, parames...)
	}
	if err != nil {
		return nil, err
	}
	//fmt.Println(hex.EncodeToString(ap))
	//create callMsg
	return &ethereum.CallMsg{
		From:     from,
		To:       &to,
		Data:     ap,
		Gas:      Gas,
		GasPrice: Gasprice,
	},nil
}


func createSendTxArgs(from common.Address,to *common.Address,nonce *hexutil.Uint64,value *hexutil.Big,message,abistr string,parames ...interface{} )(*rpctypes.SendTxArgs,error) {
	gas := hexutil.Uint64(Gas)
	args := &rpctypes.SendTxArgs{
		From: from,
		To: to,
		Gas: &gas,
		GasPrice: (*hexutil.Big)(Gasprice),
		Nonce: nonce,
	}
	//step1 get contractabi string and contract address based on the contractName
	var ap []byte
	if abistr!="" {
		abi, err := ethabi.JSON(strings.NewReader(abistr))
		if err != nil {
			return nil, err
		}
		if parames == nil {
			ap, err = abi.Pack(message)
		} else {
			ap, err = abi.Pack(message, parames...)
		}
		if err != nil {
			return nil, err
		}
		app := hexutil.Bytes(ap)
		args.Input = &app
	}else{
		args.Value = value
	}

//	fmt.Println(hex.EncodeToString(ap))
	//create callMsg
	return args,nil
}

func unpack(message,abiStr string,data []byte ) ([]interface{}, error){
	abi, err := ethabi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil,err
	}
	method, _ := abi.Methods[message]
	res,err := method.Outputs.Unpack(data)
	return res,err
}