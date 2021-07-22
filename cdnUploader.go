package main

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"log"
)

func uploadToCdn() string {
	var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams("dvmo50ocz", "661814449754765", "YCG2pplXo1wQ8eNXAqgA--RXjH4")
	resp, _ := cld.Upload.Upload(ctx, "videos/client.png", uploader.UploadParams{})
	log.Println(resp.SecureURL)
	return resp.SecureURL
}
