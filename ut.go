package main

import (
	"encoding/json"
	log "github.com/asiainfoLDP/datahub/utils/clog"
	"net/http"
	"strconv"
)

func JsonResult(w http.ResponseWriter, statusCode int, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := Result{Code: code, Msg: msg, Data: data}
	jsondata, err := json.Marshal(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getJsonBuildingErrorJson()))
	} else {
		w.WriteHeader(statusCode)
		w.Write(jsondata)
	}
}

func getJsonBuildingErrorJson() []byte {

	return []byte(log.Infof(`{"code": %d, "msg": %s}`, ErrorMarshal, "Json building error"))

}

type QueryListResult struct {
	Total   int64       `json:"total"`
	Results interface{} `json:"results"`
}

func newQueryListResult(count int64, results interface{}) *QueryListResult {
	return &QueryListResult{Total: count, Results: results}
}

func validateOffsetAndLimit(count int64, offset *int64, limit *int) {
	if *limit < 1 {
		*limit = 1
	}
	if *offset >= count {
		*offset = count - int64(*limit)
	}
	if *offset < 0 {
		*offset = 0
	}
	if *offset+int64(*limit) > count {
		*limit = int(count - *offset)
	}
}

func optionalOffsetAndSize(r *http.Request, defaultSize int64, minSize int64, maxSize int64) (int64, int) {
	size := optionalIntParamInQuery(r, "size", defaultSize)
	if size == -1 {
		return 0, -1
	}
	page := optionalIntParamInQuery(r, "page", 0)
	if page < 1 {
		page = 1
	}
	page -= 1

	if minSize < 1 {
		minSize = 1
	}
	if maxSize < 1 {
		maxSize = 1
	}
	if minSize > maxSize {
		minSize, maxSize = maxSize, minSize
	}

	if size < minSize {
		size = minSize
	} else if size > maxSize {
		size = maxSize
	}

	return page * size, int(size)
}

func optionalIntParamInQuery(r *http.Request, paramName string, defaultInt int64) int64 {
	if r.Form.Get(paramName) == "" {
		log.Debug("paramName nil", paramName, r.Form)
		return defaultInt
	}

	i, err := strconv.ParseInt(r.Form.Get(paramName), 10, 64)
	if err != nil {
		log.Debug("ParseInt", err)
		return defaultInt
	} else {
		return i
	}
}

func get(err error) {
	if err != nil {
		log.Error(err)
	}
}
