package zabbix

// For `ItemPrototypeObject` field: `ItemPrototypeType`
const (
	ItemPrototypeTypeZabbixAgent       = 0
	ItemPrototypeTypeZabbixTrapper     = 2
	ItemPrototypeTypeSimpleCheck       = 3
	ItemPrototypeTypeZabbixInternal    = 5
	ItemPrototypeTypeZabbixAgentActive = 7
	ItemPrototypeTypeExternalCheck     = 10
	ItemPrototypeTypeDBMonitor         = 11
	ItemPrototypeTypeIPMIAgent         = 12
	ItemPrototypeTypeSSHAgent          = 13
	ItemPrototypeTypeTelnetAgent       = 14
	ItemPrototypeTypeCalculated        = 15
	ItemPrototypeTypeJMXAgent          = 16
	ItemPrototypeTypeSNMPTrap          = 17
	ItemPrototypeTypeDependentItem     = 18
	ItemPrototypeTypeHTTPAgent         = 19
	ItemPrototypeTypeSNMPAgent         = 20
	ItemPrototypeTypeScript            = 21
	ItemPrototypeTypeBrowser           = 22
)

// For `ItemPrototypeObject` field: `ValueType`
const (
	ItemPrototypeValueTypeNumericFloat    = 0
	ItemPrototypeValueTypeCharacter       = 1
	ItemPrototypeValueTypeLog             = 2
	ItemPrototypeValueTypeNumericUnsigned = 3
	ItemPrototypeValueTypeText            = 4
	ItemPrototypeValueTypeBinary          = 5
)

// For `ItemPrototypeObject` field: `Status`
const (
	ItemPrototypeStatusEnabled  = 0
	ItemPrototypeStatusDisabled = 1
)

// For `ItemPrototypeGetParams` field: `Evaltype`
const (
	ItemPrototypeEvaltypeAndOr = 0
	ItemPrototypeEvaltypeOr    = 2
)

// TemplateObject struct is used to store template operations results
//
// see: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/itemprototype/object
type ItemPrototypeObject struct {
	ItemPrototypeID   int    `json:"itemid,omitempty"`
	Delay             string `json:"delay,omitempty"`
	HostID            int    `json:"hostid,omitempty"`
	InterfaceID       int    `json:"interfaceid,omitempty"`
	Key               string `json:"key_,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	ItemPrototypeType int    `json:"type,omitempty"`
	ValueType         int    `json:"value_type,omitempty"`
	SnmpOID           string `json:"snmp_oid,omitempty"`
	Status            int    `json:"status,omitempty"`
	TemplateID        int    `json:"templateid,omitempty"`
	UUID              string `json:"uuid,omitempty"`

	Groups          []HostgroupObject        `json:"groups,omitempty"`
	Tags            []ItemPrototypeTagObject `json:"tags,omitempty"`
	Templates       []TemplateObject         `json:"templates,omitempty"`
	ParentTemplates []TemplateObject         `json:"parentTemplates,omitempty"`
	Macros          []UsermacroObject        `json:"macros,omitempty"`
	Hosts           []HostObject             `json:"hosts,omitempty"`
}

// ItemPrototypeTagObject struct is used to store item tag data
//
// see: https://www.zabbix.com/documentation/7.0/manual/api/reference/template/object#template_tag
type ItemPrototypeTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// TemplateGetParams struct is used for template get requests
//
// see: https://www.zabbix.com/documentation/7.0/en/manual/api/reference/itemprototype/get#parameters
type ItemPrototypeGetParams struct {
	GetParameters

	ItemPrototypeIDs  []int `json:"itemids,omitempty"`
	DiscoveryIDs      []int `json:"discoveryids,omitempty"`
	GroupIDs          []int `json:"groupids,omitempty"`
	TemplateIDs       []int `json:"templateids,omitempty"`
	HostIDs           []int `json:"hostids,omitempty"`
	ProxyIDs          []int `json:"proxyids,omitempty"`
	GraphIDs          []int `json:"graphids,omitempty"`
	TriggerIDs        []int `json:"triggerids,omitempty"`

	Inherited          bool                     `json:"inherited,omitempty"`
	Templated          bool                     `json:"templated,omitempty"`
	Monitored          bool                     `json:"monitored,omitempty"`
	Group              string                   `json:"group,omitempty"`
	Host               string                   `json:"host,omitempty"`
	WithItemPrototypes bool                     `json:"with_items,omitempty"`
	WithTriggers       bool                     `json:"with_triggers,omitempty"`
	Evaltype           int                      `json:"evaltype,omitempty"` // has defined consts, see above
	Tags               []ItemPrototypeTagObject `json:"tags,omitempty"`

	SelectTags            SelectQuery `json:"selectTags,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectInterfaces      SelectQuery `json:"selectInterfaces,omitempty"`
	selectTriggers        SelectQuery `json:"selectTriggers,omitempty"`

	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItemPrototypes  SelectQuery `json:"selectItemPrototypes,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
}

// Structure to store creation result
type itemPrototypeCreateResult struct {
	ItemPrototypeIDs []int `json:"itemids"`
}

type itemPrototypeUpdateResult itemPrototypeCreateResult
type itemPrototypeDeleteResult itemPrototypeCreateResult

/*// Structure to store deletion result
type itemDeleteResult struct {
	ItemPrototypeIDs []int `json:"itemids"`
}*/

// ItemPrototypeGet gets items
func (z *Context) ItemPrototypeGet(params ItemPrototypeGetParams) ([]ItemPrototypeObject, int, error) {

	var result []ItemPrototypeObject

	status, err := z.request("itemprototype.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// ItemPrototypeCreate creates items
func (z *Context) ItemPrototypeCreate(params []ItemPrototypeObject) ([]int, int, error) {

	var result itemPrototypeCreateResult

	status, err := z.request("itemprototype.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemPrototypeIDs, status, nil
}

// ItemPrototypeUpdate updates items
func (z *Context) ItemPrototypeUpdate(params []ItemPrototypeObject) ([]int, int, error) {

	var result itemPrototypeUpdateResult

	status, err := z.request("itemprototype.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemPrototypeIDs, status, nil
}

// TemplateDelete deletes templates
func (z *Context) ItemPrototypeDelete(itemIDs []int) ([]int, int, error) {

	var result itemPrototypeDeleteResult

	status, err := z.request("itemprototype.delete", itemIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemPrototypeIDs, status, nil
}
