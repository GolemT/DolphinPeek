package screens

import (
	"fmt"
	"runtime"
)

func GetNeofetchInfo() string {
	hostname := "golem-workshop" // Subtle GolemT reference
	username := "Tim"            // TODO: get from system

	return fmt.Sprintf(`%s@%s
-----------
OS: %s
Arch: %s
Go: %s
CPU: Intel i7 (TODO)
GPU: NVIDIA RTX (TODO)
Memory: 8GB / 16GB (TODO)
Disk: 500GB / 1TB (TODO)

    .--.
   |o_o |
   |:_/ |
  //   \ \
 (|     | )
/'\_   _/\
\___)=(___/`,
		username, hostname,
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version())
}
