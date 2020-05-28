package interface_v1

import "api/dao"

type apiToken struct {
}

func GetToken() *apiToken {
	return new(apiToken)
}

func (a *apiToken) Address() (data []dao.Token) {
	return
}
