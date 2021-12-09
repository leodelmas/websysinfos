package main

import (
	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func getMemInfos() (Mem, error) {
	rs, error := mem.VirtualMemory()
	if error != nil {
		return Mem{}, error
	}

	var m Mem
	m.Total = rs.Total
	m.Free = rs.Free
	m.Used = rs.Used
	m.UsedPercent = rs.UsedPercent

	return m, nil
}
