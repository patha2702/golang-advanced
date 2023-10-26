package main

import(
	"fmt"
	"github.com/pborman/uuid"
	"github.com/patha2702/cryptit/encrypt"
)


func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)
	msg := "Rajendra"
	encryptMsg := encrypt.Encrypt(msg)
	fmt.Println(encryptMsg)
}
