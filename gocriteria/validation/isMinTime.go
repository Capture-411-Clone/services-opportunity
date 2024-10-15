package validation

import (
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsMinTime(field string, value interface{}, min time.Time) (bool, error) {
	if value.(time.Time).Before(min) {
		return false, GetErrorMessageByFieldValue(message.Default.IsMinTime, field, value)
	}
	return true, nil
}
