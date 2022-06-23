package models

import (
	"fmt"
	"testing"
)

var helloContractstring = `[{"constant":false,"inputs":[{"name":"_greeting","type":"string"}],"name":"setGreeting","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"getGreeting","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_greeting","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`

func TestInitDB(t *testing.T) {
	InitDB()
	account := &Account{
		Name: "mykey",
		Address: "0x9321A2F1DC2B417B7165870A27873D2A7813A616",
	}
	res,err := account.Insert()
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(res)
	acc,err := GetAccountByName("mykey1")
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(acc)
}

func TestContract_Insert(t *testing.T) {
	InitDB()
	contract := &Contract{
		Address: "0x3bfC61cbA24a835F99a5855e613b239c7110B619",
		Name: "testContract",
		ABI: helloContractstring,
	}
	res,err := contract.Insert()
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(res)
	con,err := GetContractByName("testContract")
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(con)
}
