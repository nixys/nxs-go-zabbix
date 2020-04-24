package zabbix

import (
	"reflect"
	"testing"
)

const (
	testHostmacroMacro = "{$TEST_USER_MACRO}"
	testHostmacroValue = "testUsermacroValue"
)

func TestHostmacroCRUD(t *testing.T) {

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
	hmCreatedIDs := testHostmacroCreate(t, z, hCreatedIDs[0])
	defer testHostmacroDelete(t, z, hmCreatedIDs)

	// Get
	testHostmacroGet(t, z, hmCreatedIDs)
}

func testHostmacroCreate(t *testing.T, z Context, hCreatedID int) []int {

	hmCreatedIDs, _, err := z.HostmacroCreate([]UsermacroObject{
		{
			HostID: hCreatedID,
			Macro:  testHostmacroMacro,
			Value:  testHostmacroValue,
		},
	})
	if err != nil {
		t.Fatal("Hostmacro create error:", err)
	}

	if len(hmCreatedIDs) == 0 {
		t.Fatal("Hostmacro create error: empty IDs array")
	}

	t.Logf("Hostmacro create: success")

	return hmCreatedIDs
}

func testHostmacroDelete(t *testing.T, z Context, hmCreatedIDs []int) []int {

	hmDeletedIDs, _, err := z.HostmacroDelete(hmCreatedIDs)
	if err != nil {
		t.Fatal("Hostmacro delete error:", err)
	}

	if len(hmDeletedIDs) == 0 {
		t.Fatal("Hostmacro delete error: empty IDs array")
	}

	if reflect.DeepEqual(hmDeletedIDs, hmCreatedIDs) == false {
		t.Fatal("Hostmacro delete error: IDs arrays for created and deleted hostmacro are mismatch")
	}

	t.Logf("Hostmacro delete: success")

	return hmDeletedIDs
}

func testHostmacroGet(t *testing.T, z Context, hmCreatedIDs []int) []UsermacroObject {

	hmObjects, _, err := z.UsermacroGet(UsermacroGetParams{
		HostmacroIDs:    hmCreatedIDs,
		SelectGroups:    SelectExtendedOutput,
		SelectHosts:     SelectExtendedOutput,
		SelectTemplates: SelectExtendedOutput,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"macro": testHostmacroMacro,
				"value": testHostmacroValue,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Hostmacro get error:", err)
	} else {
		if len(hmObjects) == 0 {
			t.Error("Hostmacro get error: unable to find created hostmacro")
		} else {
			t.Logf("Hostmacro get: success")
		}
	}

	return hmObjects
}
