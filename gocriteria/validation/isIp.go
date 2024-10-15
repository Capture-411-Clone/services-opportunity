package validation

import (
	"net"

	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsIp(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsIp, label, value)
	str := value.(string)
	isValid := net.ParseIP(str) != nil
	if isValid {
		return true, nil
	}
	return false, err
}
