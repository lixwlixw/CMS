package main

import (
	log "github.com/asiainfoLDP/datahub/utils/clog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func createContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("create a content")

	req.ParseForm()

	ctID := ps.ByName("content_type_id")
	log.Info(ctID)

	JsonResult(w, http.StatusOK, ResultOK, "OK", nil)
}

func updateContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("get a content type")

	req.ParseForm()

	ctID := ps.ByName("content_type_id")
	log.Info(ctID)

	JsonResult(w, http.StatusOK, ResultOK, "OK", nil)
}

func getContentTypesHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	/*var fields []ModelField = []ModelField{
		ModelField{
			Name: "",
			Id:   "",
		},
		ModelField{},
	}*/
	log.Info("get all content types")
	var cts []ContentType = []ContentType{
		ContentType{
			Name:        "汽车信息",
			Id:          "car_info",
			Description: "汽车的型号、动力、配置等参数",
			UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
			//FieldsCount: 2,
			//Fields:      fields,
		},
		ContentType{
			Name:        "房产信息",
			Id:          "house_info",
			Description: "各城市的楼市信息",
			UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
		}}

	JsonResult(w, http.StatusOK, ResultOK, "OK", newQueryListResult(2, cts))
}

func getOneContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()

	ctID := ps.ByName("content_type_id")
	log.Info(ctID)

	var fields = []ModelField{
		ModelField{
			Name: "型号",
			Id:   "model",
			Type: "string",
		},
		ModelField{
			Name: "变速器",
			Id:   "transmission",
			Type: "string",
		},
		ModelField{
			Name: "排量",
			Id:   "displacement",
			Type: "string",
		},
	}
	var ct = ContentType{
		Name:        "汽车信息",
		Id:          ctID,
		Description: "汽车的型号、动力、配置等参数",
		UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Ct:          time.Now(),
		FieldsCount: 2,
		Fields:      fields,
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", ct)
}

func createContentHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()

	contentID := ps.ByName("content_id")
	log.Info(contentID)

	var fieldvalues []FieldValue = []FieldValue{
		FieldValue{
			Name:  "?",
			Id:    "model",
			Value: "奔驰C200L",
			Type:  "string",
		},
		{
			Name:  "?",
			Id:    "transmission",
			Value: "自动挡",
			Type:  "string",
		},
		{
			Name:  "?",
			Id:    "displacement",
			Value: "2.0L",
			Type:  "string",
		},
	}
	var con Content = Content{
		ContentId:     contentID,
		ContentTypeId: "car_info",
		Name:          "奔驰C200L信息",
		Description:   "2016款奔驰C200L",
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
		Ct:            time.Now(),
		CreateUser:    "yuanwm@asiainfo.com",
		FieldsValue:   fieldvalues,
	}

	log.Info(con)
	JsonResult(w, http.StatusOK, ResultOK, "OK", nil)
}

func getOneContentHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()

	contentID := ps.ByName("content_id")
	log.Info("contentID:", contentID)

	var fieldvalues []FieldValue = []FieldValue{
		FieldValue{
			Name:  "型号",
			Id:    "model",
			Value: "奔驰C200L",
			Type:  "string",
		},
		{
			Name:  "变速器",
			Id:    "transmission",
			Value: "自动挡",
			Type:  "string",
		},
		{
			Name:  "排量",
			Id:    "displacement",
			Value: "2.0L",
			Type:  "string",
		},
	}
	var con Content = Content{
		ContentId:     contentID,
		ContentTypeId: "car_info",
		Name:          "奔驰C200L信息",
		Description:   "2016款奔驰C200L",
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
		Ct:            time.Now(),
		CreateUser:    "yuanwm@asiainfo.com",
		FieldsValue:   fieldvalues,
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", con)
}

func getContentsHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("get contents")
	var cs = []Content{
		Content{
			ContentId:     "BenChiC200L",
			ContentTypeId: "car_info",
			Name:          "奔驰C200L信息",
			UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
			CreateUser:    "me",
		},
		Content{
			ContentId:     "LAND_ROVER",
			ContentTypeId: "car_info",
			Name:          "路虎揽胜",
			UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
			CreateUser:    "me",
		},
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", newQueryListResult(2, cs))
}
