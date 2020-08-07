// Package xo contains the types for schema 'public'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Station represents a row from 'public.station'.
type Station struct {
	StationCd int `json:"station_cd"` // station_cd
	StationGCd int `json:"station_g_cd"` // station_g_cd
	StationName string `json:"station_name"` // station_name
	StationNameK sql.NullString `json:"station_name_k"` // station_name_k
	StationNameR sql.NullString `json:"station_name_r"` // station_name_r
	LineCd sql.NullInt64 `json:"line_cd"` // line_cd
	PrefCd sql.NullInt64 `json:"pref_cd"` // pref_cd
	Post sql.NullString `json:"post"` // post
	Address sql.NullString `json:"address"` // address
	Lon sql.NullFloat64 `json:"lon"` // lon
	Lat sql.NullFloat64 `json:"lat"` // lat
	OpenYmd sql.NullString `json:"open_ymd"` // open_ymd
	CloseYmd sql.NullString `json:"close_ymd"` // close_ymd
	EStatus sql.NullInt64 `json:"e_status"` // e_status
	ESort sql.NullInt64 `json:"e_sort"` // e_sort

	// xo fields
	_exists, _deleted bool

}


// Exists determines if the Station exists in the database.
func (s *Station) Exists() bool {
	return s._exists
}

// Deleted provides information if the Station has been deleted from the database.
func (s *Station) Deleted() bool {
	return s._deleted
}

// Insert inserts the Station to the database.
func (s *Station) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}


	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.station (` +
		`station_cd, station_g_cd, station_name, station_name_k, station_name_r, line_cd, pref_cd, post, address, lon, lat, open_ymd, close_ymd, e_status, e_sort` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)`

	// run query
	XOLog(sqlstr, s.StationCd, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort)
	err = db.QueryRow(sqlstr, s.StationCd, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort).Scan(&s.StationCd)
	if err != nil {
		return err
	}


	// set existence
	s._exists = true

	return nil
}


	// Update updates the Station in the database.
	func (s *Station) Update(db XODB) error {
		var err error

		// if doesn't exist, bail
		if !s._exists {
			return errors.New("update failed: does not exist")
		}

		// if deleted, bail
		if s._deleted {
			return errors.New("update failed: marked for deletion")
		}

		
			// sql query
			const sqlstr = `UPDATE public.station SET (` +
				`station_g_cd, station_name, station_name_k, station_name_r, line_cd, pref_cd, post, address, lon, lat, open_ymd, close_ymd, e_status, e_sort` +
				`) = ( ` +
				`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
				`) WHERE station_cd = $15`

			// run query
			XOLog(sqlstr, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort, s.StationCd)
			_, err = db.Exec(sqlstr, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort, s.StationCd)
			return err
	}

	// Save saves the Station to the database.
	func (s *Station) Save(db XODB) error {
		if s.Exists() {
			return s.Update(db)
		}

		return s.Insert(db)
	}

	// Upsert performs an upsert for Station.
	//
	// NOTE: PostgreSQL 9.5+ only
	func (s *Station) Upsert(db XODB) error {
		var err error

		// if already exist, bail
		if s._exists {
			return errors.New("insert failed: already exists")
		}

		// sql query
		const sqlstr = `INSERT INTO public.station (` +
			`station_cd, station_g_cd, station_name, station_name_k, station_name_r, line_cd, pref_cd, post, address, lon, lat, open_ymd, close_ymd, e_status, e_sort` +
			`) VALUES (` +
			`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
			`) ON CONFLICT (station_cd) DO UPDATE SET (` +
			`station_cd, station_g_cd, station_name, station_name_k, station_name_r, line_cd, pref_cd, post, address, lon, lat, open_ymd, close_ymd, e_status, e_sort` +
			`) = (` +
			`EXCLUDED.station_cd, EXCLUDED.station_g_cd, EXCLUDED.station_name, EXCLUDED.station_name_k, EXCLUDED.station_name_r, EXCLUDED.line_cd, EXCLUDED.pref_cd, EXCLUDED.post, EXCLUDED.address, EXCLUDED.lon, EXCLUDED.lat, EXCLUDED.open_ymd, EXCLUDED.close_ymd, EXCLUDED.e_status, EXCLUDED.e_sort` +
			`)`

		// run query
		XOLog(sqlstr, s.StationCd, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort)
		_, err = db.Exec(sqlstr, s.StationCd, s.StationGCd, s.StationName, s.StationNameK, s.StationNameR, s.LineCd, s.PrefCd, s.Post, s.Address, s.Lon, s.Lat, s.OpenYmd, s.CloseYmd, s.EStatus, s.ESort)
		if err != nil {
			return err
		}

		// set existence
		s._exists = true

		return nil
}


// Delete deletes the Station from the database.
func (s *Station) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	
		// sql query
		const sqlstr = `DELETE FROM public.station WHERE station_cd = $1`

		// run query
		XOLog(sqlstr, s.StationCd)
		_, err = db.Exec(sqlstr, s.StationCd)
		if err != nil {
			return err
		}

	// set deleted
	s._deleted = true

	return nil
}

// StationByStationCd retrieves a row from 'public.station' as a Station.
//
// Generated from index 'station_pkey'.
func StationByStationCd(db XODB, stationCd int) (*Station, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`station_cd, station_g_cd, station_name, station_name_k, station_name_r, line_cd, pref_cd, post, address, lon, lat, open_ymd, close_ymd, e_status, e_sort ` +
		`FROM public.station ` +
		`WHERE station_cd = $1`

	// run query
	XOLog(sqlstr, stationCd)
	s := Station{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, stationCd).Scan(&s.StationCd, &s.StationGCd, &s.StationName, &s.StationNameK, &s.StationNameR, &s.LineCd, &s.PrefCd, &s.Post, &s.Address, &s.Lon, &s.Lat, &s.OpenYmd, &s.CloseYmd, &s.EStatus, &s.ESort)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

