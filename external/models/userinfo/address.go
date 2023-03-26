package userinfo

type Addresses []Address

type Address struct {
	Addr1     string `json:"addr1,omitempty" db:"addr1"`
	Addr2     string `json:"addr2,omitempty" db:"addr2"`
	City      string `json:"city,omitempty" db:"city"`
	State     string `json:"state,omitempty" db:"state"`
	Zip       string `json:"zip,omitempty" db:"zip"`
	IsPrimary bool   `json:"isPrimary,omitempty" db:"is_primary"`
}
