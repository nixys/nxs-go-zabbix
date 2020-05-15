package zabbix

import "fmt"

// For `HistoryGetParams` field: `History`
const (
	HistoryObjectTypeFloat           = 0
	HistoryObjectTypeCharacter       = 1
	HistoryObjectTypeLog             = 2
	HistoryObjectTypeNumericUnsigned = 3
	HistoryObjectTypeText            = 4
)

// HistoryFloatObject struct is used to store history float operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/object#float_history
type HistoryFloatObject struct {
	Clock  int     `json:"clock,omitempty"`
	ItemID int     `json:"itemid,omitempty"`
	NS     int     `json:"ns,omitempty"`
	Value  float64 `json:"value,omitempty"`
}

// HistoryIntegerObject struct is used to store history integer operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/object#integer_history
type HistoryIntegerObject struct {
	Clock  int `json:"clock,omitempty"`
	ItemID int `json:"itemid,omitempty"`
	NS     int `json:"ns,omitempty"`
	Value  int `json:"value,omitempty"`
}

// HistoryStringObject struct is used to store history string operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/object#string_history
type HistoryStringObject struct {
	Clock  int    `json:"clock,omitempty"`
	ItemID int    `json:"itemid,omitempty"`
	NS     int    `json:"ns,omitempty"`
	Value  string `json:"value,omitempty"`
}

// HistoryTextObject struct is used to store history text operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/object#text_history
type HistoryTextObject struct {
	ID     int    `json:"id,omitempty"`
	Clock  int    `json:"clock,omitempty"`
	ItemID int    `json:"itemid,omitempty"`
	NS     int    `json:"ns,omitempty"`
	Value  string `json:"value,omitempty"`
}

// HistoryLogObject struct is used to store history log operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/object#log_history
type HistoryLogObject struct {
	ID         int    `json:"id,omitempty"`
	Clock      int    `json:"clock,omitempty"`
	ItemID     int    `json:"itemid,omitempty"`
	LogeventID int    `json:"logeventid,omitempty"`
	NS         int    `json:"ns,omitempty"`
	Severity   int    `json:"severity,omitempty"`
	Source     int    `json:"source,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
	Value      string `json:"value,omitempty"`
}

// HistoryGetParams struct is used for history get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/history/get#parameters
type HistoryGetParams struct {
	GetParameters

	History  int   `json:"history"` // has defined consts, see above
	HostIDs  []int `json:"hostids,omitempty"`
	ItemIDs  []int `json:"itemids,omitempty"`
	TimeFrom int   `json:"time_from,omitempty"`
	TimeTill int   `json:"time_till,omitempty"`

	Sortfield string `json:"sortfield,omitempty"`
}

// HistoryGet gets history
func (z *Context) HistoryGet(params HistoryGetParams) (interface{}, int, error) {

	var result interface{}

	switch params.History {
	case HistoryObjectTypeFloat:
		result = &([]HistoryFloatObject{})
	case HistoryObjectTypeCharacter:
		result = &([]HistoryStringObject{})
	case HistoryObjectTypeLog:
		result = &([]HistoryLogObject{})
	case HistoryObjectTypeNumericUnsigned:
		result = &([]HistoryIntegerObject{})
	case HistoryObjectTypeText:
		result = &([]HistoryTextObject{})
	default:
		return nil, 0, fmt.Errorf("Unknown history type")
	}

	status, err := z.request("history.get", params, result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
