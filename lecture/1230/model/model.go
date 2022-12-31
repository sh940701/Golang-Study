package model

type MnemonicResponse struct {
	Mnemonic string
}

type WalletCreateRequest struct {
	Mnemonic string `json:"mnemonic" binding:"required"`
}

type WalletResponse struct {
	PrivateKey string `json:"privatekey"`
	Address string `json:"address"`
}

type WalletCreateRequestWithPassword struct {
	Mnemonic string `json:"mnemonic" binding:"required"`
	Password string `json:"password" binding:"required"`
}