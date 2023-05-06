package decrypt

func Decryption(encryptedString string) string {
	decryptedString := ""

	for _, c := range encryptedString {
		decryptedString += string(int(c) - 5)
	}
	return decryptedString
}
