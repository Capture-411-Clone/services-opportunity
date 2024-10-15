package faker

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
)

func SQLNow() dtp.NullTime {
	return dtp.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}
