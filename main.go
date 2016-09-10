package main

import (
	log "github.com/asiainfoLDP/datahub/utils/clog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

var (
	SERVICE_IPPORT string

	MONGODB_ADDR string = "MONGODB_ADDR"
	MONGODB_PORT string = "MONGODB_PORT"

	DB_NAMESPACE_MONGO = "CMS"
	DB_NAME            = "CMS"

	//Username                     = Env("ADMIN_API_USERNAME", true)
	//Password                     = Env("ADMIN_API_USER_PASSWORD", true)

	db DB = DB{*connect()}
)

const (
	ResultOK     = 0
	ErrorMarshal = 1001
)

func init() {
	SERVICE_IPPORT = os.Getenv("SERVICE_IPPORT")
}

func main() {

	initDB()

	router := httprouter.New()

	router.GET("/", rootHandler)
	//router.GET("/debug", rootHandler)
	router.GET("/content_types", getContentTypesHandler)
	router.POST("/content_types/:content_type_id", createContentTypeHandler)
	router.PUT("/content_types/:content_type_id", updateContentTypeHandler)
	router.GET("/content_types/:content_type_id", getOneContentTypeHandler)

	router.POST("/contents/:content_id", createContentHandler)
	router.GET("/contents/:content_id", getOneContentHandler)
	router.GET("/contents", getContentsHandler)
	// router.GET("/daemon/ep/:user", userEntryPointHandler)
	// router.GET("/daemon/log/:index", getDaemonLogsHandler)
	// router.GET("/daemon/status", DaemonStatusHandler)
	// router.GET("/daemon/tags/status", getTagStatusHandler)
	// router.POST("/heartbeat", heartbeatHandler)
	// router.GET("/heartbeat/status/:user", heartbeatStatusHandler)
	router.NotFound = &mux{}
	//router.MethodNotAllowed = &mux{}

	log.Info("listening on", SERVICE_IPPORT)
	err := http.ListenAndServe(SERVICE_IPPORT, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type mux struct {
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Info("from", req.RemoteAddr, req.Method, req.URL.RequestURI(), req.Proto)

	JsonResult(w, http.StatusNotFound, 1404, "uri not found", `{"show":"nothing"}`)
}

func rootHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	for k, v := range req.Header {
		log.Infof("[%s]=[%s]\n", k, v)
	}

	JsonResult(w, http.StatusForbidden, 1403, "uri not forbidden", `{"todo":"login"}`)
}
