package gql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ks6088ts/sandbox-go/pkg/gql/model"
	"github.com/ks6088ts/sandbox-go/pkg/gql/xo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	todos    []*model.Todo
	stations []*xo.Station
	Db       *sql.DB
}

func (r *Resolver) getStationByCD(ctx context.Context, stationCd *int) (*model.Station, error) {
	station, err := xo.StationByStationCd(r.Db, *stationCd)
	if err != nil {
		return nil, errors.New("Failed to retrieve station")
	}

	return &model.Station{
		Address:     &station.Address.String,
		StationCd:   station.StationCd,
		StationName: station.StationName,
	}, nil
}
