package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"peacefulhack/goPexel/app/shared/datastruct"
	"strings"
	"time"
)

func GetImages(pAuth datastruct.PexelsAuth, query, perPage, width, height, currentPage, size, directory string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", pAuth.Url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", pAuth.Api)
	q := req.URL.Query()

	q.Add("query", query)
	q.Add("per_page", perPage)
	q.Add("width", width)
	q.Add("height", height)
	q.Add("page", currentPage)

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var pexelResp datastruct.PexelsResponse
	err = json.Unmarshal(body, &pexelResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	var img string
	tinySize := strings.ToLower(size)
	switch tinySize {
	case "tiny":
		img = pexelResp.Photos[0].Src.Tiny
	case "small":
		img = pexelResp.Photos[0].Src.Small
	case "portrait":
		img = pexelResp.Photos[0].Src.Portrait
	case "medium":
		img = pexelResp.Photos[0].Src.Medium
	case "original":
		img = pexelResp.Photos[0].Src.Original
	case "landscape":
		img = pexelResp.Photos[0].Src.Landscape
	case "large":
		img = pexelResp.Photos[0].Src.Large
	case "large2x":
		img = pexelResp.Photos[0].Src.Large2X
	}

	err = extractImage(img, time.Now().Format("20060102150405-temp.jpg"), directory)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func extractImage(url, filename, directory string) error {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	file, err := os.Create(directory + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
