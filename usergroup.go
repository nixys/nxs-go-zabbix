package zabbix

// For `UsergroupObject` field: `DebugMode`
const (
	UsergroupDebugModeDisabled = 0
	UsergroupDebugModeEnabled  = 1
)

// For `UsergroupObject` field: `GuiAccess`
const (
	UsergroupGuiAccessSystemDefaultAuth = 0
	UsergroupGuiAccessInternalAuth      = 1
	UsergroupGuiAccessLDAPAuth          = 2
	UsergroupGuiAccessDisableFrontend   = 3
)

// For `UsergroupObject` field: `UsersStatus`
const (
	UsergroupUsersStatusEnabled  = 0
	UsergroupUsersStatusDisabled = 1
)

// For `UsergroupPermissionObject` field: `Permission`
const (
	UsergroupPermissionDenied = 0
	UsergroupPermissionRO     = 2
	UsergroupPermissionRW     = 3
)

// UsergroupObject struct is used to store usergroup operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usergroup/object#user_group
type UsergroupObject struct {
	UsrgrpID    int    `json:"usrgrpid,omitempty"`
	Name        string `json:"name,omitempty"`
	DebugMode   int    `json:"debug_mode,omitempty"`   // has defined consts, see above
	GuiAccess   int    `json:"gui_access,omitempty"`   // has defined consts, see above
	UsersStatus int    `json:"users_status,omitempty"` // has defined consts, see above

	Users      []UserObject                        `json:"users,omitempty"`
	Rights     []UsergroupPermissionObject         `json:"rights,omitempty"`
	TagFilters []UsergroupTagBasedPermissionObject `json:"tag_filters,omitempty"`

	// used when new usergroup created or updated
	UserIDs []int `json:"userids,omitempty"`
}

// UsergroupPermissionObject struct is used to store usergroup permissions
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usergroup/object#permission
type UsergroupPermissionObject struct {
	ID         int `json:"id"`
	Permission int `json:"permission"` // has defined consts, see above
}

// UsergroupTagBasedPermissionObject struct is used to store usergroup tag based permission
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usergroup/object#tag_based_permission
type UsergroupTagBasedPermissionObject struct {
	GroupID int    `json:"groupid,omitempty"`
	Tag     string `json:"tag,omitempty"`
	Value   string `json:"value,omitempty"`
}

// UsergroupGetParams struct is used for usergroup get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usergroup/get#parameters
type UsergroupGetParams struct {
	GetParameters

	Status        []int `json:"status,omitempty"`
	UserIDs       []int `json:"userids,omitempty"`
	UsrgrpIDs     []int `json:"usrgrpids,omitempty"`
	WithGuiAccess []int `json:"with_gui_access,omitempty"`

	SelectTagFilters SelectQuery `json:"selectTagFilters,omitempty"`
	SelectUsers      SelectQuery `json:"selectUsers,omitempty"`
	SelectRights     SelectQuery `json:"selectRights,omitempty"`
}

// Structure to store creation result
type usergroupCreateResult struct {
	UsrgrpIDs []int `json:"usrgrpids"`
}

// Structure to store updation result
type usergroupUpdateResult struct {
	UsrgrpIDs []int `json:"usrgrpids"`
}

// Structure to store deletion result
type usergroupDeleteResult struct {
	UsrgrpIDs []int `json:"usrgrpids"`
}

// UsergroupGet gets usergroups
func (z *Context) UsergroupGet(params UsergroupGetParams) ([]UsergroupObject, int, error) {

	var result []UsergroupObject

	status, err := z.request("usergroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// UsergroupCreate creates usergroups
func (z *Context) UsergroupCreate(params []UsergroupObject) ([]int, int, error) {

	var result usergroupCreateResult

	status, err := z.request("usergroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrgrpIDs, status, nil
}

// UsergroupUpdate updates usergroups
func (z *Context) UsergroupUpdate(params []UsergroupObject) ([]int, int, error) {

	var result usergroupUpdateResult

	status, err := z.request("usergroup.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrgrpIDs, status, nil
}

// UsergroupDelete deletes usergroups
func (z *Context) UsergroupDelete(usergroupIDs []int) ([]int, int, error) {

	var result usergroupDeleteResult

	status, err := z.request("usergroup.delete", usergroupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrgrpIDs, status, nil
}
