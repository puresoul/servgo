package middleware

import (
	"net/http"
	"createdata/router"
	"createdata/log"
)

func SetUp(h http.Handler) http.Handler {
	return router.ChainHandler(h, log.Log)
}