package zabbix

// GlobalmacroObject struct is used to store globalmacro get operations results (not implemented yet)
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/object#global_macro
type GlobalmacroObject struct {
	GlobalmacroID int    `json:"globalmacroid,omitempty"`
	Macro         string `json:"macro ,omitempty"`
	Value         string `json:"value ,omitempty"`
}

// HostmacroObject struct is used to store hostmacro operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/object#host_macro
type HostmacroObject struct {
	HostmacroID int    `json:"hostmacroid,omitempty"`
	HostID      int    `json:"hostid,omitempty"`
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`

	Groups    []HostgroupObject `json:"groups,omitempty"`
	Hosts     []HostObject      `json:"hosts,omitempty"`
	Templates []TemplateObject  `json:"templates,omitempty"`
}

// HostmacroGetParams struct is used for hostmacro get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/get#parameters
type HostmacroGetParams struct {
	GetParameters

	GroupIDs     []int `json:"groupids,omitempty"`
	HostIDs      []int `json:"hostids,omitempty"`
	HostmacroIDs []int `json:"hostmacroids,omitempty"`
	TemplateIDs  []int `json:"templateids,omitempty"`

	SelectGroups    SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

// Structure to store creation result
type hostmacroCreateResult struct {
	HostmacroIDs []int `json:"hostmacroids"`
}

// Structure to store deletion result
type hostmacroDeleteResult struct {
	HostmacroIDs []int `json:"hostmacroids"`
}

// HostmacroGet gets hostmacros
func (z *Context) HostmacroGet(params HostmacroGetParams) ([]HostmacroObject, int, error) {

	var result []HostmacroObject

	status, err := z.request("usermacro.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostmacroCreate creates new hostmacros
func (z *Context) HostmacroCreate(params []HostmacroObject) ([]int, int, error) {

	var result hostmacroCreateResult

	status, err := z.request("usermacro.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostmacroIDs, status, nil
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
