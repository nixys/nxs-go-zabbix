package zabbix

import (
	"reflect"
	"testing"
)

const (
	testHostgroupName = "testHostgroup"
)

func TestHostgroupCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Create and delete
	hgCreatedIDs := testHostgroupCreate(t, z)
	defer testHostgroupDelete(t, z, hgCreatedIDs)

	// Get
	testHostgroupGet(t, z, hgCreatedIDs)
}

func testHostgroupCreate(t *testing.T, z Context) []int {

	hgCreatedIDs, _, err := z.HostgroupCreate([]HostgroupObject{
		{
			Name: testHostgroupName,
		},
	})
	if err != nil {
		t.Fatal("Hostgroup create error:", err)
	}

	if len(hgCreatedIDs) == 0 {
		t.Fatal("Hostgroup create error: empty IDs array")
	}

	t.Logf("Hostgroup create: success")

	return hgCreatedIDs
}

func testHostgroupDelete(t *testing.T, z Context, hgCreatedIDs []int) []int {

	hgDeletedIDs, _, err := z.HostgroupDelete(hgCreatedIDs)
	if err != nil {
		t.Fatal("Hostgroup delete error:", err)
	}

	if len(hgDeletedIDs) == 0 {
		t.Fatal("Hostgroup delete error: empty IDs array")
	}

	if reflect.DeepEqual(hgDeletedIDs, hgCreatedIDs) == false {
		t.Fatal("Hostgroup delete error: IDs arrays for created and deleted hostgroup are mismatch")
	}

	t.Logf("Hostgroup delete: success")

	return hgDeletedIDs
}

func testHostgroupGet(t *testing.T, z Context, hgCreatedIDs []int) []HostgroupObject {

	hgObjects, _, err := z.HostgroupGet(HostgroupGetParams{
		GroupIDs: hgCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testHostgroupName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Hostgroup get error:", err)
	} else {
		if len(hgObjects) == 0 {
			t.Error("Hostgroup get error: unable to find created hostgroup")
		} else {
			t.Logf("Hostgroup get: success")
		}
	}

	return hgObjects
}
