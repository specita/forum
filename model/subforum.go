package model

import "time"

//版块
type SubForum struct {
	ID          int
	Name        string
	CreatorID   int
	CreatorName string
	CreatedOn   time.Time
	Desc        string
}
