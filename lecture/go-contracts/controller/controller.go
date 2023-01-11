package controller

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"

	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

// Controller 타입 선언
type Controller struct {}

func NewCTL() (*Controller, error) {
	r := &Controller{}
	return r, nil
}


func(ct *Controller) GetSymbol(c gin.Context) {

}

func(ct *Controller)GetAmount(c gin.Context) {

}

func(ct *Controller)SendCoin(c gin.Context) {

    client, err := ethclient.Dial("https://api.test.wemix.com")
    if err != nil {
        fmt.Println("client error")
    }

    privateKey, err := crypto.HexToECDSA("c04e38b264d4bf35625090e4764c912718f4c8be1bce5c3c0796365f9732d0b0")
    if err != nil {
        fmt.Println(err)
    }
		
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        fmt.Println("fail convert, publickey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        fmt.Println(err)
    }
		
    value := big.NewInt(1000000000000000000) // 1개
    gasLimit := uint64(21000)
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        fmt.Println(err)
    }
		
    toAddress := common.HexToAddress("0x314613c08Cb38e3d782688e86f61a563D8959574")
	// 트랜잭션 생성
    var data []byte // 전송하는 data 없음
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        fmt.Println(err)
    }
	// 트랜잭션 서명
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        fmt.Println(err)
    }

	// // RLP 인코딩 전 트랜잭션 묶음. 현재는 1개의 트랜잭션
    // ts := types.Transactions{signedTx}
	// 	// RLP 인코딩
    // rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
    // rawTxHex := hex.EncodeToString(rawTxBytes)
    // rTxBytes, err := hex.DecodeString(rawTxHex)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }

	// // RLP 디코딩
    // rlp.DecodeBytes(rTxBytes, &tx)
	// // 트랜잭션 전송
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        fmt.Println(err)
    }
    //출력된 tx.hash를 익스플로러에 조회 가능
    fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	c.JSON(200, gin.H{
		"mag" : "success",
		"hash" : tx.Hash().Hex()
	})
}

func(ct *Controller)SendCoinByPk(c gin.Context) {

}

func(ct *Controller)SendToken(c gin.Context) {

}

func(ct *Controller)SendTokenByPk(c gin.Context) {

}