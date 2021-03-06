package model

import "encoding/json"

type Profile struct {
	Name string
	Gender string
	Age string
	Height string
	Education string
	Place string
	Marriage string
	Income string
	Id string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil{
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
