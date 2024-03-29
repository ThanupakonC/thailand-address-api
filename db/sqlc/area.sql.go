// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: area.sql

package db

import (
	"context"
)

const listDistrictTambon = `-- name: ListDistrictTambon :many
SELECT tambon_id, tambon_thai, tambon_eng, tambon_thai_short, tambon_thai_eng, district_id, province_id, region_id
FROM tambon
WHERE district_id = $1
ORDER BY tambon_thai
`

func (q *Queries) ListDistrictTambon(ctx context.Context, districtID int32) ([]Tambon, error) {
	rows, err := q.db.QueryContext(ctx, listDistrictTambon, districtID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tambon{}
	for rows.Next() {
		var i Tambon
		if err := rows.Scan(
			&i.TambonID,
			&i.TambonThai,
			&i.TambonEng,
			&i.TambonThaiShort,
			&i.TambonThaiEng,
			&i.DistrictID,
			&i.ProvinceID,
			&i.RegionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProvinces = `-- name: ListProvinces :many
SELECT province_id, province_thai, province_eng, region_id
FROM province
ORDER BY province_thai
`

func (q *Queries) ListProvinces(ctx context.Context) ([]Province, error) {
	rows, err := q.db.QueryContext(ctx, listProvinces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Province{}
	for rows.Next() {
		var i Province
		if err := rows.Scan(
			&i.ProvinceID,
			&i.ProvinceThai,
			&i.ProvinceEng,
			&i.RegionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProvincesDistrict = `-- name: ListProvincesDistrict :many
SELECT district_id, district_thai, district_eng, district_thai_short, district_eng_short, province_id, region_id
FROM district
WHERE province_id = $1
ORDER BY district_thai
`

func (q *Queries) ListProvincesDistrict(ctx context.Context, provinceID int32) ([]District, error) {
	rows, err := q.db.QueryContext(ctx, listProvincesDistrict, provinceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []District{}
	for rows.Next() {
		var i District
		if err := rows.Scan(
			&i.DistrictID,
			&i.DistrictThai,
			&i.DistrictEng,
			&i.DistrictThaiShort,
			&i.DistrictEngShort,
			&i.ProvinceID,
			&i.RegionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTambonPostcode = `-- name: ListTambonPostcode :many
SELECT postcode, tambon_id
FROM postcode
WHERE tambon_id = $1
ORDER BY postcode
`

func (q *Queries) ListTambonPostcode(ctx context.Context, tambonID int32) ([]Postcode, error) {
	rows, err := q.db.QueryContext(ctx, listTambonPostcode, tambonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Postcode{}
	for rows.Next() {
		var i Postcode
		if err := rows.Scan(&i.Postcode, &i.TambonID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
