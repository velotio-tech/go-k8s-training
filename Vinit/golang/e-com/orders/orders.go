package orders

type Order struct {
	OrderId int `json:"OrderId",bson:"OrderId"`
	Content string `json:"content",bson:"content"`
	LastModified string `json:"lastModified",bson:"lastModified"`
}


