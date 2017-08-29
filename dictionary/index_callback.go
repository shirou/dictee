package dictionary

import (
	"fmt"
	"math"
)

var callbackCalledNum = 0

func IndexCallback(current int, max int) {
	callbackCalledNum++
	if callbackCalledNum%100 == 0 {
		p := float64(current) / float64(max)
		t := int(math.Trunc(p * 100))
		if t%10 == 0 {
			fmt.Printf("%d%%,", t)
		}
	}

}
