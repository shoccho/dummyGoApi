package middleware

import (
	"errors"
	"github.com/shoccho/goCoinApi/api"
	"github.com/shoccho/goCoinApi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var ErrUnAuthorizedError = errors.New("invalid Username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		if username == "" || token == "" {
			log.Error(ErrUnAuthorizedError)
			api.RequestErrorHandler(w, ErrUnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrUnAuthorizedError)
			api.RequestErrorHandler(w, ErrUnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})

}
