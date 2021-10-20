package Model

import "time"

type GetAdmins struct {
	CreatedAtFrom time.Time `json:"createdAtFrom"`
	CreatedAtTo   time.Time `json:"createdAtTo"`
	UpdatedAtFrom time.Time `json:"updatedAtFrom"`
	UpdatedAtTo   time.Time `json:"updatedAtTo"`
	ChannelId     string    `json:"channelId"`
	AdminIds      string    `json:"adminIds"`
	ExternalId    string    `json:"externalId"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
}
