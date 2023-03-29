package helpers

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

func ConvertPrivateKeyToPublic(privateKey string) (string, error) {
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	publicKey := signer.PublicKey()
	return strings.TrimSuffix(string(ssh.MarshalAuthorizedKey(publicKey)), "\n"), nil
}
