# Glasnostic

## Goals

1. Read in the following CSV file and convert it to a JSON file that contains an array of JSON objects - one CSV line
   corresponds to one JSON object. (Hint: you need to convert to data objects (e.g. Employee) first and then serialize
   to JSON format).

## How to use

```bash
go run ./main.go convert -f ${SOURCE_FILE}
```

## Sample 

```bash
$ go run ./main.go convert -f ./test/data/raw.csv
[{"id":"1","first_name":"Marc","last_name":"Smith","email":"marc@glasnostic.com","description":"Writer of Java","role":"Dev","phone":"541-754-3010"},{"id":"2","first_name":"John","last_name":"Young","email":"john@glasnostic.com","description":"Interested in MHW","role":"HR","phone":"541-75-3010"},{"id":"3","first_name":"Peter","last_name":"Scott","email":"peter@glasnostic.com","description":"amateur boxer","role":"Dev","phone":"541-754-3010"}]
```
