package Controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.docus.tech/kdl12138/DocusServer/Api"
	"git.docus.tech/kdl12138/DocusServer/Model"
	"git.docus.tech/kdl12138/DocusServer/Storage"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var delMessage Template.DelMessage
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO log
	}
	str := string(con)
	if err := json.Unmarshal([]byte(str), &delMessage); err != nil {
		// TODO log
	}
	delData, err := Model.Find(delMessage.Uuid)
	if err != nil {
		// TODO log
	}
	// TODO write
	if err := Model.Update(delData); err != nil {
		// TODO log
	}
	fmt.Fprint(w, "delete complete")
}
func Update(w http.ResponseWriter, r *http.Request) {
	var updateMessage Template.UpdataMessage
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO log
	}
	str := string(con)
	if err := json.Unmarshal([]byte(str), &updateMessage); err != nil {
		// TODO log
	}
	delData, err := Model.Find(updateMessage.Uuid)
	if err != nil {
		// TODO log
	}
	delData.Backup = 0
	// TODO write
	if err := Model.Update(delData); err != nil {
		// TODO log
	}
	fmt.Fprint(w, "Update complete")
}
func Read(w http.ResponseWriter, r *http.Request){
	if !Api.CheckWhite(r.RemoteAddr) {
		// TODO log
	} else {
		var getMessage Template.GetMessage
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &getMessage); err != nil {
			// TODO log
		}
		// TODO find cache
		getData, err:= Model.Find(getMessage.Uuid)
		data := Template.Data{

		}
		if err != nil {
			// TODO log
		}
		jsonData, _ := json.Marshal(data)
		url := ""
		contentType := "application/json;charset=utf-8"
		body := bytes.NewBuffer(jsonData)
		res, err := http.Post(url, contentType, body)
		if res != nil{

		}
	}

}

func Add(w http.ResponseWriter, r *http.Request)  {
	if !Api.CheckWhite(r.RemoteAddr) {
		// TODO log
	} else {
		var addMessage Template.AddMessage
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &addMessage); err != nil {
			// TODO log
		}
		// TODO find cache
		addData, err:= Model.Find(addMessage.Uuid)
		if err == nil {
			// TODO log
		} else {
			size := int64(len([]byte(addMessage.Data)))
			node, block, flag, _ := Api.Find(size)
			storageData := Template.StorageData{
				block,
				Storage.StorageMap[node][block].RestOffset - size - 1,
				Storage.StorageMap[node][block].RestOffset - 1,
				flag,
			}
			jsonData, _ := json.Marshal(storageData)
			url := ""
			contentType := "application/json;charset=utf-8"
			body := bytes.NewBuffer(jsonData)
			res, err := http.Post( url, contentType, body)
			if res != nil {

			}
		}

	}


}
