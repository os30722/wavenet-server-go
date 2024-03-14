package middleware

import (
	"context"
	"net/http"

	"github.com/hepa/wavenet/utils"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// if req.Header["Authorization"] != nil {
		// 	t := strings.Split(req.Header["Authorization"][0], " ")[1]
		// 	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 			return nil, fmt.Errorf("Illegal SigningMethod")
		// 		}

		// 		return env.AccessKey, nil
		// 	})

		// 	if err != nil {
		// 		return util.HttpError(err, "", 401)
		// 	}

		// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims := make(map[string]any)
		claims["UID"] = 1
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.GetUserContextKey(utils.ClaimsKey), claims)
		next.ServeHTTP(res, req.WithContext(ctx))
		// 	}

		// } else {
		// 	return util.HttpError(nil, "", 401)
		// }

	})

}
