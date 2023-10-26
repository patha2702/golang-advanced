package main

import (
	"fmt"
	"github.com/patha2702/cryptit/encrypt"
	"github.com/patha2702/cryptit/decrypt"
)


func main() {
	message := "Hello World"
	encryptedData := encrypt.Encrypt(message)
	decryptedData := decrypt.Decrypt(encryptedData)
	fmt.Println("Original message: ", message)
	fmt.Println("Encrypted Data: ", encryptedData)
	fmt.Println("Decrypted Data: ", decryptedData)
}
