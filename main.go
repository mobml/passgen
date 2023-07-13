package main

import (
	"crypto/rand"
	"fmt"
	"github.com/atotto/clipboard"
	"math/big"
	"os"
	"strconv"
)

func generatePassword(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_-+{}[]"
	password := make([]byte, length)

	len_charset := big.NewInt(int64(len(charset)))

	//Generate a random password using a specified set of characters
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, len_charset)

		if err != nil {
			fmt.Printf("%s\n", red("Error ocurred"))
			return "", err
		}
		password[i] = charset[num.Int64()]
	}

	return string(password), nil

}

func copyToClipboard(password string) {
	err := clipboard.WriteAll(password)

	if err != nil {
		fmt.Printf("%s\n", red("The password could not be copied to the clipboard"))
		return
	}

	fmt.Printf("%s\n", blue("The password has been copied to the clipboard"))
}

func run() {
	if len(os.Args) < 2 {
		fmt.Printf("%s\n", red("Missing arguments"))
		return
	}

	length, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%s\n", red("Conversion error"))
		return
	}

	password, _ := generatePassword(length)

	if err != nil {
		fmt.Printf("%s\n", red("Error generating the password"))
		return
	}

	fmt.Printf("%s: %s\n", yellow("Password"), green("%s", password))
	copyToClipboard(password)
}

func main() {
	run()
}
