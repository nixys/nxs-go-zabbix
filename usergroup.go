package zabbix

// For `UsergroupObject` field: `DebugMode`
const (
	UsergroupDebugModeDisabled = "0"
	UsergroupDebugModeEnabled  = "1"
)

// For `UsergroupObject` field: `GuiAccess`
const (
	UsergroupGuiAccessSystemDefaultAuth = "0"
	UsergroupGuiAccessInternalAuth      = "1"
	UsergroupGuiAccessDisableFrontend   = "2"
)

// For `UsergroupObject` field: `UsersStatus`
const (
	UsergroupUsersStatusEnabled  = "0"
	UsergroupUsersStatusDisabled = "1"
)

// For `UsergroupPermissionObject` field: `Permission`
const (
	UsergroupPermissionDenied = "0"
	UsergroupPermissionRO     = "2"
	UsergroupPermissionRW     = "3"
)

// UsergroupObject struct is used to store usergroup operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usergroup/object#user_group
type UsergroupObject struct {
	UsrgrpID    string `json:"usrgrpid,omitempty"`
	Name        string `json:"name,omitempty"`
	DebugMode   string `json:"debug_mode,omitempty"`   // has defined consts, see above
	GuiAccess   string `json:"gui_access,omitempty"`   // has defined consts, see above
	UsersStatus string `json:"users_status,omitempty"` // has defined consts, see above

	Users  []UserObject                `json:"users,omitempty"`
	Rights []UsergroupPermissionObject `json:"rights,omitempty"`

	// used when new usergroup created or updated
	UserIDs []string `json:"userids,omitempty"`
}

// UsergroupPermissionObject struct is used to store usergroup permissions
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usergroup/object#permission
type UsergroupPermissionObject struct {
	ID         string `json:"id"`
	Permission string `json:"permission"` // has defined consts, see above
}

// UsergroupGetParams struct is used for usergroup get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usergroup/get#parameters
type UsergroupGetParams struct {
	GetParameters

	Status        []string `json:"status,omitempty"`
	UserIDs       []string `json:"userids,omitempty"`
	UsrgrpIDs     []string `json:"usrgrpids,omitempty"`
	WithGuiAccess []string `json:"with_gui_access,omitempty"`

	SelectUsers  SelectQuery `json:"selectUsers,omitempty"`
	SelectRights SelectQuery `json:"selectRights,omitempty"`
}

// Structure to store creation result
type usergroupCreateResult struct {
	UsrgrpIDs []string `json:"usrgrpids"`
}

// Structure to store deletion result
type usergroupDeleteResult struct {
	UsrgrpIDs []string `json:"usrgrpids"`
}

// UsergroupGet gets usergroups
func (z *Zabbix) UsergroupGet(params UsergroupGetParams) ([]UsergroupObject, int, error) {

	var result []UsergroupObject

	status, err := z.request("usergroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// UsergroupCreate creates usergroups
func (z *Zabbix) UsergroupCreate(params []UsergroupObject) ([]string, int, error) {

	var result usergroupCreateResult

	status, err := z.request("usergroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrgrpIDs, status, nil
}

// UsergroupDelete deletes usergroups
func (z *Zabbix) UsergroupDelete(usergroupIDs []string) ([]string, int, error) {

	var result usergroupDeleteResult

	status, err := z.request("usergroup.delete", usergroupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrgrpIDs, status, nil
}
