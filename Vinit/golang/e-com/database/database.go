package database

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
)

var err error
var connection *mgo.Session

func createConnection() (*mgo.Session, error) {
	//dialInfo := mgo.DialInfo{
	//	Addrs: []string{"dbservice-shard-00-00.ovs8d.mongodb.net:27017",
	//		"dbservice-shard-00-01.ovs8d.mongodb.net:27017",
	//		"dbservice-shard-00-02.ovs8d.mongodb.net:27017 ",
	//	},
	//	//Database: "velotiodb",
	//	Username: "vtuser",
	//	Password: "vtsecretpassword",
	//}
	//tlsConfig := &tls.Config{}
	//dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	//	conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
	//	return conn, err
	//}
	//return mgo.DialWithInfo(&dialInfo)
	conn, err := mgo.Dial("db:27017")
	if err != nil {
		panic(err)
	}
	return conn, err
}

func init() {
	connection, err = createConnection()
	if err != nil {
		log.Fatal("Database connection Failed")
	}
}

func GetAllUsers(Uid string) [] byte{
	session := connection.Copy()
	defer session.Close()
	type dbResponse struct {
		Data  []*interface{} `json:"data" bson:"data"`
	}
	data := dbResponse{}
	dbOps := session.DB("velotiodb").C("data")
	var param interface{}
	if Uid != "" {
		userId, _ := strconv.Atoi(Uid)
		param = bson.M{"Uid":userId}
	}
	err = dbOps.Find(param).Select(bson.M{"Uid":1, "username":1, "_id":0}).All(&data.Data)
	if err != nil {
		log.Println("Error retrieving data from the database")
		return []byte(`{"StatusCode": 401, "Data": "Unable to get Data"}`)
	}
	users, _ := json.Marshal(data)
	return users
}

func WriteToDB(data []byte) []byte{
	session := connection.Copy()
	defer session.Close()
	dbOps := session.DB("velotiodb").C("data")
	var doc interface{}
	json.Unmarshal(data, &doc)
	err := dbOps.Insert(doc)
	if err != nil {
		log.Println("Error Adding data to the database")
		return []byte(`{"StatusCode": 401, "Data": "Unable to Add Data"}`)
	}
	return []byte(`{"StatusCode": 200, "Message": "Data Added successfully", "Data":`+string(data)+`}`)
}