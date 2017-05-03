package model

import "time"

//版块=企业咕咚的企业
type SubForum struct {
	ID          int
	Name        string
	CreatorID   int
	CreatorName string
	CreatedOn   time.Time
	Desc        string
}
