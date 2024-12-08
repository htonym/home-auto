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
	dbConnRpi4  = "host=rpi4.local dbname=home-auto user=postgres sslmode=disable password=password port=9000"
)

func Build() error {
	mg.Deps(Clean)
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "arm64",
	}

	fmt.Println("Building binary...")
	return sh.RunWith(env, "go", "build", "-o", "bin/server", "./cmd/server/main.go")
}

func BuildLocal() error {
	mg.Deps(Clean)

	fmt.Println("Building binary...")
	return sh.Run("go", "build", "-o", "bin/server", "./cmd/server/main.go")
}

func Clean() error {
	fmt.Println("Removing bin/server...")
	return sh.Rm("bin/server")
}

func Deploy() error {
	mg.Deps(Build)
	user := "tony"
	host := "rpi4.local"
	remoteBin := "/home/tony/dev/dht-server/bin/server"

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

func CleanClient() error {
	fmt.Println("Removing bin/client...")
	return sh.Rm("bin/client")
}

func BuildClient() error {
	mg.Deps(CleanClient)
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "arm64",
	}

	fmt.Println("Building client binary...")
	return sh.RunWith(env, "go", "build", "-o", "bin/client", "./cmd/client/main.go")
}

func DeployClient() error {
	mg.Deps(BuildClient)
	user := "tony"
	host := "rpi4.local"
	remoteBin := "/home/tony/dev/dht-server/bin/client"

	fmt.Println("Removing old binary...")
	rmBinCmd := fmt.Sprintf("ssh %s@%s 'rm -f %s'", user, host, remoteBin)
	if err := sh.Run("sh", "-c", rmBinCmd); err != nil {
		return err
	}

	fmt.Println("Copying new binary to remote location...")
	copyCmd := fmt.Sprintf("scp bin/client %s@%s:%s", user, host, remoteBin)
	if err := sh.Run("sh", "-c", copyCmd); err != nil {
		return err
	}

	// fmt.Println("Restarting remote server...")
	// restartCmd := fmt.Sprintf("ssh %s@%s 'sudo systemctl restart dht-server.service'", user, host)
	// return sh.Run("sh", "-c", restartCmd)

	return nil
}

func BuildClientRpi2() error {
	mg.Deps(CleanClient)
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "arm",
		"GOARM":  "7",
	}

	fmt.Println("Building client binary...")
	return sh.RunWith(env, "go", "build", "-o", "bin/client", "./cmd/client/main.go")
}

func DeployClientRpi2() error {
	mg.Deps(BuildClientRpi2)
	user := "tony"
	host := "rpi2.local"
	remoteBin := "/home/tony/home-auto-client/client"

	fmt.Println("Removing old binary...")
	rmBinCmd := fmt.Sprintf("ssh %s@%s 'rm -f %s'", user, host, remoteBin)
	if err := sh.Run("sh", "-c", rmBinCmd); err != nil {
		return err
	}

	fmt.Println("Copying new binary to remote location...")
	copyCmd := fmt.Sprintf("scp bin/client %s@%s:%s", user, host, remoteBin)
	if err := sh.Run("sh", "-c", copyCmd); err != nil {
		return err
	}

	fmt.Println("Restarting remote client...")
	restartCmd := fmt.Sprintf("ssh %s@%s 'sudo systemctl restart dht-client.service'", user, host)
	return sh.Run("sh", "-c", restartCmd)

	return nil
}

func DbStatus() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q status", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}

func DbUp() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q up", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}

func DbDown() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q down", dbSchemaDir, dbConn)
	return sh.Run("sh", "-c", cmd)
}

func DbStatusRpi4() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q status", dbSchemaDir, dbConnRpi4)
	return sh.Run("sh", "-c", cmd)
}

func DbUpRpi4() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q up", dbSchemaDir, dbConnRpi4)
	return sh.Run("sh", "-c", cmd)
}

func DbDownRpi4() error {
	cmd := fmt.Sprintf("goose -dir %s postgres %q down", dbSchemaDir, dbConnRpi4)
	return sh.Run("sh", "-c", cmd)
}
