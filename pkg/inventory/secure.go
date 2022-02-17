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
