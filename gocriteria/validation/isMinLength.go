package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria/convertor"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsMinLength(field string, params ...interface{}) (bool, error) {
	fieldLabel := field
	dataValue := params[0]
	min := params[1].(int)

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

	errMessage := strings.ReplaceAll(message.Default.IsMinLength, "{field}", fieldLabel)
	errMessage = strings.ReplaceAll(errMessage, "{value}", value)
	errMessage = strings.ReplaceAll(errMessage, "{min}", strconv.Itoa(min))
	err := errors.New(errMessage)

	if len(value) < min {
		return false, err
	}

	return true, nil
}
