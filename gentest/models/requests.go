package models

type Request struct {
	ID                      int64                  `json:"id" xorm:"id pk autoincr"`
	Uuid                    string                 `json:"uuid" xorm:"uuid null"`
}

type RequestQuery struct {
	OperatorID               int64
	CreatorID                int64
	RequesterID              int64
	RequestTypeIDs           []int64
	RequestTypeNames         []string
	NameLike                 string
	CurrentStatus            int64
	IsUnclosed               sql.NullBool
	IsOverdued               sql.NullBool
	IsSuspend                sql.NullBool
	StartAt, EndAt           time.Time
	OverdueStart, OverdueEnd time.Time
}