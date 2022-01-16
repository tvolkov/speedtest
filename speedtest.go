package speedtest

import (
	"errors"
	"fmt"

	"github.com/tvolkov/speedtest/fastcom"
	"github.com/tvolkov/speedtest/speedtestnet"
)

type SpeedTestResult struct {
	provider string
	download float64
	upload   float64
}

type TestSpeedFunc func() (float64, float64, error)

const PROVIDER_SPEEDTEST string = "speedtest"
const PROVIDER_FASTCOM string = "fastcom"

/*
    this is the entrypoint for the library, it accepts a testing provider name as a string (either of 'speedtest' or 'fastcom')
	and returns SpeedTestResult object which
	For the sake of simplicity SpeedTestResult contains calculated download/upload speeds as float64 values which denote Mbps
*/

func Speedtest(providerName string) (SpeedTestResult, error) {

	if providerName == "" {
		return SpeedTestResult{}, errors.New("empty provider name")
	}

	providerMap := map[string]TestSpeedFunc{PROVIDER_SPEEDTEST: TestSpeedFunc(speedtestnet.TestSpeed),
		PROVIDER_FASTCOM: TestSpeedFunc(fastcom.TestSpeed)}

	provider := providerMap[providerName]

	if provider == nil {
		return SpeedTestResult{}, errors.New("Testing provider not found for name " + providerName)
	}

	a, b, e := providerMap[providerName]()
	fmt.Println(e)
	return SpeedTestResult{provider: providerName, download: a, upload: b}, nil
}
