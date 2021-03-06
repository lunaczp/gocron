// +build node
// 任务节点

package main

import (
	"flag"
	"fmt"
	"github.com/ouqiang/gocron/modules/rpc/auth"
	"github.com/ouqiang/gocron/modules/rpc/server"
	"github.com/ouqiang/gocron/modules/utils"
	"os"
	"runtime"
	"strings"
)

const AppVersion = "1.3"

func main() {
	var serverAddr string
	var allowRoot bool
	var version bool
	var CAFile string
	var certFile string
	var keyFile string
	var enableTLS bool
	flag.BoolVar(&allowRoot, "allow-root", false, "./gocron-node -allow-root")
	flag.StringVar(&serverAddr, "s", "0.0.0.0:5921", "./gocron-node -s ip:port")
	flag.BoolVar(&version, "v", false, "./gocron-node -v")
	flag.BoolVar(&enableTLS, "enable-tls", false, "./gocron-node -enable-tls")
	flag.StringVar(&CAFile, "ca-file", "", "./gocron-node -ca-file path")
	flag.StringVar(&certFile, "cert-file", "", "./gocron-node -cert-file path")
	flag.StringVar(&keyFile, "key-file", "", "./gocron-node -key-file path")
	flag.Parse()

	if version {
		fmt.Println(AppVersion)
		return
	}

	if enableTLS {
		if !utils.FileExist(CAFile) {
			fmt.Printf("failed to read ca cert file: %s", CAFile)
			return
		}
		if !utils.FileExist(certFile) {
			fmt.Printf("failed to read server cert file: %s", certFile)
			return
		}
		if !utils.FileExist(keyFile) {
			fmt.Printf("failed to read server key file: %s", keyFile)
			return
		}
	}

	certificate := auth.Certificate{
		CAFile:   strings.TrimSpace(CAFile),
		CertFile: strings.TrimSpace(certFile),
		KeyFile:  strings.TrimSpace(keyFile),
	}

	if runtime.GOOS != "windows" && os.Getuid() == 0 && !allowRoot {
		fmt.Println("Do not run gocron-node as root user")
		return
	}

	server.Start(serverAddr, enableTLS, certificate)
}
