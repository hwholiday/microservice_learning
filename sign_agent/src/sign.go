package src

type SingAgent interface {
	Encryption([]byte) []byte          //加密
	Decrypt([]byte) []byte             //解密
	Verification([]byte) (bool, error) //验证
}
