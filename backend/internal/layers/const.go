package enums

type ChgeInstStatus string

const (
	INST_NEW     ChgeInstStatus = "NEW"
	INST_ACTIVE  ChgeInstStatus = "ACTIVE"
	INST_EXPIRED ChgeInstStatus = "EXPIRED"
)

type AssignmentStatus string

const (
	ASGN_IN_PROGRESS AssignmentStatus = "IN_PROGRESS"
	ASGN_SUBMITTED   AssignmentStatus = "SUBMITTED"
)
