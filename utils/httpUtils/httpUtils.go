package httpUtils

import (
	"net/http"

	"github.com/coding-CEO/go-backend-test/utils/domainUtils"
)

func AddValidAccessControlAllowOrigin(w http.ResponseWriter, r *http.Request) {
	domain := r.Header.Get("origin")
	if !domainUtils.IsAllowedDomain(domain){ 
		return 
	}
	w.Header().Add("Access-Control-Allow-Origin", domain)
}

func AddAuthenticationRouteHeaders(w http.ResponseWriter, r *http.Request) {
	AddValidAccessControlAllowOrigin(w, r);
	w.Header().Add("Access-Control-Allow-Credentials","true")
}