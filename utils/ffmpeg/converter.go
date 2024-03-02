package ffmpeg

import (
	"log"
	"os"
	"os/exec"
)

func EncodeAudioFile(fileName string, extension string) {
	outDir := "./assets/" + fileName
	err := os.MkdirAll(outDir, os.ModeDir)
	if err != nil {
		log.Panic("dedede  \n", err)
	}
	cmd := exec.Command("ffmpeg", "-loglevel", "quiet", "-i", "./temp/"+fileName+"."+extension, "-map",
		"0", "-map", "0", "-b:a:0", "128k", "-b:a:1", "32k", "-f", "dash", outDir+"/out.mpd")
	if err := cmd.Run(); err != nil {
		log.Panic("dededdedededede33e  \n", err)
	}

}
