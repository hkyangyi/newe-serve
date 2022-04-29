package utils

type PageList struct {
	List     interface{} `json:"list"`
	Page     int         `json:"page" form:"page"`         //当前页
	PageSize int         `json:"pageSize" form:"pageSize"` //每页数量
	Total    int64       `json:"total"`                    //总条数
}

func (p *PageList) GetOffice() int {
	result := 0
	if p.Page > 0 {
		result = (p.Page - 1) * p.PageSize
	}
	return result
}
