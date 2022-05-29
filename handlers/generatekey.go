package handlers

import (
	"crypto/ecdsa"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo"
)

// GenerateKeyResp ...
type GenerateKeyResp struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
}

// GenerateKey ...
func GenerateKey(c echo.Context) error {
	// private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		c.Logger().Warnf("GenerateKey failed, error: %v", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	c.Logger().Debugf("privateKeyHex: %s", privateKeyHex)

	// public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.Logger().Error("casting public key to ECDSA failed")
		return c.JSON(http.StatusInternalServerError, nil)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes)
	c.Logger().Debugf("publicKeyHex: %s", publicKeyHex)

	// address
	addressHex := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	c.Logger().Debugf("addressHex: %s", addressHex)

	// hash := sha3.NewKeccak256()	// github.com/ethereum/go-ethereum/crypto/sha3 gone
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// t := hash.Sum(nil)[12:]
	// hashHex := hexutil.Encode(t)
	// c.Logger().Debugf("hashHex: %+v", hashHex)

	// resp
	resp := GenerateKeyResp{
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
		Address:    addressHex,
	}
	return c.JSON(http.StatusOK, resp)
}
