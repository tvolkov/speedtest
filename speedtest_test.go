package speedtest_test

import (
	"fmt"

	"testing"

	"github.com/tvolkov/speedtest"
)

func TestFastCom(t *testing.T) {

	s, e := speedtest.Speedtest("fastcom")

	fmt.Println(s)
	fmt.Println(e)
}

func TestSpeedtestNet(t *testing.T) {
	s, e := speedtest.Speedtest("speedtest")

	fmt.Println(s)
	fmt.Println(e)
}

func BenchmarkFastCom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, e := speedtest.Speedtest("fastcom")

		fmt.Println(s)
		fmt.Println(e)
	}
}

func BenchmarkSpeedtestNet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, e := speedtest.Speedtest("speedtest")

		fmt.Println(s)
		fmt.Println(e)
	}
}
