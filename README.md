[![Stories in Ready](https://badge.waffle.io/mattstratton/blondie.svg?label=ready&title=Ready)](http://waffle.io/mattstratton/blondie) [![Stories in Progress](https://badge.waffle.io/mattstratton/blondie.svg?label=in%progress&title=In%20Progress)](http://waffle.io/mattstratton/blondie) [![Needs Review](https://badge.waffle.io/mattstratton/blondie.svg?label=needs-review&title=Needs%20Review)](http://waffle.io/mattstratton/blondie)
[![Build Status](https://travis-ci.org/mattstratton/blondie.svg?branch=master)](https://travis-ci.org/mattstratton/blondie)
[![Coverage Status](https://coveralls.io/repos/github/mattstratton/blondie/badge.svg?branch=master)](https://coveralls.io/github/mattstratton/blondie?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattstratton/blondie)](https://goreportcard.com/report/github.com/mattstratton/blondie)
[![license](https://img.shields.io/github/license/mattstratton/blondie.svg)]()

You can see progress on tasks at http://waffle.io/mattstratton/blondie

[![Throughput Graph](https://graphs.waffle.io/mattstratton/blondie/throughput.svg)](https://waffle.io/mattstratton/blondie/metrics)
# blondie

`blondie` is a web application for submitting and reviewing talks, built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).

Requirements
===========

* Docker 1.12
* Docker Compose 1.8

```

Starting services
==============================

```
docker-compose up -d
```

Stopping services
==============================

```
docker-compose stop
```

Including new changes
==============================

If you need change some source code you can deploy it typing:

```
docker-compose build
```

## Documentation
### Talk Service
This service is used to get information about a talk. It provides the talk title, the abstract, and other details.

*Routes:*

* GET - /talks : Get all talks
* POST - /talks : Create talk
* GET - /talks/{id} : Get talk by id
* DELETE - /talks/{id} : Remove talk by id

### Account Service
This service is used to manage user accounts, specifically around login, etc.

*Routes*

* GET - /accounts : Get all accounts
* POST - /accounts: Create  account
* GET - /accounts/{id} : Get account by id
* DELETE - /accounts/{id} : Remove account by id

### Speaker Service
This service manages speaker profiles.

*Routes*

* GET - /speakers : Get all speakers
* POST - /speakers : Create speaker
* GET - /speakers/{id} : Get speaker by id
* DELETE - /speakers/{id} : Remove speaker by id

### Event Service
This service manages events, including who has access to them, etc

*Routes*

* GET - /events : Get all events
* POST - /events : Create new event
* GET - /events/{i} : Get event by id
* DELETE - /events/{i} : Remove event by id
