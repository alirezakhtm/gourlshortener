package structs

type Url struct {
	Address string `json:"address"`
	Username string `json:"username"`
	ShortUrl string `json:"shorturl"`
}

func (u Url) GetShortUrl() string {
	return u.ShortUrl
}

func (u *Url) SetShortUrl(shortUrl string) {
	u.ShortUrl =  shortUrl
}

func (u Url) GetUsername() string {
	return u.Username
}

func (u *Url) SetUsername(username string) {
	u.Username = username
}

func (u Url) GetAddress() string {
	return u.Address
}

func (u *Url) SetAddress(address string) {
	u.Address = address
}



