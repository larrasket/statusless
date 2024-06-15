package plugins

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const diskUsageIsActive = true
const procPartitionsFile = "/proc/partitions"

func init() {
	List = append(List, plugin{
		Getter: func() (string, error) {
			const kBToGB = 1024 * 1024

			data, err := os.ReadFile(procPartitionsFile)
			if err != nil {
				return "", err
			}

			lines := strings.Split(string(data), "\n")
			var totalSize int64

			for _, line := range lines {
				fields := strings.Fields(line)
				if len(fields) < 4 {
					continue
				}

				if !strings.HasPrefix(fields[3], "sd") &&
					!strings.HasPrefix(fields[3], "nvme") {
					continue
				}

				sizeKB, err := strconv.ParseInt(fields[2], 10, 64)
				if err != nil {
					continue
				}

				totalSize += sizeKB
			}

			totalSizeGB := float64(totalSize) / kBToGB

			return fmt.Sprintf(" %.1fG", totalSizeGB), nil
		},
		Trigger: time.Second * 120,
		Active:  diskUsageIsActive,
	})
}
