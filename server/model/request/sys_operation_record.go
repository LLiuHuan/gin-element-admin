package request

import "gin-element-admin/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
