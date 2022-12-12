package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, userid primitive.ObjectID) (string, error) {
	if primitive.IsValidObjectID(userid.Hex()) {
		initialLink += userid.Hex()
		urlHashBytes := sha256Of(initialLink)
		generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
		finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
		return finalString[:9], nil
	} else {
		return "", errors.New("the userId used to create the shortened link is invalid")
	}
}
