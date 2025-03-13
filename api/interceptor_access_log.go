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
		log.Println(ip, r.Method, r.URL.String(), box.GetBoxContext(ctx).Action.Name)
		next(ctx)
	}
}
