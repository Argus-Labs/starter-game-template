//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Check verifies that various prerequisites are installed or configured on your machine
func Check() error {
	return checkPrereq(true)
}

// Clear deletes all the docker volumes
func Clear() error {
	if err := sh.RunV("docker", "compose", "down", "--volumes"); err != nil {
		return err
	}
	return nil
}

// Test runs the test suite
func Test() error {
	mg.Deps(exitMagefilesDir)
	mg.Deps(Clear)

	if err := prepareDirs("testsuite", "cardinal", "nakama"); err != nil {
		return err
	}
	if err := sh.RunV("docker", "compose", "up", "--build", "--abort-on-container-exit", "--exit-code-from", "testsuite", "--attach", "testsuite"); err != nil {
		return err
	}
	return nil
}

// Stop stops Nakama and cardinal
func Stop() error {
	return sh.Run("docker", "compose", "stop")
}

// Restart restarts ONLY cardinal.
func Restart() error {
	mg.Deps(exitMagefilesDir)
	if err := sh.Run("docker", "compose", "stop", "cardinal"); err != nil {
		return err
	}
	if err := sh.Run("docker", "compose", "up", "cardinal", "--build", "-d"); err != nil {
		return err
	}
	return nil
}

// Nakama starts just the Nakama server. The game server needs to be started some other way.
func Nakama() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("nakama"); err != nil {
		return err
	}
	env := map[string]string{
		"CARDINAL_ADDR": "http://host.docker.internal:3333",
	}
	if err := sh.RunWithV(env, "docker", "compose", "up", "--build", "nakama"); err != nil {
		return err
	}
	return nil
}

// Start starts Nakama and cardinal
func Start() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDirs("cardinal", "nakama"); err != nil {
		return err
	}
	if err := sh.RunV("docker", "compose", "up", "--build", "cardinal", "nakama"); err != nil {
		return err
	}
	return nil
}

// StartDetach starts Nakama and cardinal with detach and wait-timeout 60s (suit for CI workflow)
func StartDetach() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("cardinal"); err != nil {
		return err
	}
	if err := prepareDir("nakama"); err != nil {
		return err
	}
	if err := sh.RunV("docker", "compose", "up", "--detach", "--wait", "--wait-timeout", "60"); err != nil {
		return err
	}
	return nil
}

// Build only Nakama and cardinal
func Build() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("cardinal"); err != nil {
		return err
	}
	if err := prepareDir("nakama"); err != nil {
		return err
	}
	if err := sh.RunV("docker", "compose", "build"); err != nil {
		return err
	}
	return nil
}

// Start cardinal in dev mode without Nakama
func Dev() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("cardinal"); err != nil {
		return err
	}
	if err := os.Chdir("cardinal"); err != nil {
		return err
	}

	// Set environment variables for dev mode
	os.Setenv("REDIS_MODE", "normal")
	os.Setenv("CARDINAL_PORT", "3333")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("DEPLOY_MODE", "development")

	// Run redis in a docker container because Miniredis doesn't work with Retool
	// NOTE: this is because it doesn't implement CLIENT
	err := sh.RunV("docker", "run", "-d", "-p", "6379:6379", "-e", "LOCAL_REDIS=true", "--name", "cardinal-dev-redis", "redis")
	if err != nil {
		fmt.Println("Failed to create Redis container:", err)
		os.Exit(1)
	}
	fmt.Println("Dev redis container created successfully and exposed on port 6379!")

	// Make it possible to query Redis from HTTP
	err = sh.RunV("docker", "run", "-d", "-p", "7379:7379", "--link", "cardinal-dev-redis:redis", "--name", "cardinal-dev-webdis", "anapsix/webdis")
	if err != nil {
		fmt.Println("Failed to create Webdis container:", err)
		os.Exit(1)
	}
	fmt.Println("Dev webdis container created successfully and exposed on port 7379!")

	// We are going to run cardinal in dev mode without docker
	// and use Ngrok to expose the Miniredis to the internet
	// so that we can use tools like Retool to inspect the game state
	runCardinal := sh.RunCmd("go", "run", ".")

	// Run Cardinal as a goroutine
	errCh1 := make(chan error, 1)
	go func() {
		errCh1 <- runCardinal()
	}()

	// Wait for a signal to stop
	err1 := <-errCh1
	if err1 != nil {
		err = sh.RunV("docker", "rm", "-f", "cardinal-dev-redis")
		if err != nil {
			fmt.Println("Failed to delete Redis container:", err)
			fmt.Println("Please delete it manually with `docker rm -f cardinal-dev-redis`")
		}

		err = sh.RunV("docker", "rm", "-f", "cardinal-dev-webdis")
		if err != nil {
			fmt.Println("Failed to delete Webdis container:", err)
			fmt.Println("Please delete it manually with `docker rm -f cardinal-dev-webdis`")
		}

		return nil
	}

	return nil
}
