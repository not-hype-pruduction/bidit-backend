package dds

import (
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/dds"
	ddsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/dds.v1"
)

func ToGetDDTableResponse(in *dds.DDTable) *ddsv1.GetDDTableResponse {
	return &ddsv1.GetDDTableResponse{
		Results: map[string]*ddsv1.PlayerResults{
			"NT": {
				North: int32(in.Results["NT"][0]),
				East:  int32(in.Results["NT"][1]),
				South: int32(in.Results["NT"][2]),
				West:  int32(in.Results["NT"][3]),
			},
			"S": {
				North: int32(in.Results["S"][0]),
				East:  int32(in.Results["S"][1]),
				South: int32(in.Results["S"][2]),
				West:  int32(in.Results["S"][3]),
			},
			"H": {
				North: int32(in.Results["H"][0]),
				East:  int32(in.Results["H"][1]),
				South: int32(in.Results["H"][2]),
				West:  int32(in.Results["H"][3]),
			},
			"D": {
				North: int32(in.Results["D"][0]),
				East:  int32(in.Results["D"][1]),
				South: int32(in.Results["D"][2]),
				West:  int32(in.Results["D"][3]),
			},
			"C": {
				North: int32(in.Results["C"][0]),
				East:  int32(in.Results["C"][1]),
				South: int32(in.Results["C"][2]),
				West:  int32(in.Results["C"][3]),
			},
		},
	}
}
