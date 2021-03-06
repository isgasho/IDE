// File generated by Gopher Sauce
// DO NOT EDIT!!
package main

import (
	"github.com/thestrukture/IDE/api/globals"
	"github.com/thestrukture/IDE/api/methods"
	"github.com/thestrukture/IDE/types"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cheikhshift/gos/core"
	"github.com/fatih/color"
)

func LaunchServer() {

	globals.Dfd = os.ExpandEnv("$GOPATH")
	globals.Windows = strings.Contains(runtime.GOOS, "windows")
	if globals.Dfd == "" {
		fmt.Println("Using temporary $GOPATH")
		if globals.Windows {
			os.Chdir(os.ExpandEnv("$USERPROFILE"))
		} else {
			os.Chdir(os.ExpandEnv("$HOME"))
		}

		err := os.MkdirAll("workspace/", 0700)
		if err != nil {
			fmt.Println(err.Error())

		} else {
			//download go
			os.MkdirAll("workspace/src", 0700)
			os.MkdirAll("workspace/bin", 0700)
			cwd, _ := os.Getwd()
			cwd = cwd + "/workspace"
			os.Setenv("GOPATH", cwd)
			pathbin := os.ExpandEnv("$PATH")
			if globals.Windows {
				os.Setenv("PATH", pathbin+":"+strings.Replace(cwd+"/bin", "/", "\\", -1))
			} else {
				os.Setenv("PATH", pathbin+":"+cwd+"/bin")
			}
			_, goinvs := core.RunCmdSmart("go help build")
			if goinvs != nil {
				fmt.Println("If you do not have GO, remember to download and install GO to complete installation. Find Go here : https://golang.org/dl/ . ***Installing GO requires sudo permission.")

			}
			globals.Dfd = cwd

		}
	}

	dir := os.ExpandEnv("$GOPATH") + "/src/github.com/cheikhshift/gos"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Downloading GoS")
		_, err := core.RunCmdSmart("go get github.com/cheikhshift/gos")
		if err != nil {
			color.Red("Please, install GO : https://golang.org/dl/ ")
		} else {
			core.RunCmdSmart("gos deps")
		}
	}

	if _, err := os.Stat(os.ExpandEnv("$GOPATH") + "/src/github.com/nsf/gocode/"); os.IsNotExist(err) {
		fmt.Println("Go code completion not present, installing from github.com/nsf/gocode")
		core.RunCmdSmart("go get github.com/nsf/gocode")
	}

	if _, err := os.Stat(os.ExpandEnv("$GOPATH") + "/src/github.com/golang/dep"); os.IsNotExist(err) {
		fmt.Println("Go dep not present, installing from github.com/golang/dep")
		core.RunCmdSmart("go get github.com/golang/dep/cmd/dep")
	}

	globals.AutocompletePath = filepath.Join(os.ExpandEnv("$GOPATH"), "strukture-autocomplete")

	if _, err := os.Stat(globals.AutocompletePath); os.IsNotExist(err) {
		fmt.Println("Creating autocomplete resource folder at " + globals.AutocompletePath)
		os.MkdirAll(globals.AutocompletePath, 0700)
	}

	apps := methods.GetApps()
	newapps := []types.App{}
	for _, app := range apps {
		if app.Name != "" {
			app.Pid = ""
			newapps = append(newapps, app)
		}
	}
	methods.SaveApps(newapps)

	log.Println("Strukture up on port 8884")
	if len(os.Args) == 1 && !globals.Windows {
		if isMac := strings.Contains(runtime.GOOS, "arwin"); isMac {
			core.RunCmd("open http://localhost:8884/index")
		} else {
			core.RunCmd("xdg-open http://localhost:8884/index")
		}
	} else if len(os.Args) == 1 && globals.Windows {
		core.RunCmd("cmd /C start http://localhost:8884/index")
	}

}
