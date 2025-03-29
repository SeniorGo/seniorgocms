package api

import (
	"context"
	"github.com/SeniorGo/seniorgocms/logger"
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
		logger.GetLog(ctx).Info("Access",
			"ip", ip,
			"method", r.Method,
			"url", r.URL.String(),
			"action", action)
		next(ctx)
	}
}
