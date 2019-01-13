package zabbix

import (
	"os"
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

	if err := z.Login(zbxHost, zbxUsername, zbxPassword); err != nil {
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
