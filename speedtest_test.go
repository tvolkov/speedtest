package speedtest_test

import (
	"fmt"

	"testing"

	"github.com/tvolkov/speedtest"
)

func Test(t *testing.T) {

	s, e := speedtest.Speedtest("fastcom")

	fmt.Println(s)
	fmt.Println(e)
}
