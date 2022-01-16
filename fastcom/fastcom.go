package fastcom

import (
	"fmt"
)

func TestSpeed() (int, int, error) {
	fmt.Println("Hello from fast.com gauge")
	return 1, 1, nil
}
