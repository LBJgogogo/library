package test

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"go_backend/test/project"
	"log"
	time2 "time"
)

var session *test.ElectronicSealSession //实例化对象

func StartBlockchain() {
	configs, err := conf.ParseConfigFile("test/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	config := &configs[0]
	clientInstance, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	// 初始化合约实例
	contractAddress := common.HexToAddress("0x26a6dc6c3c6671fd8a3b2fb55bf7dca091e82be5") // 正确的合约地址
	instance, err := test.NewElectronicSeal(contractAddress, clientInstance)
	if err != nil {
		log.Fatal(err)
	}
	session = &test.ElectronicSealSession{
		Contract:     instance,
		CallOpts:     *clientInstance.GetCallOpts(),
		TransactOpts: *clientInstance.GetTransactOpts(),
	}
	log.Println("WeBase平台连接成功")
}

func AddSealAccount(account, name, cardId string) {
	//1.//currentTime := time2.Now()
	//currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	//2.//strconv.FormatInt(time2.Now().Unix(), 10)
	currentTime := time2.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	accounts := common.HexToAddress(account)
	tx, _, err := session.AddSealAccount(accounts, name, cardId, currentTimeString)
	if err != nil {
		log.Fatal("用户信息上链失败", err)
	}
	fmt.Println("信息上链成功,交易号为:", tx.Hash().Hex())
}
func SealSignature(account, hash, code string) {
	accounts := common.HexToAddress(account)
	currentTime := time2.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	tx, _, err := session.SealSignature(accounts, hash, code, currentTimeString)
	if err != nil {
		log.Fatal("用户信息上链失败", err)
	}
	fmt.Println("信息上链成功,交易号为:", tx.Hash().Hex())
}

func GetAllHash(account string) []string {
	accounts := common.HexToAddress(account)
	hx, err := session.QueryAllHash(accounts)
	if err != nil {
		log.Fatal("链上哈希信息查询失败", err)
	}
	return hx
}
func GetSignature(account, hx string) (string, string) {
	accounts := common.HexToAddress(account)
	code, time, err := session.QuerySignature(accounts, hx)
	if err != nil {
		log.Fatal("链上签章记录查询失败", err)
	}
	return code, time
}
