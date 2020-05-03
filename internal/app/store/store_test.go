package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/reed665/http-rest-api/internal/app/store"
)

var (
	config *store.Config
)

func TestMain(m *testing.M) {
	config = store.NewConfig()
	config.DatabaseURL = os.Getenv("DATABASE_URL")
	if config.DatabaseURL == "" {
		_, err := toml.DecodeFile("../../../configs/apiserver_test.toml", config)
		if err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(m.Run())
}
