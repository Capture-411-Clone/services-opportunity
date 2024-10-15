package validation

import (
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria/message"
)

func IsMaxDateTime(field string, value interface{}, max time.Time) (bool, error) {
	if value.(time.Time).After(max) {
		return false, GetErrorMessageByFieldValue(message.Default.IsMaxDateTime, field, value)
	}
	return true, nil
}
