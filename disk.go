package main

import (
	"strings"

	"github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Mountpoint string `json:"mountpoint"`
	Device     string `json:"device"`
	FsType     string `json:"fsType"`
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
	Total      uint64 `json:"total"`
}

func getDiskInfos() ([]Disk, error) {
	prs, error := disk.Partitions(false)
	if error != nil {
		return nil, error
	}

	var ds []Disk
	for _, r := range prs {
		if strings.Contains(r.Mountpoint, "/snap/") {
			continue
		}
		var d Disk
		d.Mountpoint = r.Mountpoint
		d.Device = r.Device
		d.FsType = r.Fstype

		rs, error := disk.Usage(d.Device)
		if error != nil {
			return nil, error
		}
		d.Free = rs.Free
		d.Used = rs.Used
		d.Total = rs.Total
		ds = append(ds, d)
	}

	return ds, nil
}
