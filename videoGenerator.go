package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Clip struct {
	videosPath      []string
	concatListCache string
}

func editVideo() {
	//buf := &bytes.Buffer{}
	//fmt.Println("generating")
	//_ = fluentffmpeg.NewCommand("").
	//	InputPath("videos/test.mp4").
	//	OutputFormat("avi").
	//	OutputPath("videos/r.avi").
	//	Overwrite(true).
	//	Run()
	//fmt.Println("done")
	//out, _ := ioutil.ReadAll(buf) // read logs
	//fmt.Println(string(out))
	fmt.Println("generating")
	var videos []string
	videos = append(videos, "videos/test.mp4")
	videos = append(videos, "videos/test2.mp4")

	clip, _ := NewClip(videos)
	//clip.saveConcatenateList()
	clip.Concatenate("concat.mp4")

}
func NewClip(videoPath []string) (*Clip, error) {
	var clip Clip

	for _, path := range videoPath {
		if _, err := os.Stat(path); err != nil {
			return nil, errors.New("cinema.Load: unable to load file: " + err.Error())
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
		fmt.Fprintf(f, "file '%s'\n%s\n", filepath.Join(video), maxLengthPerVideo)
	}
	return nil
}

func (c *Clip) Concatenate(output string) error {
	return c.ConcatenateWithStreams(output, nil, nil)
}
func (c *Clip) ConcatenateWithStreams(output string, os io.Writer, es io.Writer) error {
	c.saveConcatenateList()
	defer c.deleteConcatenateList()
	line := c.CommandLine(output)
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
