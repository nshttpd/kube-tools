package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&keyName, "name", "", "name of secret value to update")
	updateCmd.Flags().StringVar(&keyValue, "value", "", "data of secret value to update")
}

func updateSecret(secret string) {

	kc := loadClient()

	s, err := kc.CoreV1().Secrets(Namespace).Get(secret, metav1.GetOptions{})

	if err != nil {
		log.WithFields(log.Fields{
			"secret": secret,
		}).Error("error fetching secret")
		log.Error(err)
	}

	if keyName != "" {
		m := createStringDataMap()
		if m != nil {
			// set the StringData to what we are going to update. It will overwrite already
			// existing values if they exist.
			s.StringData = m
		} else {
			return
		}
	} else {
		log.Error("missing secret value name to update")
	}

	ns, err := kc.CoreV1().Secrets(Namespace).Update(s)
	if err != nil {
		log.WithFields(log.Fields{
			"secret": secret,
		}).Error("unable to update secret")
		log.Error(err)
	} else {
		log.WithFields(log.Fields{
			"secret": ns.Name,
		}).Debug("updated data in secret")
	}

	return
}

var updateCmd = &cobra.Command{
	Use:   "update [secret]",
	Args:  cobra.ExactArgs(1),
	Short: "update an item in a secret",
	Long:  `Update or add an item to a secret in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		updateSecret(args[0])
	},
}
