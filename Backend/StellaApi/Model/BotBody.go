package Model

type BotRequestBody struct {
	ChannelId     string   `json:"channelId"`
	FbUserRef     string   `json:"fbUserRef"`
	FbLoginId     string   `json:"fbLoginId"`
	MemberId      string   `json:"memberId"`
	RecipientId   string   `json:"recipientId"`
	MessagingType string   `json:"messagingType"`
	MessagingTag  string   `json:"messagingTag"`
	Response      struct{} `jons:"response"`
}

type BotRedirectNode struct {
	ChannelId   string `json:"channelId"`
	FbUserRef   string `json:"fbUserRef"`
	MemberId    string `json:"memberId"`
	RecipientId string `json:"recipientId"`
	Redirect    struct{}
	Meta        struct{}
}
