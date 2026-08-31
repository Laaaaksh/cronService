# cronService

A small Go backend for scheduling and running HTTP-triggered cron jobs, built as a group practice project in July 2022 (commits from multiple contributors between 2022-07 and 2022-07-22, merged via pull requests).

## What it is

A REST API (Gin + GORM/MySQL) that lets users register organizations, log in, and create "cron jobs" — each one a URL plus an HTTP method (GET/POST) and a schedule expression. A background goroutine in `main.go` polls the database for jobs whose `next_time` has passed, fires the HTTP request, logs the result (`CronExecutionResult`), and reschedules the job (with basic retry-then-backoff on failure). There's also a simple user/organization/permission model alongside the cron functionality.

This reads as a multi-person learning project (routes/controllers split across contributors, commits like "added controller and routes" merged via separate branches) rather than a maintained service.

## Stack

- Go 1.16
- [Gin](https://github.com/gin-gonic/gin) for HTTP routing
- [GORM](https://github.com/jinzhu/gorm) (v1) with a MySQL driver, auto-migrated on startup
- JWT (`dgrijalva/jwt-go`) is a dependency but token validation isn't wired into the routes shown here

## Endpoints

```
/cronjobs/create-cronjob/        POST
/cronjobs/change-status/:id      POST
/cronjobs/edit/:id               POST
/cronjobs/delete/:id             DELETE
/cronjobs/get-all-cronjobs       GET
/cronjobs/get-cronjob/:id        GET
/cronjobs/get-cronlogs/:id       GET

/user/create-user/               POST
/user/login/                     POST
/user/delete/:id                 DELETE
/user/update-user/                POST

/admin/create-org/               POST
/admin/create-permission/        POST
/admin/get-permissions/          GET
```

## Running it

`go build ./...` succeeds (verified). Running it needs a reachable MySQL instance configured in `Database/Config.go` — that part was not verified since it requires a live database.

## Status

Written by a small group in July 2022 as a practice project (branches per contributor, merged via PRs). Not maintained. The cron-scheduling logic (`extractTime` in `main.go`) parses a fixed-width custom expression format rather than standard cron syntax, and job execution runs in a tight polling loop rather than a proper scheduler.
