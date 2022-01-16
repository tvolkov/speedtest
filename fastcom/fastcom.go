package fastcom

import (
	"fmt"
)

func TestSpeed() (float64, float64, error) {
	fmt.Println("Hello from fast.com gauge")
	return 1, 1, nil
}
