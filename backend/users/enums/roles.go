package enums

type Role string

const (
	None     Role = "NONE"
	Admin    Role = "ADMIN"
	Assignee Role = "ASSIGNEE"
	User     Role = "USER"
)
