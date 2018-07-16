package account

type Achievement struct{
	Id string `json:"id"`
	Current int `json:"current,omitempty"`
	Max int `json:"max,omitempty"`
	Done bool `json:"done,omitempty"`
	IsRepeatable int `json:"repeated,omitempty"`
	Bits []int `json:"bits,omitempty"`
}
