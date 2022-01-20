package zabbix

// For `TemplateGetParams` field: `Evaltype`
const (
	TemplateEvaltypeAndOr = 0
	TemplateEvaltypeOr    = 2
)

// For `TemplateTag` field: `Operator`
const (
	TemplateTagOperatorContains = 0
	TemplateTagOperatorEquals   = 1
)

// TemplateObject struct is used to store template operations results
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/object
type TemplateObject struct {
	TemplateID   string `json:"templateid,omitempty"`
	Host         string `json:"host"`
	Description  string `json:"description,omitempty"`
	Name         string `json:"name,omitempty"`
	TemplateUUID string `json:"uuid,omitempty"`

	Groups          []HostgroupObject   `json:"groups,omitempty"`
	Tags            []TemplateTagObject `json:"tags,omitempty"`
	Templates       []TemplateObject    `json:"templates,omitempty"`
	ParentTemplates []TemplateObject    `json:"parentTemplates,omitempty"`
	Macros          []UsermacroObject   `json:"macros,omitempty"`
	Hosts           []HostObject        `json:"hosts,omitempty"`
}

// TemplateTagObject struct is used to store template tag data
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/object#template_tag
type TemplateTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TemplateGetParams struct is used for template get requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/get#parameters
type TemplateGetParams struct {
	GetParameters

	TemplateIDs       []string `json:"templateids,omitempty"`
	GroupIDs          []string `json:"groupids,omitempty"`
	ParentTemplateIDs []string `json:"parentTemplateids,omitempty"`
	HostIDs           []string `json:"hostids,omitempty"`
	GraphIDs          []string `json:"graphids,omitempty"`
	ItemIDs           []string `json:"itemids,omitempty"`
	TriggerIDs        []string `json:"triggerids,omitempty"`

	WithItems     bool                `json:"with_items,omitempty"`
	WithTriggers  bool                `json:"with_triggers,omitempty"`
	WithGraphs    bool                `json:"with_graphs,omitempty"`
	WithHttptests bool                `json:"with_httptests,omitempty"`
	Evaltype      int                 `json:"evaltype,omitempty"` // has defined consts, see above
	Tags          []TemplateTagObject `json:"tags,omitempty"`

	SelectGroups          SelectQuery `json:"selectGroups,omitempty"`
	SelectTags            SelectQuery `json:"selectTags,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates       SelectQuery `json:"selectTemplates,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	SelectMacros SelectQuery `json:"selectMacros,omitempty"`
	// SelectDashboards      SelectQuery `json:"selectDashboards,omitempty"` // not implemented yet
	// SelectValueMaps    SelectQuery `json:"selectValueMaps,omitempty"` // not implemented yet
}

// Structure to store creation result
type templateCreateResult struct {
	TemplateIDs []string `json:"templateids"`
}

// Structure to store deletion result
type templateDeleteResult struct {
	TemplateIDs []string `json:"templateids"`
}

// TemplateGet gets templates
func (z *Context) TemplateGet(params TemplateGetParams) ([]TemplateObject, int, error) {

	var result []TemplateObject

	status, err := z.request("template.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// TemplateCreate creates templates
func (z *Context) TemplateCreate(params []TemplateObject) ([]string, int, error) {

	var result templateCreateResult

	status, err := z.request("template.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}

// TemplateDelete deletes templates
func (z *Context) TemplateDelete(templateIDs []string) ([]string, int, error) {

	var result templateDeleteResult

	status, err := z.request("template.delete", templateIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}
