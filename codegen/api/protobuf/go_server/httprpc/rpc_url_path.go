package httprpc

import "fmt"

func URLPath(serviceName string, rpcName string) string {
	return fmt.Sprintf("%s/%s", serviceName, rpcName)
}
