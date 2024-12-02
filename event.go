package zabbix

import "fmt"

type EventAcknowledgeActionType int64

const (
	EventAcknowledgeActionTypeClose                  EventAcknowledgeActionType = 1
	EventAcknowledgeActionTypeAck                    EventAcknowledgeActionType = 2
	EventAcknowledgeActionTypeAddMessage             EventAcknowledgeActionType = 4
	EventAcknowledgeActionTypeChangeSeverity         EventAcknowledgeActionType = 8
	EventAcknowledgeActionTypeUnack                  EventAcknowledgeActionType = 16
	EventAcknowledgeActionTypeSuppress               EventAcknowledgeActionType = 32
	EventAcknowledgeActionTypeUnsuppress             EventAcknowledgeActionType = 64
	EventAcknowledgeActionTypeChangeEventRankCause   EventAcknowledgeActionType = 128
	EventAcknowledgeActionTypeChangeEventRankSymptom EventAcknowledgeActionType = 256
)

// EventAcknowledgeParams struct is used for event acknowledge requests.
type EventAcknowledgeParams struct {
	EventIDs     []int64                    `json:"eventids"`
	Action       EventAcknowledgeActionType `json:"action"`
	Message      string                     `json:"message,omitempty"`
	Severity     ProblemSeverityType        `json:"severity,omitempty"`
	SupressUntil int64                      `json:"suppress_until,omitempty"`
}

// EventAcknowledgeResult struct is used to store the result of an event acknowledge request.
type EventAcknowledgeResult struct {
	EventIDs []int64 `json:"eventids"`
}

// EventAcknowledge acknowledges an event.
func (z *Context) EventAcknowledge(params EventAcknowledgeParams) (EventAcknowledgeResult, error) {
	var result EventAcknowledgeResult

	_, err := z.request("event.acknowledge", params, &result)
	if err != nil {
		return EventAcknowledgeResult{}, fmt.Errorf("error acknowledging event: %v", err)
	}

	return result, nil
}

// CombineEventAcknowledgeActions helper function to combine EventAcknowledgeActionType values.
func CombineEventAcknowledgeActions(actions ...EventAcknowledgeActionType) EventAcknowledgeActionType {
	var combined EventAcknowledgeActionType
	for _, action := range actions {
		combined |= action
	}
	return combined
}
