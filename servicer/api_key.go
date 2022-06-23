package servicer

import (
    "encoding/base64"
    "time"

    "github.com/gateway-server-go/common"
    "github.com/gateway-server-go/config"
    "github.com/gateway-server-go/models"

    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
)

func CreateApiKey(c echo.Context) error {
    now := time.Now()
    data := c.RealIP()
    req := &reqCreateApiKey{}
    if err := c.Bind(req); err != nil {
        return c.JSON(400, restResponse{400, err.Error(), data})
    }
    //if err := c.Validate(req); err != nil {
    //    return c.JSON(400, restResponse{400, err.Error(), ""})
    //}

    // validate request body
    if err := req.basicValidate(); err != nil {
        return c.JSON(400, restResponse{400, err.Error(), ""})
    }

    var (
        apikey = &models.ApiKey{
            Memo:      req.Memo,
            AccessIPs: req.AccessIPs,
            AccessKey: common.GenerateAccessKey(),
            CreatedAt: now.Unix(),
        }
        plaintextSk = common.GenerateSecretKey()
        err         error
    )

    plainbytesSk, err := base64.StdEncoding.DecodeString(plaintextSk)
    if err != nil {
        return c.JSON(500, restResponse{500, err.Error(), ""})
    }
    cipherbytesSk, err := common.Encrypter(plainbytesSk, config.AesSecretKey)
    if err != nil {
        log.Error(err)
        return c.JSON(500, restResponse{500, err.Error(), ""})
    }
    apikey.SecretKey = base64.StdEncoding.EncodeToString(cipherbytesSk)

    _, err = apikey.Insert()
    if err != nil {
        log.Error(err)
        return c.JSON(500, restResponse{500, err.Error(), ""})
    }
    apikey.SecretKey = plaintextSk

    return c.JSON(200, restResponse{Code: 200, Err: "", Data: apikey})
}
