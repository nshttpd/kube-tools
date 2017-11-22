package actions

import (
	"github.com/nshttpd/kube-tools/configmap"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cmRootCmd)
	cmRootCmd.AddCommand(cmGetCmd)
	cmRootCmd.AddCommand(cmCreateCmd)
	cmRootCmd.AddCommand(cmDeleteCmd)
	cmGetCmd.Flags().StringVar(&keyName, "name", "", "name of configmap item to fetch")
	cmCreateCmd.Flags().StringVar(&keyName, "name", "", "name of item to add on creation")
	cmCreateCmd.Flags().StringVar(&keyValue, "value", "", "value of item to add on creation")
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

var cmCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create configmap object",
	Args:  cobra.ExactArgs(1),
	Long:  `create a configmap either bare or with data`,
	Run: func(cmd *cobra.Command, args []string) {
		configmap.CreateConfigMap(args[0], cmd.Flag("namespace").Value.String(),
			keyName, keyValue, loadClient())
	},
}

var cmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete configmap object",
	Args:  cobra.ExactArgs(1),
	Long:  `delete the specified ConfigMap object from the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		configmap.DeleteConfigMap(args[0], cmd.Flag("namespace").Value.String(), loadClient())
	},
}

// update
