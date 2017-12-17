# service-agenda

[![Build Status](https://travis-ci.org/James-Yip/service-agenda.svg?branch=master)](https://travis-ci.org/James-Yip/service-agenda)

service-agenda is a simple client-server application based on the concept of microservices, where the client and the server are run in two different containers respectively.

This application supports various operations for agenda management including user register, meeting creation & query, etc.

## API Design

See [https://agenda29.docs.apiary.io](https://agenda29.docs.apiary.io)

## Usage

```
Agenda is a CLI meeting manager program based on cobra.

This application is a tool to support various operations on meetings
including user register, meeting creation & query, etc.

Usage:
  agenda [command]

Available Commands:
  createMeeting Create a meeting
  help          Help about any command
  listMeetings  List all meetings
  listUsers     List all registered users
  login         User login
  logout        User logout
  register      Register user

Flags:
  -h, --help   help for agenda

Use "agenda [command] --help" for more information about a command.

```

## Usage Examples


## Mock Test

Test client using mock server.

```
$ go test -v ./cli/service/
=== RUN   TestRegister
Register success as:
{
    "username":"james",
    "email":"james@qq.com",
    "phone":"12345678912"
}
--- PASS: TestRegister (2.23s)
=== RUN   TestLogin
Logout: success.
Login: success.
--- PASS: TestLogin (0.69s)
=== RUN   TestLogout
Logout: success.
--- PASS: TestLogout (0.29s)
=== RUN   TestListUsers
send get users
after send get users
[
    {
        "username":"zhang3"
        "email":"zhang3@mail2.sysu.edu.cn"
        "phone":"12345678912"
    },{
        "username":"li4"
        "email":"li4@mail2.sysu.edu.cn"
        "phone":"14785236912"
    }
]--- PASS: TestListUsers (0.29s)
=== RUN   TestCreateMeeting
Creat meeting success as:
{
    "title":"Meeting1",
    "sponsor":"zhang3",
    "participators":["li4"],
    "startTime":"2017-11-11/11:11",
    "endTime":"2017-11-11/11:22"
}
--- PASS: TestCreateMeeting (0.31s)
=== RUN   TestListMeetings
[
    {### Change Participators []
        "title":"Meeting1",
        "sponsor":"zhang3",
        "participators":["li4"],
        "startTime":"2017-11-11/11:11",
        "endTime":"2017-11-11/11:22"
    },{
        "title":"Meeting2",
        "sponsor":"li4",
        "participators":["zhang3"],
        "startTime":"2017-12-12/12:12",
        "endTime":"2017-12-12/12:22"
    }
]--- PASS: TestListMeetings (0.32s)
PASS
ok  	github.com/James-Yip/service-agenda/cli/service	4.122s
```
