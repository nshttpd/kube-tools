package cmd

import (
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	keyName  string
	keyValue string
)

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&keyName, "name", "", "secret name to add on creation")
	createCmd.Flags().StringVar(&keyValue, "value", "", "secret value to add to 'name' on creation")
}

// creates a map that is the Secret key, value. When applied to the Secret
// as v1.Secret.StringData and not v1.Secret.Data, the API server handles
// the base64 encoding of the value data. This can only be used for non-binary
// secret values.
func createStringDataMap() map[string]string {
	// if the value starts with an @ it means read the data from a file
	if strings.HasPrefix(keyValue, "@") {
		// read it and reset the value to the actual file contents
		if f, err := ioutil.ReadFile(keyValue[1:]); err != nil {
			log.WithFields(log.Fields{
				"keyName":  keyName,
				"keyValue": keyValue,
			}).Error("error reading file for secret")
			log.Error(err)
			return nil
		} else {
			keyValue = string(f[:])
		}
	}
	// add the actual data to the secret
	m := map[string]string{
		keyName: keyValue,
	}

	return m
}

func createSecret(secret string) {
	kc := loadClient()

	s := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret,
			Namespace: Namespace,
		},
		Type: "Opaque",
	}

	// we've got a KeyName so we're going to create the
	// Secret with some data.
	if keyName != "" {
		m := createStringDataMap()
		if m != nil {
			s.StringData = m
		} else {
			return
		}
	}

	_, err := kc.CoreV1().Secrets(Namespace).Create(&s)
	if err != nil {
		log.WithFields(log.Fields{
			"secret": secret,
		}).Error("error creating secret")
	} else {
		log.WithFields(log.Fields{
			"secret": secret,
		}).Info("created secret")
	}
	return
}

var createCmd = &cobra.Command{
	Use:   "create [secret]",
	Args: cobra.ExactArgs(1),
	Short: "create a secret in the cluster",
	Long:  `Create a secret in the cluster either bare or with data`,
	Run: func(cmd *cobra.Command, args []string) {
		createSecret(args[0])
	},
}
