package speedtest

import (
	"errors"
	"fmt"

	"github.com/tvolkov/speedtest/fastcom"
	"github.com/tvolkov/speedtest/speedtestnet"
)

type SpeedTestResult struct {
	provider string
	download int
	upload   int
}

type TestSpeedFunc func() (int, int, error)

const PROVIDER_SPEEDTEST string = "speedtest"
const PROVIDER_FASTCOM string = "fastcom"

func Speedtest(providerName string) (SpeedTestResult, error) {

	if providerName == "" {
		return SpeedTestResult{}, errors.New("empty provider name")
	}

	providerMap := map[string]TestSpeedFunc{PROVIDER_SPEEDTEST: TestSpeedFunc(speedtestnet.TestSpeed), PROVIDER_FASTCOM: TestSpeedFunc(fastcom.TestSpeed)}

	a, b, e := providerMap[providerName]()
	fmt.Println(e)
	return SpeedTestResult{provider: "test", download: a, upload: b}, nil
}
