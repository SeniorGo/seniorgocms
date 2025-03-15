package api

import (
	"context"
	"log"
	"net"

	"github.com/fulldump/box"
)

func InterceptorAccessLog(next box.H) box.H {
	return func(ctx context.Context) {
		r := box.GetRequest(ctx)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if forwardedIp := r.Header.Get("X-Forwarded-For"); forwardedIp != "" {
			ip = forwardedIp
		}
		action := "<no-action>"
		if boxAction := box.GetBoxContext(ctx).Action; boxAction != nil {
			action = boxAction.Name
		}
		log.Println(ip, r.Method, r.URL.String(), action)
		next(ctx)
	}
}
