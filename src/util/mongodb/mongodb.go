package mongodb
import (
	mgo "gopkg.in/mgo.v2"
	"util/logs"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"util/message"
)

var (
	log = logs.New("util/mongo")
)

type ConnectOpt struct {
	Address string
	Port string
	Database string
	Collections map[string]string
}

type Instance struct {
	opts ConnectOpt
	session *mgo.Session
}

func NewInstance(opts ConnectOpt) (ins *Instance, err error)  {
	ins = &Instance{
		opts: opts,
	}

	ins.session, err = mgo.Dial(opts.Address)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (this *Instance) DB() *mgo.Database  {
	return this.session.DB(this.opts.Database)
}

func (this* Instance) Collection(name string) *mgo.Collection {
	return this.session.DB(this.opts.Database).C(this.opts.Collections[name])
}

func (this*Instance) CheckObjectId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return bson.NewObjectId(), errors.New(message.INVALID_ID)
	}

	return bson.ObjectIdHex(id), nil
}