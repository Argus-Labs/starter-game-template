package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

func prepareDirs(dirs ...string) error {
	for _, d := range dirs {
		if err := prepareDir(d); err != nil {
			return fmt.Errorf("failed to prepare dir %d: %w", d, err)
		}
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

func exitCardinalDir() error {
	curr, err := os.Getwd()
	if err != nil {
		return err
	}
	curr = filepath.Base(curr)
	if curr == "cardinal" {
		if err := os.Chdir(".."); err != nil {
			return err
		}

	}
	return nil
}
