package secrets

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nshttpd/kube-tools/common"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateSecret(secret string, namespace string, keyName string, keyValue string, kc *kubernetes.Clientset) {

	s := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret,
			Namespace: namespace,
		},
		Type: "Opaque",
	}

	// we've got a KeyName so we're going to create the
	// Secret with some data.
	// When applied to the Secret as v1.Secret.StringData and not v1.Secret.Data,
	// the API server handles the base64 encoding of the value data. This can
	// only be used for non-binary secret values.
	if keyName != "" {
		m := common.CreateStringDataMap(keyName, keyValue)
		if m != nil {
			s.StringData = m
		} else {
			return
		}
	}

	_, err := kc.CoreV1().Secrets(namespace).Create(&s)
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
