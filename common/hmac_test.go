package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	testSecretKey = "M2E5MmU4MjItOTE3MS00NGI0LWEyODEtMDllMTAxOWU1OTE1MTM1NzIxOTcwMTI3MjM1MDcyMA=="
	jsonData = `{
  "from":"mykey",
  "to":"testContract",
  "message":"getGreeting"}`

)


func TestHmacSign(t *testing.T) {
	mapParams:= make(map[string]string)
	err := json.Unmarshal([]byte(jsonData),&mapParams)
	if err!=nil{
		t.Error(err)
	}
	method:= "POST"
	hostname := "127.0.0.1:8000"
	path := "/v1/eth/base/getEstimateGas"
	secretKey :=  testSecretKey
	nowTs := time.Now().Unix()
	mapParams["timestamp"] = strconv.FormatInt(nowTs,10)
	fmt.Println(mapParams)
	sign := HmacSign(mapParams,method,hostname,path,secretKey)
	fmt.Println(sign)
}