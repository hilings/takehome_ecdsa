package handlers

import (
	"myapp/util"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo"
)

// SignReq ...
type SignReq struct {
	Message    string `json:"message"`
	PrivateKey string `json:"privateKey"`
}

// SignResp ...
type SignResp struct {
	SignedMessage string `json:"signedMessage"`
}

// Sign ...
func Sign(c echo.Context) error {
	req := &SignReq{}
	util.LoadRequestBody(c, req)

	// validate
	keyHex := req.PrivateKey
	if len(keyHex) != 64 && len(keyHex) != 66 {
		c.Logger().Warnf("invalid privateKey: %s", keyHex)
		return c.JSON(http.StatusBadRequest, nil)
	} else if len(keyHex) == 66 {
		keyHex = keyHex[2:]
	}

	// prepare privateKey
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		c.Logger().Warnf("HexToECDSA failed, error: %v, privateKey: %s", err, req.PrivateKey)
		return c.JSON(http.StatusBadRequest, nil)
	}

	// prepare message
	msgBytes := []byte(req.Message)
	msgHash := crypto.Keccak256Hash(msgBytes)

	// sign
	signatureBytes, err := crypto.Sign(msgHash.Bytes(), privateKey)
	if err != nil {
		c.Logger().Warnf("Sign failed, error: %v, privateKey: %+v", err, privateKey)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	signatureHex := hexutil.Encode(signatureBytes)
	c.Logger().Debugf("signatureHex: %s", signatureHex)

	resp := &SignResp{
		SignedMessage: signatureHex,
	}
	return c.JSON(http.StatusOK, resp)
}
