package api

type CpuInfo struct {
	Count int     `json:"count"`
	Usage float64 `json:"usage"`
}

type MemInfo struct {
	Total int `json:"total"`
	Used  int `json:"used"`
}

type HostMetric struct {
	Cpu *CpuInfo `json:"cpu"`
	Mem *MemInfo `json:"mem"`
}
