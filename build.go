package main

import (
	"log"
	"flag"
	"os"
	"runtime"
	"strings"
	"os/exec"
)

const (
	Binary = "easyproxy"
	PKG = "./src"
)

func main() {
	log.Println("start building...")
	flag.Parse()
	if flag.NArg() == 0 {
		log.Println("please use go run build.go to build")
		return
	}
	for _, cmd := range flag.Args() {
		switch cmd {
		case "build":
			clean()
			build()
		case "clean":
			clean()
		default:
			log.Println("unknown command %s", cmd)
		}
	}
}

func execFile() string {
	return "./bin/" + execName()
}

func execName() string {
	name := Binary
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return name
}

func build() {
	args := []string{"build", "-ldflags", "-w -s"}
	args = append(args, "-o", execFile())
	args = append(args, PKG)
	runCommand("go", args...)
}

func runCommand(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		log.Println(err)
	}
}

func clean() {
	runCommand("rm", "-f", execFile())
}
