package src

import (
	"io/ioutil"
	"os"
)

type Sign struct {
	rsaPublicKey  []byte
	rsaPrivateKey []byte
}

func NewSignAgent() *Sign {
	checkErr := func(err error) {
		os.Exit(1)
	}
	sign := &Sign{}
	var err error
	sign.rsaPrivateKey, err = ioutil.ReadFile("../conf/rsa_private_key.pem")
	checkErr(err)
	sign.rsaPublicKey, err = ioutil.ReadFile("../conf/rsa_public_key.pem")
	checkErr(err)
	return sign
}

func (s *Sign) Encryption(data []byte) []byte {

   return nil
}
