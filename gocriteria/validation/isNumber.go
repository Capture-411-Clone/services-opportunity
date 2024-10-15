package validation

import (
	"github.com/Capture-411/services-opportunity/gocriteria/convertor"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsNumber(field string, params ...interface{}) (bool, error) {
	label, value := GetFieldLabelAndValue(field, params)
	err := GetErrorMessageByFieldValue(message.Default.IsNumber, label, value)
	number, numberError := convertor.ToNumber(value)
	if numberError != nil || number == nil {
		return false, err
	}
	return true, nil
}
