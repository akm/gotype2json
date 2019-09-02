# gotype2json

## Usage

```golang

tm := gotype2json.TypeMap{}
tm.Start(
	(*http.Request)(nil),
	(*http.Response)(nil),
)

// Output typemap...
```

See https://github.com/akm/gotype2json/blob/master/cmd/example1/main.go for more detail

Result JSON is https://github.com/akm/gotype2json/blob/master/cmd/example1/example.json
