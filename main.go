package main

import (
	"context"
	"fmt"
	"github.com/ibrokethecloud/rancher-rbac-lister/pkg/lister"
	"os"
)

func main() {
	root, err := lister.NewReportCommand(context.TODO())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
