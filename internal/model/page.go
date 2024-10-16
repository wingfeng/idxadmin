package model

import "strings"

type PageReq struct { //每页显示记录数
	PageSize int             `json:"pageSize" form:"pageSize"`
	Cols     map[string]bool `form:"cols"`
	//当前页
	Page int `json:"page" form:"page"`
	//过滤条件过滤条件只支持'field operator ?'格式
	Filters []string `json:"filters"`
	//多个参数用分号分隔 如: 内容1;内容2;
	Args []string `json:"args"`
	//排序类型
	SortOrder string `json:"sortOrder"`
	//排序名称
	SortField string `json:"sortField"`
}

func (req *PageReq) GetSortOrder() string {
	return strings.TrimSuffix(req.SortOrder, "end")
}

type PageRes struct {
	//记录总数
	Total    int         `json:"total" `
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize" form:"pageSize"`
	List     interface{} `json:"list"`
}
