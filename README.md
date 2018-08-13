# Holidaze - A tool for syncing holidays from orangehrm to google calendar

## What's it for?
This project is in it's very early stages but the goal of the project is to produce a tool which can sync holidays from orangehrm to google calendar.

## What it does
Right now running the application will
- output upcoming events from your primary calendar
- create a google calendar event in your primary calendar and invite test@test.com

## How to contribute
Check out the [trello board](https://trello.com/b/1sTpCBdI/holidaze) for tasks/next steps. Feel free to pick up, update or add tasks there.

## Running Locally

### Clone

- Fork the repository.
- Clone it using `git clone git@github.com:<your-username>/holidaze.git`

### Run

- Create a github issue on this repo requesting to be added to the google project
- Follow the steps to download credentials as described [here](https://developers.google.com/calendar/quickstart/go)
- Download dependencies using `make prepare`
- Run using `make run`