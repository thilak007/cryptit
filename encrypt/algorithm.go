package encrypt

func SecretEncryption(str string) string {
	encryptedString := ""

	for _, c := range str {
		encryptedString += string(int(c) + 5)
	}
	return encryptedString
}
