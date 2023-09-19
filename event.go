package zabbix

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
