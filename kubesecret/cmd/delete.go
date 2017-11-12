package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	RootCmd.AddCommand(deleteCmd)
}

func deleteSecret(secret string) {
	kc := loadClient()

	err := kc.CoreV1().Secrets(Namespace).Delete(secret,
		&v1.DeleteOptions{
			TypeMeta: v1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1"},
		},
	)

	if err != nil {
		log.Error("error deleting secret")
		log.Error(err)
	}
	return
}

var deleteCmd = &cobra.Command{
	Use:   "delete [secret]",
	Short: "delete secret from cluster",
	Args: cobra.ExactArgs(1),
	Long:  `delete specified secret from cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteSecret(args[0])
	},
}
