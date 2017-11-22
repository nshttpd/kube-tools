package actions

import (
	"github.com/nshttpd/kube-tools/secrets"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(secretRootCmd)
	secretRootCmd.AddCommand(secretGetCmd)
	secretRootCmd.AddCommand(secretCreateCmd)
	secretRootCmd.AddCommand(secretDeleteCmd)
	secretRootCmd.AddCommand(secretUpdateCmd)
	secretCreateCmd.Flags().StringVar(&keyName, "name", "", "secret name to add on creation")
	secretCreateCmd.Flags().StringVar(&keyValue, "value", "", "secret value to add to 'name' on creation")
	secretUpdateCmd.Flags().StringVar(&keyName, "name", "", "secret name to add on creation")
	secretUpdateCmd.Flags().StringVar(&keyValue, "value", "", "secret value to add to 'name' on creation")
}

var secretRootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret CRUD level manipulation",
	Long:  `create, delete, get, update of secret objects in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var secretGetCmd = &cobra.Command{
	Use:   "get [secret]",
	Short: "get secret from cluster",
	Args:  cobra.ExactArgs(1),
	Long:  `get specified secret from cluster and save it to a JSON file in 'cwd'`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets.GetSecret(args[0], cmd.Flag("namespace").Value.String(), loadClient())
	},
}

var secretCreateCmd = &cobra.Command{
	Use:   "create [secret]",
	Args:  cobra.ExactArgs(1),
	Short: "create a secret in the cluster",
	Long:  `Create a secret in the cluster either bare or with data`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets.CreateSecret(args[0], cmd.Flag("namespace").Value.String(),
			keyName, keyValue, loadClient())
	},
}

var secretDeleteCmd = &cobra.Command{
	Use:   "delete [secret]",
	Short: "delete secret from cluster",
	Args:  cobra.ExactArgs(1),
	Long:  `delete specified secret from cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets.DeleteSecret(args[0], cmd.Flag("namespace").Value.String(), loadClient())
	},
}

var secretUpdateCmd = &cobra.Command{
	Use:   "update [secret]",
	Args:  cobra.ExactArgs(1),
	Short: "update an item in a secret",
	Long:  `Update or add an item to a secret in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		secrets.UpdateSecret(args[0], cmd.Flag("namespace").Value.String(),
			keyName, keyValue, loadClient())
	},
}
