# Slack Bot Project

## Interface

* Sending a Message

``` bash
slackbot post \
    --token "${SLACK_AUTH_TOKEN}" \
    --channel "channel name" \ 
    message \
    --title "Message Title" \
    --text "Message text here"
```

* Sending a File as attachment

``` bash
slackbot post \
    --token "${SLACK_AUTH_TOKEN}" \
    --channel "${CHANNEL_ID}" \
    file \
    --file /Users/yyovkov/Downloads/report-users-20230526.xlsx \
    --initialComment "Test me"
```

## Setup Environment

* Init project

``` bash
go mod init githu.com/yyovkov/slackbot
```

* Setup Cobra

``` bash
go get -u github.com/spf13/cobra@v1.7.0
go install github.com/spf13/cobra-cli@v1.3.0
```

* Init cobra

``` bash
cobra-cli init
```

* Add subcommand *send*

``` bash
cobra-cli add send
```

* Add subcommand *message* of a command *send*

``` bash
cobra-cli add message -p send
```

## Building the project

``` bash
go build
```

## Run Against S3

``` bash
export AWS_PROFILE=sandbox
```
