package report

import (
	"context"

	"github.com/spf13/cobra"
)

func NewReportCommand(ctx context.Context) (*cobra.Command, error) {
	var err error
	rc := ReportCommandConfig{}
	rootCmd := &cobra.Command{
		Use:   "kubectl-rancher-rbac-report",
		Short: "kubectl plugin to generate Rancher rbac report",
		Long: `The plugin interacts with the k8s api of the cluster where Rancher is installed (local) cluster, and attempts to 
generate a list of all rbac settings being managed by Rancher.`,
		Run: func(cmd *cobra.Command, args []string) {
			rc.Client, err = CreateClientset()
			if err != nil {
				panic(err)
			}
			rc.Context = ctx
			rc.GenerateReport()
		},
	}

	rootCmd.PersistentFlags().StringVarP(&rc.Cluster, "cluster", "c", "", "Generate report for specific cluster only")

	return rootCmd, nil
}
