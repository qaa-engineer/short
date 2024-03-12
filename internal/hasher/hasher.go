package hasher

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

func GetShortLink(input string) (string, error) {

	if len(input) == 0 {
		return "", errors.New("input should not be empty")
	}

	urlHashBytes := sha256Of(input)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}
	return finalString[:8], nil
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return string(encoded), err
}
