package secrets

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteSecret(secret string, namespace string, kc *kubernetes.Clientset) {

	err := kc.CoreV1().Secrets(namespace).Delete(secret,
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
