package plugins

import (
	"fmt"
	"syscall"
	"time"
)

const diskUsageIsActive = true
const procPartitionsFile = "/proc/partitions"

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			var stat syscall.Statfs_t

			// Use the root directory ("/") to get the file system statistics
			err := syscall.Statfs("/", &stat)
			if err != nil {
				return "", err
			}

			// Calculate free space in gigabytes
			freeSpaceGB := float64(stat.Bavail*uint64(stat.Bsize)) / (1024 * 1024 * 1024)

			// Format the free space as a string with one decimal place

			return fmt.Sprintf("   %.1fG", freeSpaceGB), nil
		},
		Span:   time.Second * 120,
		Active: diskUsageIsActive,
		Order:  8,
	})
}
