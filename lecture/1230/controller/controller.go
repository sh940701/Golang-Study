package controller

import (
	"io/ioutil"
	"lecture/1230/model"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func NewMnemonic(c *gin.Context) {
	entropy, err := hdwallet.NewEntropy(256)
	if err != nil {
		panic(err)
	}

	mnemonic, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		panic(err)
	}

	response := model.MnemonicResponse{Mnemonic : mnemonic}

	c.IndentedJSON(200, response)
}

func NewWallet(c *gin.Context) {
	var body model.WalletCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
	}

	mnemonic := body.Mnemonic

	seed, _ := hdwallet.NewSeedFromMnemonic(mnemonic)
	wallet, _ := hdwallet.NewFromSeed(seed)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	account, _ := wallet.Derive(path, false)
	privateKey, _ := wallet.PrivateKeyHex(account)

	address := account.Address.Hex()

	result := model.WalletResponse{
		PrivateKey : privateKey,
		Address : address,
	}

	c.IndentedJSON(200, result)
}


func NewWalletWithKeystore(c *gin.Context) {
	var body model.WalletCreateRequestWithPassword
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
		return
	}
	mnemonic := body.Mnemonic
	password := body.Password

	seed, _ := hdwallet.NewSeedFromMnemonic(mnemonic)
	wallet, _ := hdwallet.NewFromSeed(seed)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	account, _ := wallet.Derive(path, false)
	privateKey, _ := wallet.PrivateKey(account)

	address := account.Address.Hex()

	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("Could not create random uuid: %v", err))
	}

	ks := &keystore.Key {
		Id: id,
		Address: crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	fmt.Println(ks.Address)

	keyjson, err := keystore.EncryptKey(ks, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		log.Fatalf(err.Error())
	}

	keystoreName := strings.Join([]string{address, "json"}, ".")
	keystoreFile := strings.Join([]string{"./tmp", keystoreName}, "/")
	if err := ioutil.WriteFile(keystoreFile, keyjson, 0700); err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"msg" : "ok"})
}