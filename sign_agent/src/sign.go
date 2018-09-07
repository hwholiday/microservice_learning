package src

type SingAgent interface {
	Encryption()[]byte//加密
	Decrypt()[]byte//解密
	Verification()(bool,error)//验证
}
