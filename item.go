package zabbix

// For `ItemObject` field: `ItemType`
const (
	ItemTypeZabbixAgent       = 0
	ItemTypeZabbixTrapper     = 2
	ItemTypeSimpleCheck       = 3
	ItemTypeZabbixInternal    = 5
	ItemTypeZabbixAgentActive = 7
	ItemTypeWebItem           = 9
	ItemTypeExternalCheck     = 10
	ItemTypeDBMonitor         = 11
	ItemTypeIPMIAgent         = 12
	ItemTypeSSHAgent          = 13
	ItemTypeTelnetAgent       = 14
	ItemTypeCalculated        = 15
	ItemTypeJMXAgent          = 16
	ItemTypeSNMPTrap          = 17
	ItemTypeDependentItem     = 18
	ItemTypeHTTPAgent         = 19
	ItemTypeSNMPAgent         = 20
	ItemTypeScript            = 21
	ItemTypeBrowser           = 22
)

// For `ItemObject` field: `ValueType`
const (
	ItemValueTypeNumericFloat    = 0
	ItemValueTypeCharacter       = 1
	ItemValueTypeLog             = 2
	ItemValueTypeNumericUnsigned = 3
	ItemValueTypeText            = 4
	ItemValueTypeBinary          = 5
)

// For `ItemObject` field: `Status`
const (
	ItemStatusEnabled  = 0
	ItemStatusDisabled = 1
)

// For `ItemGetParams` field: `Evaltype`
const (
	ItemEvaltypeAndOr = 0
	ItemEvaltypeOr    = 2
)

// TemplateObject struct is used to store template operations results
//
// see: https://www.zabbix.com/documentation/7.0/manual/api/reference/template/object
type ItemObject struct {
	ItemID      int    `json:"itemid,omitempty"`
	Delay       string `json:"delay,omitempty"`
	HostID      int    `json:"hostid,omitempty"`
	InterfaceID int    `json:"interfaceid,omitempty"`
	Key         string `json:"key_,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ItemType    int    `json:"type,omitempty"`
	ValueType   int    `json:"value_type,omitempty"`
	SnmpOID     string `json:"snmp_oid,omitempty"`
	Status      int    `json:"status,omitempty"`
	TemplateID  int    `json:"templateid,omitempty"`
	UUID        string `json:"uuid,omitempty"`

	Groups          []HostgroupObject   `json:"groups,omitempty"`
	Tags            []ItemTagObject     `json:"tags,omitempty"`
	Templates       []TemplateObject    `json:"templates,omitempty"`
	ParentTemplates []TemplateObject    `json:"parentTemplates,omitempty"`
	Macros          []UsermacroObject   `json:"macros,omitempty"`
	Hosts           []HostObject        `json:"hosts,omitempty"`
}

// ItemTagObject struct is used to store item tag data
//
// see: https://www.zabbix.com/documentation/7.0/manual/api/reference/template/object#template_tag
type ItemTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// TemplateGetParams struct is used for template get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/template/get#parameters
type ItemGetParams struct {
	GetParameters

	ItemIDs           []int `json:"itemids,omitempty"`
	GroupIDs          []int `json:"groupids,omitempty"`
	TemplateIDs       []int `json:"templateids,omitempty"`
	HostIDs           []int `json:"hostids,omitempty"`
	ProxyIDs          []int `json:"proxyids,omitempty"`
	GraphIDs          []int `json:"graphids,omitempty"`
	TriggerIDs        []int `json:"triggerids,omitempty"`

	Inherited     bool                `json:"inherited,omitempty"`
	Templated     bool                `json:"templated,omitempty"`
	Monitored     bool                `json:"monitored,omitempty"`
	Group         string              `json:"group,omitempty"`
	Host          string              `json:"host,omitempty"`
	WithItems     bool                `json:"with_items,omitempty"`
	WithTriggers  bool                `json:"with_triggers,omitempty"`
	Evaltype      int                 `json:"evaltype,omitempty"` // has defined consts, see above
	Tags          []ItemTagObject     `json:"tags,omitempty"`

	SelectTags            SelectQuery `json:"selectTags,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectInterfaces      SelectQuery `json:"selectInterfaces,omitempty"`
	selectTriggers        SelectQuery `json:"selectTriggers,omitempty"`

	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
}

// Structure to store creation result
type itemCreateResult struct {
	ItemIDs []int `json:"itemids"`
}

type itemUpdateResult itemCreateResult
type itemDeleteResult itemCreateResult

/*// Structure to store deletion result
type itemDeleteResult struct {
	ItemIDs []int `json:"itemids"`
}*/

// ItemGet gets items
func (z *Context) ItemGet(params ItemGetParams) ([]ItemObject, int, error) {

	var result []ItemObject

	status, err := z.request("item.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// ItemCreate creates items
func (z *Context) ItemCreate(params []ItemObject) ([]int, int, error) {

	var result itemCreateResult

	status, err := z.request("item.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}

// ItemUpdate updates items
func (z *Context) ItemUpdate(params []ItemObject) ([]int, int, error) {

	var result itemUpdateResult

	status, err := z.request("item.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}

// TemplateDelete deletes templates
func (z *Context) ItemDelete(itemIDs []int) ([]int, int, error) {

	var result itemDeleteResult

	status, err := z.request("item.delete", itemIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}
