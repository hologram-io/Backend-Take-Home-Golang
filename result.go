package usage_parser

import "net"

type UsageResult struct {
	Id        int
	BytesUsed int
	CellId    int
	Dmcc      string
	Ip        net.IP
	Mnc       int
}

func NewBasicUsage(id int, bytesused int) *UsageResult {
	return &UsageResult{Id: id, BytesUsed: bytesused}
}

func NewExtendedUsage(id int, bytesused int, cellid int, dmcc string, mnc int) *UsageResult {
	return &UsageResult{
		Id:        id,
		BytesUsed: bytesused,
		CellId:    cellid,
		Dmcc:      dmcc,
		Mnc:       mnc,
	}
}

func NewHexUsage(id int, bytesused int, cellid int, ip net.IP, mnc int) *UsageResult {
	return &UsageResult{
		Id:        id,
		BytesUsed: bytesused,
		CellId:    cellid,
		Ip:        ip,
		Mnc:       mnc,
	}
}
