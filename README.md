# JSON Timestamp Filter


## Summary

A tool for extracting portions of data from JSON streams using an ISO 8601 range as an input.

This is suitable for filtering typical Kafka or Druid datasets.

## To build

```
go build -o jtsf jsontimestampfilter.go
```

## To test

```
go test
```

## To use

The filter takes data from standard input and requires the ISO8601 range as a parameter.

An example ISO8601 range might be

```
"2017-11-29T14:00Z/2017-11-29T22:00Z"
```

Typical data might consist of a stream with lines like
```
{"timestamp":1511290852857,"animal":"rabbit","car":"mustang","fruit":"apple" }
```

An example of using the tool would be

```
cat largejsonfile.input | jtsf "2017-11-29T14:00Z/2017-11-29T22:00Z" > filteredjsonfile.output
