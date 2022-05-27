package inventory

import (
	"errors"

	vault "github.com/sosedoff/ansible-vault-go"
)

func (inventory *Inventory) readEncryptedFile(filePath string, password string) (str string, err error) {
	str, err = vault.DecryptFile(filePath, password)
	if err != nil {
		return "", errors.New("file could not be decrypted")
	}
	return
}

func (inventory *Inventory) writeEncryptedFile(filePath string, password string, invToWrite []byte) (err error) {
	err = vault.EncryptFile(filePath, string(invToWrite), password)
	if err != nil {
		return errors.New("file could not be encrypted")
	}
	return
}
