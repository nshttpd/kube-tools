package secrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetSecret(secret string, namespace string, kc *kubernetes.Clientset) {

	s, err := kc.CoreV1().Secrets(namespace).Get(secret, metav1.GetOptions{})
	if err != nil {
		log.Error("error getting secret from cluster")
		log.Error(err)
	} else {
		if b, err := json.Marshal(s); err != nil {
			log.Error("error marshalling secret")
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
