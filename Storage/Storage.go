package Storage

import (
	"git.docus.tech/kdl12138/DocusServer/Api"
	"net/http"
)
// TODO new a block
func NewBlock(w http.ResponseWriter, r *http.Request)  (err error){
	if Api.CheckWhite(r.RemoteAddr){

	} else {

	}

	return nil
}