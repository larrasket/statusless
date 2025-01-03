package plugins

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const weatherIsActive = false
const wttrInURL = "https://wttr.in/?format=%l:+%t"

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			resp, err := http.Get(wttrInURL)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return "", fmt.Errorf("weather request failed with status code %d", resp.StatusCode)
			}

			var weather string
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			weather = string(data)

			return "  " + weather, nil
		},
		Span:      time.Hour * 2,
		ErrorSpan: 10 * time.Second,
		Active:    weatherIsActive,
		Order:     10,
	})
}
