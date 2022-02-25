package model

type Connect struct {
	CID   string `json:"cid"`   //company id
	UID   string `json:"uid"`   // user id
	CName string `json:"cname"` // channel name
	UName string `json:"uname"` // username
	TID   string `json:"tid"`   //team id
}

type RedisChannelData struct {
	NName    string //NodeName
	CType    string //channel type
	Init     bool
	Status   string
	Url      string
	UID      string //User id
	SSection string //store section
	QR       string
	LA       string //last Activity
}
