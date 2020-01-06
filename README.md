# GoCal
 
A lightweight and extensible Google Calendar CLI

## Motivation

During the course of my last coop I often lamented to my coworkers both how disruptive to my workflow having to go to Google Calendar, and how long it takes to actually check. With that in mind, I decided to make the next command and used it to great effect during the coop term. I then started to extend the initial next command and this is what I ended up with. This does not have the full functionality of Google Calendar (yet) but I figured  that this was a good way both to continue learning Go, and a good way to get more comfortable with APIs. 

## Setup
- clone the repo locally
- navigate to the root
- run `go install`
- run `gocal next` and follow the instructions to authenticate with your own google calendar


## Commands
`next` - the og command, lists your next ten events
`list` - gives a list of all the available calendars

## Contributions

I would love to accept contributions as I build this out. Open any PRs and I'll review them posthaste. There are todos littered throughout the code. I also need to write tests for several of the commands, but I'm unsure how to proceed as of now, so will continue to do research into it.

Potential Feature list:
- edit events
- agenda: get events for a certain time period
