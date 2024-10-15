package validation

import (
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsMinDate(field string, value interface{}, min time.Time) (bool, error) {
	if value.(time.Time).Before(min) {
		return false, GetErrorMessageByFieldValue(message.Default.IsMinDate, field, value)
	}
	return true, nil
}
