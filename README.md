# service-agenda

[![Build Status](https://travis-ci.org/James-Yip/service-agenda.svg?branch=master)](https://travis-ci.org/James-Yip/service-agenda)

service-agenda is a simple client-server application based on the concept of microservices, where the client and the server are run in two different containers respectively.

This application supports various operations for agenda management including user register, meeting creation etc.

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

1. pull the service-agenda image

```
$ sudo docker pull yezh7/service-agenda
Using default tag: latest
latest: Pulling from yezh7/service-agenda
f49cf87b52c1: Pull complete
7b491c575b06: Pull complete
b313b08bab3b: Pull complete
215a2061b8a4: Pull complete
04fa9dcc5f7d: Pull complete
044102b3b4a1: Pull complete
37e616f9e3fe: Pull complete
792a3ef002db: Pull complete
3546448acc4c: Pull complete
01387abcb16f: Pull complete
Digest: sha256:ec65115dbdfcf18714c0b01fab4fbc95661c358c6c92dd0971fefa72eea884af
Status: Downloaded newer image for yezh7/service-agenda:latest
```

2. run the server

```
$ sudo docker run --rm -p 8080:8080 yezh7/service-agenda service
[negroni] listening on :8080
```

3. run the client

```
$ sudo docker run -it --rm --name agenda-cli --net host yezh7/service-agenda
root@james-X555LPB:/#
```


## Test

### Service Test

register
```
/# cli register -u james -p 123456 -e james@qq.com -t 13623887454
Register success as:
"james"
```

login
```
/# cli login -u james -p 123456
Login: success.
```

logout
```
/# cli logout
Logout: success.
```

listUsers
```
/# cli listUsers
[
  {
    "UserName": "james",
    "Password": "123456",
    "Email": "james@qq.com",
    "Phone": "13623887454"
  },
  {
    "UserName": "alice",
    "Password": "123456",
    "Email": "alice@qq.com",
    "Phone": "13623887455"
  }
]

```

createMeeting
```
/# cli createMeeting -t singleDog -p "alice" -s 2017-11-11/11:11 -e 2017-11-11/11:22
Creat meeting success as:
{
  "Title": "singleDog",
  "Sponsor": "james",
  "StartTime": "2017-11-11/11:11",
  "EndTime": "2017-11-11/11:22",
  "Participators": [
    "alice"
  ]
}
```

listMeeting
```
/# cli listMeetings -s 2017-11-11/11:11 -e 2017-11-11/11:22
[
  {
    "Title": "singleDog",
    "Sponsor": "james",
    "StartTime": "2017-11-11/11:11",
    "EndTime": "2017-11-11/11:22",
    "Participators": [
      "alice"
    ]
  }
]
```


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
