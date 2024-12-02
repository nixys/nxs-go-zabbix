package zabbix

import (
	"os"
	"strconv"
	"testing"
)

func loginTest(z *Context, t *testing.T) {

	zbxHost := os.Getenv("ZABBIX_HOST")
	if zbxHost == "" {
		t.Fatal("Login error: undefined env var `ZABBIX_HOST`")
	}

	zbxUsername := os.Getenv("ZABBIX_USERNAME")
	if zbxUsername == "" {
		t.Fatal("Login error: undefined env var `ZABBIX_USERNAME`")
	}

	zbxPassword := os.Getenv("ZABBIX_PASSWORD")
	if zbxPassword == "" {
		t.Fatal("Login error: undefined env var `ZABBIX_PASSWORD`")
	}

	zbxInsecureSkipVerify := os.Getenv("ZABBIX_INSECURE_SKIP_VERIFY")
	if zbxInsecureSkipVerify == "" {
		t.Fatal("Login error: undefined env var `ZABBIX_INSECURE_SKIP_VERIFY`")
	}

	// Convert to bool
	zbxInsecureSkipVerifyBool, err := strconv.ParseBool(zbxInsecureSkipVerify)
	if err != nil {
		t.Fatal("Login error: invalid env var `ZABBIX_INSECURE_SKIP_VERIFY`")
	}

	loginParams := LoginParams{
		Host:               zbxHost,
		User:               zbxUsername,
		Password:           zbxPassword,
		InsecureSkipVerify: zbxInsecureSkipVerifyBool,
	}

	if err := z.Login(loginParams); err != nil {
		t.Fatal("Login error: ", err)
	} else {
		t.Logf("Login: success")
	}
}

func logoutTest(z *Context, t *testing.T) {

	if err := z.Logout(); err != nil {
		t.Fatal("Logout error: ", err)
	} else {
		t.Logf("Logout: success")
	}
}
