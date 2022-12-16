package models

type CaesarSimpleInput struct {
	UserId     string `json:"user_id"`
	PlaintText string `json:"plain_text"`
	Key        int    `json:"key"`
}
type CaesarWithPermutationInput struct {
	UserId     string `json:"user_id"`
	PlaintText string `json:"plain_text"`
	Key        int    `json:"key"`
	Word       string `json:"word"`
}

type VigenereInput struct {
	UserId     string `json:"user_id"`
	PlaintText string `json:"plain_text"`
	Key        string `json:"key"`
}

type Playfair struct {
	UserId     string `json:"user_id"`
	PlaintText string `json:"plain_text"`
	Key        string `json:"key"`
}
