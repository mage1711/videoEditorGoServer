package main

import (
	"errors"
	"fmt"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Clip struct {
	videosPath      []string
	concatListCache string
}

func editVideo(videoList []map[string]interface{}) {
	fmt.Println("generating")
	var videos []string
	var videoPath string
	var convetedVideoPath string
	for _, element := range videoList {
		var name = element["filename"].(string)

		videoLength := 15 * 30
		videoPath = folderName + "/" + name + ".mp4"
		convetedVideoPath = folderName + "/" + name + " converted " + strconv.Itoa(videoLength) + ".mp4"
		if _, err := os.Stat(convetedVideoPath); os.IsNotExist(err) {
			fmt.Println("converting:" + folderName)
			_ = fluentffmpeg.NewCommand("").
				InputPath(videoPath).
				OutputFormat("mp4").
				OutputPath(convetedVideoPath).VideoCodec("libx264").Preset("ultrafast").FrameRate(30).Resolution("1920x1080").VFrames(videoLength).
				Overwrite(true).Run()
			fmt.Println("converted:" + folderName)
		} else {
			fmt.Println("already converted")
		}
		videos = append(videos, convetedVideoPath)

	}
	fmt.Println(videos)

	clip, _ := NewClip(videos)
	clip.Concatenate(resVideoName)

}
func NewClip(videoPath []string) (*Clip, error) {
	var clip Clip

	for _, path := range videoPath {
		if _, err := os.Stat(path); err != nil {
			return nil, errors.New("unable to load file: " + err.Error())
		}
	}
	dir := filepath.Dir(videoPath[0])
	clip = Clip{videosPath: videoPath, concatListCache: filepath.Join(dir, "concat.txt")}
	return &clip, nil
}

func (c *Clip) saveConcatenateList() error {
	var maxLengthPerVideo = "outpoint 15"
	f, err := os.Create(c.concatListCache)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, video := range c.videosPath {
		fmt.Fprintf(f, "file '%s'\n%s\n", filepath.Base(video), maxLengthPerVideo)
	}
	return nil
}

func (c *Clip) Concatenate(output string) error {
	return c.ConcatenateWithStreams(output, nil, nil)
}
func (c *Clip) ConcatenateWithStreams(output string, os io.Writer, es io.Writer) error {
	c.saveConcatenateList()
	//defer c.deleteConcatenateList()
	line := c.CommandLine(output)
	fmt.Println(line)
	cmd := exec.Command(line[0], line[1:]...)
	cmd.Stderr = es
	cmd.Stdout = os

	err := cmd.Run()
	if err != nil {
		return errors.New("cinema.Video.Concatenate: ffmpeg failed: " + err.Error())
	}
	return nil
}

func (c *Clip) deleteConcatenateList() error {
	if err := os.Remove(c.concatListCache); err != nil {
		return err
	}
	return nil
}

func (c *Clip) CommandLine(output string) []string {
	cmdline := []string{
		"ffmpeg",
		"-y",
		"-f", "concat",
		"-safe", "0",
		"-i", c.concatListCache,
		"-c", "copy",
	}
	cmdline = append(cmdline, "-fflags", "+genpts", filepath.Join(filepath.Dir(c.videosPath[0]), output))
	return cmdline
}
