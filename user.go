package zabbix

// For `UserObject` field: `AutoLogin`
const (
	UserAutoLoginDisabled = 0
	UserAutoLoginEnabled  = 1
)

// For `UserObject` field: `Theme`
const (
	UserThemeDefault = "default"
	UserThemeBlue    = "blue-theme"
	UserThemeDark    = "dark-theme"
)

// For `UserObject` field: `Type`
const (
	UserTypeUser       = 1
	UserTypeAdmin      = 2
	UserTypeSuperAdmin = 3
)

// For `MediaObject` field: `Active`
const (
	MediaActiveEnabled  = 0
	MediaActiveDisabled = 1
)

// UserObject struct is used to store user operations results
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/user/object
type UserObject struct {
	UserID        string `json:"userid,omitempty"`
	Username      string `json:"username"`
	AttemptClock  int    `json:"attempt_clock,omitempty"`
	AttemptFailed int    `json:"attempt_failed,omitempty"`
	AttemptIP     string `json:"attempt_ip,omitempty"`
	AutoLogin     int    `json:"autologin,omitempty"` // has defined consts, see above
	AutoLogout    string `json:"autologout"`
	Lang          string `json:"lang,omitempty"`
	Name          string `json:"name,omitempty"`
	Refresh       string `json:"refresh,omitempty"`
	RowsPerPage   int    `json:"rows_per_page,omitempty"`
	Surname       string `json:"surname,omitempty"`
	Theme         string `json:"theme,omitempty"` // has defined consts, see above
	URL           string `json:"url,omitempty"`
	Timezone      string `json:"timezone,omitempty"`
	RoleID        string `json:"roleid"`

	// used for user.login
	UserDataObject

	Medias     []MediaObject     `json:"medias,omitempty"`
	Mediatypes []MediatypeObject `json:"mediatypes,omitempty"`
	Usrgrps    []UsergroupObject `json:"usrgrps,omitempty"`

	// used when new user created
	UserMedias []MediaObject `json:"user_medias,omitempty"`
	Passwd     string        `json:"passwd,omitempty"`
}

// MediaObject struct is used to store media operations results
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/user/object#media
type MediaObject struct {
	MediaTypeID string   `json:"mediatypeid"`
	SendTo      []string `json:"sendto"`
	Active      int      `json:"active,omitempty"` // has defined consts, see above
	Severity    int      `json:"severity,omitempty"`
	Period      string   `json:"period,omitempty"`
}

// UserLoginParams struct is used for login requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/user/login#parameters
type UserLoginParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UserData string `json:"userData,omitempty"`
}

// UserDataObject struct is used to store authenticated user additional info
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/user/login#return_values
type UserDataObject struct {
	DebugMode bool   `json:"debug_mode,omitempty"`
	GUIAccess int    `json:"gui_access,omitempty"`
	SessionID string `json:"sessionid,omitempty"`
	UserIP    string `json:"userip,omitempty"`
}

// UserGetParams struct is used for user get requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/user/get#parameters
type UserGetParams struct {
	GetParameters

	MediaIDs     []string `json:"mediaids,omitempty"`
	NediatypeIDs []string `json:"mediatypeids,omitempty"`
	UserIDs      []string `json:"userids,omitempty"`
	UsrgrpIDs    []string `json:"usrgrpids,omitempty"`

	GetAccess        bool        `json:"getAccess,omitempty"`
	SelectMedias     SelectQuery `json:"selectMedias,omitempty"`
	SelectMediatypes SelectQuery `json:"selectMediatypes,omitempty"`
	SelectUsrgrps    SelectQuery `json:"selectUsrgrps,omitempty"`
}

// Structure to store creation result
type userCreateResult struct {
	UserIDs []string `json:"userids"`
}

// Structure to store deletion result
type userDeleteResult struct {
	UserIDs []string `json:"userids"`
}

// UserGet gets users
func (z *Context) UserGet(params UserGetParams) ([]UserObject, int, error) {

	var result []UserObject

	status, err := z.request("user.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// UserCreate creates users
func (z *Context) UserCreate(params []UserObject) ([]string, int, error) {

	var result userCreateResult

	status, err := z.request("user.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UserIDs, status, nil
}

// UserDelete deletes users
func (z *Context) UserDelete(userIDs []string) ([]string, int, error) {

	var result userDeleteResult

	status, err := z.request("user.delete", userIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UserIDs, status, nil
}

func (z *Context) userLogin(params UserLoginParams) (string, int, error) {

	var result string

	status, err := z.request("user.login", params, &result)
	if err != nil {
		return "", status, err
	}

	return result, status, nil
}

func (z *Context) userLogout() (bool, int, error) {

	var result bool

	status, err := z.request("user.logout", []string{}, &result)
	if err != nil {
		return false, status, err
	}

	return result, status, nil
}
