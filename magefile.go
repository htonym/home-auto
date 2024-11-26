//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	dbSchemaDir = "./db/schema"
	dbConn      = "host=localhost dbname=home-auto user=postgres sslmode=disable password=password port=9000"
)

func Build() error {
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "arm64",
	}

	fmt.Println("Building binary...")
	return sh.RunWith(env, "go", "build", "-o", "bin/server", "./cmd/server/main.go")
}

func Clean() error {
	return sh.Rm("bin/server")
}

func Deploy() error {
	mg.Deps(Build)
	user := "tony"
	host := "rpi4.local"
	remoteBin := "/home/tony/dev/dht-server/bin/dhtserver"

	rmBinCmd := fmt.Sprintf("ssh %s@%s 'rm -f %s'", user, host, remoteBin)
	fmt.Println("Removing old binary...")
	if err := sh.Run("sh", "-c", rmBinCmd); err != nil {
		return err
	}

	copyCmd := fmt.Sprintf("scp bin/server %s@%s:%s", user, host, remoteBin)
	fmt.Println("Copying new binary to remote location...")
	if err := sh.Run("sh", "-c", copyCmd); err != nil {
		return err
	}

	fmt.Println("Restarting remote server...")
	restartCmd := fmt.Sprintf("ssh %s@%s 'sudo systemctl restart dht-server.service'", user, host)
	return sh.Run("sh", "-c", restartCmd)
}

func DbUp() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q up", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}

func DbStatus() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q status", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}

func DbDown() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q down", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}
