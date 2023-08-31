package zabbix

type ZabbixProblemObject struct {
	EventID       string              `json:"eventid,omitempty"`
	Source        string              `json:"source,omitempty"`
	Object        string              `json:"object,omitempty"`
	ObjectID      string              `json:"objectid,omitempty"`
	Clock         string              `json:"clock,omitempty"`
	Ns            string              `json:"ns,omitempty"`
	REventID      string              `json:"r_eventid,omitempty"`
	RClock        string              `json:"r_clock,omitempty"`
	RNs           string              `json:"r_ns,omitempty"`
	CorrelationID string              `json:"correlationid,omitempty"`
	UserID        string              `json:"userid,omitempty"`
	Name          string              `json:"name,omitempty"`
	Acknowledged  string              `json:"acknowledged,omitempty"`
	Severity      string              `json:"severity,omitempty"`
	CauseEventID  string              `json:"cause_eventid,omitempty"`
	OpData        string              `json:"opdata,omitempty"`
	Acknowledges  []ZabbixAcknowledge `json:"acknowledges,omitempty"`
	Suppression   []ZabbixSuppression `json:"suppression_data,omitempty"`
	Suppressed    string              `json:"suppressed,omitempty"`
	Tags          []ZabbixProblemTag  `json:"tags,omitempty"`
}

type ZabbixAcknowledge struct {
	AcknowledgeID string `json:"acknowledgeid,omitempty"`
	UserID        string `json:"userid,omitempty"`
	EventID       string `json:"eventid,omitempty"`
	Clock         string `json:"clock,omitempty"`
	Message       string `json:"message,omitempty"`
	Action        string `json:"action,omitempty"`
	OldSeverity   string `json:"old_severity,omitempty"`
	NewSeverity   string `json:"new_severity,omitempty"`
	SuppressUntil string `json:"suppress_until,omitempty"`
	TaskID        string `json:"taskid,omitempty"`
}

type ZabbixSuppression struct {
	MaintenanceID string `json:"maintenanceid,omitempty"`
	SuppressUntil string `json:"suppress_until,omitempty"`
	UserID        string `json:"userid,omitempty"`
}

type ZabbixProblemTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// ProblemGetParams struct is used for problem get requests
type ProblemGetParams struct {
	GetParameters

	ObjectIDs []string `json:"objectids,omitempty"`
	// ... Add other fields as needed
}

// ProblemGet gets problems
func (z *Context) ProblemGet(params ProblemGetParams) ([]ZabbixProblemObject, int, error) {
	var result []ZabbixProblemObject

	status, err := z.request("problem.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
