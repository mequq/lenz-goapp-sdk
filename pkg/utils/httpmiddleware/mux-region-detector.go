package httpmiddleware

import (
	"context"
	"net"
	"net/http"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"
)

// region detector middleware
func (m *GorillaMuxMiddleware) RegionDetectorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isIRAN := false
		isMTNI := false

		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.Header.Get("X-Forwarded-For")
		}
		if ip == "" {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
		region, _ := m.regionDetector.GetRegion(ctx, ip)
		if region != nil {

			// check if the region is Iran
			isIRAN = region.IsIran()
			// check if the region is MTNI
			isMTNI = region.IsMTNI()
		}
		ctx = context.WithValue(ctx, appcontext.REQUEST_IS_IRAN, isIRAN)
		ctx = context.WithValue(ctx, appcontext.REQUEST_IS_MTNI, isMTNI)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
