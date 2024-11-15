package connection

import (
	"fmt"
)

func getUrl() string {
	fmt.Println("Initializing connection package...")

	ip := GetEnvVariable("HOST_NODE_1")
	port := GetEnvVariable("PORT_NODE_2")
	url := fmt.Sprint("http:{ip}:{port}", ip, port)
	return url
}
