package inventory

import (
	"errors"

	"github.com/jklaiber/jumper/pkg/common"

	vault "github.com/sosedoff/ansible-vault-go"
)

const service = "jumper"

func (inventory *Inventory) readEncryptedFile(filePath string) (str string, err error) {
	str, err = vault.DecryptFile(filePath, common.GetSecretFromKeyring())
	if err != nil {
		return "", errors.New("file could not be decrypted")
	}
	return
}
