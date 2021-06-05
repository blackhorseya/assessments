# IP Rate Limit

[![Build Status](https://travis-ci.com/blackhorseya/ip-rate-limit.svg?branch=main)](https://travis-ci.com/blackhorseya/ip-rate-limit)
[![codecov](https://codecov.io/gh/blackhorseya/ip-rate-limit/branch/main/graph/badge.svg?token=0WF53W6ZWH)](https://codecov.io/gh/blackhorseya/ip-rate-limit)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/ip-rate-limit)](https://goreportcard.com/report/github.com/blackhorseya/ip-rate-limit)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/ip-rate-limit)](https://github.com/blackhorseya/ip-rate-limit/blob/main/LICENSE)

## Goals

請實作一個基本的web server 並滿足以下要求:

- 每個 `IP` 每分鐘僅能接受 60 個 requests。
- 在首頁顯示目前的 request 量,超過限制的話則顯示 `Error`, 例如在一分鐘內第 30 個 request 則顯示 30,第 61 個request 則顯示 `Error`。
- 可以使用任意資料庫,也可以自行設計 in-memory 資料結構,並在文件中說明理由。
- 請附上測試。

你不需要實作:

- 資料持久化。
- 設計網頁。

## Proposal

- 姓名: Sean Cheng/鄭晉丞
- 開發耗用時間(最小單位為分鐘):
- 開發採用程式語言: `golang`
- 請提供 repository 連結或 zip 檔: [Source Code](https://github.com/blackhorseya/ip-rate-limit)
- 請在文件內說明如何啟動,或者也可以 deploy 到 hosting platform(例如 Heroku):
    - [IP Rate Limit](https://ip-rate-limit.seancheng.space/)
    - [Swagger](https://ip-rate-limit.seancheng.space/api/docs/index.html)
- 其他說明:

## How to use

Clone source code

```bash
git clone git@github.com:blackhorseya/ip-rate-limit.git
```

### Run with local

```bash
make run-with-local
```

### Run with Docker

```bash
make run-with-docker
```

## Reports

```bash
vegeta attack -duration=10s -rate=5 -targets=vegeta.conf | vegeta report

Requests      [total, rate, throughput]         50, 5.10, 5.07
Duration      [total, attack, wait]             9.866s, 9.803s, 63.057ms
Latencies     [min, mean, 50, 90, 95, 99, max]  54.971ms, 98.164ms, 67.556ms, 206.54ms, 267.183ms, 349.651ms, 349.651ms
Bytes In      [total, mean]                     111, 2.22
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:50  
Error Set:
```

## Tech

### Development WorkFlow

using [Trunk-based development](https://blog.seancheng.space/posts/what-is-trunk-based-development)
Reference [official documents](https://cloud.google.com/solutions/devops/devops-tech-trunk-based-development)

### Dependencies

- [gin](https://github.com/gin-gonic/gin)
- [wire](https://github.com/google/wire)
- [logrus](https://github.com/sirupsen/logrus)
- [swaggo](https://github.com/swaggo/swag)
- [testify](https://github.com/stretchr/testify)
- [mockery](https://github.com/vektra/mockery)

### CI/CD

- [Travis-CI](https://travis-ci.com/blackhorseya/ip-rate-limit)

### Infrastructure

- [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine?hl=zh-tw)
- [Helm 3](https://helm.sh/)
- [Cloudflare](https://www.cloudflare.com/zh-tw/)

### Testing

- [vegeta](https://github.com/tsenart/vegeta)
