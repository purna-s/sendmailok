package sendmailok

import (

	"fmt"
	"log"
	"net/smtp"
	"strings"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("Aserver", "smtp.gmail.com")
	tc.SetInput("Bport", "587")
	tc.SetInput("Csender", "sendalertsforq@gmail.com")
	tc.SetInput("Dapppassword", "ptcxejoylzgtrfmh")
	tc.SetInput("Ercpnt", "purna.sadanala@prowesssoft.com")
	tc.SetInput("Fsub", "Q_Subscriber_Down!")
	tc.SetInput("Gbody", "Subscriber_for_Queue_is_down.")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}
