package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	RootCmd.AddCommand(getCmd)
}

func getSecret(secret string) {
	kc := loadClient()

	s, err := kc.CoreV1().Secrets(Namespace).Get(secret, metav1.GetOptions{})
	if err != nil {
		log.Error("error getting secret from cluster")
		log.Error(err)
	} else {
		if b, err := json.Marshal(s); err != nil {
			log.Error("error msarshalling secret")
			log.Error(err)
			return
		} else {
			fn := fmt.Sprintf("%s.json", secret)
			err := ioutil.WriteFile(fn, b, 0644)
			if err != nil {
				log.Error("error writing secret file")
				log.Error(err)
			}
			return
		}
	}
}

var getCmd = &cobra.Command{
	Use:   "get [secret]",
	Short: "get secret from cluster",
	Args:  cobra.ExactArgs(1),
	Long:  `get specified secret from cluster and save it to a JSON file in 'cwd'`,
	Run: func(cmd *cobra.Command, args []string) {
		getSecret(args[0])
	},
}
