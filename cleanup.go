package registry

import "strings"

func cleanup(name string) string {
	name = strings.Replace(name, "p3srservices.", "p3sr/services/", -1)
	name = strings.Replace(name, "cermineproto.", "p3sr/services/", -1)
	name = strings.Replace(name, ".", "/", -1)
	name = strings.TrimSuffix(name, "Service")
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}
