package speedtestnet

import (
	"fmt"
)

func TestSpeed() (int, int, error) {
	fmt.Println("Hello from speedtest.net gauge")
	return 2, 2, nil
}
