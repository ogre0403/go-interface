package provider


type UserInfo struct {
	Username    string
}


type Provider interface {
	Name() string
	QueryUser(token string) (*UserInfo, error)
}