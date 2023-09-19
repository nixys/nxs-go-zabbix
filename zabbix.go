package zabbix

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// Zabbix select constants
const (
	SelectExtendedOutput = "extend"
	SelectCount          = "count"
)

// For `GetParameters` field: `SortOrder`
const (
	GetParametersSortOrderASC  = "ASC"
	GetParametersSortOrderDESC = "DESC"
)

// Context struct is used for store settings to communicate with Zabbix API
type Context struct {
	sessionKey string
	host       string
}

// GetParameters struct is used as embedded struct for some other structs within package
//
// see for details: https://www.zabbix.com/documentation/5.0/manual/api/reference_commentary#common_get_method_parameters
type GetParameters struct {
	CountOutput            bool                   `json:"countOutput,omitempty"`
	Editable               bool                   `json:"editable,omitempty"`
	ExcludeSearch          bool                   `json:"excludeSearch,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	Limit                  int                    `json:"limit,omitempty"`
	Output                 SelectQuery            `json:"output,omitempty"`
	PreserveKeys           bool                   `json:"preservekeys,omitempty"`
	Search                 map[string]string      `json:"search,omitempty"`
	SearchByAny            bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string               `json:"sortfield,omitempty"`
	SortOrder              []string               `json:"sortorder,omitempty"` // has defined consts, see above
	StartSearch            bool                   `json:"startSearch,omitempty"`
}

// SelectQuery is used as field type in some structs
type SelectQuery interface{}

// SelectFields is used as field type in some structs
type SelectFields []string

type requestData struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Auth    string      `json:"auth,omitempty"`
	ID      int         `json:"id"`
}

type responseData struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	ID int `json:"id"`
}

// Login gets the Zabbix session
func (z *Context) Login(host, user, password string) error {

	var err error

	z.host = host

	r := UserLoginParams{
		User:     user,
		Password: password,
	}

	if z.sessionKey, _, err = z.userLogin(r); err != nil {
		return err
	}

	return nil
}

// Logout destroys the Zabbix session
func (z *Context) Logout() error {

	_, _, err := z.userLogout()

	z.sessionKey = ""

	if err != nil {
		return err
	}

	return nil
}

func (z *Context) request(method string, params interface{}, result interface{}) (int, error) {

	resp := responseData{
		Result: result,
	}

	req := requestData{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    z.sessionKey,
		ID:      1,
	}

	status, err := z.httpPost(req, &resp)
	if err != nil {
		return status, err
	}

	if resp.Error.Code != 0 {
		return status, errors.New(resp.Error.Data + " " + resp.Error.Message)
	}

	return status, nil
}

func (z *Context) httpPost(in interface{}, out interface{}) (int, error) {

	s, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", z.host, strings.NewReader(string(s)))
	if err != nil {
		return 0, err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json-rpc")

	// Make request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		if bodyBytes, err := io.ReadAll(res.Body); err == nil {
			return res.StatusCode, errors.New(string(bodyBytes))
		}
	} else {
		if out != nil {

			rawConf := make(map[string]interface{})

			dJ := json.NewDecoder(res.Body)
			if err := dJ.Decode(&rawConf); err != nil {
				return res.StatusCode, fmt.Errorf("json decode error: %v", err)
			}

			dM, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				Result:           out,
				TagName:          "json",
			})
			if err != nil {
				return res.StatusCode, fmt.Errorf("mapstructure create decoder error: %v", err)
			}

			if err := dM.Decode(rawConf); err != nil {
				return res.StatusCode, fmt.Errorf("mapstructure decode error: %v", err)
			}
		}
	}

	return res.StatusCode, nil
}
