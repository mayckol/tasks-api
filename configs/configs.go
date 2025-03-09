package configs

import (
	"fmt"
	"github.com/mayckol/envsnatch"
	"os"
	"path/filepath"
)

type EnvVars struct {
	WebServerPort     int    `env:"WEB_SERVER_PORT"`
	MysqlRootPassword string `env:"MYSQL_ROOT_PASSWORD"`
	MysqlDatabase     string `env:"MYSQL_DATABASE"`
	MysqlUser         string `env:"MYSQL_USER"`
	MysqlPassword     string `env:"MYSQL_PASSWORD"`
	MysqlHost         string `env:"MYSQL_HOST"`
	MysqlPort         string `env:"MYSQL_PORT"`
}

// Config loads the configuration from the .env file or .env.test file and returns the configuration and the invalid variables
func Config(envName string) (*EnvVars, *[]envsnatch.UnmarshalingErr, error) {
	env := ".env"
	es, _ := envsnatch.NewEnvSnatch()
	fileDir, _ := filepath.Split(dir(env))
	es.AddPath(fileDir)
	es.AddFileName(envName)
	var envs EnvVars
	invalidVars, err := es.Unmarshal(&envs)
	return &envs, invalidVars, err
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
