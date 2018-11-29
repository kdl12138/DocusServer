package Model

import (
	"git.docus.tech/kdl12138/DocusServer/Template"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type SessionStore struct {
	session *mgo.Session
}

func Init(dialinfo *mgo.DialInfo) {
	var err error
	session, err = mgo.DialWithInfo(dialinfo)
	if err != nil {
		// TODO log
	}
	session.SetMode(mgo.Monotonic, true)
}

func (d *SessionStore) Connect(table string) *mgo.Collection {
	return d.session.DB(Template.Database).C(table)
}
func NewSessionStore() *SessionStore {
	ds := &SessionStore{
		session: session.Copy(),
	}
	return ds
}
func (d *SessionStore) Close() {
	d.session.Close()
}
func GetErrNotFound() error {
	return mgo.ErrNotFound
}
func Find(name string) (GetData Template.MasterData, err error) {
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.Connect(Template.Database)
	if err := con.Find(bson.M{"uuid": name}).One(&GetData); err != nil {
		if err.Error() != GetErrNotFound().Error() {
			return GetData, err
		}

	}
	return GetData, nil
}
func FindAll()  (GetData[] Template.MasterData, err error){
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.Connect(Template.Database)
	if err := con.Find(bson.M{}).All(&GetData); err != nil {
		if err.Error() != GetErrNotFound().Error() {
			return GetData, err
		}

	}
	return GetData, nil
}
func Delete(name string) (err error) {
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.Connect(Template.Database)
	if err := con.Remove(bson.M{"uuid": name}); err != nil {
		// TODO log
	}
	//DelData = Template.DelData{Identity: tempData.Identity,
	//	SessionId:   tempData.SessionId,
	//	Action:      tempData.Action,
	//	Uuid:        tempData.Uuid,
	//	Backup:      tempData.Backup,
	//	BlockStatus: tempData.BlockStatus,
	//	Node:        tempData.Node,
	//	Block:       tempData.Block,
	//	OffsetStart: tempData.OffsetStart,
	//	OffsetEnd:   tempData.OffsetEnd,
	//	Gzip:        tempData.Gzip,
	//	Timestamp:   tempData.Timestamp,
	//	Md5:         tempData.Md5}
	return nil
}
func Update(data Template.MasterData) (err error) {
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.Connect(Template.Database)
	if err := con.Update(bson.M{"uuid": data.Uuid}, &data); err != nil {
		// TODO log
	}
	return nil
}
func Add(data Template.MasterData) (err error) {
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.Connect(Template.Database)
	if err := con.Insert(&data); err != nil {
		// TODO log
	}
	return nil
}
