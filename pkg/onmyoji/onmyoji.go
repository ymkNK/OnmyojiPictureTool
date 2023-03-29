package onmyoji

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	onmyojiMediaURL  = "https://yys.163.com/media/picture.html"
	onmyojiPicPrefix = "https://yys.res.netease.com/pc/zt/"
)

var (
	picRegex, _          = regexp.Compile(".*yys.res.netease.com.*")
	onmyojiMediaRegex, _ = regexp.Compile(".*yys.163.com.*")
)

type LinkInfo struct {
	X    int
	Y    int
	Link string
}

func (i *LinkInfo) Compute() int {
	return i.X * i.Y
}

func RefreshOnmyojiMedia(existedFileMap map[string]bool, downloadPath string) (err error) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains
		colly.URLFilters(picRegex, onmyojiMediaRegex),
	)

	picLinkMap := make(map[string]*LinkInfo)
	picNameMap := make(map[string]bool)

	var newPicCount, totalPicCount = 0, 0

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if !picRegex.MatchString(link) {
			return
		}

		name, x, y := getLinkInfo(link)
		if !picNameMap[name] {
			picNameMap[name] = true

			totalPicCount++
		}

		if existedFileMap[name] {
			return
		}

		tempLink := LinkInfo{
			X:    x,
			Y:    y,
			Link: link,
		}

		if picLinkMap[name] == nil {
			picLinkMap[name] = &tempLink

			newPicCount++
			fmt.Printf("New Pic found: %q -> %s\n", e.Text, link)
		}

		if tempLink.Compute() > picLinkMap[name].Compute() {
			picLinkMap[name] = &tempLink

			fmt.Printf("Pic update : %s -> %s\n", name, e.Text)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err = c.Visit(onmyojiMediaURL)

	DownloadPic(picLinkMap, downloadPath)

	fmt.Println(fmt.Sprintf("New pic count: %d, total count: %d", newPicCount, totalPicCount))

	return
}

func getLinkInfo(link string) (name string, x, y int) {
	linkSimple := strings.ReplaceAll(link, onmyojiPicPrefix, "")
	linkSimple = strings.ReplaceAll(linkSimple, ".jpg", "")
	linkSplits := strings.Split(linkSimple, "/")
	xyArray := strings.Split(linkSplits[len(linkSplits)-1], "x")

	x, _ = strconv.Atoi(xyArray[0])
	y, _ = strconv.Atoi(xyArray[1])

	linkSplits = linkSplits[:len(linkSplits)-1]

	name = strings.Join(linkSplits, "_") + ".jpg"

	return
}

func DownloadPic(newPicMap map[string]*LinkInfo, downloadPath string) {
	c := colly.NewCollector(
		// Visit only domains
		colly.URLFilters(picRegex, onmyojiMediaRegex),
		colly.Async(false),
	)

	c.OnResponse(func(response *colly.Response) {
		s := response.Request.URL.String()
		name, _, _ := getLinkInfo(s)
		SavePicture(response, name, downloadPath)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("DownloadPic visiting", r.URL.String())
	})

	for i := range newPicMap {
		time.Sleep(time.Second)

		err := c.Visit(newPicMap[i].Link)
		if err != nil {
			fmt.Println(fmt.Sprintf("visit [%s] err: %s", newPicMap[i].Link, err.Error()))

			continue
		}
	}
}

func SavePicture(response *colly.Response, name, downloadPath string) {
	fmt.Println(fmt.Sprintf("[%s/%s] is saving", downloadPath, name))

	filePath := downloadPath + "/" + name
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) // ignore_security_alert
	if err != nil {
		fmt.Println("file open failed", err)

		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("file close failed", err)
		}
	}()

	//写入文件时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	_, err = writer.Write(response.Body)
	if err != nil {
		fmt.Println("writer write failed", err)

		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("writer flush failed", err)

		return
	}

	fmt.Println(fmt.Sprintf("[%s/%s] has saved", downloadPath, name))
}
