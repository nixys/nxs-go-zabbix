package zabbix

// For `MediatypeObject` field: `Type`
const (
	MediatypeEmail     = 0
	MediatypeScript    = 1
	MediatypeSMS       = 2
	MediatypeJabber    = 3
	MediatypeEzTexting = 100
)

// For `MediatypeObject` field: `Status`
const (
	MediatypeStatusEnabled  = 0
	MediatypeScriptDisabled = 1
)

// MediatypeObject struct is used to store mediatype operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/mediatype/object
type MediatypeObject struct {
	MediatypeID int    `json:"mediatypeid,omitempty"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"` // has defined consts, see above
	ExecPath    string `json:"exec_path,omitempty"`
	GsmModem    string `json:"gsm_modem,omitempty"`
	Passwd      string `json:"passwd,omitempty"`
	SMTPEmail   string `json:"smtp_email,omitempty"`
	SMTPHelo    string `json:"smtp_helo,omitempty"`
	SMTPServer  string `json:"smtp_server,omitempty"`
	Status      int    `json:"status,omitempty"` // has defined consts, see above
	Username    string `json:"username,omitempty"`

	Users []UserObject `json:"users,omitempty"`
}

// MediatypeGetParams struct is used for mediatype get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/mediatype/get#parameters
type MediatypeGetParams struct {
	GetParameters

	MediatypeIDs []int `json:"mediatypeids,omitempty"`
	MediaIDs     []int `json:"mediaids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`

	SelectUsers SelectQuery `json:"selectUsers,omitempty"`
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
