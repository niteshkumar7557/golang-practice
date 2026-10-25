package main

import "fmt"

type EmailNotification struct {
	Email string
	Subject string
	Body string
}

type SMSNotification struct {
	Phone int
	Content string
}

func (e EmailNotification) Send() string{
	return "Email sent to " + e.Email
}
func (e SMSNotification) Send() string{
	return "SMS sent to " + fmt.Sprint(e.Phone)
}

type Notifier interface {
	Send() string
}

func notifyUser(n Notifier){
	fmt.Println(n.Send())
}

func main(){
	e := EmailNotification{Email: "nitesh@gmail.com", Subject: "Hello!", Body: "Body content Yayayaya!"}
	s := SMSNotification{Phone: 929292, Content: "You are dead bro!"}

	notifyUser(e)
	notifyUser(s)

}