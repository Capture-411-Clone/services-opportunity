package validation

import (
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsLowerCase(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsLowerCase, label, value)
	str := value.(string)
	isValid := str == strings.ToLower(str)
	if isValid {
		return true, nil
	}
	return false, err
}
