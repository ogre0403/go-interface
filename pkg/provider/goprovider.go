package provider

type GoProvider struct {
}

func (g *GoProvider) Name() string {
	return "aaa"
}

func (g *GoProvider) QueryUser(token string) (*UserInfo, error) {
	return &UserInfo{
		Username: "jjjj",
	}, nil
}
