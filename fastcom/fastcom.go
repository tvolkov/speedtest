package fastcom

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("fast.com")

const uploadMBSize = 10
const parallelismLevel = 10

func TestSpeed() (float64, float64, error) {
	log.Info("Testing speed using fast.com")
	url := fast.GetDlUrls(1)

	if len(url) == 0 {
		log.Warning("Unable to get download url")
		return 0, 0, errors.New("unable to get download url")
	}

	log.Info("Testing download speed")
	downloadSpeed, downloadErr := testSpeed(url[0], downloadFile)

	log.Info("Testing upload speed")
	uploadSpeed, uploadErr := testSpeed(url[0], uploadFile)

	if downloadErr != nil {
		log.Warning("Error testing download speed")
		return 0, 0, downloadErr
	}

	if uploadErr != nil {
		log.Warning("Erro testing upload speed")
		return 0, 0, uploadErr
	}

	return downloadSpeed, uploadSpeed, nil
}

func testSpeed(url string, action func(url string) (int, error)) (float64, error) {

	startTime := time.Now()

	var totalDataAmountBytes = 0

	var waitGroup sync.WaitGroup
	waitGroup.Add(parallelismLevel)

	for i := 0; i < parallelismLevel; i++ {
		currentFunc := func() {
			data, e := action(url)
			totalDataAmountBytes += data
			if e != nil {
				log.Warning(e)
			}
		}
		go func(executable func()) {
			executable()
			waitGroup.Done()
		}(currentFunc)
	}

	waitGroup.Wait()

	duration := time.Since(startTime)
	speed := float64(totalDataAmountBytes/1024/1024/10) * 8 * parallelismLevel / duration.Seconds()
	speedStr := fmt.Sprintf("%f", speed)
	r, _ := strconv.ParseFloat(string(speedStr), 64)

	return r, nil
}

func downloadFile(url string) (int, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Warning("error while executing GET for url: %s", url)
		return 0, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Warning("Received non 200 response code while downloading file")
		return 0, errors.New("received non 200 response code")
	}

	byteArray, err := io.ReadAll(response.Body)
	log.Debug("download response byte array of length ", len(byteArray))

	if err != nil {
		log.Warning("Error while downloading file, %s", err)
		return 0, err
	}

	return len(byteArray), nil
}

func uploadFile(myurl string) (int, error) {
	var payload = strings.Repeat("1", uploadMBSize*1024*1024)
	response, err := http.PostForm(myurl, url.Values{"key": {payload}})

	if err != nil {
		log.Warning("Error while uploading file: %s", err)
		return 0, err
	}

	defer response.Body.Close()

	byteArray, err := io.ReadAll(response.Body)
	log.Debug("response byte array of length ", len(byteArray))

	if err != nil {
		log.Warning("error while uploading file, %s", err)
		return 0, err
	}
	return len(payload), nil
}
