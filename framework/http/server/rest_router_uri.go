package server

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"strings"
)

func getResourceUriIndex(name, base string) string {
	return getResourceUri(name, base)
}

func getResourceUriCreate(name, base string) string {
	return fmt.Sprintf("%s/create", getResourceUri(name, base))
}

func getResourceUriStore(name, base string) string {
	return getResourceUri(name, base)
}

func getResourceUriShow(name, base string) string {
	return fmt.Sprintf("%s/:%s", getResourceUri(name, base), base)
}

func getResourceUriEdit(name, base string) string {
	return fmt.Sprintf("%s/edit", getResourceUri(name, base))
}

func getResourceUriUpdate(name, base string) string {
	return fmt.Sprintf("%s/:%s", getResourceUri(name, base), base)
}

func getResourceUriDestroy(name, base string) string {
	return fmt.Sprintf("%s/:%s", getResourceUri(name, base), base)
}

func getResourceUri(resource, base string) string {
	if !strings.Contains(resource, ".") {
		return resource
	}
	segments := strings.Split(resource, ".")
	uri := getNestedResourceUri(segments)

	current := fmt.Sprintf("/:%s", base)
	return strings.Replace(uri, current, "", -1)
}

func getNestedResourceUri(segments []string) string {
	var uri []string
	for _, segment := range segments {
		resource := fmt.Sprintf("%s/:%s", segment, getResourceWildcard(segment))
		uri = append(uri, resource)
	}
	return strings.Join(uri, "/")
}

func getResourceWildcard(value string) string {
	val := inflection.Singular(value)
	return strings.Replace(val, "_", "_", -1)
}

func getBaseName(resource string) string {
	uri := strings.Split(resource, ".")
	return getResourceWildcard(uri[len(uri)-1])
}
