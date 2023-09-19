package zabbix

import (
	"testing"
)

const (
	testHistoryItemID = 45503
	testHistoryType   = 0
)

func TestHistoryCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Get
	testHistoryGet(t, z)
}

func testHistoryGet(t *testing.T, z Context) []HistoryFloatObject {

	r := []HistoryFloatObject{}

	hObjects, _, err := z.HistoryGet(HistoryGetParams{
		History: HistoryObjectTypeFloat,
		//ItemIDs: []int{testHistoryItemID},
		GetParameters: GetParameters{
			Limit: 1,
		},
	})

	if err != nil {
		t.Error("History get error:", err)
	} else {
		r = *hObjects.(*[]HistoryFloatObject)
		if len(r) == 0 {
			t.Error("History get error: unable to find history")
		} else {
			t.Logf("History get: success")
		}
	}

	return r
}
