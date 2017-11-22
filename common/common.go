package common

import (
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateStringDataMap(keyName string, keyValue string) map[string]string {
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

	m := map[string]string{
		keyName: keyValue,
	}

	return m
}
