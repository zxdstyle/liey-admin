package requests

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

type (
	Resource string

	Resources []Resource
)

func (res Resource) String() string {
	return string(res)
}

func parseWith(r *ghttp.Request) (resources Resources) {
	with := r.GetQuery("_with")
	if with == nil {
		return
	}

	var withs []string
	if with.IsSlice() {
		withs = with.Strings()
	} else {
		withs = strings.Split(with.String(), ",")
	}

	for _, resource := range withs {
		resources = append(resources, Resource(gstr.UcFirst(resource)))
	}
	return
}
