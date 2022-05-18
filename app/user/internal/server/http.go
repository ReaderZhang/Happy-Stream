package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "github.com/qqz/Happy-Stream/api/user/v1"
	"github.com/qqz/Happy-Stream/app/user/internal/conf"
	"github.com/qqz/Happy-Stream/app/user/internal/service"
	http2 "net/http"
	"time"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/order.v1.User/Register"] = struct{}{}
	whiteList["/order.v1.User/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//selector.Server(
			//	jwt.Server(func(token *jwt2.Token) (interface{}, error) {
			//		return []byte("qqz"), nil
			//	}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
			//		return &jwt2.MapClaims{}
			//	})),
			//).
			//	Match(NewWhiteListMatcher()).
			//	Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ResponseEncoder(func(w http2.ResponseWriter, r *http2.Request, i interface{}) error {
		type response struct {
			Code int
			Data interface{}
			Time string
		}
		reply := &response{
			Code: 200,
			Data: i,
			Time: time.Now().String(),
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(reply)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	}))
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, user)
	return srv
}
