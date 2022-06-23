package servicer

import "C"
import (
	"github.com/gateway-server-go/models"
	"github.com/labstack/echo/v4"
)

func AddAccount(c echo.Context)error{
	acc := &Account{}
	if err := c.Bind(acc);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := acc.basicValidate(); err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	macc := models.Account{
		Address: acc.Address,
		Name: acc.Name,
	}
	if _,err := macc.Insert(); err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: "add new account "+acc.Name+" successful "})
}

func GetAccount(c echo.Context)error{
	name := c.QueryParam("name")
	address := c.QueryParam("address")
	if name == "" && address == ""{
		return c.JSON(400, restResponse{400, "name or address cannot be empty", ""})
	}
	var acc *models.Account
	var err error
	if name != ""{
		acc,err = models.GetAccountByName(name)
	}else if address != ""{
		acc,err = models.GetAccountByAddress(address)
	}
	if err!= nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: acc})
}

func AddContract(c echo.Context)error{
	con := &Contract{}
	if err := c.Bind(con);err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	if err := con.basicValidate(); err != nil {
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	mcon := models.Contract{
		ABI: con.ABI,
		Address: con.Address,
		Name: con.Name,
	}
	if _,err := mcon.Insert(); err != nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: "add new contract "+con.Name+" successful "})
}

func GetContract(c echo.Context)error{
	name := c.QueryParam("name")
	address := c.QueryParam("address")
	if name == "" && address == ""{
		return c.JSON(400, restResponse{400, "name or address cannot be empty", ""})
	}
	var con *models.Contract
	var err error
	if name != ""{
		con,err = models.GetContractByName(name)
	}else if address != ""{
		con,err = models.GetContractByAddr(address)
	}
	if err!= nil{
		return c.JSON(400, restResponse{400, err.Error(), ""})
	}
	return c.JSON(200, restResponse{Code: 200, Err: "", Data: con})
}