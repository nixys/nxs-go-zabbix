package zabbix

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Zabbix struct {
	sessionKey string
	host       string
}

type requestData struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Auth    string      `json:"auth,omitempty"`
	ID      int         `json:"id"`
}

type responseData struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	ID int `json:"id"`
}

/* see for details: https://www.zabbix.com/documentation/2.4/manual/api/reference_commentary#data_types */
type GetParameters struct {
	Editable               bool                   `json:"editable,omitempty"`
	ExcludeSearch          bool                   `json:"excludeSearch,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	Limit                  int                    `json:"limit,omitempty"`
	NodeIDs                []string               `json:"nodeids,omitempty"`
	Output                 SelectQuery            `json:"output,omitempty"`
	PreserveKeys           bool                   `json:"preservekeys,omitempty"`
	Search                 map[string]string      `json:"search,omitempty"`
	SearchByAny            bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string               `json:"sortfield,omitempty"`
	SortOrder              string                 `json:"sortorder,omitempty"`
	StartSearch            bool                   `json:"startSearch,omitempty"`
}

type SelectQuery interface{}

type SelectFields []string

const (
	SelectExtendedOutput = "extend"
	SelectCount          = "count"
)

func (r *responseData) getResult(result interface{}) error {

	if err := json.Unmarshal(r.Result, result); err != nil {
		return err
	}

	return nil
}

func (z *Zabbix) Login(host, user, password string) error {

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

func (z *Zabbix) Logout() error {

	_, _, err := z.userLogout()

	z.sessionKey = ""

	if err != nil {
		return err
	}

	return nil
}

func (z *Zabbix) request(method string, params interface{}, result interface{}) (int, error) {

	resp := responseData{}

	r := requestData{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    z.sessionKey,
		ID:      1,
	}

	status, err := z.httpPost(r, &resp)
	if err != nil {
		return status, err
	}

	if resp.Error.Code != 0 {
		err = errors.New(resp.Error.Data + " " + resp.Error.Message)
		return status, err
	}

	if err := resp.getResult(result); err != nil {
		return status, err
	}

	return status, nil
}

func (z *Zabbix) httpPost(in interface{}, out interface{}) (int, error) {

	var bodyBytes []byte

	s, err := json.Marshal(in)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", z.host, strings.NewReader(string(s)))
	if err != nil {
		return 0, err
	}

	/* Set headers */
	req.Header.Add("Content-Type", "application/json-rpc")

	/* Make request */
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, err = ioutil.ReadAll(res.Body)
		if err == nil {
			err = errors.New(string(bodyBytes))
		}
	} else {
		if out != nil {
			decoder := json.NewDecoder(res.Body)
			err = decoder.Decode(out)
		}
	}

	if err != nil {
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}
