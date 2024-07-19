package mail

import "testing"

func TestMail(t *testing.T) {

	mock := &SMTPMock{}
	SendEmail(mock, "Test Email from Go", "chosh901@naver.com", "Hello,\n\nThis is a test email sent from a Go program.")
}

type SMTPMock struct {
	server   string
	port     string
	user     string
	password string
}

func (s SMTPMock) Server() string {
	return s.server
}

func (s SMTPMock) Port() string {
	return s.port
}

func (s SMTPMock) User() string {
	return s.user
}

func (s SMTPMock) Password() string {
	return s.password
}
