package defs

type UserCredential struct {
	UserName string `json:"user_name",bson:"username"`
	Password string `json:"pwd",bson:"password"`
}
