package database

import (
	"bytes"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var err error
var connection *mgo.Session

type Order struct {
	OrderId int `json:"OrderId",bson:"OrderId"`
	Content string `json:"content",bson:"content"`
	LastModified string `json:"lastModified",bson:"lastModified"`
}

type User struct {
	Username string `json:"username",bson:"username"`
	UserId int `json:"Uid",bson:"Uid"`
	Orders []Order `json:"orders",bson:"orders"`
}

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

func getOrderOfId(orderId string, data interface{}) []byte{
	userData, _ := json.Marshal(data)
	var user User
	reader := ioutil.NopCloser(strings.NewReader(string(userData)))
	_ = json.NewDecoder(reader).Decode(&user)
	response := new(bytes.Buffer)
	for i,v := range user.Orders{
		oid, _ := strconv.Atoi(orderId)
		if oid == v.OrderId{
			json.NewEncoder(response).Encode(user.Orders[i])
			break
		}
	}
	return response.Bytes()
}

func init() {
	connection, err = createConnection()
	if err != nil {
		log.Fatal("Database connection Failed")
	}
}

func AddOrder(uid string, detail []byte) []byte{
	session := connection.Copy()
	defer session.Close()
	dbOps := session.DB("velotiodb").C("data")
	userId, _ := strconv.Atoi(uid)
	var doc interface{}
	json.Unmarshal(detail, &doc)
	filter := bson.M{"Uid":userId}
	err = dbOps.Update(filter,bson.M{"$push":bson.M{"orders":doc}})
	if err != nil {
		log.Println("Error retrieving data from the database")
		return []byte(`{"StatusCode": 401, "Data": "Unable to Add Data Check UserId"}`)
	}
	return []byte(`{"StatusCode": 200, "Message": "Data Added successfully", "Data":`+string(detail)+`}`)
}

func DeleteOrders(uid string, oid string) []byte {
	session := connection.Copy()
	defer session.Close()
	dbOps := session.DB("velotiodb").C("data")
	userId, _ := strconv.Atoi(uid)
	filter := bson.M{"Uid":userId}
	if oid =="" {
		emptyOrder := make([]interface{},0)
		err = dbOps.Update(filter,bson.M{"$set":bson.M{"orders":emptyOrder}})
	} else {
		orderId, _ := strconv.Atoi(oid)
		err = dbOps.Update(filter, bson.M{"$pull":bson.M{"orders":bson.M{"OrderId":orderId}}})
	}
	if err != nil {
		log.Println("Error Updating data in the database")
		return []byte(`{"StatusCode": 401, "Data": "Unable to delete Data Check UserId and order Id"}`)
	}
	return []byte(`{"StatusCode": 200, "Message": "Data Updated successfully"}`)
}

func GetOrders(uid string, oid string) []byte {
	type dbResponse struct {
		Data []*interface{} `json:"data" bson:"data"`
	}
	data := dbResponse{}
	session := connection.Copy()
	defer session.Close()
	dbOps := session.DB("velotiodb").C("data")
	userId, _ := strconv.Atoi(uid)
	filter := bson.M{"Uid":userId}
	err = dbOps.Find(filter).Select(bson.M{"_id":0}).All(&data.Data)
	if err != nil {
		log.Println("Error retrieving data from the database")
		return []byte(`{"StatusCode": 401, "Data": "Unable to get Data"}`)
	}
	if oid == "" {
		orderData, _ := json.Marshal(data)
		return orderData
	} else {
		return getOrderOfId(oid, data.Data[0])
	}
}