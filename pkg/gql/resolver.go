package gql

import (
	"context"
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
	DbConfig DatabaseConfig
}

func (r *Resolver) getStationByCD(ctx context.Context, stationCd *int) (*model.Station, error) {
	db, err := getDatabase(r.DbConfig)
	defer db.Close()

	if err != nil {
		return nil, errors.New("Failed to connect database")
	}

	station, err := xo.StationByStationCd(db, *stationCd)
	if err != nil {
		return nil, errors.New("Failed to retrieve station")
	}

	return &model.Station{
		Address:     &station.Address.String,
		StationCd:   station.StationCd,
		StationName: station.StationName,
	}, nil
}
