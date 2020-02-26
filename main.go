package main

import (
	"fmt"

	"github.com/google/go-tpm/tpm"
)

func main() {
	rwc, err := tpm.OpenTPM("/dev/tpm0")
	if err != nil {
		fmt.Printf("Error. Cannot open /dev/tpm0, %s", err)
		return
	}
	pubEK, err := tpm.ReadPubEK(rwc)
	if err != nil {
		fmt.Printf("Error. Cannot read endorsemen public key, %s", err)
		return
	}

	pubEKStr := string(pubEK)
	fmt.Println(pubEKStr)
}
