package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"markettracker.com/pkg/event"
)

type Asset struct {
	id           AssetID
	date         Date
	exchangeName ExchangeName
	price        Price

	events []event.Event
}

// TODO: Revies if it is needed to pass all the parameters or a DTO is enough
func NewAsset(assetDTO AssetDTO) (Asset, error) {
	id, err := NewAssetID(assetDTO.ID)
	if err != nil {
		return Asset{}, err
	}
	date, err := NewDate(assetDTO.Date)
	if err != nil {
		return Asset{}, err
	}
	name, err := NewExchangeName(assetDTO.Exchange)
	if err != nil {
		return Asset{}, err
	}
	price, err := NewPrice(assetDTO.Price)
	if err != nil {
		return Asset{}, err
	}
	asset := Asset{
		id:           id,
		date:         date,
		exchangeName: name,
		price:        price,
	}
	asset.Record(NewAssetRecordedEvent(assetDTO.ID, assetDTO.Date, assetDTO.Exchange, assetDTO.Price))
	return asset, nil
}

var ErrInvalidAssetID = errors.New("ErrInvalidAssetID")

type AssetID struct {
	value string
}

func NewAssetID(value string) (AssetID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return AssetID{}, ErrInvalidAssetID
	}
	return AssetID{
		value: v.String(),
	}, nil
}

type Date struct {
	value time.Time
}

func NewDate(date time.Time) (Date, error) {
	return Date{
		value: date,
	}, nil
}

type ExchangeName struct {
	value string
}

func NewExchangeName(name string) (ExchangeName, error) {
	return ExchangeName{
		value: name,
	}, nil
}

type Price struct {
	value float32
}

func NewPrice(value float32) (Price, error) {
	return Price{
		value: value,
	}, nil
}

func (a Asset) Record(evt event.Event) {
	a.events = append(a.events, evt)
}

func (a Asset) PullEvents() []event.Event {
	evts := a.events
	a.events = []event.Event{}
	return evts
}
