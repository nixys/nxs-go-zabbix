package zabbix

import (
	"reflect"
	"strconv"
	"testing"
)

const (
	testActionName         = "testAction"
	testActionEscPeriod    = 300
	testActionDefShortdata = "{HOST.NAME1} [{TRIGGER.STATUS}]: {TRIGGER.NAME}"
	testActionDefLongdata  = "Trigger: {TRIGGER.NAME}\r\nTrigger status: {TRIGGER.STATUS}\r\nTrigger severity: {TRIGGER.SEVERITY}\r\nTrigger URL: {TRIGGER.URL}\r\nEvent type: E1\r\n\r\nItem values:\r\n\r\n1. {ITEM.NAME1} ({HOST.NAME1}:{ITEM.KEY1}): {ITEM.VALUE1}\r\n2. {ITEM.NAME2} ({HOST.NAME2}:{ITEM.KEY2}): {ITEM.VALUE2}\r\n3. {ITEM.NAME3} ({HOST.NAME3}:{ITEM.KEY3}): {ITEM.VALUE3}\r\n\r\nOriginal event ID: {EVENT.ID}"
	testMediaTypeID        = 1
)

func TestActionCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Preparing auxiliary data
	hgCreatedIDs := testHostgroupCreate(t, z)
	defer testHostgroupDelete(t, z, hgCreatedIDs)

	// Create and delete
	aCreatedIDs := testActionCreate(t, z, hgCreatedIDs[0], 7)
	defer testActionDelete(t, z, aCreatedIDs)

	// Get
	testActionGet(t, z, aCreatedIDs)
}

func testActionCreate(t *testing.T, z Context, hostgrpID, usergrpID int) []int {

	aCreatedIDs, _, err := z.ActionCreate([]ActionObject{
		{
			Name:        testActionName,
			Eventsource: 0,
			Status:      ActionStatusEnabled,
			EscPeriod:   testActionEscPeriod,
			Filter: ActionFilterObject{
				EvalType: ActionFilterEvalTypeAndOr,
				Conditions: []ActionFilterConditionObject{
					{
						ConditionType: ActionFilterConditionTypeProblemIsSuppressed,
						Operator:      ActionFilterConditionOperatorNo,
					},
					{
						ConditionType: ActionFilterConditionTypeHostroup,
						Operator:      ActionFilterConditionOperatorEQ,
						Value:         strconv.Itoa(hostgrpID),
					},
				},
			},
			Operations: []ActionOperationObject{
				{
					OperationType: ActionOperationTypeSendMsg,
					EscPeriod:     0,
					EscStepFrom:   2,
					EscStepTo:     2,
					EvalType:      ActionOperationEvalTypeAndOR,
					Opconditions: []ActionOperationConditionObject{
						{
							ConditionType: ActionOperationConditionTypeEventAcknowledged,
							Operator:      ActionOperationConditionOperatorEq,
							Value:         "0",
						},
					},
					Opmessage: ActionOperationMessageObject{
						DefaultMsg:  ActionOperationMessageDefaultMsgFromOperation,
						MediatypeID: testMediaTypeID,
					},
					OpmessageGrp: []ActionOpmessageGrpObject{
						{
							UsrgrpID: usergrpID,
						},
					},
				},
			},
		},
	})

	if err != nil {
		t.Fatal("Action create error:", err)
	}

	if len(aCreatedIDs) == 0 {
		t.Fatal("Action create error: empty IDs array")
	}

	t.Logf("Action create: success")

	return aCreatedIDs
}

func testActionDelete(t *testing.T, z Context, aCreatedIDs []int) []int {

	aDeletedIDs, _, err := z.ActionDelete(aCreatedIDs)
	if err != nil {
		t.Fatal("Action delete error:", err)
	}

	if len(aDeletedIDs) == 0 {
		t.Fatal("Action delete error: empty IDs array")
	}

	if reflect.DeepEqual(aDeletedIDs, aCreatedIDs) == false {
		t.Fatal("Action delete error: IDs arrays for created and deleted action are mismatch")
	}

	t.Logf("Action delete: success")

	return aDeletedIDs
}

func testActionGet(t *testing.T, z Context, aCreatedIDs []int) []ActionObject {

	aObjects, _, err := z.ActionGet(ActionGetParams{
		SelectOperations: SelectExtendedOutput,
		SelectFilter:     SelectExtendedOutput,
		ActionIDs:        aCreatedIDs,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testActionName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Action get error:", err)
	} else {
		if len(aObjects) == 0 {
			t.Error("Action get error: unable to find created action")
		} else {
			t.Logf("Action get: success")
		}
	}

	return aObjects
}
