//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Check verifies that various prerequisites are installed or configured on your machine
func Check() error {
	return checkPrereq(true)
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

// Start starts Nakama and cardinal
func Start() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("cardinal"); err != nil {
		return err
	}
	if err := prepareDir("nakama"); err != nil {
		return err
	}
	if err := sh.RunV("docker", "compose", "up", "--build"); err != nil {
		return err
	}
	return nil
}

// Start cardinal in dev mode
func Dev() error {
	mg.Deps(exitMagefilesDir)
	if err := prepareDir("cardinal"); err != nil {
		return err
	}
	if err := os.Chdir("cardinal"); err != nil {
		return err
	}
	os.Setenv("CARDINAL_PORT", "4200")
	os.Setenv("REDIS_ADDR", "localhost:6379")

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
	fmt.Println("\nhttps://editor.world.dev\n")

	// We are going to run cardinal in dev mode without docker
	// and use Ngrok to expose the Miniredis to the internet
	// so that we can use tools like Retool to inspect the game state
	runCardinal := sh.RunCmd("go", "run", "main.go")

	// Run Cardinal as a goroutine
	errCh1 := make(chan error, 1)
	go func() {
		errCh1 <- runCardinal()
	}()
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

		return err1
	}

	return nil
}

func prepareDir(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return err
	}
	if err := sh.Rm("./vendor"); err != nil {
		return err
	}
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	if err := sh.Run("go", "mod", "vendor"); err != nil {
		return err
	}
	if err := os.Chdir(".."); err != nil {
		return err
	}
	return nil
}

func exitMagefilesDir() error {
	curr, err := os.Getwd()
	if err != nil {
		return err
	}
	curr = filepath.Base(curr)
	if curr == "magefiles" {
		if err := os.Chdir(".."); err != nil {
			return err
		}

	}
	return nil
}
