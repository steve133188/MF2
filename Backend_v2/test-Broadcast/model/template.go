package model

type SendTemp struct {
	Customer     []CInfo `json:"customer"`
	TemplateName string  `json:"template_name"`
	LanguageCode string  `json:"language_code"`
}

type CInfo struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type BBody struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type Header struct {
	Type    string `json:"type"`
	Body    string `json:"body"`
	Example string `json:"example"`
}

type Template struct {
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	ChannelId string  `json:"channel_id"`
	Language  string  `json:"language"`
	Header    Header  `json:"header"`
	Body      string  `json:"body"`
	BodyEX    string  `json:"body_example"`
	Footer    string  `json:"footer"`
	BType     string  `json:"buttons_type"`
	BBody     []BBody `json:"buttons_body"`
}

type CreateWhatsAppMessageTemplateInput struct {
	Name      string
	Category  string
	ChannelId string
	Language  string
}
