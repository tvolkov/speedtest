package speedtestnet

import (
	"errors"

	"github.com/op/go-logging"

	"github.com/showwin/speedtest-go/speedtest"
)

var log = logging.MustGetLogger("speedtest.net")

func TestSpeed() (float64, float64, error) {
	log.Info("Testing speed using speedtest.net")
	user, err := speedtest.FetchUserInfo()

	if err != nil {
		log.Warningf("Error fetching user info: %s", err)
		return 0, 0, errors.New("Error fetching user info")
	}

	serverList, err := speedtest.FetchServerList(user)

	if err != nil {
		log.Warningf("error fetching server list: %s", err)
		return 0, 0, errors.New("error fetching server list")
	}

	targets, err := serverList.FindServer([]int{})

	if err != nil {
		log.Warningf("error finding server: %s")
		return 0, 0, errors.New("error finding server")
	}

	downloadSpeedSum := 0.0
	uploadSpeedSum := 0.0
	n := targets.Len()

	for _, s := range targets {
		err = s.PingTest()
		if err != nil {
			log.Warningf("unable to test ping")
		}

		err = s.DownloadTest(false)
		if err != nil {
			log.Warningf("unable to test download speed")
		}

		err = s.UploadTest(false)
		if err != nil {
			log.Warningf("unable to test upload speed")
		}

		log.Infof("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		downloadSpeedSum += s.DLSpeed
		uploadSpeedSum += s.ULSpeed
	}

	if downloadSpeedSum == 0 || uploadSpeedSum == 0 {
		log.Warningf("could not test connection speed, try again")
	}

	return downloadSpeedSum / float64(n), uploadSpeedSum / float64(n), nil
}
