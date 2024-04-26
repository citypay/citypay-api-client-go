package citypay

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"
)

func GenerateApiKey(clientID string, licenseKey string) string {
	nonce := make([]byte, 16)
	_, _ = rand.Read(nonce)
	return GenerateApiKeyFor(clientID, licenseKey, nonce, time.Now().UTC())
}

func GenerateApiKeyFor(clientID string, licenseKey string, nonce []byte, dt time.Time) string {
	// Formatting the datetime as "YYYYMMDDHHMM"
	ds := dt.Format("200601021504")
	dsBytes, _ := hex.DecodeString(ds)

	// Creating the message for HMAC
	var message []byte
	message = append(message, []byte(clientID)...)
	message = append(message, nonce...)
	// Convert datetime string to bytes and append
	message = append(message, dsBytes...)

	// Creating HMAC with SHA256
	h := hmac.New(sha256.New, []byte(licenseKey))
	h.Write(message)
	digest := h.Sum(nil)

	// Preparing the final byte slice to encode
	var dest []byte
	dest = append(dest, []byte(clientID)...)
	dest = append(dest, ':')
	dest = append(dest, []byte(strings.ToUpper(hex.EncodeToString(nonce)))...)
	dest = append(dest, ':')
	dest = append(dest, digest...)

	// Base64 Encoding
	return base64.StdEncoding.EncodeToString(dest)
}
