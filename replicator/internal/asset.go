package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"markettracker.com/pkg/event"
)

type AssetRecordedEventDTO struct {
	event.EventDTO
	Data Data                   `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type Data struct {
	AggregateId string  `json:"aggregateId"`
	Date        string  `json:"date"`
	Exchange    string  `json:"exchange"`
	Price       float32 `json:"price"`
}

type AssetRepository interface {
	Save(ctx context.Context, asset Asset) error
}

type Asset struct {
	id           AssetID
	date         Date
	exchangeName ExchangeName
	price        Price
}

// TODO: Revies if it is needed to pass all the parameters or a DTO is enough
func NewAsset(id string, date string, exchange string, price float32) (Asset, error) {
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

func (a *AssetID) String() string {
	return a.value
}

var ErrDateDoesNotMatchTheRightFormat = errors.New("ErrDateDoesNotMatchTheRightFormat")

type Date struct {
	value  time.Time
	format string
}

func NewDate(date string) (Date, error) {
	format := time.RFC3339Nano
	value, err := time.Parse(format, date)
	if err != nil {
		return Date{}, ErrDateDoesNotMatchTheRightFormat
	}
	return Date{
		value:  value,
		format: format,
	}, nil
}

func (d *Date) String() string {
	return d.value.Format(d.format)
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

func (a *Asset) ID() string {
	return a.id.String()
}

func (a *Asset) Date() string {
	return a.date.String()
}
