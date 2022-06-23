package servicer

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-server-go/controller"
	"github.com/labstack/echo/v4"
	"math/big"
	"strings"
)

func ReadErc20(c echo.Context)error{
	name := c.QueryParam("name")
	contract_name := c.QueryParam("contract_name")
	paths := strings.Split(c.Request().URL.Path,"/")
	if strings.Trim(name," ") == "" || strings.Trim(contract_name," ") == ""{
		return c.JSON(400, restResponse{400, "contract name cannot be empty", ""})
	}
	parameters := make([]interface{},0)
	message := paths[len(paths)-1]
	switch message {
	    case "balanceOf":
			owner := c.QueryParam("owner")
			if strings.Trim(owner," ") == "" {
				return c.JSON(400, restResponse{400, "owner name cannot be empty", ""})
			}
			ownerAdd := common.HexToAddress(owner)
			parameters = append(parameters,ownerAdd)
	    case "allowance":
			owner := c.QueryParam("owner")
			spender := c.QueryParam("spender")
			if strings.Trim(owner," ") == ""|| strings.Trim(spender," ") == ""{
				return c.JSON(400, restResponse{400, "owner or spender cannot be empty", ""})
			}
			ownerAdd := common.HexToAddress(owner)
			spenderAdd := common.HexToAddress(spender)
			parameters = append(parameters,ownerAdd)
			parameters = append(parameters,spenderAdd)
	}
	res,err := controller.NewEmController().Call(name,contract_name,message,true,parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: res})
}

func WriteErc20Pre(c echo.Context)error{
	tx,message,parameters,err  := writeErc20(c)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	res,err := controller.RawTransaction(tx.Name,tx.ContractName,message,0,parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	hexRes := hex.EncodeToString(res)
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: hexRes})
}

func WriteErc20Do(c echo.Context)error{
	tx,message,parameters,err  := writeErc20(c)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	sign,err := hex.DecodeString(tx.Sign)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if len(sign) == 0 {
		return c.JSON(400, restResponse{400, "tx sign cannot be empty", ""})
	}
	res,err := controller.NewEmController().SendRawTransaction(tx.Name,tx.ContractName,message,0,sign,parameters...)
	if err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: string(res)})
}

func writeErc20(c echo.Context)(*reqWriteErc20,string,[]interface{},error){
	tx := &reqWriteErc20{}
	if err := c.Bind(tx);err != nil {
		return nil,"",nil,err
	}
	if err := tx.basicValidate(); err != nil {
		return nil,"",nil,err
	}
	paths := strings.Split(c.Request().URL.Path,"/")
	parameters := make([]interface{},0)
	message := paths[len(paths)-2]
	value := big.NewInt(tx.Value)
	switch message {
	case "transfer":
		if strings.Trim(tx.To," ") == "" {
			return nil,"",nil,errors.New("to address cannot be empty")
		}
		to := common.HexToAddress(tx.To)
		parameters = append(parameters,to)
		parameters = append(parameters,value)
	case "approve":
		if strings.Trim(tx.Spender," ") == "" {
			return nil,"",nil,errors.New("spender address cannot be empty")
		}
		spender := common.HexToAddress(tx.Spender)
		parameters = append(parameters,spender)
		parameters = append(parameters,value)
	case "transferFrom":
		if strings.Trim(tx.From," ") == "" ||strings.Trim(tx.To," ") == ""{
			return nil,"",nil,errors.New("from or to address cannot be empty")
		}
		to := common.HexToAddress(tx.To)
		from := common.HexToAddress(tx.From)
		parameters = append(parameters,from)
		parameters = append(parameters,to)
		parameters = append(parameters,value)
	}
	return tx,message,parameters,nil
}



