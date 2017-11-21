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
	getCmd.Flags().StringVar(&keyName, "name", "", "name of configmap item to fetch")
}

func getConfigMap(cm string) {
	kc := loadClient()

	configMap, err := kc.CoreV1().ConfigMaps(Namespace).Get(cm, metav1.GetOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"configmap": cm,
		}).Error("error getting config map from cluster")
		log.Error(err)
	} else {
		if keyName == "" {
			log.WithFields(log.Fields{
				"configmap": cm,
			}).Info("saving configmap")
			if b, err := json.Marshal(configMap); err != nil {
				log.Error("error marshalling config map")
				log.Error(err)
			} else {
				fn := fmt.Sprintf("%s.json", cm)
				err := ioutil.WriteFile(fn, b, 0644)
				if err != nil {
					log.Error("eror writing configmap file")
					log.Error(err)
				}
			}
		} else {
			if d, ok := configMap.Data[keyName]; ok {
				err := ioutil.WriteFile(keyName, []byte(d), 0664)
				if err != nil {
					log.WithFields(log.Fields{
						"file": keyName,
					}).Error("error writing file")
					log.Error(err)
				}
			} else {
				log.WithFields(log.Fields{
					"item": keyName,
				}).Error("config map item does not exist")
			}
		}
	}
}

var getCmd = &cobra.Command{
	Use:   "get [configmap]",
	Short: "get configmap or configmap data from cluster",
	Args:  cobra.ExactArgs(1),
	Long: `get specified config map from cluster and save it as the same name in 'cwd'
if -name option specified get that specific piece of configmap data.`,
	Run: func(cmd *cobra.Command, args []string) {
		getConfigMap(args[0])
	},
}
