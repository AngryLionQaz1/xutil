package snow

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	log "github.com/sirupsen/logrus"
	pb "gopkg.in/cheggaaa/pb.v1"
)

var videoDownloadFloder = "videos"
var defaultSource = "1080p"

var sources = map[string]int{
	"480p":  1000,
	"720p":  2000,
	"1080p": 4000,
}

func download(id string) {
	partName := fmt.Sprintf("%s/%s.mp4", videoDownloadFloder, id)
	if _, err := os.Stat(partName); err == nil {
		log.Info("video file exists: ", id)
		return
	}

	url := fmt.Sprintf("https://c1.cdnjav.com/share/trailers/%s/1min_%d.mp4", id, sources[defaultSource])
	fmt.Println("Downloading: ", url)

	// get file size
	headResp, err := http.Head(url)
	if err != nil {
		log.Errorf("failed to get header: %s", id)
		return
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		log.Errorf("failed to get Content-Length: %s", id)
		return
	}

	// start progress bar
	bar := pb.New(size)
	bar.SetRefreshRate(time.Second)
	bar.ShowSpeed = true
	bar.SetUnits(pb.U_BYTES)
	bar.Start()

	// create get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("failed to split NewRequest for get: %s", id)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("failed to get requests: %s", id)
		return
	}
	defer res.Body.Close()

	output, err := os.OpenFile(partName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("failed to create %s", partName)
		return
	}
	defer output.Close()

	// create proxy reader
	reader := bar.NewProxyReader(res.Body)

	// and copy from pb reader
	io.Copy(output, reader)

	bar.FinishPrint("Done: " + id)
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
	// colly.Async(true),
	)

	// set 1hour for downlaod video
	c.SetRequestTimeout(3600 * time.Second)

	// Limit the number of threads started by colly to two
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referrer(c)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Error("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		r.Request.Retry()
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// On every a element which has href attribute call callback
	c.OnHTML(".thumb-content a", func(e *colly.HTMLElement) {
		id := e.Attr("clickitem")
		download(id)
	})

	urlFmt := "https://javhd.com/zh/japanese-porn-videos/justadded/all/%d"
	for i := 1; i <= 30; i++ {
		c.Visit(fmt.Sprintf(urlFmt, i))
	}

	// Wait until threads are finished
	c.Wait()
}
