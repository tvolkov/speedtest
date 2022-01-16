package fastcom

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("fast.com")

const uploadMBSize = 100
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

func testSpeed(url string, action func(url string) error) (float64, error) {

	startTime := time.Now()

	var waitGroup sync.WaitGroup
	waitGroup.Add(parallelismLevel)
	defer waitGroup.Wait()

	for i := 0; i < parallelismLevel; i++ {
		currentFunc := func() {
			action(url)
		}
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(currentFunc)
	}

	duration := time.Since(startTime)
	return uploadMBSize * 8 * float64(parallelismLevel) / duration.Seconds(), nil
}

func downloadFile(url string) error {
	response, err := http.Get(url)

	if err != nil {
		log.Warning("error while executing GET for url: %s", url)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Warning("Received non 200 response code while downloading file")
		return errors.New("received non 200 response code")
	}

	byteArray, err := io.ReadAll(response.Body)
	log.Debug("download response byte array of length %d", len(byteArray))

	if err != nil {
		log.Warning("Error while downloading file, %s", err)
		return err
	}

	return nil
}

func uploadFile(myurl string) error {
	response, err := http.PostForm(myurl, url.Values{"key": {strings.Repeat("string(rand.Int31())", uploadMBSize*1024*1024)}})

	if err != nil {
		log.Warning("Error while uploading file: %s", err)
		return err
	}

	defer response.Body.Close()

	byteArray, err := io.ReadAll(response.Body)
	log.Debug("response byte array of length %d", len(byteArray))

	if err != nil {
		log.Warning("error while uploading file, %s", err)
		return err
	}
	return nil
}
