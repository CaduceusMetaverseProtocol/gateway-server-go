package servicer

import (
	"encoding/hex"
	"github.com/gateway-server-go/controller"
	"github.com/labstack/echo/v4"
	"math/big"
	"strconv"
)

func GetBalance(c echo.Context)error{
	name := c.QueryParam("name")
	number := c.QueryParam("block_number")
	var blockNumber  *big.Int
	if number != ""{
		n ,err := strconv.ParseInt(number,10,64)
		if err!=nil{
			return c.JSON(400, restResponse{400, "block_number must is int64 string", ""})
		}
		blockNumber = big.NewInt(n)
	}
	if name==""{
		return c.JSON(400, restResponse{400, "name cannot be empty", ""})
	}
	res,err := controller.NewEmController().GetBalance(name,blockNumber)
	if err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res.String()})
}

func GetBlockByNumber(c echo.Context)error{
	height := c.QueryParam("block_height")
	var blockNumber  *big.Int
	if height != ""{
		n ,err := strconv.ParseInt(height,10,64)
		if err!=nil{
			return c.JSON(400, restResponse{400, "block_number must is int64 string", ""})
		}
		blockNumber = big.NewInt(n)
	}
	res,err := controller.NewEmController().GetBlockByNumber(blockNumber)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func GetBlockByHash(c echo.Context)error{
	hash := c.QueryParam("block_hash")
	if hash==""{
		return c.JSON(400, restResponse{400, "block_hash cannot be empty", ""})
	}
	res,err := controller.NewEmController().GetBlockByHash(hash)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func GetBlockNumber(c echo.Context)error{
	res,err := controller.NewEmController().GetBlockNumber()
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func GetTransactionReceipt(c echo.Context)error{
	name := c.QueryParam("tx_id")
	if name==""{
		return c.JSON(400, restResponse{400, "tx_id cannot be empty", ""})
	}
	res,err := controller.NewEmController().GetTransactionReceipt(name)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func GetTransactionByHash(c echo.Context)error{
	name := c.QueryParam("tx_id")
	if name==""{
		return c.JSON(400, restResponse{400, "tx_id cannot be empty", ""})
	}
	res,isPending,err := controller.NewEmController().GetTransactionByHash(name)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	resp := respTx{
		Tx: res,
		Ispending: isPending,
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: resp})
}

func GetNonceAt(c echo.Context)error{
	name := c.QueryParam("name")
	number := c.QueryParam("block_number")
	var blockNumber  *big.Int
	if number != ""{
		n ,err := strconv.ParseInt(number,10,64)
		if err!=nil{
			return c.JSON(400, restResponse{400, "block_number must is int64 string", ""})
		}
		blockNumber = big.NewInt(n)
	}
	if name==""{
		return c.JSON(400, restResponse{400, "name cannot be empty", ""})
	}
	res,err := controller.NewEmController().GetNonceAt(name,blockNumber)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}


func GetCode(c echo.Context)error{
	name := c.QueryParam("name")
	if name==""{
		return c.JSON(400, restResponse{400, "name cannot be empty", ""})
	}
	number := c.QueryParam("block_number")
	var blockNumber  *big.Int
	if number != ""{
		n ,err := strconv.ParseInt(number,10,64)
		if err!=nil{
			return c.JSON(400, restResponse{400, "block_number must is int64 string", ""})
		}
		blockNumber = big.NewInt(n)
	}
	res,err := controller.NewEmController().GetCode(name,blockNumber)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func Call(c echo.Context)error{
	tx := &reqTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}

	res,err := controller.NewEmController().Call(tx.FromName,tx.ToName,tx.Message,false,tx.Parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func ViewCall(c echo.Context)error{
	tx := &reqTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}

	res,err := controller.NewEmController().Call(tx.FromName,tx.ToName,tx.Message,true,tx.Parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}
func GetEstimateGas(c echo.Context)error{
	tx := &reqTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	res,err := controller.NewEmController().EstimateGas(tx.FromName,tx.ToName,tx.Message,tx.Parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func RawTransaction(c echo.Context)error{
	tx := &reqTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}

	res,err := controller.RawTransaction(tx.FromName,tx.ToName,tx.Message,tx.Value,tx.Parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	hexres := hex.EncodeToString(res)
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: hexres})
}
func BaseRawTransaction(c echo.Context)error{
	tx := &reqBaseTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}

	res,err := controller.RawTransaction(tx.FromName,tx.ToName,"",tx.Value,nil)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	hexres := hex.EncodeToString(res)
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: hexres})
}

func SendRawTransaction(c echo.Context)error{
	tx := &baseTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	res,err := controller.NewEmController().SendBaseRawTransaction(tx.Tx)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func RawTransfer(c echo.Context)error{
	tx := &reqRawTx{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	sign,err := hex.DecodeString(tx.Sign)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	res,err := controller.NewEmController().SendRawTransaction(tx.FromName,tx.ToName,tx.Message,tx.Value,sign,tx.Parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}
func BaseRawTransfer(c echo.Context)error{
	tx := &reqBaseTransfer{}
	if err := c.Bind(tx);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := tx.basicValidate();err!=nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	sign,err := hex.DecodeString(tx.Sign)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	res,err := controller.NewEmController().SendRawTransaction(tx.FromName,tx.ToName,"",tx.Value,sign,nil)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}
//TODO
func SendTransaction(){

}