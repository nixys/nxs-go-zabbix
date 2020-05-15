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
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/template/object
type TemplateObject struct {
	TemplateID  int    `json:"templateid,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`

	Groups          []HostgroupObject   `json:"groups,omitempty"`
	Tags            []TemplateTagObject `json:"tags,omitempty"`
	Templates       []TemplateObject    `json:"templates,omitempty"`
	ParentTemplates []TemplateObject    `json:"parentTemplates,omitempty"`
	Macros          []UsermacroObject   `json:"macros,omitempty"`
	Hosts           []HostObject        `json:"hosts,omitempty"`
}

// TemplateTagObject struct is used to store template tag data
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/template/object#template_tag
type TemplateTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TemplateGetParams struct is used for template get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/template/get#parameters
type TemplateGetParams struct {
	GetParameters

	TemplateIDs       []int `json:"templateids,omitempty"`
	GroupIDs          []int `json:"groupids,omitempty"`
	ParentTemplateIDs []int `json:"parentTemplateids,omitempty"`
	HostIDs           []int `json:"hostids,omitempty"`
	GraphIDs          []int `json:"graphids,omitempty"`
	ItemIDs           []int `json:"itemids,omitempty"`
	TriggerIDs        []int `json:"triggerids,omitempty"`

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
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`

	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
}

// Structure to store creation result
type templateCreateResult struct {
	TemplateIDs []int `json:"templateids"`
}

// Structure to store deletion result
type templateDeleteResult struct {
	TemplateIDs []int `json:"templateids"`
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
func (z *Context) TemplateCreate(params []TemplateObject) ([]int, int, error) {

	var result templateCreateResult

	status, err := z.request("template.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}

// TemplateDelete deletes templates
func (z *Context) TemplateDelete(templateIDs []int) ([]int, int, error) {

	var result templateDeleteResult

	status, err := z.request("template.delete", templateIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}
