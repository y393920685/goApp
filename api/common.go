package api

type PageListReq struct {
	PageNum  int    `p:"pageNum" d:"1" v:"integer|min:1#起始页必须是整数|起始页最小{min}" dc:"页码"`                    // 页码
	PageSize int    `p:"pageSize" d:"30" v:"integer|between:1,300#每页数量必须是整数|每页数量{min}-{max}"  dc:"每页数量"` // 每页数量
	Search   string `p:"search" v:"regex:^\\S{1,30}$#查询1-30非空格字符" dc:"模糊搜索内容"`                           // 搜索内容
}
type PageListRes struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
