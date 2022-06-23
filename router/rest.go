package routes

import (
    "github.com/gateway-server-go/middleware"
    "github.com/gateway-server-go/servicer"
    "github.com/labstack/echo/v4"
    echoMW "github.com/labstack/echo/v4/middleware"
)

func RegisterRestAPI(e *echo.Echo) {
    r := e.Group("/v1")
    echoMW.DefaultKeyAuthConfig.AuthScheme = ""
    r.Use(echoMW.KeyAuthWithConfig(echoMW.KeyAuthConfig{
       KeyLookup: "header:X-ACCESS-KEY",
       Validator: middleware.VerifyHmacSign,
    }))
    r.POST("/eth/base/getEstimateGas", servicer.GetEstimateGas)
    r.POST("/eth/base/call", servicer.Call)
    r.POST("/eth/base/prepare", servicer.BaseRawTransaction)
    r.POST("/eth/base/sendRawTransaction", servicer.SendRawTransaction)
    r.POST("/eth/base/transfer", servicer.BaseRawTransfer)

    //eth contract
    r.POST("/eth/contract/prepare", servicer.RawTransaction)
    r.POST("/eth/contract/do", servicer.RawTransfer)
    r.POST("/eth/contract/view", servicer.Call)

    //erc20 API
    r.GET("/eth/contract/erc20/name", servicer.ReadErc20)
    r.GET("/eth/contract/erc20/symbol", servicer.ReadErc20)
    r.GET("/eth/contract/erc20/decimals", servicer.ReadErc20)
    r.GET("/eth/contract/erc20/totalSupply", servicer.ReadErc20)
    r.GET("/eth/contract/erc20/balanceOf", servicer.ReadErc20)
    r.GET("/eth/contract/erc20/allowance", servicer.ReadErc20)
    //write
    r.POST("/eth/contract/erc20/transfer/prepare", servicer.WriteErc20Pre)
    r.POST("/eth/contract/erc20/approve/prepare", servicer.WriteErc20Pre)
    r.POST("/eth/contract/erc20/transferFrom/prepare", servicer.WriteErc20Pre)
    r.POST("/eth/contract/erc20/transfer/do", servicer.WriteErc20Do)
    r.POST("/eth/contract/erc20/approve/do", servicer.WriteErc20Do)
    r.POST("/eth/contract/erc20/transferFrom/do", servicer.WriteErc20Do)
    //for eth get
    r = e.Group("/v1")
    r.GET("/eth/base/getBalance", servicer.GetBalance)
    r.GET("/eth/base/getBlockByNumber", servicer.GetBlockByNumber)
    r.GET("/eth/base/getBlockByHash", servicer.GetBlockByHash)
    r.GET("/eth/base/getBlockNumber", servicer.GetBlockNumber)
    r.GET("/eth/base/getTransactionReceipt", servicer.GetTransactionReceipt)
    r.GET("/eth/base/getTransactionByHash", servicer.GetTransactionByHash)
    r.GET("/eth/base/getNonceAt", servicer.GetNonceAt)
    r.GET("/eth/base/getCode", servicer.GetCode)


    // for account and contract
    r = e.Group("/v1")
    r.POST("/user/addAccount", servicer.AddAccount)
    r.GET("/user/getAccount", servicer.GetAccount)
    r.POST("/contract/add", servicer.AddContract)
    r.GET("/contract/get", servicer.GetContract)
    r.POST("/apiKey", servicer.CreateApiKey)
}
