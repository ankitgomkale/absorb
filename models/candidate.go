package models

import (
	"time"
)

type Candidate struct {
	ID        int64
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	InterestedRole string `json:"interested_role"`
	YearsOfExp string `json:"years_of_exp"`
	CandidateLocation string `json:"candidate_location"`
	LastOrgs string `json:"last_org"`
	NotableOrg string `json:"notable_org"`
	CollegeName string `json:"college_name"`
	Remote string `json:"remote"`
	Relocate string `json:"relocate"`
	Description string `json:"description"`
	Salary	int64	`json:"salary"`
	Entrepreneur	string `json:"entrepreneur"`
	CommunicationMethod	string `json:"communication_method"`
	ResumeLink	string `json:"resume_link"`
	Linkedin	string `json:"linkedin"`
	Website	string `json:"website"`
	Github	string `json:"github"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

var Candidates []Candidate