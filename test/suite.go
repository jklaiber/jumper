package test

import (
	"testing"

	"github.com/jklaiber/jumper/internal/config"
)

const (
	testConfigPath = "./.jumper.yaml"
)

func SetupTestConfig(tb testing.TB) {
	config.GetInstance().SetConfigFile(testConfigPath)
	if err := config.Parse(); err != nil {
		tb.Fatalf("could not parse test config: %v", err)
	}
}
