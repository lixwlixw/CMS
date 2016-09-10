package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func getCollectionsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	JsonResult(rw, http.StatusNotFound, 1404, "uri not found", `{"name":"test"}`)
}
