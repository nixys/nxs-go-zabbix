package zabbix

// For `MediatypeObject` field: `Type`
const (
	MediatypeEmail   = 0
	MediatypeScript  = 1
	MediatypeSMS     = 2
	MediatypeWebhook = 4
)

// For `MediatypeObject` field: `SMTPSecurity`
const (
	MediatypeSMTPSecurityNone     = 0
	MediatypeSMTPSecuritySTARTTLS = 1
	MediatypeSMTPSecuritySSLTLS   = 2
)

// For `MediatypeObject` field: `SMTPVerifyHost`
const (
	MediatypeSMTPVerifyHostNo  = 0
	MediatypeSMTPVerifyHostYes = 1
)

// For `MediatypeObject` field: `SMTPVerifyPeer`
const (
	MediatypeSMTPVerifyPeerNo  = 0
	MediatypeSMTPVerifyPeerYes = 1
)

// For `MediatypeObject` field: `SMTPAuthentication`
const (
	MediatypeSMTPAuthenticationNone           = 0
	MediatypeSMTPAuthenticationNormalPassword = 1
)

// For `MediatypeObject` field: `Status`
const (
	MediatypeStatusEnabled  = 0
	MediatypeScriptDisabled = 1
)

// For `MediatypeObject` field: `ContentType`
const (
	MediatypeContentTypePlainText = 0
	MediatypeContentTypeHTML      = 1
)

// For `MediatypeObject` field: `ProcessTags`
const (
	MediatypeProcessTagsNo  = 0
	MediatypeProcessTagsYes = 1
)

// For `MediatypeObject` field: `ShowEventMenu`
const (
	MediatypeShowEventMenuNo  = 0
	MediatypeShowEventMenuYes = 1
)

// For `MediatypeMessageTemplateObject` field: `EventSource`
const (
	MediatypeMessageTemplateEventSourceTriggers         = 0
	MediatypeMessageTemplateEventSourceDiscovery        = 1
	MediatypeMessageTemplateEventSourceAutoregistration = 2
	MediatypeMessageTemplateEventSourceInternal         = 3
)

// For `MediatypeMessageTemplateObject` field: `Recovery`
const (
	MediatypeMessageTemplateRecoveryOperations         = 0
	MediatypeMessageTemplateRecoveryRecoveryOperations = 1
	MediatypeMessageTemplateRecoveryUpdateOperations   = 2
)

// MediatypeObject struct is used to store mediatype operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/mediatype/object
type MediatypeObject struct {
	MediatypeID        int                                `json:"mediatypeid,omitempty"`
	Name               string                             `json:"name,omitempty"`
	Type               int                                `json:"type,omitempty"` // has defined consts, see above
	ExecPath           string                             `json:"exec_path,omitempty"`
	GsmModem           string                             `json:"gsm_modem,omitempty"`
	Passwd             string                             `json:"passwd,omitempty"`
	SMTPEmail          string                             `json:"smtp_email,omitempty"`
	SMTPHelo           string                             `json:"smtp_helo,omitempty"`
	SMTPServer         string                             `json:"smtp_server,omitempty"`
	SMTPPort           int                                `json:"smtp_port,omitempty"`
	SMTPSecurity       int                                `json:"smtp_security,omitempty"`       // has defined consts, see above
	SMTPVerifyHost     int                                `json:"smtp_verify_host,omitempty"`    // has defined consts, see above
	SMTPVerifyPeer     int                                `json:"smtp_verify_peer,omitempty"`    // has defined consts, see above
	SMTPAuthentication int                                `json:"smtp_authentication,omitempty"` // has defined consts, see above
	Status             int                                `json:"status,omitempty"`              // has defined consts, see above
	Username           string                             `json:"username,omitempty"`
	ExecParams         string                             `json:"exec_params,omitempty"`
	MaxSessions        int                                `json:"maxsessions,omitempty"`
	MaxAttempts        int                                `json:"maxattempts,omitempty"`
	AttemptInterval    string                             `json:"attempt_interval,omitempty"`
	ContentType        int                                `json:"content_type,omitempty"` // has defined consts, see above
	Script             string                             `json:"script,omitempty"`
	Timeout            string                             `json:"timeout,omitempty"`
	ProcessTags        int                                `json:"process_tags,omitempty"`    // has defined consts, see above
	ShowEventMenu      int                                `json:"show_event_menu,omitempty"` // has defined consts, see above
	EventMenuURL       string                             `json:"event_menu_url,omitempty"`
	EventMenuName      string                             `json:"event_menu_name,omitempty"`
	Parameters         []MediatypeWebhookParametersObject `json:"parameters,omitempty"`
	Description        string                             `json:"description,omitempty"`
	MessageTemplates   []MediatypeMessageTemplateObject   `json:"message_templates,omitempty"`

	Users []UserObject `json:"users,omitempty"`
}

// MediatypeWebhookParametersObject struct is used for mediatypes webhook parameters
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/mediatype/object#webhook_parameters
type MediatypeWebhookParametersObject struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// MediatypeMessageTemplateObject struct is used for mediatypes message template
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/mediatype/object#message_template
type MediatypeMessageTemplateObject struct {
	EventSource int    `json:"eventsource"` // has defined consts, see above
	Recovery    int    `json:"recovery"`    // has defined consts, see above
	Subject     string `json:"subject,omitempty"`
	Message     string `json:"message,omitempty"`
}

// MediatypeGetParams struct is used for mediatype get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/mediatype/get#parameters
type MediatypeGetParams struct {
	GetParameters

	MediatypeIDs []int `json:"mediatypeids,omitempty"`
	MediaIDs     []int `json:"mediaids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`

	SelectMessageTemplates SelectQuery `json:"selectMessageTemplates,omitempty"`
	SelectUsers            SelectQuery `json:"selectUsers,omitempty"`
}

// Structure to store creation result
type mediatypeCreateResult struct {
	MediatypeIDs []int `json:"mediatypeids"`
}

// Structure to store deletion result
type mediatypeDeleteResult struct {
	MediatypeIDs []int `json:"mediatypeids"`
}

// MediatypeGet gets mediatypes
func (z *Context) MediatypeGet(params MediatypeGetParams) ([]MediatypeObject, int, error) {

	var result []MediatypeObject

	status, err := z.request("mediatype.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// MediatypeCreate creates mediatypes
func (z *Context) MediatypeCreate(params []MediatypeObject) ([]int, int, error) {

	var result mediatypeCreateResult

	status, err := z.request("mediatype.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.MediatypeIDs, status, nil
}

// MediatypeDelete deletes mediatypes
func (z *Context) MediatypeDelete(mediatypeIDs []int) ([]int, int, error) {

	var result mediatypeDeleteResult

	status, err := z.request("mediatype.delete", mediatypeIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.MediatypeIDs, status, nil
}
