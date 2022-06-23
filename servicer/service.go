package servicer

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-server-go/models"
	"net"
	"strings"
)

type restResponse struct {
	Code int         `json:"code"`
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}
type restResponse2 struct {
	Code int         `json:"code"`
	Err  string      `json:"err"`
	Data string      `json:"data"`
}
type Account struct {
	Address string `json:"address" `
	Name    string `json:"name" `  //userName
}

type Contract struct {
	Address string `json:"address" `
	ABI     string `json:"abi" `
	Name    string `json:"name" `
}

type reqTx struct {
	FromName 	 	 string `json:"from"`
	ToName		 	 string `json:"to"`
	Message      	 string `json:"message"`
	Parameters		 []interface{} `json:"parameters"`
	Value 		 	 int64  `json:"value"`
}

type reqBaseTx struct {
	FromName 	 	 string `json:"from"`
	ToName		 	 string `json:"to"`
	Value 		 	 int64  `json:"value"`
}

type baseTx struct {
	Tx 	 	 string `json:"tx"`
}

type reqRawTx struct {
	FromName 	 	 string `json:"from"`
	ToName		 	 string `json:"to"`
	Message      	 string `json:"message"`
	Parameters		 []interface{} `json:"parameters"`
	Value 		 	 int64  `json:"value"`
	Sign			 string `json:"sign"`
}
type reqBaseTransfer struct {
	FromName 	 	 string `json:"from"`
	ToName		 	 string `json:"to"`
	Value 		 	 int64  `json:"value"`
	Sign			 string `json:"sign"`
}

type respTx struct{
	Tx 		interface{} 	`json:"tx"`
	Ispending bool `json:"ispending"`
}

type reqReadErc20 struct {
	Name    	 string `json:"name"`
	ContractName string `json:"contract_name"`
	Parameters   []interface{} `json:"parameters"`
}

type reqWriteErc20 struct {
	Name         string `json:"name"`
	ContractName string `json:"contract_name"`
	Sign			 string `json:"sign"`
	From 	 	 string `json:"from"`
	To		 	 string `json:"to"`
	Value 		 	 int64  `json:"value"`
	Spender			 string `json:"spender"`
}

type reqCreateApiKey struct {
	Memo      string `json:"memo"`
	AccessIPs string `json:"access_ips" validate:"required"`
}

func (r *reqCreateApiKey) basicValidate() error {
	// validate memo length
	if len(r.Memo) > 100 {
		return errors.New("\"memo\" up to 100 characters")
	}

	// validate access IPs
	if r.AccessIPs == "" {
		return nil
	}
	ips := strings.Split(r.AccessIPs, ",")

	if len(ips) > 10 {
		return errors.New("up to 10 IPs")
	} else {
		for _, ip := range ips {
			if net.ParseIP(ip) == nil {
				return errors.New("invalid IP")
			}
		}
	}

	return nil
}

func (acc *Account)basicValidate()error{

	if acc.Address == "" || acc.Name == ""{
		return errors.New("address or name connot be empty")
	}
	if len(common.HexToAddress(acc.Address))==0{
		return errors.New("address must is hex ")
	}
	res,_ := models.GetAccountByName(acc.Name)
	if res.Name != ""{
		return errors.New("account already exist ")
	}
	return nil
}

func (c *Contract)basicValidate()error{

	if c.Address == "" || c.Name == "" ||c.ABI == ""{
		return errors.New("address or name or abi connot be empty")
	}
	if len(common.HexToAddress(c.Address))==0{
		return errors.New("address must is hex ")
	}
	res,_ := models.GetContractByName(c.Name)
	if res.Name != ""{
		return errors.New("account already exist ")
	}
	return nil
}

func (tx *reqTx)basicValidate()error{
	if tx.FromName == "" || tx.ToName == "" {
		return errors.New("from_name or to_name connot be empty")
	}else if tx.Value == 0{
		if tx.Message == "" {
			return errors.New("message connot be empty")
		}
	}

	return nil
}

func (tx *reqRawTx)basicValidate()error{
	if tx.FromName == "" || tx.ToName == "" || len(tx.Sign) == 0{
		return errors.New("from_name or to_name or sign connot be empty")
	}else if tx.Value == 0{
		if tx.Message == "" {
			return errors.New("message connot be empty")
		}
	}
	return nil
}

func (tx *reqBaseTransfer)basicValidate()error{
	if tx.FromName == "" || tx.ToName == "" || len(tx.Sign) == 0{
		return errors.New("from_name or to_name or sign connot be empty")
	}else if tx.Value == 0{
			return errors.New("value connot be empty")
	}
	return nil
}

func (tx *reqWriteErc20)basicValidate()error{
	if tx.Name == "" || tx.ContractName == "" {
		return errors.New("user_name or contract_name connot be empty")
	}else if !(tx.Value > 0){
		return errors.New("value must be greater than zero")
	}
	return nil
}

func (tx *reqBaseTx)basicValidate()error{
	if tx.FromName == "" || tx.ToName == "" {
		return errors.New("from_name or to_name connot be empty")
	}else if tx.Value == 0{
		return errors.New("value connot be empty")
	}
	return nil
}
func (tx *baseTx)basicValidate()error{
	if !strings.HasPrefix(tx.Tx,"0x"){
		return errors.New("tx must be the hex data from which the '0x' begins ")
	}
	return nil
}
