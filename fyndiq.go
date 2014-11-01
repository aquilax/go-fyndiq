package fyndiq

type FyndiqApi struct {
	user string
	token string
}

func NewFyndiqApi(user string, token string) (*FyndiqApi) {
	return &FyndiqApi{user, token}
}
