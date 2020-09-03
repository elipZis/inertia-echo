package trait

import (
	"database/sql/driver"
	"time"
)

// A time.Time wrapper for Echo and Gorm to be able to use models in both cases
type Timestamp struct {
	time.Time
}

// Try to parse some common date formats
func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	if err != nil {
		ts, err = time.Parse(time.RFC822, src)
		if err != nil {
			ts, err = time.Parse(time.UnixDate, src)
			if err != nil {
				ts, err = time.Parse(time.RFC1123, src)
				if err != nil {
					ts, err = time.Parse(time.RFC850, src)
					if err != nil {
						ts, err = time.Parse(time.ANSIC, src)
						if err != nil {
							ts, err = time.Parse(time.RFC1123Z, src)
							if err != nil {
								ts, err = time.Parse(time.RFC822Z, src)
								if err != nil {
									ts, err = time.Parse(time.RFC3339Nano, src)
								}
							}
						}
					}
				}
			}
		}
	}
	*t = Timestamp{
		ts,
	}
	return err
}

// Implement go.sql interface to allow gorm to get the time out of this
func (t Timestamp) Value() (driver.Value, error) {
	return t.Time, nil
}

// Implement go.sql interface to allow gorm to get the time out of this
func (t *Timestamp) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return t.UnmarshalParam(v)
	case time.Time:
		*t = Timestamp{
			src.(time.Time),
		}
		return nil
	default:
		return nil
	}
}
