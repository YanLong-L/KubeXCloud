package response

type PodListItem struct {
	Name     string `json:"name"`     // pod 名称
	Ready    string `json:"ready"`    // 状态 1/2
	Status   string `json:"status"`   // status Running/Error
	Restarts int32  `json:"restarts"` // n 次
	Age      int64  `json:"age"`      // 运行时间
	IP       string `json:"IP"`       // pod ip
	Node     string `json:"node"`     // pod被调度到了哪台node
}
