package secret

import (
	"log"
	"os/user"

	"github.com/jklaiber/jumper/internal/common"
	"github.com/zalando/go-keyring"
)

func getUsername() (username string) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("could not get username fromt the system")
	}
	return currentUser.Username
}

func SetSecretInKeyring(secret string) {
	err := keyring.Set(common.ServiceName, getUsername(), secret)
	if err != nil {
		log.Fatal("could not set secret in keyring")
	}
}

func GetSecretFromKeyring() (secret string) {
	secret, err := keyring.Get(common.ServiceName, getUsername())
	if err != nil {
		log.Fatal("could not get secret from keyring")
	}
	return
}

func DeleteSecretFromKeyring() {
	err := keyring.Delete(common.ServiceName, getUsername())
	if err != nil {
		log.Fatal("could not delete secret from keyring")
	}
}

func SecretAvailableFromKeyring() (available bool) {
	_, err := keyring.Get(common.ServiceName, getUsername())
	return err == nil
}
