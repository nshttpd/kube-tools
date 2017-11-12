package main

import (
	"fmt"
	"os"

	"github.com/nshttpd/kube-tools/kubesecret/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
