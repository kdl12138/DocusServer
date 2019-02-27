package Controller

import (
	"bytes"
	"encoding/json"
	"git.docus.tech/kdl12138/DocusServer/Api"
	"git.docus.tech/kdl12138/DocusServer/Model"
	"git.docus.tech/kdl12138/DocusServer/Storage"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"io/ioutil"
	"net/http"
)
// TODO check identify and md5
// TODO exception choose another storage server and update the error one
// TODO find in cache
var msg string = "success"

// get delete message and post message to storage server
func Delete(w http.ResponseWriter, r *http.Request) {
	var delMessage Template.DelMessage
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg = err.Error()
		// TODO log
	}
	str := string(con)
	if err := json.Unmarshal([]byte(str), &delMessage); err != nil {
		msg = err.Error()
		// TODO log
	}
	// TODO write
	if err := Model.Delete(delMessage.Uuid); err != nil {
		msg = err.Error()
		// TODO log
	}
	Api.JsonReturn(w, r, "controller", "delete", Template.STATUS_TRUE, msg)
}

// get update message and post message to storage server
func Update(w http.ResponseWriter, r *http.Request) {
	var updateMessage Template.UpdataMessage
	//var findData Template.MasterData
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO log
	}
	str := string(con)
	if err := json.Unmarshal([]byte(str), &updateMessage); err != nil {
		// TODO log
	}
	findData, err := Model.Find(updateMessage.Uuid)
	if err != nil {
		// TODO log
	}
	data := Template.Data{
		Identity: updateMessage.Identity,
		//SessionId:updateMessage.SessionId,
		Uuid: updateMessage.Uuid,
		Backup: findData.Backup,
		//BlockStatus: findData.BlockStatus,
		Node: findData.Node,
		Block: findData.Block,
		OffsetStart: findData.OffsetStart,
		OffsetEnd: findData.OffsetEnd,
		Gzip: findData.Gzip,
		Timestamp: updateMessage.Timestamp,
		Md5: updateMessage.Md5,
	}
	jsonData, _ := json.Marshal(data)
	url := ""
	contentType := "application/json;charset=utf-8"
	body := bytes.NewBuffer(jsonData)
	res, err := http.Post(url, contentType, body)
	if res != nil{

	}
	if err := Model.Update(updateMessage); err != nil {
		// TODO log
	}
	Api.JsonReturn(w, r, "controller", "update", Template.STATUS_TRUE,msg )
}

// get read message and post message to storage server
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
		findData, err:= Model.Find(getMessage.Uuid)
		data := Template.Data{
			Identity: getMessage.Identity,
			//SessionId:updateMessage.SessionId,
			Uuid: getMessage.Uuid,
			Backup: findData.Backup,
			//BlockStatus: findData.BlockStatus,
			Node: findData.Node,
			Block: findData.Block,
			OffsetStart: findData.OffsetStart,
			OffsetEnd: findData.OffsetEnd,
			Gzip: findData.Gzip,
			Timestamp: getMessage.Timestamp,
			Md5: findData.Md5,
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
// get add message and post message to storage server
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
		if _, err:= Model.Find(addMessage.Uuid);err == nil {
			// TODO log
		} else {
			size := int64(len([]byte(addMessage.Data)))
			node, block, flag, _ := Api.Find(size)
			storageData := Template.StorageData{
				block,
				Storage.StorageMap[node][block].RestOffset - size - 1,
				Storage.StorageMap[node][block].RestOffset - 1,
				flag,
				addMessage.Data,
			}
			jsonData, _ := json.Marshal(storageData)
			url := ""
			contentType := "application/json;charset=utf-8"
			body := bytes.NewBuffer(jsonData)
			res, _ := http.Post( url, contentType, body)
			if res != nil {

			}
		}

	}


}
