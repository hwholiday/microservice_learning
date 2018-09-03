package src

import (
	"io/ioutil"
	"os"
	"github.com/lunny/log"
)

type SignAgent struct {
	rsaPublicKey  []byte
	rsaPrivateKey []byte
}

func NewSignAgent() *SignAgent {
	checkErr := func(err error) {
		log.Error(err)
		os.Exit(1)
	}
	sign := &SignAgent{}
	var err error
	sign.rsaPrivateKey, err = ioutil.ReadFile("../conf/rsa_private_key.pem")
	checkErr(err)
	sign.rsaPublicKey, err = ioutil.ReadFile("../conf/rsa_public_key.pem")
	checkErr(err)
	return sign
}

func Run()  {
	
}
