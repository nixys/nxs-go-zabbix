package zabbix

import ()

/* For `HostinterfaceObject` field: `Main` */
const (
	HostinterfaceMainNotDefault = "0"
	HostinterfaceMainDefault    = "1"
)

/* For `HostinterfaceObject` field: `Type` */
const (
	HostinterfaceTypeAgent = "1"
	HostinterfaceTypeSNMP  = "2"
	HostinterfaceTypeIPMI  = "3"
	HostinterfaceTypeJMX   = "4"
)

/* For `HostinterfaceObject` field: `UseIP` */
const (
	HostinterfaceUseipDNS = "0"
	HostinterfaceUseipIP  = "1"
)

/* For `HostinterfaceObject` field: `Bulk` */
const (
	HostinterfaceBulkDontUse = "0"
	HostinterfaceBulkUse     = "1"
)

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostinterface/object#hostinterface */
type HostinterfaceObject struct {
	InterfaceID string `json:"interfaceid,omitempty"`
	DNS         string `json:"dns"`
	HostID      string `json:"hostid,omitempty"`
	IP          string `json:"ip"`
	Main        string `json:"main"` /* has defined consts, see above */
	Port        string `json:"port"`
	Type        string `json:"type"`           /* has defined consts, see above */
	UseIP       string `json:"useip"`          /* has defined consts, see above */
	Bulk        string `json:"bulk,omitempty"` /* has defined consts, see above */

	/* Items []ItemObject `json:"items,omitempty"` /* not implemented yet */
	Hosts []HostObject `json:"hosts,omitempty"`
}

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostinterface/get#parameters */
type HostinterfaceGetParams struct {
	GetParameters

	HostIDs      []string `json:"hostids,omitempty"`
	InterfaceIDs []string `json:"interfaceids,omitempty"`
	ItemIDs      []string `json:"itemids,omitempty"`
	TriggerIDs   []string `json:"triggerids,omitempty"`

	/* SelectItems SelectQuery `json:"selectItems,omitempty"` /* not implemented yet */
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

/* Structure to store creation result */
type hostinterfaceCreateResult struct {
	InterfaceIDs []string `json:"interfaceids"`
}

/* Structure to store deletion result */
type hostinterfaceDeleteResult struct {
	InterfaceIDs []string `json:"interfaceids"`
}

func (z *Zabbix) HostinterfaceGet(params HostinterfaceGetParams) ([]HostinterfaceObject, int, error) {

	var result []HostinterfaceObject

	status, err := z.request("hostinterface.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Zabbix) HostinterfaceCreate(params []HostinterfaceObject) ([]string, int, error) {

	var result hostinterfaceCreateResult

	status, err := z.request("hostinterface.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}

func (z *Zabbix) HostinterfaceDelete(hostinterfaceIDs []string) ([]string, int, error) {

	var result hostinterfaceDeleteResult

	status, err := z.request("hostinterface.delete", hostinterfaceIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}
