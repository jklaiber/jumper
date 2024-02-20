package inventory

import (
	"fmt"

	"github.com/jklaiber/jumper/internal/common"

	vault "github.com/sosedoff/ansible-vault-go"
)

func (inventory *Inventory) readEncryptedFile(filePath string) (str string, err error) {
	str, err = vault.DecryptFile(filePath, common.GetSecretFromKeyring())
	if err != nil {
		return "", fmt.Errorf("file could not be decrypted")
	}
	return
}
