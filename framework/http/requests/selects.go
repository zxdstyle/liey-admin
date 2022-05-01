package requests

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

type Selects []string

func parseSelects(r *ghttp.Request) (selects Selects) {
	query := r.GetQuery("_selects")
	if query == nil {
		return
	}

	if query.IsSlice() {
		selects = query.Strings()
	} else {
		selects = strings.Split(query.String(), ",")
	}

	return
}
