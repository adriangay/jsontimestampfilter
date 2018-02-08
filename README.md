# JSON Timestamp Filter


## Summary

A simple tool for extracting a portion of data from a json stream using ISO 8601 ranges as an input

## To build

Run 

```
go build -o jtsf jsontimestampfilter.go
```

## To Test

Run

```
go test
```
in the top level directory

## To use

The filter takes data from standard input and requires the ISO8601 range as a parameter.

An example date might be

```
"2017-11-29T14:00Z/2017-11-29T22:00Z"
```

An example of using the tool would be

```
cat largejsonfile.input | jtsf "2017-11-29T14:00Z/2017-11-29T22:00Z" > filteredjsonfile.output