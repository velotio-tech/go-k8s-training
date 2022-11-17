package auth

type Auth struct {
	AuthId   string `bson:"authId" json:"authId"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type AuthDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
