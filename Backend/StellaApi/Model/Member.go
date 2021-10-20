package Model

type GetMembers struct {
	CreatedAtFrom  string   `json:"createdAtFrom"`
	CreatedAtTo    string   `json:"createdAtTo"`
	UpdatedAtFrom  string   `json:"updatedAtFrom"`
	UpdatedAtTo    string   `json:"updatedAtTo"`
	ChannelId      string   `json:"channelId"`
	MemberIds      struct{} `json:"memberIds"`
	ExternalIds    struct{} `json:"externalIds"`
	Firstname      string   `json:"firstName"`
	Lastname       string   `json:"lastName"`
	Gender         string   `json:"gender"`
	Locales        struct{} `json:"locales"`
	LocaleOperator string   `json:"localeOperator"`
	TagFilters     struct{} `json:"tagFilters"`
}

// type GetMember struct {
// 	MemberId   string `json:"member_id"`
// 	ExternalId string `json:"external_id"`
// }

type ToggleLiveChat struct {
	MemberId   string `json:"member_id"`
	ExternalId string `json:"external_id"`
	LiveChat   bool   `json:"live_chat"`
}
