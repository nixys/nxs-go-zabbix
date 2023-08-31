package zabbix

import (
	"fmt"
	"os"
	"testing"
)

func TestProblemGet(t *testing.T) {
	z := &Context{}
	err := z.Login(os.Getenv("ZABBIX_HOST"), os.Getenv("ZABBIX_USERNAME"), os.Getenv("ZABBIX_PASSWORD"))
	if err != nil {
		t.Fatalf("Login failed: %s", err)
	}

	params := ProblemGetParams{
		ObjectIDs: []string{"20143"},
		// ... Add other fields as needed
	}

	problems, _, err := z.ProblemGet(params)
	if err != nil {
		t.Fatalf("ProblemGet failed: %s", err)
	}

	if len(problems) == 0 {
		t.Fatalf("No problems found")
	}

	fmt.Println("Problems list:")
	for _, p := range problems {
		fmt.Println(p)
	}

	// ... Add more tests to validate the results
}
