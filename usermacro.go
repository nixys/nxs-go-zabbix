package zabbix

// For `UsermacroObject` field: `Type`
const (
	UsermacroTypeText   = 0
	UsermacroTypeSecret = 1
)

// UsermacroObject struct is used to store hostmacro and globalmacro operations results.
// In API docs Global and Host it is a two different object types that are joined in this package
// into one object `UsermacroObject` that includes fields form both API objects.
// The reason is the some other objects do not separates this types.
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usermacro/object#host_macro
// and https://www.zabbix.com/documentation/5.0/manual/api/reference/usermacro/object#global_macro
type UsermacroObject struct {

	// Gobal macro fields only
	GlobalmacroID int `json:"globalmacroid,omitempty"`

	// Host macro fields only
	HostmacroID int `json:"hostmacroid,omitempty"`
	HostID      int `json:"hostid,omitempty"`

	// Common fields
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`
	Type        int    `json:"type,omitempty"` // has defined consts, see above
	Description string `json:"description,omitempty"`

	Groups    []HostgroupObject `json:"groups,omitempty"`
	Hosts     []HostObject      `json:"hosts,omitempty"`
	Templates []TemplateObject  `json:"templates,omitempty"`
}

// UsermacroGetParams struct is used for hostmacro get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/usermacro/get#parameters
type UsermacroGetParams struct {
	GetParameters

	Globalmacro    bool  `json:"globalmacro,omitempty"`
	GlobalmacroIDs []int `json:"globalmacroids,omitempty"`
	GroupIDs       []int `json:"groupids,omitempty"`
	HostIDs        []int `json:"hostids,omitempty"`
	HostmacroIDs   []int `json:"hostmacroids,omitempty"`
	TemplateIDs    []int `json:"templateids,omitempty"`

	SelectGroups    SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

// Structure to store creation result
type hostmacroCreateResult struct {
	HostmacroIDs []int `json:"hostmacroids"`
}

// Structure to store creation global macros result
type globalmacroCreateResult struct {
	GlobalmacroIDs []int `json:"globalmacroids"`
}

// Structure to store deletion result
type hostmacroDeleteResult struct {
	HostmacroIDs []int `json:"hostmacroids"`
}

// Structure to store deletion global macros result
type globalmacroDeleteResult struct {
	GlobalmacroIDs []int `json:"globalmacroids"`
}

// UsermacroGet gets global or host macros according to the given parameters
func (z *Context) UsermacroGet(params UsermacroGetParams) ([]UsermacroObject, int, error) {

	var result []UsermacroObject

	status, err := z.request("usermacro.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostmacroCreate creates new hostmacros
func (z *Context) HostmacroCreate(params []UsermacroObject) ([]int, int, error) {

	var result hostmacroCreateResult

	status, err := z.request("usermacro.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostmacroIDs, status, nil
}

// GlobalmacroCreate creates new globalmacros
func (z *Context) GlobalmacroCreate(params []UsermacroObject) ([]int, int, error) {

	var result globalmacroCreateResult

	status, err := z.request("usermacro.createglobal", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GlobalmacroIDs, status, nil
}

// HostmacroDelete deletes hostmacros
func (z *Context) HostmacroDelete(hostmacroIDs []int) ([]int, int, error) {

	var result hostmacroDeleteResult

	status, err := z.request("usermacro.delete", hostmacroIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostmacroIDs, status, nil
}

// GlobalmacroDelete deletes globalmacros
func (z *Context) GlobalmacroDelete(globalmacroIDs []int) ([]int, int, error) {

	var result globalmacroDeleteResult

	status, err := z.request("usermacro.deleteglobal", globalmacroIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GlobalmacroIDs, status, nil
}
