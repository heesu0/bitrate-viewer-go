package ffprobe

import (
	"encoding/json"
	"os/exec"
)

func GetProbeData(videoFile string) (ProbeData, error) {
	cmd := exec.Command("ffprobe",
		"-select_streams", "v:0",
		"-show_format",
		"-show_streams",
		"-show_frames",
		"-print_format", "json",
		videoFile)

	output, err := cmd.Output()
	if err != nil {
		return ProbeData{}, err
	}

	var probeData ProbeData
	err = json.Unmarshal(output, &probeData)
	if err != nil {
		return ProbeData{}, err
	}

	return probeData, nil
}
