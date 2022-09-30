package helpers

import "time"

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("01/02/2006", stringDate)
	return err == nil
}
