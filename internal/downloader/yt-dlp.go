package downloader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"yt-go/internal/types"
)

// yt-dlp -f "bestvideo[height<=144]+bestaudio/best[height<=144]" -P "" --cookies-from-browser chromium --progress-template "{\"id\":\"%(info.id)s\",\"title\":\"%(info.title)s\",\"progress\":\"%(progress._percent_str)s\",\"speed\":\"%(progress._speed_str)s\",\"eta\":\"%(progress._eta_str)s\",\"total\":\"%(progress._total_bytes_str)s\",\"res\":\"%(info.height)s\"}" "https://www.youtube.com/watch?v=VQRLujxTm3c"
var videoAudioResolution string = "bestvideo[height<=144]+bestaudio/best[height<=144]"
var destinationFolder string = ""
var progressTemplate string = `download:{"id":"%(info.id)s","title":"%(info.title)s","duration":"%(info.duration)s","resolution":"%(info.width)sx%(info.height)s","format_note":"%(info.format_note)s","uploader":"%(info.uploader)s","status":"%(progress.status)s","percent":"%(progress._percent_str)s","downloaded":"%(progress.downloaded_bytes)s","total":"%(progress.total_bytes)s","speed":"%(progress._speed_str)s","eta":"%(progress._eta_str)s"}` + "\n"
var browser string = "chromium"

func DownloadVideo(urlString string) {
	var cmd = exec.Command(
		"yt-dlp",
		"-f", videoAudioResolution,
		"-P", destinationFolder,
		"--cookies-from-browser", browser,
		"--progress-template", progressTemplate,
		urlString,
	)

	cmd.Env = append(os.Environ(), "PYTHONUNBUFFERED=1")

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	var scanner *bufio.Scanner = bufio.NewScanner(stdOut)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "download:") {
			jsonPart := strings.TrimPrefix(line, "download:")
			var download types.DownloadInfo

			if err := json.Unmarshal([]byte(jsonPart), &download); err == nil {
				fmt.Printf("Progress : %+v\n", download)
			} else {
				fmt.Println("JSON parse error : ", err)
			}
		} else {
			fmt.Println("OTHER : ", line)
		}

	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Download finished with error : ", err)
	}
}
