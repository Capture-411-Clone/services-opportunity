package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria/convertor"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsEndWith(field string, dataValue interface{}, str string) (bool, error) {
	fieldLabel := field

	label, ok := (*message.FieldLabels)[field]
	if ok {
		fieldLabel = label
	}

	value := ""
	number, errConv := convertor.ToNumber(dataValue)
	if errConv == nil && number != nil {
		value = fmt.Sprintf("%v", *number)
	} else {
		value = dataValue.(string)
	}

	if strings.Index(value, str) == len(value)-1 {
		return true, nil
	}

	errMessage := strings.ReplaceAll(message.Default.IsMaxLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{str}", str)
	err := errors.New(errMessage)

	return false, err
}
