package store

import (
	"util/mongodb"
	"model"
	"util/logs"
	"gopkg.in/mgo.v2/bson"
)

var log = logs.New("guestStore")

type GuestStore struct {
	mgoDB *mongodb.Instance
}

func NewGuestStore(db *mongodb.Instance) *GuestStore {
	return &GuestStore{
		mgoDB: db,
	}
}

func (this *GuestStore) List() (guests []model.Guest, err error)  {
	err = this.mgoDB.Collection(mongodb.GuestCollection).Find(nil).All(&guests)

	return guests, err
}

func (this *GuestStore) Create(guest *model.Guest) error {
	return this.mgoDB.Collection(mongodb.GuestCollection).Insert(guest)
}

func (this *GuestStore) Get(guestId string) (guest *model.Guest, err error) {
	guestObjId, err := this.mgoDB.CheckObjectId(guestId)
	if err != nil {
		return nil, err
	}

	err = this.mgoDB.Collection(mongodb.GuestCollection).FindId(guestObjId).One(&guest)

	return guest, err
}

func (this *GuestStore) Update(guestId string, guest *model.Guest) error {

	guestServer, err := this.Get(guestId)
	if (err != nil) {
		return err
	}

	guestServer.FullName = guest.FullName
	guestServer.Address = guest.Address
	guestServer.Email = guest.Email
	guestServer.HomeTown = guest.HomeTown
	guestServer.Phone = guest.Phone

	err = this.mgoDB.Collection(mongodb.GuestCollection).UpdateId(bson.ObjectIdHex(guestId), guestServer)

	return err
}

func (this *GuestStore) Delete(guestId string) error{
	_, err := this.Get(guestId)
	if (err != nil) {
		return err
	}

	err = this.mgoDB.Collection(mongodb.GuestCollection).RemoveId(bson.ObjectIdHex(guestId))

	return err
}