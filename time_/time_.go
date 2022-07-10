package time_

import (
	"fmt"
	"time"
)

func Main() {
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2022, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(future.Format(time.Layout))
	fmt.Println(future.Format(time.ANSIC))
	fmt.Println(future.Format(time.UnixDate))
	fmt.Println(future.Format(time.RubyDate))
	fmt.Println(future.Format(time.RFC822))
	fmt.Println(future.Format(time.RFC822Z))
	fmt.Println(future.Format(time.RFC850))
	fmt.Println(future.Format(time.RFC1123))
	fmt.Println(future.Format(time.RFC1123Z))
	fmt.Println(future.Format(time.RFC3339))
	fmt.Println(future.Format(time.RFC3339Nano))
	fmt.Println(future.Format(time.Kitchen))
}
