package zabbix

type ProblemSourceType int64

const (
	ProblemSourceTypeTrigger             ProblemSourceType = 0
	ProblemSourceTypeInternal            ProblemSourceType = 3
	ProblemSourceTypeServiceStatusUpdate ProblemSourceType = 4
)

type ProblemObjectType int64

const (
	ProblemObjectTypeTrigger ProblemObjectType = 0
	ProblemObjectTypeItem    ProblemObjectType = 4
	ProblemObjectTypeLLDRule ProblemObjectType = 5
	ProblemObjectTypeService ProblemObjectType = 6
)

type ProblemAcknowledgeType int64

const (
	ProblemAcknowledgeTypeFalse ProblemAcknowledgeType = 0
	ProblemAcknowledgeTypeTrue  ProblemAcknowledgeType = 1
)

type ProblemSeverityType int64

const (
	ProblemSeverityTypeNotClassified ProblemSeverityType = 0
	ProblemSeverityTypeInformation   ProblemSeverityType = 1
	ProblemSeverityTypeWarning       ProblemSeverityType = 2
	ProblemSeverityTypeAverage       ProblemSeverityType = 3
	ProblemSeverityTypeHigh          ProblemSeverityType = 4
	ProblemSeverityTypeDisaster      ProblemSeverityType = 5
)

type ProblemSuppressedType int64

const (
	ProblemSuppressedTypeNormalState ProblemSuppressedType = 0
	ProblemSuppressedTypeSuppressed  ProblemSuppressedType = 1
)

type ProblemEvalType int64

const (
	ProblemEvalTypeAndOr ProblemEvalType = 0
	ProblemEvalTypeOR    ProblemEvalType = 2
)

// ProblemObject struct is used to store problem operations results
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/object
type ProblemObject struct {
	EventID       int64                       `json:"eventid,omitempty"`
	Source        ProblemSourceType           `json:"source,omitempty"`
	Object        ProblemObjectType           `json:"object,omitempty"`
	ObjectID      int64                       `json:"objectid,omitempty"`
	Clock         int64                       `json:"clock,omitempty"`
	Ns            int64                       `json:"ns,omitempty"`
	REventID      int64                       `json:"r_eventid,omitempty"`
	RClock        int64                       `json:"r_clock,omitempty"`
	RNs           int64                       `json:"r_ns,omitempty"`
	CauseEventID  int64                       `json:"cause_eventid,omitempty"`
	CorrelationID int64                       `json:"correlationid,omitempty"`
	UserID        int64                       `json:"userid,omitempty"`
	Name          string                      `json:"name,omitempty"`
	Acknowledged  ProblemAcknowledgeType      `json:"acknowledged,omitempty"`
	Severity      ProblemSeverityType         `json:"severity,omitempty"`
	Suppressed    ProblemSuppressedType       `json:"suppressed,omitempty"`
	OpData        string                      `json:"opdata,omitempty"`
	URLs          []ProblemMediaTypeURLObject `json:"urls,omitempty"`
	Acknowledges  []ProblemAcknowledgeObject  `json:"acknowledges,omitempty"`
	Tags          []ProblemTagObject          `json:"tags,omitempty"`
	Suppression   []ProblemSuppressionObject  `json:"suppression_data,omitempty"`
}

type ProblemAcknowledgeObject struct {
	AcknowledgeID int64                      `json:"acknowledgeid,omitempty"`
	UserID        int64                      `json:"userid,omitempty"`
	EventID       int64                      `json:"eventid,omitempty"`
	Clock         int64                      `json:"clock,omitempty"`
	Message       string                     `json:"message,omitempty"`
	Action        EventAcknowledgeActionType `json:"action,omitempty"`
	OldSeverity   ProblemSeverityType        `json:"old_severity,omitempty"`
	NewSeverity   ProblemSeverityType        `json:"new_severity,omitempty"`
	SuppressUntil int64                      `json:"suppress_until,omitempty"`
	TaskID        int64                      `json:"taskid,omitempty"`
}

type ProblemMediaTypeURLObject struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type ProblemTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type ProblemSuppressionObject struct {
	MaintenanceID int64 `json:"maintenanceid,omitempty"`
	UserID        int64 `json:"userid,omitempty"`
	SuppressUntil int64 `json:"suppress_until,omitempty"`
}

// ProblemGetParams struct is used for problem get requests
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/get#parameters
type ProblemGetParams struct {
	GetParameters

	EventIDs              []int64               `json:"eventids,omitempty"`
	GroupIDs              []int64               `json:"groupids,omitempty"`
	HostIDs               []int64               `json:"hostids,omitempty"`
	ObjectIDs             []int64               `json:"objectids,omitempty"`
	ApplicationIDs        []int64               `json:"applicationids,omitempty"`
	Source                ProblemSourceType     `json:"source,omitempty"`
	Object                ProblemObjectType     `json:"object,omitempty"`
	Acknowledged          bool                  `json:"acknowledged,omitempty"`
	Suppressed            bool                  `json:"suppressed,omitempty"`
	Severities            []ProblemSeverityType `json:"severities,omitempty"`
	Evaltype              ProblemEvalType       `json:"evaltype,omitempty"`
	Tags                  []ProblemTagObject    `json:"tags,omitempty"`
	Recent                bool                  `json:"recent,omitempty"`
	EventIDFrom           int64                 `json:"eventid_from,omitempty"`
	EventIDTill           int64                 `json:"eventid_till,omitempty"`
	TimeFrom              int64                 `json:"time_from,omitempty"`
	TimeTill              int64                 `json:"time_till,omitempty"`
	SelectAcknowledges    SelectQuery           `json:"selectAcknowledges,omitempty"`
	SelectTags            SelectQuery           `json:"selectTags,omitempty"`
	SelectSuppressionData SelectQuery           `json:"selectSuppressionData,omitempty"`
	SortField             []string              `json:"sortfield,omitempty"`
}

// ProblemGet gets problems
func (z *Context) ProblemGet(params ProblemGetParams) ([]ProblemObject, int, error) {

	var result []ProblemObject

	status, err := z.request("problem.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
