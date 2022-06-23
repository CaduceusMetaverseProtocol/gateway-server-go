package controller

import (
	"bytes"
	"fmt"
	rpctypes "github.com/cosmos/ethermint/rpc/types"
	evmtypes "github.com/cosmos/ethermint/x/evm/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gateway-server-go/common"
	"github.com/gateway-server-go/ethclient"
	"github.com/gateway-server-go/models"
	"math/big"
)

type EmController struct {
	emClient  *ethclient.Client
}


func NewEmController()*EmController{
	return &EmController{
		emClient: common.NewEthClient(),
	}
}

func RawTransaction(fromName ,contractName,message string,value int64,parameters ...interface{} )([]byte,error){
	account,err := models.GetAccountByName(fromName)
	if err !=nil{
		return nil, err
	}
	from := ethcommon.HexToAddress(account.Address)
	var abi string
	var toAdd string
	if message != ""{
		contract,err := models.GetContractByName(contractName)
		if err !=nil{
			return nil, err
		}
		abi = contract.ABI
		toAdd = contract.Address
	}else{
		acc,err := models.GetAccountByName(contractName)
		if err !=nil{
			return nil, err
		}
		toAdd = acc.Address
	}
	to := ethcommon.HexToAddress(toAdd)
	vb := hexutil.Big(*(big.NewInt(value)))
	nonce,err := NewEmController().GetNonceAt(fromName,nil)
	if err !=nil{
		return nil, err
	}
	nonceN := hexutil.Uint64(nonce)
	arg,err := createSendTxArgs(from,&to,&nonceN,&vb,message,abi,parameters...)
	if err!=nil{
		return nil,err
	}
	txHash ,err := RlpTxHash(*arg)
	if err!=nil{
		return nil,err
	}
	return txHash,nil
}

func RlpTxHash(args rpctypes.SendTxArgs)([]byte,error){
	if args.Nonce == nil {
		return nil,fmt.Errorf("Nonce can not is nil ")
	}
	// Assemble transaction from fields
	tx, err := generateFromArgs(args)
	if err != nil {
		fmt.Println("failed to generate tx", "error", err)
		return nil, err
	}
	if err := tx.ValidateBasic(); err != nil {
		fmt.Println("tx failed basic validation", "error", err)
		return nil, err
	}
	// Sign transaction
	rlpHash := tx.RLPSignBytes(ChainIDEpoch)
	return rlpHash.Bytes(),nil
}

// generateFromArgs populates tx message with args (used in RPC API)
func  generateFromArgs(args rpctypes.SendTxArgs) (*evmtypes.MsgEthereumTx, error) {
	var (
		nonce, gasLimit uint64
		err             error
	)

	amount := (*big.Int)(args.Value)
	gasPrice := (*big.Int)(args.GasPrice)

	if args.GasPrice == nil {
		// Set default gas price
		// TODO: Change to min gas price from context once available through server/daemon
		gasPrice = Gasprice
	}

	if args.Nonce == nil {
		// get the nonce from the account retriever and the pending transactions
		return nil, fmt.Errorf("Nonce can not is nil ")
	} else {
		nonce = (uint64)(*args.Nonce)
	}

	if err != nil {
		return nil, err
	}

	if args.Data != nil && args.Input != nil && !bytes.Equal(*args.Data, *args.Input) {
		return nil, fmt.Errorf("both 'data' and 'input' are set and not equal. Please use 'input' to pass transaction call data")
	}

	// Sets input to either Input or Data, if both are set and not equal error above returns
	var input []byte
	if args.Input != nil {
		input = *args.Input
	} else if args.Data != nil {
		input = *args.Data
	}

	if args.To == nil && len(input) == 0 {
		// Contract creation
		return nil, fmt.Errorf("contract creation without any data provided")
	}

	if args.Gas == nil {
		return nil, fmt.Errorf("Gas can not is zero ")
	} else {
		gasLimit = (uint64)(*args.Gas)
	}
	msg := evmtypes.NewMsgEthereumTx(nonce, args.To, amount, gasLimit, gasPrice, input)

	return &msg, nil
}

func addSign(msg *evmtypes.MsgEthereumTx,chainID *big.Int, sig []byte)error{
	if len(sig) != 65 {
		return fmt.Errorf("wrong size for signature: got %d, want 65", len(sig))
	}

	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])

	var v *big.Int

	if chainID.Sign() == 0 {
		v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	} else {
		v = big.NewInt(int64(sig[64] + 35))
		chainIDMul := new(big.Int).Mul(chainID, big.NewInt(2))

		v.Add(v, chainIDMul)
	}

	msg.Data.V = v
	msg.Data.R = r
	msg.Data.S = s
	return nil
}