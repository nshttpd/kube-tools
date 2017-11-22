package configmap

import (
	"github.com/nshttpd/kube-tools/common"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateConfigMap(cmname string, namespace string, keyName string, keyValue string, kc *kubernetes.Clientset) {

	cm := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmname,
			Namespace: namespace,
		},
	}

	if keyName != "" {
		m := common.CreateStringDataMap(keyName, keyValue)
		if m != nil {
			cm.Data = m
		} else {
			return
		}
	}

	_, err := kc.CoreV1().ConfigMaps(namespace).Create(&cm)
	if err != nil {
		log.WithFields(log.Fields{
			"configmap": cmname,
		}).Error("error creating configmap")
		log.Error(err)
	} else {
		log.WithFields(log.Fields{
			"configmap": cmname,
		}).Info("created configmap")
	}
	return
}
