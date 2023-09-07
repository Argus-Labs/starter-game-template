//go:build mage

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	ignore "github.com/sabhiram/go-gitignore"
	"golang.org/x/mod/modfile"
)

func makeGitIgnore() (*ignore.GitIgnore, error) {
	return ignore.CompileIgnoreFileAndLines(".gitignore",
		".git",
		"go.mod",
		"go.sum")
}

// Copy copies this sample project to the <target> directory and initializes it with 'go mod init <modulePath>'.
// The module path parameter should be set to your code's repository. See https://golang.org/ref/mod#go-mod-init
// for more info about go mod.
func Copy(target, modulePath string) error {
	ignore, err := makeGitIgnore()
	if err != nil {
		return err
	}

	mg.Deps(exitMagefilesDir)
	if err := os.MkdirAll(target, os.ModePerm); err != nil {
		return err
	}
	fmt.Printf("copying files to %q\n", target)
	walkErr := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ignore.MatchesPath(path) {
			return nil
		}
		if info.IsDir() {
			if err := os.MkdirAll(filepath.Join(target, path), os.ModePerm); err != nil {
				return mg.Fatalf(1, "failed to make path %q and %q", target, path)
			}
			return nil
		}
		source := filepath.Join(".", path)
		dest := filepath.Join(target, path)
		if err := sh.Copy(dest, source); err != nil {
			return mg.Fatalf(1, "copy failure: %v", err)
		}
		return nil
	})

	if walkErr != nil {
		fmt.Printf("error during file walk: %v", walkErr)
		return walkErr
	}

	replaceDirectives, err := getReplaceDirectives("cardinal/go.mod")
	if err != nil {
		return mg.Fatalf(1, "failed to get replace directives: %v", err)
	}

	if err := os.Chdir(target); err != nil {
		return mg.Fatalf(1, "failed to change directory to %q: %v\n", target, err)
	}

	if err := goModInitAndTidy(modulePath, "nakama", nil); err != nil {
		return err
	}
	if err := goModInitAndTidy(modulePath, "magefiles", nil); err != nil {
		return err
	}

	if err := goModInitAndTidy(modulePath, "cardinal", replaceDirectives); err != nil {
		return err
	}

	fmt.Println("All done. SUCCESS!")
	return nil
}

func goModInitAndTidy(modulePath, component string, replace []*modfile.Replace) error {
	fmt.Printf("running 'go mod init' for %q\n", component)
	fmt.Println()
	if err := os.Chdir(component); err != nil {
		return mg.Fatalf(1, "chdir for %q failure: %v", component, err)
	}
	if err := sh.Run("go", "mod", "init", modulePath+"/"+component); err != nil {
		return mg.Fatalf(1, "go mod init for %q failure: %v", component, err)
	}
	if replace != nil {
		fmt.Println("adding custom replace directives")
		if err := addReplaceDirectiveToGoMod("go.mod", replace); err != nil {
			return mg.Fatalf(1, "adding replace directive for %q failed: %v", component, err)
		}
		fmt.Println("successfully added custom replace directives")
	}
	fmt.Println("successfully ran 'go mod init'")
	fmt.Printf("running 'go mod tidy' for %q\n", component)
	fmt.Println()
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return mg.Fatalf(1, "go mod tidy for %q failure: %v", component, err)
	}
	if err := os.Chdir(".."); err != nil {
		return mg.Fatalf(1, "chdir .. for %q failure: %v", component, err)
	}
	fmt.Println("successfully ran 'go mod tidy'")
	return nil
}

func getReplaceDirectives(file string) ([]*modfile.Replace, error) {
	mod, err := getModFile(file)
	if err != nil {
		return nil, err
	}
	return mod.Replace, nil
}

func addReplaceDirectiveToGoMod(file string, replace []*modfile.Replace) error {
	mod, err := getModFile(file)
	if err != nil {
		return err
	}
	for _, r := range replace {
		o, n := r.Old, r.New
		if err := mod.AddReplace(o.Path, o.Version, n.Path, n.Version); err != nil {
			return err
		}
	}
	buf, err := mod.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(file, buf, 0)
}

func getModFile(file string) (*modfile.File, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return modfile.Parse(file, buf, nil)
}
