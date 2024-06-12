package response

type PodListItem struct {
	Name     string `json:"name"`
	Ready    string `json:"ready"`
	Status   string `json:"status"`
	Restarts int32  `json:"restarts"`
	Age      int64  `json:"age"`
	IP       string `json:"IP"`
	Node     string `json:"node"`
}
