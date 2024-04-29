package citypay

import "crypto/rand"

func getValidCardNumber() string { return "5525864776075446" }

func generateRandomId() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, 10)
	rand.Read(b)

	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return "test-" + string(b)
}
