package plugins

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const brighntessIsActive = true
const brightnessFile = "/sys/class/backlight/intel_backlight/brightness"
const maxBrightnessFile = "/sys/class/backlight/intel_backlight/max_brightness"

func init() {
	List = append(List, plugin{
		Getter: func() (string, error) {

			brightnessData, err := func() ([]byte, error) {
				f, err := os.Open(brightnessFile)
				if err != nil {
					return nil, err
				}
				defer f.Close()
				var size int
				if info, err := f.Stat(); err == nil {
					size64 := info.Size()
					if int64(int(size64)) == size64 {
						size = int(size64)
					}
				}
				size++
				if size < 512 {
					size = 512
				}
				data := make([]byte, 0, size)
				for {
					n, err := f.Read(data[len(data):cap(data)])
					data = data[:len(data)+n]
					if err != nil {
						if err == io.EOF {
							err = nil
						}
						return data, err
					}
					if len(data) >= cap(data) {
						d := append(data[:cap(data)], 0)
						data = d[:len(data)]
					}
				}
			}()
			if err != nil {
				return "", err
			}

			maxBrightnessData, err := os.ReadFile(maxBrightnessFile)
			if err != nil {
				return "", err
			}

			brightnessStr := strings.TrimSpace(string(brightnessData))
			maxBrightnessStr := strings.TrimSpace(string(maxBrightnessData))

			brightness, err := strconv.Atoi(brightnessStr)
			if err != nil {
				return "", err
			}

			maxBrightness, err := strconv.Atoi(maxBrightnessStr)
			if err != nil {
				return "", err
			}

			brightnessPercentage := (float64(brightness) / float64(maxBrightness)) * 100
			return fmt.Sprintf(" %.2f%%", brightnessPercentage), nil
		},
		Trigger: time.Millisecond * 180,
		Active:  brighntessIsActive,
	})
}
