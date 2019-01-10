package zabbix

// For `HostObject` field: `Available`
const (
	HostAvailableUnknown     = "0"
	HostAvailableAvailable   = "1"
	HostAvailableUnavailable = "2"
)

// For `HostObject` field: `Flags`
const (
	HostFlagsPlain      = "0"
	HostFlagsDiscovered = "4"
)

// For `HostObject` field: `Flags`
const (
	HostInventoryModeDisabled  = "-1"
	HostInventoryModeManual    = "0"
	HostInventoryModeAutomatic = "1"
)

// For `HostObject` field: `IpmiAuthtype`
const (
	HostIpmiAuthtypeDefault  = "-1"
	HostIpmiAuthtypeNone     = "0"
	HostIpmiAuthtypeMD2      = "1"
	HostIpmiAuthtypeMD5      = "2"
	HostIpmiAuthtypeStraight = "4"
	HostIpmiAuthtypeOEM      = "5"
	HostIpmiAuthtypeRMCP     = "6"
)

// For `HostObject` field: `IpmiAvailable`
const (
	HostIpmiAvailableUnknown     = "0"
	HostIpmiAvailableAvailable   = "1"
	HostIpmiAvailableUnavailable = "2"
)

// For `HostObject` field: `IpmiPrivilege`
const (
	HostIpmiPrivilegeCallback = "1"
	HostIpmiPrivilegeUser     = "2"
	HostIpmiPrivilegeOperator = "3"
	HostIpmiPrivilegeAdmin    = "4"
	HostIpmiPrivilegeOEM      = "5"
)

// For `HostObject` field: `JmxAvailable`
const (
	HostJmxAvailableUnknown     = "0"
	HostJmxAvailableAvailable   = "1"
	HostJmxAvailableUnavailable = "2"
)

// For `HostObject` field: `MaintenanceStatus`
const (
	HostMaintenanceStatusDisable = "0"
	HostMaintenanceStatusEnable  = "1"
)

// For `HostObject` field: `MaintenanceType`
const (
	HostMaintenanceTypeDataCollectionEnabled  = "0"
	HostMaintenanceTypeDataCollectionDisabled = "1"
)

// For `HostObject` field: `SnmpAvailable`
const (
	HostSnmpAvailableUnknown     = "0"
	HostSnmpAvailableAvailable   = "1"
	HostSnmpAvailableUnavailable = "2"
)

// For `HostObject` field: `Status`
const (
	HostStatusMonitored   = "0"
	HostStatusUnmonitored = "1"
)

// HostObject struct is used to store host operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/host/object#host
type HostObject struct {
	HostID            string `json:"hostid,omitempty"`
	Host              string `json:"host,omitempty"`
	Available         string `json:"available,omitempty"` // has defined consts, see above
	Description       string `json:"description,omitempty"`
	DisableUntil      string `json:"disable_until,omitempty"`
	Error             string `json:"error,omitempty"`
	ErrorsFrom        string `json:"errors_from,omitempty"`
	Flags             string `json:"flags,omitempty"`          // has defined consts, see above
	InventoryMode     string `json:"inventory_mode,omitempty"` // has defined consts, see above
	IpmiAuthtype      string `json:"ipmi_authtype,omitempty"`  // has defined consts, see above
	IpmiAvailable     string `json:"ipmi_available,omitempty"` // has defined consts, see above
	IpmiDisableUntil  string `json:"ipmi_disable_until,omitempty"`
	IpmiError         string `json:"ipmi_error,omitempty"`
	IpmiErrorsFrom    string `json:"ipmi_errors_from,omitempty"`
	IpmiPassword      string `json:"ipmi_password,omitempty"`
	IpmiPrivilege     string `json:"ipmi_privilege,omitempty"` // has defined consts, see above
	IpmiUsername      string `json:"ipmi_username,omitempty"`
	JmxAvailable      string `json:"jmx_available,omitempty"` // has defined consts, see above
	JmxDisableUntil   string `json:"jmx_disable_until,omitempty"`
	JmxError          string `json:"jmx_error,omitempty"`
	JmxErrorsFrom     string `json:"jmx_errors_from,omitempty"`
	MaintenanceFrom   string `json:"maintenance_from,omitempty"`
	MaintenanceStatus string `json:"maintenance_status,omitempty"` // has defined consts, see above
	MaintenanceType   string `json:"maintenance_type,omitempty"`   // has defined consts, see above
	MaintenanceID     string `json:"maintenanceid,omitempty"`
	Name              string `json:"name,omitempty"`
	ProxyHostID       string `json:"proxy_hostid,omitempty"`
	SnmpAvailable     string `json:"snmp_available,omitempty"` // has defined consts, see above
	SnmpDisableUntil  string `json:"snmp_disable_until,omitempty"`
	SnmpError         string `json:"snmp_error,omitempty"`
	SnmpErrorsFrom    string `json:"snmp_errors_from,omitempty"`
	Status            string `json:"status,omitempty"` // has defined consts, see above

	Groups     []HostgroupObject     `json:"groups,omitempty"`
	Interfaces []HostinterfaceObject `json:"interfaces,omitempty"`
	Macros     []HostmacroObject     `json:"macros,omitempty"`

	ParentTemplates []TemplateObject `json:"parentTemplates,omitempty"` // Used to store result for `get` operations
	Templates       []TemplateObject `json:"templates,omitempty"`       // Used for `create` operations
}

// HostGetParams struct is used for host get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/host/get#parameters
type HostGetParams struct {
	GetParameters

	GroupIDs       []string `json:"groupids,omitempty"`
	ApplicationIDs []string `json:"applicationids,omitempty"`
	DserviceIDs    []string `json:"dserviceids,omitempty"`
	GraphIDs       []string `json:"graphids,omitempty"`
	HostIDs        []string `json:"hostids,omitempty"`
	HttptestIDs    []string `json:"httptestids,omitempty"`
	InterfaceIDs   []string `json:"interfaceids,omitempty"`
	ItemIDs        []string `json:"itemids,omitempty"`
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`
	MonitoredHosts bool     `json:"monitored_hosts,omitempty"`
	ProxyHosts     bool     `json:"proxy_hosts,omitempty"`
	ProxyIDs       []string `json:"proxyids,omitempty"`
	TemplatedHosts bool     `json:"templated_hosts,omitempty"`
	TemplateIDs    []string `json:"templateids,omitempty"`
	TriggerIDs     []string `json:"triggerids,omitempty"`

	WithItems              bool `json:"with_items,omitempty"`
	WithApplications       bool `json:"with_applications,omitempty"`
	WithGraphs             bool `json:"with_graphs,omitempty"`
	WithHttptests          bool `json:"with_httptests,omitempty"`
	WithMonitoredHttptests bool `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems     bool `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers  bool `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems   bool `json:"with_simple_graph_items,omitempty"`
	WithTriggers           bool `json:"with_triggers,omitempty"`
	WithInventory          bool `json:"withInventory,omitempty"`

	SelectGroups          SelectQuery `json:"selectGroups,omitempty"`
	SelectApplications    SelectQuery `json:"selectApplications,omitempty"`
	SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule   SelectQuery `json:"selectDiscoveryRule ,omitempty"`
	SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"`
	SelectHostDiscovery   SelectQuery `json:"selectHostDiscovery ,omitempty"`
	SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"`
	SelectInterfaces      SelectQuery `json:"selectInterfaces,omitempty"`
	SelectInventory       SelectQuery `json:"selectInventory,omitempty"`
	SelectItems           SelectQuery `json:"selectItems,omitempty"`
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectScreens         SelectQuery `json:"selectScreens,omitempty"`
	SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"`
}

// Structure to store creation result
type hostCreateResult struct {
	HostIDs []string `json:"hostids"`
}

// Structure to store deletion result
type hostDeleteResult struct {
	HostIDs []string `json:"hostids"`
}

// HostGet gets hosts
func (z *Zabbix) HostGet(params HostGetParams) ([]HostObject, int, error) {

	var result []HostObject

	status, err := z.request("host.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostCreate creates hosts
func (z *Zabbix) HostCreate(params []HostObject) ([]string, int, error) {

	var result hostCreateResult

	status, err := z.request("host.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostIDs, status, nil
}

// HostDelete deletes hosts
func (z *Zabbix) HostDelete(hostIDs []string) ([]string, int, error) {

	var result hostDeleteResult

	status, err := z.request("host.delete", hostIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostIDs, status, nil
}
