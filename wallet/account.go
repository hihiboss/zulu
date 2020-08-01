package wallet

type Account struct {
	pub string
	pvt string
}

func (a *Account) PubKey() string {
	return ""
}
