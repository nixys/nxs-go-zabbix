package zabbix

import (
	"reflect"
	"testing"
)

const (
	testMediatypeName        = "testMediatypeName"
	testMediatypeDescription = "testMediatypeDescription"
	testMediatypeExecPath    = "test_script.sh"
)

func TestMediatypeCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Create and delete
	mtCreatedIDs := testMediatypeCreate(t, z)
	defer testMediatypeDelete(t, z, mtCreatedIDs)

	// Get
	testMediatypeGet(t, z, mtCreatedIDs)
}

func testMediatypeCreate(t *testing.T, z Context) []int {

	hiCreatedIDs, _, err := z.MediatypeCreate([]MediatypeObject{
		{
			Name:        testMediatypeName,
			Description: testMediatypeDescription,
			Type:        MediatypeScript,
			ExecPath:    testMediatypeExecPath,
			MessageTemplates: []MediatypeMessageTemplateObject{
				{
					EventSource: MediatypeMessageTemplateEventSourceTriggers,
					Recovery:    MediatypeMessageTemplateRecoveryOperations,
					Subject:     "SSS",
					Message:     "MMM",
				},
			},
		},
	})

	if err != nil {
		t.Fatal("Mediatype create error:", err)
	}

	if len(hiCreatedIDs) == 0 {
		t.Fatal("Mediatype create error: empty IDs array")
	}

	t.Logf("Mediatype create: success")

	return hiCreatedIDs
}

func testMediatypeDelete(t *testing.T, z Context, mtCreatedIDs []int) []int {

	mtDeletedIDs, _, err := z.MediatypeDelete(mtCreatedIDs)
	if err != nil {
		t.Fatal("Mediatype delete error:", err)
	}

	if len(mtDeletedIDs) == 0 {
		t.Fatal("Mediatype delete error: empty IDs array")
	}

	if reflect.DeepEqual(mtDeletedIDs, mtCreatedIDs) == false {
		t.Fatal("Mediatype delete error: IDs arrays for created and deleted mediatype are mismatch")
	}

	t.Logf("Mediatype delete: success")

	return mtDeletedIDs
}

func testMediatypeGet(t *testing.T, z Context, mtCreatedIDs []int) []MediatypeObject {

	mtObjects, _, err := z.MediatypeGet(MediatypeGetParams{
		SelectUsers:  SelectExtendedOutput,
		MediatypeIDs: mtCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name":      testMediatypeName,
				"exec_path": testMediatypeExecPath,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Mediatype get error:", err)
	} else {
		if len(mtObjects) == 0 {
			t.Error("Mediatype get error: unable to find created mediatype")
		} else {
			t.Logf("Mediatype get: success")
		}
	}

	return mtObjects
}
