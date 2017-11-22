package secrets

import (
	"github.com/nshttpd/kube-tools/common"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateSecret(secret string, namespace string, keyName string, keyValue string, kc *kubernetes.Clientset) {

	s, err := kc.CoreV1().Secrets(namespace).Get(secret, metav1.GetOptions{})

	if err != nil {
		log.WithFields(log.Fields{
			"secret": secret,
		}).Error("error fetching secret")
		log.Error(err)
		return
	}

	if keyName != "" {
		m := common.CreateStringDataMap(keyName, keyValue)
		if m != nil {
			// set the StringData to what we are going to update. It will overwrite already
			// existing values if they exist.
			s.StringData = m
		} else {
			return
		}
	} else {
		log.Error("missing secret value name to update")
		return
	}

	ns, err := kc.CoreV1().Secrets(namespace).Update(s)
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
