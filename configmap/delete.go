package configmap

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteConfigMap(cm string, namespace string, kc *kubernetes.Clientset) {

	err := kc.CoreV1().ConfigMaps(namespace).Delete(cm, &v1.DeleteOptions{})

	if err != nil {
		log.WithFields(log.Fields{
			"configmap": cm,
		}).Error("error deleting configmap")
		log.Error(err)
	}
	return
}
