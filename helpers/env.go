package helpers

import (
	"os"
)

func SetEnv() {
	err := os.Setenv("DATABASE_URL", "postgres://postgres:12345@localhost/fgmotoru")
	ErrDefaultDetect(err, "Ошибка установки ENV")
}