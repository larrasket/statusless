package plugins

import (
	"net"
	"time"
)

const localIPIsActive = false
const networkInterface = "wlp1s0"

func init() {
	PList = append(PList, Plugin{
		Getter:    getLocalIP,
		Span:      time.Second * 120,
		ErrorSpan: 10 * time.Second,
		Active:    localIPIsActive,
		Order:     9,
	})
}

func getLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Name == networkInterface {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return " " + ipNet.IP.String(), nil
				}
			}
			return "  not found", nil
		}
	}

	return "  not found", nil
}
