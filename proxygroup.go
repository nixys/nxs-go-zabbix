package zabbix

// For `ProxygroupObject` field: `Status`
const (
	ProxygroupFlagsPlain       = 0
	ProxygroupFlagsDiscrovered = 4
)

// For `ProxygroupObject` field: `Internal`
const (
	ProxygroupStateUnknown    = 0
	ProxygroupStateOffline    = 1
	ProxygroupStateRecovering = 2
	ProxygroupStateOnline     = 3
	ProxygroupStateDegrading  = 4
)

// ProxygroupObject struct is used to store Proxygroup operations results
//
// see: https://www.zabbix.com/documentation/current/en/manual/api/reference/proxygroup/object
type ProxygroupObject struct {
	ProxyGroupID   int    `json:"proxy_groupid,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	FailoverDelay  string `json:"failover_delay,omitempty"`
	MinOnline      string `json:"min_online,omitempty"`
	State          int    `json:"state,omitempty"`
}

// ProxygroupGetParams struct is used for Proxygroup get requests
//
// see: https://www.zabbix.com/documentation/current/en/manual/api/reference/proxygroup/get#parameters
type ProxygroupGetParams struct {
	GetParameters

	ProxyGroupIDs  []int `json:"proxy_groupids,omitempty"`
	ProxyIDs       []int `json:"proxyids,omitempty"`
	MaintenanceIDs []int `json:"maintenanceids,omitempty"`

	// SelectDiscoveryRule  SelectQuery `json:"selectDiscoveryRule,omitempty"` // not implemented yet
	// SelectGroupDiscovery SelectQuery `json:"selectGroupDiscovery,omitempty"` // not implemented yet
	SelectProxies   SelectQuery `json:"selectProxies,omitempty"`
}

// Structure to store creation result
type ProxygroupCreateResult struct {
	ProxyGroupIDs []int `json:"ProxyGroupIDs"`
}

// Structure to store deletion result
type ProxygroupDeleteResult struct {
	ProxyGroupIDs []int `json:"ProxyGroupIDs"`
}

// ProxygroupGet gets Proxygroups
func (z *Context) ProxygroupGet(params ProxygroupGetParams) ([]ProxygroupObject, int, error) {

	var result []ProxygroupObject

	status, err := z.request("proxygroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// ProxygroupCreate creates Proxygroups
func (z *Context) ProxygroupCreate(params []ProxygroupObject) ([]int, int, error) {

	var result ProxygroupCreateResult

	status, err := z.request("proxygroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ProxyGroupIDs, status, nil
}

// ProxygroupDelete deletes Proxygroups
func (z *Context) ProxygroupDelete(ProxyGroupIDs []int) ([]int, int, error) {

	var result ProxygroupDeleteResult

	status, err := z.request("proxygroup.delete", ProxyGroupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ProxyGroupIDs, status, nil
}
