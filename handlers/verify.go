package handlers

import (
	"myapp/common"
	"myapp/util"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo"
)

// VerifyRequest ...
type VerifyRequest struct {
	Address       string `json:"address"`
	SignedMessage string `json:"signedMessage"`
}

// VerifyResponse ...
type VerifyResponse struct {
	Verified bool `json:"verified"`
}

// Verify ...
func Verify(c echo.Context) error {
	req := &VerifyRequest{}
	util.LoadRequestBody(c, req)

	publicKey := util.LoadRequestHeader(c, common.PublicKey)
	if len(publicKey) == 0 {
		c.Logger().Warn("empty publicKey")
		return c.JSON(http.StatusBadRequest, nil)
	}
	c.Logger().Debugf("publicKey: %s", publicKey)

	// original message
	msg := util.RandomMessage(publicKey) // publicKey as seed for now
	msgHash := crypto.Keccak256Hash([]byte(msg))
	c.Logger().Debugf("msgHash: %s", msgHash.Hex())

	signatureBytes, err := hexutil.Decode(req.SignedMessage) // signature in hex
	if err != nil {
		c.Logger().Errorf("Decode failed, err: %v", err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	// recover signed publicKey
	sigPublicKeyBytes, err := crypto.Ecrecover(msgHash.Bytes(), signatureBytes)
	if err != nil {
		c.Logger().Errorf("Ecrecover failed, err: %v", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	c.Logger().Debugf("sigPublicKeyHex: %s, publicKey: %s", string(hexutil.Encode(sigPublicKeyBytes)), publicKey)

	sigPublicKeyECDSA, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)
	if err != nil {
		c.Logger().Errorf("UnmarshalPubkey failed, err: %v", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// signed address
	sigAddressHex := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()

	verified := sigAddressHex == req.Address
	if !verified {
		c.Logger().Warnf("verified false, signed address: %s, req address: %s", sigAddressHex, req.Address)
	}
	resp := &VerifyResponse{
		Verified: verified,
	}
	return c.JSON(http.StatusOK, resp)
}
