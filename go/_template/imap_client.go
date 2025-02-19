package main

import (
	"log"
	"os"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func main() {
	userMail := os.Getenv("MAIL_USER")
	passMail := os.Getenv("MAIL_PASS")
	mailServer := os.Getenv("MAIL_SERVER")

	c, err := imapclient.DialTLS(mailServer+":993", nil)
	if err != nil {
		log.Fatalf("failed to dial IMAP server: %v", err)
	}
	defer c.Close()

	err = c.Login(userMail, passMail).Wait()
	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}

	mailboxes, err := c.List("", "%", nil).Collect()
	if err != nil {
		log.Fatalf("failed to list mailboxes: %v", err)
	}

	log.Printf("Found %v mailboxes", len(mailboxes))
	for _, mbox := range mailboxes {
		log.Printf(" - %v", mbox.Mailbox)
	}

	selectedMbox, err := c.Select("INBOX", nil).Wait()
	if err != nil {
		log.Fatalf("failed to select INBOX: %v", err)
	}
	log.Printf("INBOX contains %v messages", selectedMbox.NumMessages)

	listCmd := c.List("", "%", &imap.ListOptions{
		ReturnStatus: []imap.StatusItem{
			imap.StatusItemNumMessages,
			imap.StatusItemNumUnseen,
		},
	})
	for {
		mbox := listCmd.Next()
		if mbox == nil {
			break
		}
		log.Printf("Mailbox %q contains %v messages (%v unseen)",
			mbox.Mailbox, *mbox.Status.NumMessages, *mbox.Status.NumUnseen)
	}
	if err := listCmd.Close(); err != nil {
		log.Fatalf("LIST command failed: %v", err)
	}

	statusItems := []imap.StatusItem{imap.StatusItemNumMessages}
	if data, err := c.Status("INBOX", statusItems).Wait(); err != nil {
		log.Fatalf("STATUS command failed: %v", err)
	} else {
		log.Printf("INBOX contains %v messages", *data.NumMessages)
	}

}
