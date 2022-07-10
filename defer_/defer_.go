package defer_

import (
	"errors"
	"fmt"
)

func Main() {
	err := deferReturnSample()
	fmt.Printf("error=%s", err)
}

func deferReturnSample() (err error) {
	const str = "str defined in func"
	defer func() {
		err = errors.New(str)
	}()
	return
}
