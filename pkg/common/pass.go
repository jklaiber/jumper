package common

import (
	"log"
	"os/user"

	"github.com/zalando/go-keyring"
)

func getUsername() (username string) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return currentUser.Username
}

func SetSecretInKeyring(secret string) {
	err := keyring.Set(ServiceName, getUsername(), secret)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSecretFromKeyring() (secret string) {
	secret, err := keyring.Get(ServiceName, getUsername())
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DeleteSecretFromKeyring() {
	err := keyring.Delete(ServiceName, getUsername())
	if err != nil {
		log.Fatal(err)
	}
}

func SecretAvailableFromKeyring() (available bool) {
	_, err := keyring.Get(ServiceName, getUsername())
	if err != nil {
		return false
	}
	return true
}
