package Model

import "time"

type Assignment struct {
	CreatedAtFrom      time.Time `json:"createdAtFrom"`
	CreatedAtTo        time.Time `json:"createdAtTo"`
	AssignedAtTime     time.Time `json:"assignedAtTime"`
	AssignedAtTo       time.Time `json:"assignedAtTo"`
	InactivatedAtFrom  time.Time `json:"inactivatedAtFrom"`
	InactivatedAtTo    time.Time `json:"inactivatedAtTo"`
	Inlet              string    `json:"inlet"`
	MemberId           string    `json:"memberId"`
	Outlets            struct{}  `json:"outlets"`
	GroupId            string    `json:"groupId"`
	Status             string    `json:"status"`
	AssignmentIds      struct{}  `json:"assignmentIds"`
	AssigneeId         string    `json:"assigneeId"`
	ParentAssignmentId string    `json:"parentAssignmentId"`
}
