# simple-http-to-email-service
This is a very simple microservice written in Go. The service runs independently and receives data on /POST request and sends them via email. Your main Software simply needs to send a /POST request to send an email to a list of recipients mentioned in the /POST data

# Running
`go run main.go`
> Please refer to Go documentation for hosting it in production

# Configuration
Create  a file in the base directory (where main.go is located) named config.json

```
{
  "email": "",
  "password": "",
  "host": "smtp.gmail.com",
  "port": "587"
}
```
> This service has only been tested with google smtp. If you are using a google smtp, please note that the normal account password may not work diretly because of security.
> You  will need to create APP password from the google Account.
> To generate an APP password, please check https://support.google.com/accounts/answer/185833?hl=en

# Usage
/POST /send
```
{
    "to":["johndoe@email.com","janedoe@email.com"],
    "subject":"TEST SUBJECT",
    "body":"This is a test email"
}
```
> To send a fancy body, you can add HTML tags as well. This service simply sends the body as it is without any alterations.

# Expected Response
```
{
    "email": {
        "To": [
            "johndoe@email.com",
            "janedoe@email.com"
        ],
        "Subject": "TEST SUBJECT",
        "Body": "This is a test email"
    },
    "detail": "EMAIL SENT"
}
```
