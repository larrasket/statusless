package plugins

import (
	"net"
	"time"
)

const localIPIsActive = true
const networkInterface = "wlan0"

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {
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
		},
		Span:   time.Second * 120,
		Active: localIPIsActive,
		Order:  9,
	})
}
