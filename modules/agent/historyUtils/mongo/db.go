package mongo

import (
	"errors"
	"eth-agent/config"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// var (
// 	mongoURL = "mongodb://" + config.SysConf.Mongo.Username + ":" + config.SysConf.Mongo.Password + "@" + config.SysConf.Mongo.Domain + ":" + config.SysConf.Mongo.Port + "/" + config.SysConf.Mongo.DBName
// )

var mgoSession *mgo.Session

// GetMongoSession get the mongo session
func GetMongoSession() (*mgo.Session, error) {
	var err error

	info := &mgo.DialInfo{
		Addrs:    []string{config.SysConf.Mongo.Domain},
		Timeout:  10 * time.Second,
		Database: config.SysConf.Mongo.DBName,
		Username: config.SysConf.Mongo.Username,
		Password: config.SysConf.Mongo.Password,
	}

	if mgoSession == nil {
		mgoSession, err = mgo.DialWithInfo(info)

		fmt.Println(err)

		if err != nil {
			message := fmt.Sprintf("Failed to start the Mongo session")
			err = errors.New(message)

			return nil, err
		}
	}
	return mgoSession.Clone(), err
}
