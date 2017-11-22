package configmap

import (
	"github.com/nshttpd/kube-tools/common"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateConfigMap(cmname string, namespace string, keyName string, keyValue string, kc *kubernetes.Clientset) {

	cm, err := kc.CoreV1().ConfigMaps(namespace).Get(cmname, metav1.GetOptions{})

	if err != nil {
		log.WithFields(log.Fields{
			"configmap": cmname,
		}).Error("error fetching configmap")
		log.Error(err)
		return
	}

	if keyName != "" {
		m := common.CreateStringDataMap(keyName, keyValue)
		if m != nil {
			if cm.Data == nil {
				cm.Data = m
			} else {
				cm.Data[keyName] = m[keyName]
			}
		}
	} else {
		log.Error("missing configmap value name to update")
		return
	}

	ncm, err := kc.CoreV1().ConfigMaps(namespace).Update(cm)
	if err != nil {
		log.WithFields(log.Fields{
			"configmap": cmname,
		}).Error("unable to update configmap")
		log.Error(err)
	} else {
		log.WithFields(log.Fields{
			"configmap": ncm.Name,
			"keyName":   keyName,
		}).Debug("update data in configmap")
	}

	return
}
