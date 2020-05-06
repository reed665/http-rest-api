package sqlstore_test

import (
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

var (
	conf Config
)

func TestMain(m *testing.M) {
	conf.DatabaseURL = os.Getenv("DATABASE_URL")
	if conf.DatabaseURL == "" {
		_, err := toml.DecodeFile("../../../../configs/apiserver_test.toml", &conf)
		if err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(m.Run())
}
