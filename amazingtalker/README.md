# amazingtalker

[![Build Status](https://travis-ci.com/blackhorseya/amazingtalker.svg?token=ksYeJ43XxxMyjL9VYtsT&branch=main)](https://travis-ci.com/blackhorseya/amazingtalker)
[![codecov](https://codecov.io/gh/blackhorseya/amazingtalker/branch/main/graph/badge.svg?token=pQ91ZnHZtI)](https://codecov.io/gh/blackhorseya/amazingtalker)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/lobster)](https://github.com/blackhorseya/lobster/blob/main/LICENSE)

## Cover Letter

### 為什麼想加入AmazingTalker？

因為以下幾點想要加入 AmazingTalker 一起合作。

#### 公司文化

- 使用 `OKR` 訂定清楚的目標
- 可以在適度地評估風險後去挑戰事情，嘗試錯誤，retro，持續改進，找到最適合當下的 solution
- 跑得快，每天都在進步的感覺
- 可以平等的 challenge 人事物

### 您的職涯目標和人生目標是什麼？為什麼定這些目標？

#### 職涯目標

- 短期(3～6個月)：找到對工程師相對友善和尊重的環境持續成長、學習並將學習到的知識帶入公司，如 `Backend`、`SRE`。
- 中期(1～3年)：會朝向 `Staff Engineer` 或 `Associate Software Architect` 前進，並且持續強化 Soft Skills，如語言、溝通等。
- 長期(3～5年)：最終目標想成為一名能讓同事們信服也會想一起努力的 `System Architect`。

#### 人生目標

1. 我的人生目標基本上與職涯目標一致，想多接觸各式各樣的人，吸收非專業相關的知識，並且利用我的專業幫助他們成長或給出相對懶惰的 solution。
2. 以 open source 的形式做自己的 side project，希望有人因為我的設計理念改善一點他們的成長或思維。

#### 為什麼定這些目標？

1. 職涯目標：因為我小的時候表哥讓我試玩他寫的程式，我看到他的那份成就感，讓我感受到不一樣的快樂，而後我的職涯目標就變成成為一名工程師。 隨著進入相關領域後並了解了更多，加上我小時候喜歡玩**樂高**
   ，拼裝設計等的玩法，所以我把這個想法投射在 `System Architect`。

2. 人生目標：想讓我的人生對世界的一些人有那麼一小點的貢獻。

## Goals

### Description

請實作兩隻API，分別是教師列表以及教師個人資訊API並**透過快取回傳資料**

快取（cache）機制，規則如下:

1. 用戶端發送請求 -> 有快取 -> 快取建立的時間超過 10 分鐘 -> 觸發「刷新快取的工作（job）」到佇列由背景執行 -> 先回傳目前快取的資料給用戶端
2. 用戶端發送請求 -> 沒有快取 -> 從資料庫取得資料 -> 寫入到快取 -> 回傳資料給用戶端
3. 在沒有快取的情況下，避免請求併發時產生競爭情況（race condition）。例如：快取過期時同時有10個請求，但只能有一個請求觸發快取運算的內容，其餘9個請求需等待第1個請求運算完畢後，直接拿快取內的資料

### Required

- 實作上述的快取機制
- API 1 教師列表 路徑： **/api/tutors/{language.slug}**
- API 2 教師個人資訊 路徑： **/api/tutor/{tutor.slug}**
- 請使用以下任一語言撰寫：Ruby、Node.js、PHP、Python、Java、C#、Go、Rust
- 回應格式皆使用JSON
- 單元測試
