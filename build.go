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

func build() {
	file := "./bin/" + Binary
	if runtime.GOOS == "windows" {
		file += ".exe"
	}
	args := []string{"build"}
	args = append(args, "-o", file)
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
	runCommand("rm", "-r", "./bin")
}
