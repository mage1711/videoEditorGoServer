package main

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"log"
)

func uploadToCdnTest() string {
	var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams("dvmo50ocz", "661814449754765", "YCG2pplXo1wQ8eNXAqgA--RXjH4")
	resp, _ := cld.Upload.Upload(ctx, "videos/res.mp4", uploader.UploadParams{})
	log.Println(resp.SecureURL)
	return resp.SecureURL
}
func uploadToCdn(video string) string {
	var ctx = context.Background()
	cld, _ := cloudinary.NewFromParams("dvmo50ocz", "661814449754765", "YCG2pplXo1wQ8eNXAqgA--RXjH4")
	resp, _ := cld.Upload.Upload(ctx, video, uploader.UploadParams{})
	log.Println(resp.SecureURL)
	return resp.SecureURL
}
