package sendmailok

import (

	"fmt"
	"log"
	"strings"
	"net/smtp"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-sendmailok")

// MyActivity is a stub for your Activity implementation
type sendmailok struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &sendmailok{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *sendmailok) Metadata() *activity.Metadata {
	return a.metadata
}


// Eval implements activity.Activity.Eval
func (a *sendmailok) Eval(ctx activity.Context) (done bool, err error) {
	
	
	server := ctx.GetInput("Aserver").(string)
	port := ctx.GetInput("Bport").(string)
	sender := ctx.GetInput("Csender").(string)
	apppass := ctx.GetInput("Dapppassword").(string)
	ercpnt := ctx.GetInput("Ercpnt").(string)
	fsub := ctx.GetInput("Fsub").(string)
	gbody := ctx.GetInput("Gbody").(string)
	
	
	
	// Set up authentication information.
	//auth := smtp.PlainAuth(
	//	"",
	//	"sendalertsforq@gmail.com",
	//	"ptcxejoylzgtrfmh",
	//	"smtp.gmail.com",
	//)
	
	auth := smtp.PlainAuth("", sender, apppass, server,)
	
	t := []string{"To:", ercpnt}
	strings.Join(t, " ")
	
	s := []string{"Subject:", fsub}
	strings.Join(s, " ")
	
	serv := []string{server, port}
	strings.Join(serv, ":")
	
	
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	
	to := []string{ercpnt}
	msg := []byte(strings.Join(t, " ") + "\r\n" + strings.Join(s, " ") + "\r\n" + gbody + "\r\n")
	
	err = smtp.SendMail(strings.Join(serv, ":"), auth, sender, to, msg)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}
	
	fmt.Println("Mail Sent")
	log.Println("Mail Sent")


	// Set the output as part of the context
	activityLog.Debugf("Activity has sent the mail Successfully")
	fmt.Println("Activity has sent the mail Successfully")

	ctx.SetOutput("output", "Mail_Sent_Successfully")

	return true, nil
}
