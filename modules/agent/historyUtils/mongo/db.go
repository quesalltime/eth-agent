package mongo

import (
	"errors"
	"eth-agent/config"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var (
	mongoURL = "mongodb://" + config.SysConf.Mongo.Domain + ":" + config.SysConf.Mongo.Port + "/"
)

var mgoSession *mgo.Session

// GetMongoSession get the mongo session
func GetMongoSession() (*mgo.Session, error) {
	var err error

	if mgoSession == nil {
		mgoSession, err = mgo.Dial(mongoURL)
		if err != nil {
			message := fmt.Sprintf("Failed to start the Mongo session")
			err = errors.New(message)
		}
	}
	return mgoSession.Clone(), err
}
