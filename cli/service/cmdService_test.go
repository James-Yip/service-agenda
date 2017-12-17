package service

import (
	"testing"
)

func TestRegister(t *testing.T) {
	// mock server test
	err := Register("james", "123456", "james@qq.com", "12345678912")
	if err != nil {
		t.Error(err)
	}
	Logout()
	err = Register("alice", "123456", "alice@qq.com", "12345678913")
	if err != nil {
		t.Error(err)
	}
}

func TestLogout(t *testing.T) {
	// mock server test
	err := Logout()
	if err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {
	// mock server test
	err := Login("james", "123456")
	if err != nil {
		t.Error()
	}
}

func TestListUsers(t *testing.T) {
	// mock server test
	err := ListUsers()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMeeting(t *testing.T) {
	// mock server test
	err := CreateMeeting("Meeting1", "alice", "2017-11-11/11:11", "2017-11-11/11:22")
	if err != nil {
		t.Error(err)
	}
}

func TestListMeetings(t *testing.T) {
	// mock server test
	err := ListMeetings("2017-11-11/11:00", "2017-11-11/11:30")
	if err != nil {
		t.Error(err)
	}
}
