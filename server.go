package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	rating         string
	url            string
	name           string
	filename       string
	platform       string
	release_date   []string
	position       int
	video_found    bool
	playlist_found bool
}

var folderName = "videos"
var resVideoName = "res.mp4"
var resVideoPath = folderName + "/" + resVideoName
var maxSize = 309715200

const YOUTUBE_API_KEY = "AIzaSyA7s6rv1gzBHF3GmRPQuG4YrKZv6D4ig-g"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func getList(rw http.ResponseWriter, req *http.Request) {
	responseBody, _ := ioutil.ReadAll(req.Body)
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		panic(err)
	}
	getVideos(data)
	fmt.Println(data[0]["name"])
}
func generate(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	fmt.Println("Request received")
	fmt.Println(req.Method)
	if req.Method == "POST" {
		fmt.Println("Request received")
		responseBody, _ := ioutil.ReadAll(req.Body)
		var data []map[string]interface{}
		err := json.Unmarshal([]byte(responseBody), &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
		getVideos(data)
		editVideo(data)
		var link = uploadToCdn(resVideoPath)
		fmt.Fprintf(w, link)
	}
	if req.Method == "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}
func upload(rw http.ResponseWriter, req *http.Request) {
	uploadToCdnTest()
}
func main() {
	http.HandleFunc("/videolist", getList)  // Update this line of code
	http.HandleFunc("/hello", helloHandler) // Update this line of code
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/upload", upload)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
