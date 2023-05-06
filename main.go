package main

import (
	"fmt"

	"github.com/thilak007/cryptit/decrypt"
	"github.com/thilak007/cryptit/encrypt"
)

func main() {

	text := "Have a nice day!"
	fmt.Println(encrypt.SecretEncryption(text))

	var encryptedText string
	fmt.Scan(&encryptedText)

	decryptedText := decrypt.Decryption(encryptedText)
	fmt.Printf("%v decrypted value is %v\n", encryptedText, decryptedText)
}
