package client_test

import (
	"fmt"
	"testing"

	"github.com/icehubin/futu-go/client"
)

func TestRsa(t *testing.T) {
	var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDIHSwgItw3n8kwKpdP1230Tv0lF+mEnx/jBRwx2WVpMcjuqxjr
f9xZBxkPe5I5h9cUXDkw3Sb0xea4B9g8X6sNAjHFt4+SsxYHDUSF+YsSUo1bQqeN
mHb1MydCCFX06YwhiV+CTuWg+bayd6d5TQbp7M/kweTrl7Tc6kwyD7HyxQIDAQAB
AoGAOBz+B5De4XYbyzbWMRw+h9gyYdlfsMGSBm2jZ1MHfcTpDAvMNBUpa4ly/IVx
DlLJ+7qUiuAokVCBDq+ccel9o+wQgZQMXCi09hgYae+bT/T/tVq0hvjHsdQGxIWC
ofdPvMBhmsHmrG1F9M+RE5bPMfPh6bCAtipMFjRtsXutn2ECQQDuybM3VmBnmHoW
cevEA16GujfojYr1ilIZKtni7FZ5wAaNnyd48t3Zq3bQEn2mqMUI/iRzKDjlFNT/
Kq9Ka4vpAkEA1onUq3bkPdnQmIGjXHLFW7ByQyHGKPfF48Y1r6gBu9kIMIdnunGN
es80nMmVDbwYK5rX4tOBvvpZWUfK/q5SfQJBAOcT/7+0chSuhK9FzU9hp08f1EzS
9L+K/MnoIvSrmz+06WKRuwQbmz8y5AaVbk0ik5KRRjq+mNhvYXQRHZMaIkkCQQCX
mr7WjCWuL4Xgv3uZkOE2cfHMzskhmjYVR7QYdTkEbdIDuSvr7OJ1rocXZLwYAJtz
9PAqMhy1wGPzW3BvmPBhAkEAsP/qVrlZEwONsc3l4ARjyiPKqsyaYiwr5JG+ARWi
AuKkOuBf/mGqYfcZunr1GS+d86T9zItEjNN3AzTITfs+aw==
-----END RSA PRIVATE KEY-----
`)
	body := []byte(`ahahahahahahadahadfasdfasdfasdfadsfhaddfsfad
	adfadfasdfasdfasdfadfadfadfasdfadsfasdfadfasdfadsfadf
	adfadsfadsfadf3tgf;lkjaldhfop32rnldhjfoq3jrnl;sdhfpvispn;aerigk;kzhoudjnaf
	pasdufoajnt43;ou9hbnv/sghpay7syhg';sdfjlp87394iwktgjaldh9asdlbgpou'asdf
	p7afualdhp9q3lh;piaydpofh'ql4yof;jalyff9ayu'pjfvmnbandbvpa7PEo;jntbgvp8GS
	a[df-9awel.l'lfuposfug']fd.ljajh[auweho[8]]
	`)
	fmt.Println(string(body))
	encrypt_body := client.RsaEncrypt(body, privateKey)
	fmt.Println(encrypt_body)
	decrypt_body := client.RsaDecrypt(encrypt_body, privateKey)
	fmt.Println(string(decrypt_body))

	aesKey := []byte(`481AB0DC9294C443`)
	iv := []byte(`75B7A8D39FF4EFE0`)

	cbcEncrypt := client.AesCBCEncrypt(body, aesKey, iv)
	fmt.Println(cbcEncrypt)
	cbcDecrypt := client.AesCBCDecrypt(cbcEncrypt, aesKey, iv)
	fmt.Println(string(cbcDecrypt))

	ftEcbEncrypt := client.AesFTECBEncrypt(body, aesKey)
	fmt.Println(ftEcbEncrypt)
	ftEcbDecrypt := client.AesFTECBDecrypt(ftEcbEncrypt, aesKey)
	fmt.Println(string(ftEcbDecrypt))

}
