package zabbix

import (
	"testing"
)

func TestUsermediaCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Get
	testUsermediaGet(t, z)
}

func testUsermediaGet(t *testing.T, z Context) []MediaObject {

	mObjects, _, err := z.UsermediaGet(UsermediaGetParams{
		GetParameters: GetParameters{
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Usermedia get error:", err)
	} else {
		if len(mObjects) == 0 {
			t.Error("Usermedia get error: unable to find usermedia")
		} else {
			t.Logf("Usermedia get: success")
		}
	}

	return mObjects
}
