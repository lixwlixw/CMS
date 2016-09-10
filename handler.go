package main

import (
	"encoding/json"
	log "github.com/asiainfoLDP/datahub/utils/clog"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CMD_INC      = "$inc"
	CMD_ADDTOSET = "$addToSet"
	CMD_SET      = "$set"
	CMD_UNSET    = "$unset"
	CMD_IN       = "$in"
	CMD_OR       = "$or"
	CMD_REGEX    = "$regex"
	CMD_OPTION   = "$options"
	CMD_AND      = "$and"
	CMD_CASE_ALL = "$i"
	CMD_PULL     = "$pull"
	CMD_NOTEQUAL = "$ne"
)

func createContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("create a content")
	dbcopy := db.copy()
	defer dbcopy.Close()

	req.ParseForm()

	ctID := ps.ByName("content_type_id")
	log.Info(ctID)

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error("read request body err:", err)
	}
	log.Printf("%s\n", string(body))

	contentType := new(ContentType)
	if len(body) == 0 {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, err.Error(), nil)
		return
	}
	if err := json.Unmarshal(body, &contentType); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, err.Error(), nil)
		return
	}

	log.Printf("%v\n", contentType)

	now := time.Now()

	contentType.UpdateTime = now.Format("2006-01-02 15:04:05")
	contentType.Ct = now
	contentType.Id = ctID

	if err := db.DB(DB_NAME).C(C_CONTENT_TYPE).Insert(contentType); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorDataBase, err.Error(), nil)
		return
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", nil)
}

func updateContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("get a content type")
	dbcopy := db.copy()
	defer dbcopy.Close()

	req.ParseForm()

	ctID := ps.ByName("content_type_id")
	log.Info(ctID)

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error("read request body err:", err)
	}
	log.Printf("%s\n", string(body))

	if len(body) == 0 {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, SErrorBadReq, nil)
		return
	}

	sliceFields := []ModelField{}
	fields := FieldsPara{Fields: &sliceFields}
	if err := json.Unmarshal(body, &fields); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, err.Error(), nil)
		return
	}
	log.Info(sliceFields)
	lenth := len(sliceFields)
	uptime := time.Now().Format("2006-01-02 15:04:05")

	Q := bson.M{COL_ID: ctID}
	U := bson.M{CMD_SET: bson.M{COL_FIELDS: sliceFields, COL_FIELDSCOUNT: lenth, COL_UPDATETIME: uptime}}

	if err := dbcopy.DB(DB_NAME).C(C_CONTENT_TYPE).Update(Q, U); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorDataBase, err.Error(), nil)
		return
	}

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
	dbcopy := db.copy()
	defer dbcopy.Close()

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
	log.Debug(cts)

	var contenttypes = []ContentType{}
	sort := "-rank"
	get := bson.M{COL_NAME: "1", COL_ID: "1", COL_DESCRIPTION: "1", COL_UPDATETIME: "1"}

	err := dbcopy.DB(DB_NAME).C(C_CONTENT_TYPE).Find(nil).Sort(sort).Select(get).All(&contenttypes)
	if err != nil {
		JsonResult(w, http.StatusInternalServerError, ErrorDataBase, err.Error(), nil)
		return
	}

	//err = db.DB(DB_NAME).C(C_DATAITEM).Find(query).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&contents)

	JsonResult(w, http.StatusOK, ResultOK, "OK", newQueryListResult(int64(len(contenttypes)), contenttypes))
}

func getOneContentTypeHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	dbcopy := db.copy()
	defer dbcopy.Close()

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
	log.Debug(ct)

	var contenttype = ContentType{}

	Q := bson.M{COL_ID: ctID}

	err := dbcopy.DB(DB_NAME).C(C_CONTENT_TYPE).Find(Q).One(&contenttype)
	if err != nil {
		JsonResult(w, http.StatusNotFound, ErrorDataBase, err.Error(), nil)
		return
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", contenttype)
}

func createContentHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	dbcopy := db.copy()
	defer dbcopy.Close()

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

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Error("read request body err:", err)
	}
	log.Printf("%s\n", string(body))

	content := Content{}
	if len(body) == 0 {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, SErrorBadReq, nil)
		return
	}
	if err := json.Unmarshal(body, &content); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorBadReq, err.Error(), nil)
		return
	}

	now := time.Now()

	content.UpdateTime = now.Format("2006-01-02 15:04:05")
	content.Ct = now
	content.ContentId = contentID
	content.CreateUser = "me"

	log.Printf("%v\n", content)

	if err := db.DB(DB_NAME).C(C_CONTENT).Insert(content); err != nil {
		JsonResult(w, http.StatusBadRequest, ErrorDataBase, err.Error(), nil)
		return
	}
	JsonResult(w, http.StatusOK, ResultOK, "OK", nil)
}

func getOneContentHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	dbcopy := db.copy()
	defer dbcopy.Close()

	contentID := ps.ByName("content_id")
	log.Info("contentID:", contentID)

	// var fieldvalues []FieldValue = []FieldValue{
	// 	FieldValue{
	// 		Name:  "型号",
	// 		Id:    "model",
	// 		Value: "奔驰C200L",
	// 		Type:  "string",
	// 	},
	// 	{
	// 		Name:  "变速器",
	// 		Id:    "transmission",
	// 		Value: "自动挡",
	// 		Type:  "string",
	// 	},
	// 	{
	// 		Name:  "排量",
	// 		Id:    "displacement",
	// 		Value: "2.0L",
	// 		Type:  "string",
	// 	},
	// }
	// var con Content = Content{
	// 	ContentId:     contentID,
	// 	ContentTypeId: "car_info",
	// 	Name:          "奔驰C200L信息",
	// 	Description:   "2016款奔驰C200L",
	// 	UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
	// 	Ct:            time.Now(),
	// 	CreateUser:    "yuanwm@asiainfo.com",
	// 	FieldsValue:   fieldvalues,
	// }
	// log.Debug(con)

	var content = Content{}

	Q := bson.M{COL_CONTENT_ID: contentID}

	err := dbcopy.DB(DB_NAME).C(C_CONTENT).Find(Q).One(&content)
	if err != nil {
		JsonResult(w, http.StatusNotFound, ErrorDataBase, err.Error(), nil)
		return
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", content)
}

func getContentsHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("get contents")
	dbcopy := db.copy()
	defer dbcopy.Close()
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
	log.Debug(cs)

	var contents = []Content{}
	sort := "-rank"
	get := bson.M{COL_CONTENT_ID: "1", COL_CONTENT_TYPE_ID: "1", COL_NAME: "1", COL_DESCRIPTION: "1", COL_UPDATETIME: "1", COL_CREATEUSER: "1"}

	err := dbcopy.DB(DB_NAME).C(C_CONTENT).Find(nil).Sort(sort).Select(get).All(&contents)
	if err != nil {
		JsonResult(w, http.StatusInternalServerError, ErrorDataBase, err.Error(), nil)
		return
	}

	JsonResult(w, http.StatusOK, ResultOK, "OK", newQueryListResult(int64(len(contents)), contents))
}
