package zabbix

// For `ActionObject` field: `Status`
const (
	ActionStatusEnabled  = 0
	ActionStatusDisabled = 1
)

// For `ActionObject` field: `PauseSuppressed`
const (
	ActionPauseSuppressedEnabled  = 0
	ActionPauseSuppressedDisabled = 1
)

// For `ActionOperationObject` field: `OperationType`
const (
	ActionOperationTypeSendMsg              = 0
	ActionOperationTypeRemoteCmd            = 1
	ActionOperationTypeAddHost              = 2
	ActionOperationTypeRmHost               = 3
	ActionOperationTypeAddToHostGroup       = 4
	ActionOperationTypeRmFromHostGroup      = 5
	ActionOperationTypeLinkToTpl            = 6
	ActionOperationTypeUnlinkFromTpl        = 7
	ActionOperationTypeEnableHost           = 8
	ActionOperationTypeDisableHost          = 9
	ActionOperationTypeSetHostInventoryMode = 10
)

// For `ActionOperationObject` field: `EvalType`
const (
	ActionOperationEvalTypeAndOR = 0
	ActionOperationEvalTypeAnd   = 1
	ActionOperationEvalTypeOr    = 2
)

// For `ActionOperationCommandObject` field: `Type`
const (
	ActionOperationCommandTypeCustomScript = 0
	ActionOperationCommandTypeIPMI         = 1
	ActionOperationCommandTypeSSH          = 2
	ActionOperationCommandTypeTelnet       = 3
	ActionOperationCommandTypeGlobalScript = 4
)

// For `ActionOperationCommandObject` field: `AuthType`
const (
	ActionOperationCommandAuthTypePassword = 0
	ActionOperationCommandAuthTypePubKey   = 1
)

// For `ActionOperationCommandObject` field: `ExecuteOn`
const (
	ActionOperationCommandExecuteOnAgent  = 0
	ActionOperationCommandExecuteOnServer = 1
	ActionOperationCommandExecuteOnProxy  = 2
)

// For `ActionOperationMessageObject` field: `DefaultMsg`
const (
	ActionOperationMessageDefaultMsgFromOperation = 0
	ActionOperationMessageDefaultMsgFromMediaType = 1
)

// For `ActionOperationConditionObject` field: `ConditionType`
const (
	ActionOperationConditionTypeEventAcknowledged = 14
)

// For `ActionRecoveryOperationObject` field: `OperationType `
const (
	ActionRecoveryOperationTypeSendMsg           = 0
	ActionRecoveryOperationTypeRemoteCmd         = 1
	ActionRecoveryOperationTypeNotifyAllInvolved = 11
)

// For `ActionUpdateOperationObject` field: `OperationType `
const (
	ActionUpdateOperationTypeSendMsg           = 0
	ActionUpdateOperationTypeRemoteCmd         = 1
	ActionUpdateOperationTypeNotifyAllInvolved = 12
)

// For `ActionOperationConditionObject` field: `Operator`
const (
	ActionOperationConditionOperatorEq = 0
)

// For `ActionFilterObject` field: `EvalType`
const (
	ActionFilterEvalTypeAndOr  = 0
	ActionFilterEvalTypeAnd    = 1
	ActionFilterEvalTypeOr     = 2
	ActionFilterEvalTypeCustom = 3
)

// For `ActionFilterConditionObject` field: `ConditionType`
const (
	ActionFilterConditionTypeHostroup             = 0
	ActionFilterConditionTypeHost                 = 1
	ActionFilterConditionTypeTrigger              = 2
	ActionFilterConditionTypeTriggerName          = 3
	ActionFilterConditionTypeTriggerSeverity      = 4
	ActionFilterConditionTypeTriggerValue         = 5
	ActionFilterConditionTypeTimePeriod           = 6
	ActionFilterConditionTypeHostIP               = 7
	ActionFilterConditionTypeDiscoveryServiceType = 8
	ActionFilterConditionTypeDiscoveryServicePort = 9
	ActionFilterConditionTypeDiscoveryStatus      = 10
	ActionFilterConditionTypeUpdownTimeDuration   = 11
	ActionFilterConditionTypeRcvValue             = 12
	ActionFilterConditionTypeHostTemplate         = 13
	ActionFilterConditionTypeApplication          = 15
	ActionFilterConditionTypeProblemIsSuppressed  = 16
	ActionFilterConditionTypeDiscRule             = 18
	ActionFilterConditionTypeDiscCheck            = 19
	ActionFilterConditionTypeProxy                = 20
	ActionFilterConditionTypeDiscObject           = 21
	ActionFilterConditionTypeHostName             = 22
	ActionFilterConditionTypeEventType            = 23
	ActionFilterConditionTypeHostMetadata         = 24
	ActionFilterConditionTypeTag                  = 25
	ActionFilterConditionTypeTagValue             = 26
)

// For `ActionFilterConditionObject` field: `Operator`
const (
	ActionFilterConditionOperatorEQ          = 0  // =
	ActionFilterConditionOperatorNE          = 1  // <>
	ActionFilterConditionOperatorContains    = 2  // contains
	ActionFilterConditionOperatorNotrContain = 3  // does not contain
	ActionFilterConditionOperatorIN          = 4  // in
	ActionFilterConditionOperatorGE          = 5  // >=
	ActionFilterConditionOperatorLE          = 6  // <=
	ActionFilterConditionOperatorNotIn       = 7  // not in
	ActionFilterConditionOperatorMatches     = 8  // matches
	ActionFilterConditionOperatorNotMatches  = 9  // does not match
	ActionFilterConditionOperatorYes         = 10 // yes
	ActionFilterConditionOperatorNo          = 11 // no
)

// ActionObject struct is used to store action operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action
type ActionObject struct {
	ActionID        int    `json:"actionid,omitempty"`
	EscPeriod       int    `json:"esc_period"`
	Eventsource     int    `json:"eventsource"`
	Name            string `json:"name"`
	Status          int    `json:"status,omitempty"`           // has defined consts, see above
	PauseSuppressed int    `json:"pause_suppressed,omitempty"` // has defined consts, see above

	Operations            []ActionOperationObject         `json:"operations,omitempty"`
	Filter                ActionFilterObject              `json:"filter,omitempty"`
	RecoveryOperations    []ActionRecoveryOperationObject `json:"recovery_operations,omitempty"`
	AcknowledgeOperations []ActionRecoveryOperationObject `json:"acknowledge_operations,omitempty"`
}

// ActionOperationObject struct is used to store action operations
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_operation
type ActionOperationObject struct {
	OperationID   int                              `json:"operationid,omitempty"`
	OperationType int                              `json:"operationtype"` // has defined consts, see above
	ActionID      int                              `json:"actionid,omitempty"`
	EscPeriod     int                              `json:"esc_period,omitempty"`
	EscStepFrom   int                              `json:"esc_step_from,omitempty"`
	EscStepTo     int                              `json:"esc_step_to,omitempty"`
	EvalType      int                              `json:"evaltype,omitempty"` // has defined consts, see above
	Opcommand     ActionOperationCommandObject     `json:"opcommand,omitempty"`
	OpcommandGrp  []ActionOpcommandGrpObject       `json:"opcommand_grp,omitempty"`
	OpcommandHst  []ActionOpcommandHstObject       `json:"opcommand_hst,omitempty"`
	Opconditions  []ActionOperationConditionObject `json:"opconditions,omitempty"`
	Opgroup       []ActionOpgroupObject            `json:"opgroup,omitempty"`
	Opmessage     ActionOperationMessageObject     `json:"opmessage,omitempty"`
	OpmessageGrp  []ActionOpmessageGrpObject       `json:"opmessage_grp,omitempty"`
	OpmessageUsr  []ActionOpmessageUsrObject       `json:"opmessage_usr,omitempty"`
	Optemplate    []ActionOptemplateObject         `json:"optemplate,omitempty"`
	Opinventory   ActionOpinventoryObject          `json:"opinventory,omitempty"`
}

// ActionOperationCommandObject struct is used to store action operation commands
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_operation_command
type ActionOperationCommandObject struct {
	Command    string `json:"command"`
	Type       int    `json:"type"`                 // has defined consts, see above
	AuthType   int    `json:"authtype,omitempty"`   // has defined consts, see above
	ExecuteOn  int    `json:"execute_on,omitempty"` // has defined consts, see above
	Password   string `json:"password,omitempty"`
	Port       string `json:"port,omitempty"`
	PrivateKey string `json:"privatekey,omitempty"`
	PublicKey  string `json:"publickey,omitempty"`
	ScriptID   int    `json:"scriptid,omitempty"`
	UserName   string `json:"username,omitempty"`
}

// ActionOperationMessageObject struct is used to store action operation messages
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_operation_message
type ActionOperationMessageObject struct {
	DefaultMsg  int    `json:"default_msg,omitempty"` // has defined consts, see above
	MediatypeID int    `json:"mediatypeid,omitempty"`
	Message     string `json:"message,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

// ActionOperationConditionObject struct is used to store action operation conditions
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_operation_condition
type ActionOperationConditionObject struct {
	OpconditionID int    `json:"opconditionid,omitempty"`
	ConditionType int    `json:"conditiontype"` // has defined consts, see above
	Value         string `json:"value"`
	OperationID   int    `json:"operationid,omitempty"`
	Operator      int    `json:"operator,omitempty"`
}

// ActionRecoveryOperationObject struct is used to store action recovery operations
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_recovery_operation
type ActionRecoveryOperationObject struct {
	OperationID   int                          `json:"operationid"`
	OperationType int                          `json:"operationtype,omitempty"` // has defined consts, see above
	ActionID      int                          `json:"actionid,omitempty"`
	Opcommand     ActionOperationCommandObject `json:"opcommand,omitempty"`
	OpcommandGrp  []ActionOpcommandGrpObject   `json:"opcommand_grp,omitempty"`
	OpcommandHst  []ActionOpcommandHstObject   `json:"opcommand_hst,omitempty"`
	Opmessage     ActionOperationMessageObject `json:"opmessage,omitempty"`
	OpmessageGrp  []ActionOpmessageGrpObject   `json:"opmessage_grp,omitempty"`
	OpmessageUsr  []ActionOpmessageUsrObject   `json:"opmessage_usr,omitempty"`
}

// ActionUpdateOperationObject struct is used to store action update operations
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_update_operation
type ActionUpdateOperationObject struct {
	OperationID   int                          `json:"operationid"`
	OperationType int                          `json:"operationtype,omitempty"` // has defined consts, see above
	Opcommand     ActionOperationCommandObject `json:"opcommand,omitempty"`
	OpcommandGrp  []ActionOpcommandGrpObject   `json:"opcommand_grp,omitempty"`
	OpcommandHst  []ActionOpcommandHstObject   `json:"opcommand_hst,omitempty"`
	Opmessage     ActionOperationMessageObject `json:"opmessage,omitempty"`
	OpmessageGrp  []ActionOpmessageGrpObject   `json:"opmessage_grp,omitempty"`
	OpmessageUsr  []ActionOpmessageUsrObject   `json:"opmessage_usr,omitempty"`
}

// ActionFilterObject struct is used to store action filters
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_filter
type ActionFilterObject struct {
	Conditions  []ActionFilterConditionObject `json:"conditions"`
	EvalType    int                           `json:"evaltype"` // has defined consts, see above
	EvalFormula string                        `json:"eval_formula,omitempty"`
	Formula     string                        `json:"formula,omitempty"`
}

// ActionFilterConditionObject struct is used to store action filter conditions
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/object#action_filter_condition
type ActionFilterConditionObject struct {
	ConditionID   int    `json:"conditionid,omitempty"`
	ConditionType int    `json:"conditiontype"` // has defined consts, see above
	Value         string `json:"value"`
	Value2        string `json:"value2,omitempty"`
	ActionID      int    `json:"actionid,omitempty"`
	FormulaID     string `json:"formulaid,omitempty"`
	Operator      int    `json:"operator,omitempty"` // has defined consts, see above
}

// ActionOpcommandGrpObject struct is used to store action opcommand groups
type ActionOpcommandGrpObject struct {
	OpcommandGrpID int `json:"opcommand_grpid,omitempty"`
	OperationID    int `json:"operationid,omitempty"`
	GroupID        int `json:"groupid,omitempty"`
}

// ActionOpcommandHstObject struct is used to store action opcommand hosts
type ActionOpcommandHstObject struct {
	OpcommandHstID int `json:"opcommand_hstid,omitempty"`
	OperationID    int `json:"operationid,omitempty"`
	HostID         int `json:"hostid,omitempty"`
}

// ActionOpgroupObject struct is used to store action opgroups
type ActionOpgroupObject struct {
	OperationID int `json:"operationid,omitempty"`
	GroupID     int `json:"groupid,omitempty"`
}

// ActionOpmessageGrpObject struct is used to store action opmessage groups
type ActionOpmessageGrpObject struct {
	OperationID int `json:"operationid,omitempty"`
	UsrgrpID    int `json:"usrgrpid,omitempty"`
}

// ActionOpmessageUsrObject struct is used to store action opmessage users
type ActionOpmessageUsrObject struct {
	OperationID int `json:"operationid,omitempty"`
	UserID      int `json:"userid,omitempty"`
}

// ActionOptemplateObject struct is used to store action optemplates
type ActionOptemplateObject struct {
	OperationID int `json:"operationid,omitempty"`
	TemplateID  int `json:"templateid,omitempty"`
}

// ActionOpinventoryObject struct is used to store action opinventory
type ActionOpinventoryObject struct {
	OperationID   int `json:"operationid,omitempty"`
	InventoryMode int `json:"inventory_mode,omitempty"`
}

// ActionGetParams struct is used for action get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/action/get#parameters
type ActionGetParams struct {
	GetParameters

	ActionIDs    []int `json:"actionids,omitempty"`
	GroupIDs     []int `json:"groupids,omitempty"`
	HostIDs      []int `json:"hostids,omitempty"`
	TriggerIDs   []int `json:"triggerids,omitempty"`
	MediatypeIDs []int `json:"mediatypeids,omitempty"`
	UsrgrpIDs    []int `json:"usrgrpids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`
	ScriptIDs    []int `json:"scriptids,omitempty"`

	SelectFilter             SelectQuery `json:"selectFilter,omitempty"`
	SelectOperations         SelectQuery `json:"selectOperations,omitempty"`
	SelectRecoveryOperations SelectQuery `json:"selectRecoveryOperations,omitempty"`
	SelectUpdateOperations   SelectQuery `json:"selectUpdateOperations,omitempty"`
}

// Structure to store creation result
type actionCreateResult struct {
	ActionIDs []int `json:"actionids"`
}

// Structure to store deletion result
type actionDeleteResult struct {
	ActionIDs []int `json:"actionids"`
}

// ActionGet gets actions
func (z *Context) ActionGet(params ActionGetParams) ([]ActionObject, int, error) {

	var result []ActionObject

	status, err := z.request("action.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// ActionCreate creates actions
func (z *Context) ActionCreate(params []ActionObject) ([]int, int, error) {

	var result actionCreateResult

	status, err := z.request("action.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ActionIDs, status, nil
}

// ActionDelete deletes actions
func (z *Context) ActionDelete(actionIDs []int) ([]int, int, error) {

	var result actionDeleteResult

	status, err := z.request("action.delete", actionIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ActionIDs, status, nil
}
