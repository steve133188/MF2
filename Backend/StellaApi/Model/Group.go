package Model

import "time"

type GetGroups struct {
	CreatedAtFrom   time.Time `json:"createdAtFrom"`
	CreatedAtTo     time.Time `json:"createdAtTo"`
	UpdatedAtTime   time.Time `json:"updatedAtTime"`
	UpdatedAtTo     time.Time `json:"updatedAtTo"`
	Inlet           string    `json:"inlet"`
	MemberId        string    `json:"memberId"`
	Outlet          string    `json:"outlets"`
	GroupIds        struct{}  `json:"groupIds"`
	AssignmentIds   struct{}  `json:"assignmentIds"`
	AdminId         string    `json:"adminId"`
	Valid           bool      `json:"valid"`
	ExternalId      string    `json:"externalId"`
	AdminExternalId string    `json:"adminExternalId"`
	GroupType       string    `json:"type"`
	Label           string    `json:"label"`
	Active          bool      `json:"active"`
}
