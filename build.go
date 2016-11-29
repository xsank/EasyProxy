package main

import (
	"log"
	"flag"
	"os"
	"runtime"
	"strings"
	"os/exec"
	"io"
)

const (
	BinaryDir = "./bin/"
	Binary = "easyproxy"
	ConfDir = "./conf/"
	Config = "default.json"
	PKG = "./src"
)

func main() {
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
	return BinaryDir + execName()
}

func confFile() string {
	return ConfDir + Config
}

func execName() string {
	name := Binary
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return name
}

func build() {
	log.Println("start building...")
	args := []string{"build", "-ldflags", "-w -s"}
	args = append(args, "-o", execFile())
	args = append(args, PKG)
	runCommand("go", args...)
	copyConf()
}

func copyConf() {
	DstDir := BinaryDir + ConfDir
	os.Mkdir(DstDir, os.ModePerm)
	src, err := os.Open(confFile())
	if err != nil {
		log.Println("cannot open file:", err)
		return
	}
	defer src.Close()
	dst, err := os.Create(DstDir + Config)
	if err != nil {
		log.Println("create conf file failed:", err)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Println("copy conf failed:", err)
	}
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
	log.Println("start cleaning...")
	err := os.RemoveAll(BinaryDir)
	if err != nil {
		log.Println("clean files failed:", err)
	}
}
