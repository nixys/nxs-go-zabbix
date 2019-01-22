package zabbix

import (
	"reflect"
	"testing"
)

const (
	testHostinterfaceIP   = "10.1.1.2"
	testHostinterfacePort = "10151"
)

func TestHostinterfaceCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Preparing auxiliary data
	hgCreatedIDs := testHostgroupCreate(t, z)
	defer testHostgroupDelete(t, z, hgCreatedIDs)

	tCreatedIDs := testTemplateCreate(t, z, hgCreatedIDs)
	defer testTemplateDelete(t, z, tCreatedIDs)

	hCreatedIDs := testHostCreate(t, z, hgCreatedIDs, tCreatedIDs)
	defer testHostDelete(t, z, hCreatedIDs)

	// Create and delete
	hiCreatedIDs := testHostinterfaceCreate(t, z, hCreatedIDs[0])
	defer testHostinterfaceDelete(t, z, hiCreatedIDs)

	// Get
	testHostinterfaceGet(t, z, hCreatedIDs)
}

func testHostinterfaceCreate(t *testing.T, z Context, hCreatedID int) []int {

	hiCreatedIDs, _, err := z.HostinterfaceCreate([]HostinterfaceObject{
		{
			HostID: hCreatedID,
			IP:     testHostinterfaceIP,
			Main:   HostinterfaceMainNotDefault,
			Port:   testHostinterfacePort,
			Type:   HostinterfaceTypeAgent,
			UseIP:  HostinterfaceUseipIP,
		},
	})

	if err != nil {
		t.Fatal("Hostinterface create error:", err)
	}

	if len(hiCreatedIDs) == 0 {
		t.Fatal("Hostinterface create error: empty IDs array")
	}

	t.Logf("Hostinterface create: success")

	return hiCreatedIDs
}

func testHostinterfaceDelete(t *testing.T, z Context, hiCreatedIDs []int) []int {

	hiDeletedIDs, _, err := z.HostinterfaceDelete(hiCreatedIDs)
	if err != nil {
		t.Fatal("Hostinterface delete error:", err)
	}

	if len(hiDeletedIDs) == 0 {
		t.Fatal("Hostinterface delete error: empty IDs array")
	}

	if reflect.DeepEqual(hiDeletedIDs, hiCreatedIDs) == false {
		t.Fatal("Hostinterface delete error: IDs arrays for created and deleted hostinterface are mismatch")
	}

	t.Logf("Hostinterface delete: success")

	return hiDeletedIDs
}

func testHostinterfaceGet(t *testing.T, z Context, hCreatedIDs []int) []HostinterfaceObject {

	hiObjects, _, err := z.HostinterfaceGet(HostinterfaceGetParams{
		SelectHosts: SelectExtendedOutput,
		HostIDs:     hCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"ip": testHostinterfaceIP,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Hostinterface get error:", err)
	} else {
		if len(hiObjects) == 0 {
			t.Error("Hostinterface get error: unable to find created hostinterface")
		} else {
			t.Logf("Hostinterface get: success")
		}
	}

	return hiObjects
}
