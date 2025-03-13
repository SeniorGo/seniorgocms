package auth

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fulldump/box"
)

const XGlueAuthentication = "X-Glue-Authentication"

type Auth struct {
	Session struct {
		ID string `json:"id"`
	} `json:"session"`
	User User `json:"user"`
}

type User struct {
	ID      string `json:"id"`
	Nick    string `json:"nick"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}

var ErrUnauthorized = errors.New("unauthorized")

// Require the user is authenticated
func Require(next box.H) box.H {
	return func(ctx context.Context) {

		d := box.GetRequest(ctx).Header.Get(XGlueAuthentication)

		if d == "" {
			box.SetError(ctx, ErrUnauthorized)
			return
		}

		a := &Auth{}

		err := json.Unmarshal([]byte(d), &a)
		if err != nil {
			box.SetError(ctx, ErrUnauthorized)
			return
		}

		ctx = SetAuth(ctx, a)

		next(ctx)
	}
}

const key = "6fbc299a-3546-11ed-bf91-87a0b0cea4af"

func SetAuth(ctx context.Context, a *Auth) context.Context {
	return context.WithValue(ctx, key, a)
}

func GetAuth(ctx context.Context) *Auth {
	v := ctx.Value(key)
	if v == nil {
		return nil
	}
	return v.(*Auth)
}
