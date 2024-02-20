package inventory

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jklaiber/jumper/internal/common"
	vault "github.com/sosedoff/ansible-vault-go"
	"gopkg.in/yaml.v2"
)

func (inventory *Inventory) EditInventory(filePath string) error {
	secret := common.GetSecretFromKeyring()

	decryptedContents, err := vault.DecryptFile(filePath, secret)
	if err != nil {
		return fmt.Errorf("file decryption failed: %v", err)
	}

	tempFile, err := os.CreateTemp("", "vault")
	if err != nil {
		return fmt.Errorf("temporary file creation failed: %v", err)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	if err := os.WriteFile(tempFile.Name(), []byte(decryptedContents), 0644); err != nil {
		return fmt.Errorf("writing to temporary file failed: %v", err)
	}

	editor := getEditor()
	cmd := exec.Command(editor, tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed to start: %v", editor, err)
	}

	editedContents, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return fmt.Errorf("reading temporary file failed: %v", err)
	}

	convertedContents := strings.ReplaceAll(string(editedContents), "\t", "    ")

	var yamlContent interface{}
	if err := yaml.Unmarshal([]byte(convertedContents), &yamlContent); err != nil {
		if err = vault.EncryptFile(filePath, string(decryptedContents), secret); err != nil {
			return fmt.Errorf("file encription of original content fialed: %v", err)
		}
		return fmt.Errorf("YAML validation failed: %v", err)
	}

	if err := vault.EncryptFile(filePath, string(convertedContents), secret); err != nil {
		return fmt.Errorf("file encryption failed: %v", err)
	}

	return nil
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "vim"
	}
	return editor
}
