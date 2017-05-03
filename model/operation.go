package model

import "time"

const (
	ADD = iota
	DELETE
	UPDATE
)

//业务操作日志
type OperationLog struct {
	ID            int
	OperationType int
	Comment       string
	OperatorID    string
	OperationTime time.Time
}
