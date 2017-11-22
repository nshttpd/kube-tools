package actions

import (
	"github.com/nshttpd/kube-tools/configmap"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cmRootCmd)
	cmRootCmd.AddCommand(cmGetCmd)
	cmGetCmd.Flags().StringVar(&keyName, "name", "", "name of configmap item to fetch")
}

var cmRootCmd = &cobra.Command{
	Use:     "configmap",
	Aliases: []string{"cm"},
	Short:   "configmap CRUD level manipulation",
	Long:    `create, delete, get update of configmap objects in cluser`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get configmap from cluster",
	Args:  cobra.ExactArgs(1),
	Long: `get specified configmap from cluster and save it to a JSON file in 'cwd'
if -name option specified get that specific piece of configmap data`,
	Run: func(cmd *cobra.Command, args []string) {
		configmap.GetConfigMap(args[0], cmd.Flag("namespace").Value.String(),
			keyName, loadClient())
	},
}

var cmCreatecmd = &cobra.Command{
	Use:   "create",
	Short: "create configmap in cluster",
	Args:  cobra.ExactArgs(1),
	Long:  `create a configmap either bare or with data`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// create
// delete
// update
