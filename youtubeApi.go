package main

import (
	"flag"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
)

var (
	maxResults = flag.Int64("max-results", 1, "Max YouTube results")
)

func getYoutubeVideoLink(fileName string) string {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: YOUTUBE_API_KEY},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	part := []string{"id", "snippet"}

	call := service.Search.List(part).
		Q(fileName + " game trailer").
		MaxResults(*maxResults)
	response, err := call.Do()

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	for id := range videos {
		log.Println(id)
		if len(id) > 0 {
			return id
			break
		}

	}
	return ""
}
func printIDs(sectionName string, matches map[string]string) {
	log.Printf("%v:\n\n", sectionName)
	for id, title := range matches {
		log.Printf("[%v] %v\n\n", id, title)
	}
	log.Printf("\n\n")
}
