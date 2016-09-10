package main

import (
	"fmt"
	log "github.com/asiainfoLDP/datahub/utils/clog"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type ResultPages struct {
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
}

type Collections struct {
	Name        int `json:name`
	Description int `json:description`
	FieldsCount int `json:fieldscount`
}

type ContentType struct {
	Name        string       `json:"name"`
	Id          string       `json:"id"`
	Description string       `json:"description"`
	UpdateTime  string       `json:"updatetime, omitempty"`
	Ct          time.Time    `json:"-"`
	FieldsCount int          `json:"fieldscount, omitempty"`
	Fields      []ModelField `json:"fields, omitempty"`
}

type ModelField struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Content struct {
	ContentId     string       `json:"content_id"`
	ContentTypeId string       `json:"content_type_id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	UpdateTime    string       `json:"updatetime, omitempty"`
	Ct            time.Time    `json:"-"`
	CreateUser    string       `json:"createuser, omitempty"`
	FieldsValue   []FieldValue `json:"fieldsvalue, omitempty"`
}

type FieldValue struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Value string `json:"valued"`
	Type  string `json:"type"`
}

type DB struct {
	mgo.Session
}

func (db *DB) copy() *DB {
	return &DB{*db.Copy()}
}

func initDB() bool {
	var err error
	for i := 0; i < 3; i++ {
		ip, port := getMgoAddr()
		url := fmt.Sprintf(`%s:%s/CMS?maxPoolSize=500`, ip, port)
		if _, err = mgo.Dial(url); err != nil {
			time.Sleep(time.Second * 10)
			continue
		} else {
			break
		}
	}
	if err != nil {
		return false
	}

	return true
}

func refreshDB() {

	for {
		select {
		case <-time.After(time.Second * 5):
			if err := db.Ping(); err != nil {
				log.Infof("%s db connect err %s", time.Now().Format("2006-01-02 15:04:05"), err)
				db = DB{*connect()}
				db.Refresh()
			}
		}
	}
}

func connect() *mgo.Session {
	ip, port := getMgoAddr()
	if ip == "" {
		log.Error("can not init mongo ip")
	}

	if port == "" {
		log.Error("can not init mongo port")
	}

	url := fmt.Sprintf(`%s:%s/CMS?maxPoolSize=500`, ip, port)
	log.Infof("[Mongo Addr] %s", url)

	var session *mgo.Session
	var err error
	try := 0
	for {
		ip, port = getMgoAddr()
		url = fmt.Sprintf(`%s:%s/CMS?maxPoolSize=500`, ip, port)
		session, err = mgo.Dial(url)
		if err != nil {
			try++
			log.Errorf("dial mgo(%s) err %s, already try %d times", url, err.Error(), try)
			time.Sleep(time.Second * 10)
		} else {
			break
		}
	}

	initDb(session)
	return session
}

func initDb(session *mgo.Session) {
	//db := session.DB(DB_NAMESPACE_MONGO)
	//err := db.C(C_REPOSITORY).EnsureIndex(mgo.Index{Key: []string{COL_REPNAME}, Unique: true})
	//get(err)
}

func getMgoAddr() (ip, port string) {

	ip = os.Getenv(MONGODB_ADDR)
	port = os.Getenv(MONGODB_PORT)

	return
}
