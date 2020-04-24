package zabbix

import (
	"reflect"
	"testing"
)

const (
	testHostName   = "testHost"
	testHostIP     = "10.1.1.1"
	testHostPort   = "10150"
	testMacro      = "{$TEST_MACRO}"
	testMacroValue = "testMacroValue"
)

func TestHostCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Preparing auxiliary data
	hgCreatedIDs := testHostgroupCreate(t, z)
	defer testHostgroupDelete(t, z, hgCreatedIDs)

	tCreatedIDs := testTemplateCreate(t, z, hgCreatedIDs)
	defer testTemplateDelete(t, z, tCreatedIDs)

	// Create and delete
	hCreatedIDs := testHostCreate(t, z, hgCreatedIDs, tCreatedIDs)
	defer testHostDelete(t, z, hCreatedIDs)

	// Update
	testHostUpdate(t, z, hCreatedIDs)

	// Get
	testHostGet(t, z, hCreatedIDs, tCreatedIDs, hgCreatedIDs)
}

func testHostCreate(t *testing.T, z Context, hgCreatedIDs, tCreatedIDs []int) []int {

	var groups []HostgroupObject
	var templates []TemplateObject

	// Add groups to host
	for _, e := range hgCreatedIDs {
		groups = append(groups, HostgroupObject{
			GroupID: e,
		})
	}

	// Add templates to host
	for _, e := range tCreatedIDs {
		templates = append(templates, TemplateObject{
			TemplateID: e,
		})
	}

	hCreatedIDs, _, err := z.HostCreate([]HostObject{
		{
			Host:      testHostName,
			Groups:    groups,
			Templates: templates,
			Interfaces: []HostinterfaceObject{
				{
					IP:    testHostIP,
					Main:  HostinterfaceMainDefault,
					Port:  testHostPort,
					Type:  HostinterfaceTypeAgent,
					UseIP: HostinterfaceUseipIP,
				},
			},
			Macros: []UsermacroObject{
				{
					Macro: testMacro,
					Value: testMacroValue,
				},
			},
		},
	})

	if err != nil {
		t.Fatal("Host create error:", err)
	}

	if len(hCreatedIDs) == 0 {
		t.Fatal("Host create error: empty IDs array")
	}

	t.Logf("Host create: success")

	return hCreatedIDs
}

func testHostUpdate(t *testing.T, z Context, hCreatedIDs []int) []int {

	var hObjects []HostObject

	// Preparing host objects array to update
	for _, i := range hCreatedIDs {
		hObjects = append(hObjects, HostObject{
			HostID: i,
			Name:   testHostName + "_upd",
		})
	}

	hUpdatedIDs, _, err := z.HostUpdate(hObjects)
	if err != nil {
		t.Fatal("Host update error:", err)
	}

	if len(hUpdatedIDs) == 0 {
		t.Fatal("Host update error: empty IDs array")
	}

	if reflect.DeepEqual(hUpdatedIDs, hCreatedIDs) == false {
		t.Fatal("Host update error: IDs arrays for created and updated hosts are mismatch")
	}

	t.Logf("Host update: success")

	return hUpdatedIDs
}

func testHostDelete(t *testing.T, z Context, hCreatedIDs []int) []int {

	hDeletedIDs, _, err := z.HostDelete(hCreatedIDs)
	if err != nil {
		t.Fatal("Host delete error:", err)
	}

	if len(hDeletedIDs) == 0 {
		t.Fatal("Host delete error: empty IDs array")
	}

	if reflect.DeepEqual(hDeletedIDs, hCreatedIDs) == false {
		t.Fatal("Host delete error: IDs arrays for created and deleted host are mismatch")
	}

	t.Logf("Host delete: success")

	return hDeletedIDs
}

func testHostGet(t *testing.T, z Context, hCreatedIDs, tCreatedIDs, hgCreatedIDs []int) []HostObject {

	hObjects, _, err := z.HostGet(HostGetParams{
		SelectParentTemplates: SelectExtendedOutput,
		SelectMacros:          SelectExtendedOutput,
		HostIDs:               hCreatedIDs,
		TemplateIDs:           tCreatedIDs,
		GroupIDs:              hgCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testHostName + "_upd",
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Host get error:", err)
	} else {
		if len(hObjects) == 0 {
			t.Error("Host get error: unable to find created host")
		} else {

			// Check macro in each created host
			for _, h := range hObjects {

				foundMacro := false

				for _, m := range h.Macros {
					if m.Macro == testMacro && m.Value == testMacroValue {
						foundMacro = true
						break
					}
				}

				if foundMacro == false {
					t.Error("Host get error: unable to find macro in created host")
				}
			}

			t.Logf("Host get: success")
		}
	}

	return hObjects
}
