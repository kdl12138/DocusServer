package Controller

import (
	"encoding/json"
	"git.docus.tech/kdl12138/DocusServer/Api"
	"git.docus.tech/kdl12138/DocusServer/File"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"io/ioutil"
	"net/http"
)
// TODO check md5
// get master's message and do somethings
func AddData(w http.ResponseWriter, r *http.Request)  {
	if Api.CheckWhite(r.RemoteAddr){

	} else {
		var storageData Template.StorageData
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &storageData); err != nil {
			// TODO log
		}

		File.Add(storageData)
	}
}
func DeleteData(w http.ResponseWriter, r *http.Request)  {
	if Api.CheckWhite(r.RemoteAddr){

	} else {
		var storageData Template.StorageData
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &storageData); err != nil {
			// TODO log
		}
		File.Delete(storageData)
		Api.JsonReturn(w, r, "Storage", "delete", Template.STATUS_TRUE,msg )
	}
}
func ReadData(w http.ResponseWriter, r *http.Request)  {
	if Api.CheckWhite(r.RemoteAddr){

	} else {
		var storageData Template.StorageData
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &storageData); err != nil {
			// TODO log
		}
		File.Read(storageData)
		Api.JsonReturn(w, r, "Storage", "read", Template.STATUS_TRUE,msg )
	}
}
func UpdateData(w http.ResponseWriter, r *http.Request)  {
	if Api.CheckWhite(r.RemoteAddr){

	} else {
		var storageData Template.StorageData
		con, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// TODO log
		}
		str := string(con)
		if err := json.Unmarshal([]byte(str), &storageData); err != nil {
			// TODO log
		}
		File.Update(storageData)
		Api.JsonReturn(w, r, "Storage", "update", Template.STATUS_TRUE,msg )
	}
}
