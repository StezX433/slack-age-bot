# Slack Age Bot

A Go-based Slack bot that calculates age and provides
additional birthday and lifetime information.

## Features

- Calculate age from year of birth
- Calculate exact age from date of birth
- Calculate number of days lived
- Find the next birthday
- Find the next age milestone
- Input validation
- Unit tests

## Project Structure

```text
slack-age-bot/
│
├── cmd/
│   └── bot/
│       └── main.go
│
├── internal/
│   ├── age/
│   │   ├── calculator.go
│   │   ├── birthday.go
│   │   ├── milestone.go
│   │   └── calculator_test.go
│   │
│   ├── slack/
│   │   ├── client.go
│   │   └── handler.go
│   │
│   └── config/
│       └── config.go
│
├── .env.example
├── .gitignore
├── go.mod
└── README.md