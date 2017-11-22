package main

import (
	"fmt"
	"os"

	"github.com/nshttpd/kube-tools/cmd/kubetool/actions"
)

func main() {
	if err := actions.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
