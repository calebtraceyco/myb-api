package userinfo

type Addresses []Address

type Address struct {
	Addr1     string `json:"addr1,omitempty"`
	Addr2     string `json:"addr2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	IsPrimary string `json:"isPrimary,omitempty"`
}
