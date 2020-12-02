package model

type Client struct {
	ID    int64
	Name  string
	Email string
	Phone string
}

func NewClient(name string, email string, phone string) *Client {
	c := &Client{Name: name, Email: email, Phone: phone}
	return c
}

func NewEmptyClient() interface{} {
	return &Client{}
}
