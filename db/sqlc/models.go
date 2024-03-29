// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type District struct {
	DistrictID        int32  `json:"district_id"`
	DistrictThai      string `json:"district_thai"`
	DistrictEng       string `json:"district_eng"`
	DistrictThaiShort string `json:"district_thai_short"`
	DistrictEngShort  string `json:"district_eng_short"`
	ProvinceID        int32  `json:"province_id"`
	RegionID          int32  `json:"region_id"`
}

type Postcode struct {
	Postcode int32 `json:"postcode"`
	TambonID int32 `json:"tambon_id"`
}

type Province struct {
	ProvinceID   int32         `json:"province_id"`
	ProvinceThai string        `json:"province_thai"`
	ProvinceEng  string        `json:"province_eng"`
	RegionID     sql.NullInt32 `json:"region_id"`
}

type Region struct {
	RegionID   int32  `json:"region_id"`
	RegionThai string `json:"region_thai"`
	RegionEng  string `json:"region_eng"`
}

type Tambon struct {
	TambonID        int32  `json:"tambon_id"`
	TambonThai      string `json:"tambon_thai"`
	TambonEng       string `json:"tambon_eng"`
	TambonThaiShort string `json:"tambon_thai_short"`
	TambonThaiEng   string `json:"tambon_thai_eng"`
	DistrictID      int32  `json:"district_id"`
	ProvinceID      int32  `json:"province_id"`
	RegionID        int32  `json:"region_id"`
}
