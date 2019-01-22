package zabbix

// For `MediaObject` field: `Active`
const (
	MediaActiveEnabled  = 0
	MediaActiveDisabled = 1
)

// MediaObject struct is used to store media operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermedia/object
type MediaObject struct {
	MediaID     int    `json:"mediaid,omitempty"`
	Active      int    `json:"active,omitempty"` // has defined consts, see above
	MediaTypeID int    `json:"mediatypeid,omitempty"`
	Period      string `json:"period,omitempty"`
	SendTo      string `json:"sendto,omitempty"`
	Severity    int    `json:"severity,omitempty"`
	UserID      int    `json:"userid,omitempty"`
}

// UsermediaGetParams struct is used for media get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermedia/get#parameters
type UsermediaGetParams struct {
	GetParameters

	MediaIDs     []int `json:"mediaids,omitempty"`
	UsrgrpIDs    []int `json:"usrgrpids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`
	MediatypeIDs []int `json:"mediatypeids,omitempty"`
}

// UsermediaGet gets medias
func (z *Context) UsermediaGet(params UsermediaGetParams) ([]MediaObject, int, error) {

	var result []MediaObject

	status, err := z.request("usermedia.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
