package event

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const dummyEventType = Type("dummy.event.type")

type dummyAggregate struct {
	Id    string
	Value string
}

type aggregateAttributesDTO struct {
	Id    string `json:"aggregateId"`
	Value string `json:"value"`
}

type dummyEventDTO struct {
	EventDTO
}

func newDummyEventDTO(evt Event) dummyEventDTO {
	return dummyEventDTO{
		EventDTO: EventDTO{
			Data: EventDataDTO{
				EventId:    evt.Id(),
				OccurredOn: "2022-12-12",
				Type:       string(dummyEventType),
				Meta:       evt.Meta(),
				Attributes: evt.Data(),
			},
		},
	}
}

type dummyEvent struct {
	BaseEvent
	Aggregate dummyAggregate
	meta      map[string]interface{}
}

func newDummyEvent(aggregateId string) dummyEvent {
	meta := make(map[string]interface{})
	meta["meta1"] = "valuemeta1"
	return dummyEvent{
		BaseEvent: NewBaseEvent(aggregateId),
		Aggregate: dummyAggregate{
			Id:    aggregateId,
			Value: "valueOfTheAggregate",
		},
		meta: meta,
	}
}

func (e dummyEvent) DTO() interface{} {
	return newDummyEventDTO(e)
}

func (e dummyEvent) Data() interface{} {
	return aggregateAttributesDTO{
		Id:    e.Aggregate.Id,
		Value: e.Aggregate.Value,
	}
}

func (e dummyEvent) Meta() map[string]interface{} {
	return e.meta
}

func (e dummyEvent) Type() Type {
	return dummyEventType
}

func marshalEvent(evt Event) ([]byte, error) {
	dto := evt.DTO()
	buf, err := json.Marshal(&dto)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Test_Event_DTO(t *testing.T) {
	t.Run("should return a standard value of the event when use the composition with BaseEventDTO", func(t *testing.T) {
		aggregateId := "someID"
		evt := newDummyEvent(aggregateId)
		value, err := marshalEvent(evt)
		require.NoError(t, err, "should return a right value")
		log.Println(string(value))
		expected := make(map[string]interface{})
		err = json.Unmarshal(value, &expected)
		require.NoError(t, err, "should return a right value")

		assert.NotNil(t, expected["data"])

		expected, ok := expected["data"].(map[string]interface{})
		require.Equal(t, true, ok, "casting data field is ok")

		assert.Equal(t, evt.Id(), expected["eventId"])
		assert.Equal(t, "2022-12-12", expected["occurredOn"])
		assert.Equal(t, "dummy.event.type", expected["type"])

		require.NotNil(t, expected["attributes"])
		require.NotNil(t, expected["meta"])

		expectedAttributes, ok := expected["attributes"].(map[string]interface{})
		require.Equal(t, true, ok, "casting attributes field is ok")

		assert.Equal(t, aggregateId, expectedAttributes["aggregateId"])
		assert.Equal(t, "valueOfTheAggregate", expectedAttributes["value"])

		expectedMeta, ok := expected["meta"].(map[string]interface{})
		require.Equal(t, true, ok, "casting meta field is ok")

		assert.Equal(t, "valuemeta1", expectedMeta["meta1"])
	})
}
