package zabbix

import ()

/* For `HostgroupObject` field: `Status` */
const (
	HostgroupFlagsPlain       = "0"
	HostgroupFlagsDiscrovered = "4"
)

/* For `HostgroupObject` field: `Internal` */
const (
	HostgroupInternalFalse = "0"
	HostgroupInternalTrue  = "1"
)

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostgroup/object */
type HostgroupObject struct {
	GroupID  string `json:"groupid,omitempty"`
	Name     string `json:"name,omitempty"`
	Flags    string `json:"flags,omitempty"`    /* has defined consts, see above */
	Internal string `json:"internal,omitempty"` /* has defined consts, see above */

	Hosts     []HostObject     `json:"hosts,omitempty"`
	Templates []TemplateObject `json:"templates,omitempty"`
}

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostgroup/get#parameters */
type HostgroupGetParams struct {
	GetParameters

	GraphIDs       []string `json:"graphids,omitempty"`
	GroupIDs       []string `json:"groupids,omitempty"`
	HostIDs        []string `json:"hostids,omitempty"`
	MaintenanceIDs []string `json:"maintenanceids,omitempty"`
	MonitoredHosts bool     `json:"monitored_hosts,omitempty"`
	RealHosts      bool     `json:"real_hosts,omitempty"`
	TemplatedHosts bool     `json:"templated_hosts,omitempty"`
	TemplateIDs    []string `json:"templateids,omitempty"`
	TriggerIDs     []string `json:"triggerids,omitempty"`

	WithApplications       bool `json:"with_applications,omitempty"`
	WithGraphs             bool `json:"with_graphs,omitempty"`
	WithHostsAndTemplates  bool `json:"with_hosts_and_templates,omitempty"`
	WithHttptests          bool `json:"with_httptests,omitempty"`
	WithItems              bool `json:"with_items,omitempty"`
	WithMonitoredHttptests bool `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems     bool `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers  bool `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems   bool `json:"with_simple_graph_items,omitempty"`
	WithTriggers           bool `json:"with_triggers,omitempty"`

	/* SelectDiscoveryRule  SelectQuery `json:"selectDiscoveryRule,omitempty"` /* not implemented yet */
	/* SelectGroupDiscovery SelectQuery `json:"selectGroupDiscovery,omitempty"` /* not implemented yet */
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

/* Structure to store creation result */
type hostgroupCreateResult struct {
	GroupIDs []string `json:"groupids"`
}

/* Structure to store deletion result */
type hostgroupDeleteResult struct {
	GroupIDs []string `json:"groupids"`
}

func (z *Zabbix) HostgroupGet(params HostgroupGetParams) ([]HostgroupObject, int, error) {

	var result []HostgroupObject

	status, err := z.request("hostgroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Zabbix) HostgroupCreate(params []HostgroupObject) ([]string, int, error) {

	var result hostgroupCreateResult

	status, err := z.request("hostgroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GroupIDs, status, nil
}

func (z *Zabbix) HostgroupDelete(groupIDs []string) ([]string, int, error) {

	var result hostgroupDeleteResult

	status, err := z.request("hostgroup.delete", groupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GroupIDs, status, nil
}
