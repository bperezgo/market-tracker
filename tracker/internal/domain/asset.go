package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"markettracker.com/pkg/event"
)

type AssetRepository interface {
	Save(ctx context.Context, asset Asset) error
}

type Asset struct {
	id           AssetID
	date         Date
	exchangeName ExchangeName
	price        Price

	events []event.Event
}

// TODO: Revies if it is needed to pass all the parameters or a DTO is enough
func NewAsset(id string, date time.Time, exchange string, price float32) (Asset, error) {
	idVO, err := NewAssetID(id)
	if err != nil {
		return Asset{}, err
	}
	dateVO, err := NewDate(date)
	if err != nil {
		return Asset{}, err
	}
	nameVO, err := NewExchangeName(exchange)
	if err != nil {
		return Asset{}, err
	}
	priceVO, err := NewPrice(price)
	if err != nil {
		return Asset{}, err
	}
	asset := Asset{
		id:           idVO,
		date:         dateVO,
		exchangeName: nameVO,
		price:        priceVO,
	}
	asset.Record(NewAssetRecordedEvent(id, date, exchange, price))
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

func (a Asset) ID() string {
	return a.id.value
}

type Date struct {
	value time.Time
}

func NewDate(date time.Time) (Date, error) {
	return Date{
		value: date,
	}, nil
}

func (a Asset) RFC3339() string {
	return a.date.value.Format(time.RFC3339)
}

func (a Asset) Date() time.Time {
	return a.date.value
}

var ErrExchangeNameIsEmpty = errors.New("exchange name cannot be empty")

type ExchangeName struct {
	value string
}

func NewExchangeName(name string) (ExchangeName, error) {
	if name == "" {
		return ExchangeName{}, ErrExchangeNameIsEmpty
	}
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

func (a Asset) Float32Price() float32 {
	return a.price.value
}

func (a *Asset) Record(evt event.Event) {
	a.events = append(a.events, evt)
}

func (a *Asset) PullEvents() []event.Event {
	evts := a.events
	a.events = []event.Event{}
	return evts
}
