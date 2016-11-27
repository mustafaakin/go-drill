# go-drill
Go client for Apache Drill that utilizes REST API

## Query Example

```go
d, err := NewDrillbit("http://localhost:8047")

type fileQuery struct {
    Name             string
    IsDirectory      string
    IsFile           bool   `json:",string"`
    Length           uint64 `json:",string"`
    Owner            string
    Group            string
    Permissions      string
    AccessTime       string
    ModificationTime string
}

res := make([]fileQuery, 0)

err = d.Query(&res, "show files in dfs")
```

## TODO

-[X] Queries
-[] Profiles
-[] Storage plugins
-[] Stats
-[] Options
-[] Test and coverage reports
