package common

import (
	"github.com/gateway-server-go/ethclient"
	"log"
	"sync"
)

var (
	once sync.Once
	Client = &ethclient.Client{}
	EthermintNodeEndpoint = "http://182.92.150.56:8547"
)

func NewEthClient()*ethclient.Client{
	once.Do(initClient)
	return Client
}

func initClient(){
	conn, err := ethclient.Dial(EthermintNodeEndpoint)
	if err != nil {
		log.Fatalf("failed to connect ethereum node %s", err)
	}
	Client = conn
}