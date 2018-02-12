package util

import (
	"golang.org/x/crypto/ssh"
)

func ReadPem(path string) (ssh.AuthMethod, error) {
	key, err := ReadRsaPrivateKey(path)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signer), nil
}
