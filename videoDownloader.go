package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func downloadVideo(fileName string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err := os.Stat(folderName + "/" + fileName + ".mp4"); os.IsNotExist(err) {
		// Get size
		size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		downloadSize := int(size)

		if downloadSize < maxSize {
			// Create the file
			out, err := os.Create(folderName + "/" + fileName + ".mp4")
			if err != nil {
				return err
			}
			defer out.Close()
			log.Println("Downloading " + fileName)
			log.Println("Downloading ", downloadSize, " bytes")

			// Write the body to file
			_, err = io.Copy(out, resp.Body)
		} else {
			log.Println(url, " is bigger than max size")
			log.Println("Downloading youtube version")

			//downloadFromYoutube(fileName)
		}
		log.Println("Finished downloading ", fileName)
		return err
	}
	//checkVideoValidity(fileName)
	log.Println(fileName + " already downloaded")
	return err
}

func downloadFromYoutube(fileName string) {
	//var videoID = getYoutubeVideoLink(fileName)
	//log.Println("yt video link",videoID)
	//client := youtube.Client{}
	//
	//video, err := client.GetVideo(videoID)
	//if err != nil {
	//	panic(err)
	//}
	//var formatPosition int
	//for index,element :=range video.Formats {
	//	if element.Quality == "hd1080"{
	//		formatPosition = index
	//	}
	//}
	//log.Println(formatPosition)
	//log.Println(video.Formats[formatPosition].Quality)
	//resp, _, _ := client.GetStream(video,&video.Formats[formatPosition])
	////resp, _, _ := client.GetStream(video, &video.Formats[formatPosition])
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Close()
	//
	//file, err := os.Create(folderName + "/" + fileName + ".mp4")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//_, err = io.Copy(file, resp)
	//if err != nil {
	//	panic(err)
	//}
}

func getVideos(videoList []map[string]interface{}) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0755)
	}
	for _, element := range videoList {
		var name = element["filename"].(string)

		if element["video_found"].(bool) {
			//var link = fmt.Sprintf("%v", element["video_url"])
			var link = element["video_url"].(string)
			downloadVideo(name, link)
			fmt.Println(name, link)
		} else {
			fmt.Println(name)
		}

		//downloadVideo(name,link)
	}

}
