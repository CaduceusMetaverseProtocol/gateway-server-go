package servicer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReqTx(t *testing.T){
	parameters := []interface{}{"nihao, for postman"}
	req := reqTx{
		"mykey",
		"mykey2",
		"setGreeting",
		parameters,
		0,
	}
	b,err := json.Marshal(req)
	if err != nil{
		t.Error(err)
	}
	fmt.Println(string(b))
}
