package inventory

import (
	"fmt"

	"github.com/jklaiber/jumper/internal/secret"
	vault "github.com/sosedoff/ansible-vault-go"
)

type InventoryReader interface {
	ReadInventory(filePath string) (string, error)
}

type DefaultInventoryReader struct{}

func (reader *DefaultInventoryReader) ReadInventory(filePath string) (string, error) {
	str, err := vault.DecryptFile(filePath, secret.GetSecretFromKeyring())
	if err != nil {
		return "", fmt.Errorf("file could not be decrypted")
	}
	return str, nil
}
