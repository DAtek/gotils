package gotils

import "os"

type EnvConfig string

func (val EnvConfig) Load() string {
	return os.Getenv(string(val))
}
