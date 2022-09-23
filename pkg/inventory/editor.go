package inventory

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jklaiber/jumper/pkg/common"
	vault "github.com/sosedoff/ansible-vault-go"
)

func (inventory *Inventory) EditInventory(filePath string, password string) (err error) {
	result, err := vault.DecryptFile(filePath, common.GetSecretFromKeyring())
	if err != nil {
		return errors.New("file could not be decrypted")
	}

	// Create a new temp file
	var tempFile *os.File
	tempFile, err = createTempFile()
	if err != nil {
		return errors.New("temporary file could not be created")
	}

	// Write decrypted inputs to temp file
	err = ioutil.WriteFile(tempFile.Name(), []byte(result), 0644)
	if err != nil {
		return errors.New("inventory could not be written to temporary file")
	}

	// Open editor for modifications
	cmd := exec.Command(getEditor(), tempFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return errors.New("could not start the editor")
	}

	// Encrypt changed content
	var unencryptedContents []byte
	unencryptedContents, err = ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return err
	}
	err = vault.EncryptFile(filePath, string(unencryptedContents), common.GetSecretFromKeyring())
	if err != nil {
		return errors.New("could not encrypt temporary file")
	}

	err = tempFile.Close()
	if err != nil {
		return errors.New("temporary file could not be closed")
	}

	err = cleanupFile(tempFile)
	if err != nil {
		return err
	}

	return err
}

func createTempFile() (*os.File, error) {
	t, err := ioutil.TempFile("", "vault")
	if err != nil {
		return nil, err
	}
	return t, nil
}

func cleanupFile(t *os.File) error {

	_, err := os.Stat(t.Name())
	if os.IsNotExist(err) {
		return nil
	}

	// Delete the temp file
	err = os.Remove(t.Name())
	if err != nil {
		return err
	}
	return nil
}

func getEditor() string {
	var editorEnv = os.Getenv("EDITOR")
	if editorEnv == "" {
		return "vim"
	}
	return editorEnv
}
