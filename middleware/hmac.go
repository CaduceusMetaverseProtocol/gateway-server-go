package middleware

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/gateway-server-go/common"
    "github.com/gateway-server-go/config"
    "github.com/gateway-server-go/models"

    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
)

func VerifyHmacSign(accessKey string, c echo.Context) (bool, error) {
    if !config.EnableHmac(){
        return true,nil
    }
    nowTs := time.Now().Unix()
    req := c.Request()

    // parse params
    reqParams := make(map[string]interface{})
    if req.ContentLength == 0 { // Method: GET/DELETE
        if req.Method == http.MethodGet || req.Method == http.MethodDelete {
            for k, v := range c.QueryParams() {
                reqParams[k] = v[0]
            }
        } else {
            return false, nil
        }
    } else { // Method: POST
        ctype := req.Header.Get("Content-Type")
        switch {
        case strings.HasPrefix(ctype, "application/json"):
            // copy body stream
            var buf bytes.Buffer
            teeBody := io.TeeReader(req.Body, &buf)
            req.Body = ioutil.NopCloser(&buf)

            decoder := json.NewDecoder(teeBody)
            err := decoder.Decode(&reqParams)
            if err != nil {
                log.Info(err)
                return false, nil
            }
        default:
            return false, nil
        }
    }
    mapParams := make(map[string]string)
    for k, v := range reqParams {
        mapParams[k] = fmt.Sprintf("%v", v)
    }
    log.Info(mapParams)

    // check timestamp
    reqTs, err := strconv.ParseInt(mapParams["timestamp"], 10, 64)
    if err != nil {
        log.Info(err)
        return false, nil
    }
    delay := nowTs - reqTs
    if delay > 500 || delay < -5 {
        log.Infof("timestamp is invalid client: %d server: %d", reqTs, nowTs)
        return false, nil
    }

    // query db
    apiKey, err := models.GetApiKeyByAk(accessKey)
    if err != nil {
        log.Info(err)
        return false, nil
    }

    // verify access ips
    if !common.StringInSlice(c.RealIP(), strings.Split(apiKey.AccessIPs, ",")) {
        log.Infof("ip %s not in white list %s", c.RealIP(), apiKey.AccessIPs)
        return false, nil
    }

    // verify signature
    cipherbytesSk, err := base64.StdEncoding.DecodeString(apiKey.SecretKey)
    if err != nil {
        log.Info(err)
        return false, nil
    }
    plainbytesSk, err := common.Decrypter(cipherbytesSk, config.AesSecretKey)
    if err != nil {
        log.Info(err)
        return false, nil
    }
    base64Sk := base64.StdEncoding.EncodeToString(plainbytesSk)
    reqSignature := mapParams["signature"]
    delete(mapParams, "signature")
    signature := common.HmacSign(mapParams, req.Method, req.Host, req.URL.Path, base64Sk)
    if signature == reqSignature {
        return true, nil
    }
    log.Infof("req sig: %s server sig: %s", reqSignature, signature)
    return false, nil
}
