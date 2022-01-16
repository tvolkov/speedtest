package speedtestnet

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func TestSpeed() (float64, float64, error) {
	fmt.Println("Hello from speedtest.net gauge")
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	downloadSpeedSum := 0.0
	uploadSpeedSum := 0.0
	n := targets.Len()

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		downloadSpeedSum += s.DLSpeed
		uploadSpeedSum += s.ULSpeed
	}

	return downloadSpeedSum / float64(n), uploadSpeedSum / float64(n), nil
}
