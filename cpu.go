package main

import (
	"github.com/shirou/gopsutil/cpu"
)

type Cpu struct {
	VendorId  string  `json:"vendorId"`
	Family    string  `json:"family"`
	Model     string  `json:"model"`
	Mhz       float64 `json:"mhz"`
	CacheSize int32   `json:"cacheSize"`
}

func getCpuInfos() ([]Cpu, error) {
	rs, error := cpu.Info()
	if error != nil {
		return nil, error
	}

	var cs []Cpu
	for _, r := range rs {
		var c Cpu
		c.VendorId = r.VendorID
		c.Family = r.Family
		c.Model = r.Model
		c.Mhz = r.Mhz
		c.CacheSize = r.CacheSize
		cs = append(cs, c)
	}

	return cs, nil
}
