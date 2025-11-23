package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func generateHMAC(secret, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	return hex.EncodeToString(signature)
}

func verifyHMAC(secret, message, receivedSig string) bool {
	key := []byte(secret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	receivedMAC, err := hex.DecodeString(receivedSig)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedMAC, receivedMAC)
}

func signRequest(c *gin.Context) {
	psk := c.Request.Header.Get("HMAC-PSK") // For Testing purposes TO-DO
	request := c.Query("request")

	signature := generateHMAC(psk, request)

	c.JSON(200, gin.H{
		"signature": signature,
		"signedFor": request,
	})
}

func verifyRequest(c *gin.Context) {
	psk := c.Request.Header.Get("HMAC-PSK") // For Testing purposes TO-DO
	signed := c.Query("signature")
	request := c.Query("request")
	isValid := verifyHMAC(psk, request, signed)
	if !isValid {
		c.JSON(403, gin.H{
			"message": "Invalid HMAC signature",
			"success": false,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Valid HMAC signature",
			"success": true,
		})
		return
	}
}
