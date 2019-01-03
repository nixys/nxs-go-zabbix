package zabbix

import ()

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/template/object */
type TemplateObject struct {
	TemplateID  string `json:"templateid,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`

	Groups          []HostgroupObject `json:"groups,omitempty"`
	Hosts           []HostObject      `json:"hosts,omitempty"`
	Templates       []TemplateObject  `json:"templates,omitempty"`
	ParentTemplates []TemplateObject  `json:"parentTemplates,omitempty"`
	Macros          []HostmacroObject `json:"macros,omitempty"`
}

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/template/get#parameters */
type TemplateGetParams struct {
	GetParameters

	TemplateIDs       []string `json:"templateids,omitempty"`
	GroupIDs          []string `json:"groupids,omitempty"`
	ParentTemplateIDs []string `json:"parentTemplateids,omitempty"`
	HostIDs           []string `json:"hostids,omitempty"`
	GraphIDs          []string `json:"graphids,omitempty"`
	ItemIDs           []string `json:"itemids,omitempty"`
	TriggerIDs        []string `json:"triggerids,omitempty"`

	WithItems     bool `json:"with_items,omitempty"`
	WithTriggers  bool `json:"with_triggers,omitempty"`
	WithGraphs    bool `json:"with_graphs,omitempty"`
	WithHttptests bool `json:"with_httptests,omitempty"`

	SelectGroups          SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates       SelectQuery `json:"selectTemplates,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`

	/* SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` /* not implemented yet */
	/* SelectItems           SelectQuery `json:"selectItems,omitempty"` /* not implemented yet */
	/* SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` /* not implemented yet */
	/* SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` /* not implemented yet */
	/* SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` /* not implemented yet */
	/* SelectApplications    SelectQuery `json:"selectApplications,omitempty"` /* not implemented yet */
	/* SelectScreens         SelectQuery `json:"selectScreens,omitempty"` /* not implemented yet */
}

/* Structure to store creation result */
type TemplateCreateResult struct {
	TemplateIDs []string `json:"templateids"`
}

/* Structure to store deletion result */
type TemplateDeleteResult struct {
	TemplateIDs []string `json:"templateids"`
}

func (z *Zabbix) TemplateGet(params TemplateGetParams) ([]TemplateObject, int, error) {

	var result []TemplateObject

	status, err := z.request("template.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Zabbix) TemplateCreate(params []TemplateObject) ([]string, int, error) {

	var result TemplateCreateResult

	status, err := z.request("template.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}

func (z *Zabbix) TemplateDelete(templateIDs []string) ([]string, int, error) {

	var result TemplateDeleteResult

	status, err := z.request("template.delete", templateIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}
