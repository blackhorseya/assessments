# Glasnostic

## Goals

1. Read in the following CSV file and convert it to a JSON file that contains an array of JSON objects - one CSV line
   corresponds to one JSON object (hint convert first to data objects and then serialize to JSON format).
2. John still loves to use Excel. From time to time we need to convert the JSON exported from the ERP back to CSV.
   Extend the program from step 1.
3. US phone numbers are in the format XXX-XXX-XXXX. We don’t want to import invalid phone numbers so please write a
   validator on the converted data objects that ensures that only valid objects are converted (CSV to JSON and vice
   versa). Log the invalid objects, so they can easily be found in the source file.
4. Glasnostic now has thousands of employees coming from different regions. Help John to modify this program:
    - Read the input file.
    - Break up inputs into multiple workers
    - Aggregate results from workers
5. Sometimes we want to assign the order of the columns of the exported CSV fields. Extend the program from step 2 to
   support this. Order is given as CSV of the column names, e.g.: FirstName,ID,LastName,...
   
## Test Data

```csv
ID,FirstName,LastName,Email,Description,Role,Phone
1,Marc,Smith,marc@glasnostic.com,Writer of Java,Dev,541-754-3010
2,John,Young,john@glasnostic.com,Interested in MHW,HR,541-75-3010
3,Peter,Scott,peter@glasnostic.com,amateur boxer,Dev,541-754-3010
```

## Technologies

- [Travis-CI](https://travis-ci.com/blackhorseya/glasnostic) for CI/CD
- [Cobra](https://github.com/spf13/cobra) for implement command line
- [Protobuf](https://developers.google.com/protocol-buffers) for definition entities
- [Docker](https://www.docker.com/) for build binary

## How to use

```bash
glas

csvjson coding test

Usage:
  glas [command]

Available Commands:
  csv         transfer json file to csv
  help        Help about any command
  json        transfer csv file to json

Flags:
  -f, --file string     source data file path (default "./test/raw.csv")
  -h, --help            help for glas
  
  -o, --output string   output data file path
  -v, --version         version for glas

Use "glas [command] --help" for more information about a command.

```
