package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ibrokethecloud/rancher-rbac-report/pkg/report"
)

func main() {
	root, err := report.NewReportCommand(context.TODO())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
