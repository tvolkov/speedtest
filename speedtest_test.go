package speedtest_test

import (
	"testing"

	"github.com/tvolkov/speedtest"
)

func TestFastCom(t *testing.T) {

	s, e := speedtest.Speedtest("fastcom")

	if e != nil {
		t.Fatal("Error occurred ", e)
	}

	t.Log(s)
}

func TestSpeedtestNet(t *testing.T) {
	s, e := speedtest.Speedtest("speedtest")

	if e != nil {
		t.Fatal("Error occurred ", e)
	}

	t.Log(s)
}

func BenchmarkFastCom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speedtest.Speedtest("fastcom")

	}
}

func BenchmarkSpeedtestNet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speedtest.Speedtest("speedtest")
	}
}
