package model

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"testing"
)

func TestCheckIsExistModeAdmin(t *testing.T) {


	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a73f8de0d9974dd7a35c5d241e24e853")
	if err != nil {
		fmt.Println("11")
		fmt.Println(err.Error())
		return
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Println(balance)


	//tokenAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") //usDT
	//instance, err := token.NewToken(tokenAddress, client)
	//fmt.Println(instance)
	//if err != nil {
	//	fmt.Println("116")
	//	fmt.Println(err.Error())
	//	return
	//}
	//address := common.HexToAddress("TJigzvpzYkoSS3JM3HwZtfJ1zu5vUQQQQQ")
	//bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	//
	//
	//fmt.Println(bal)
	//
	//if err != nil {
	//	fmt.Println("--------------")
	//	fmt.Println(err.Error())
	//	return
	//}
	//usd := ToDecimal(bal.String(), 6)


	//fmt.Println(usd)
}


func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}