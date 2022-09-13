![DBNS](static/dbns_banner.png)

---

DBNS (DataBase for Nuclei Scanner) is a tool that allows you to keep track of the scans that are done with Nuclei in a simple way by saving the results in a database.

## Dependencies

- [Golang 1.17 or above](https://go.dev/doc/install)
- [ProjectDiscovery Nuclei](https://github.com/projectdiscovery/nuclei)
- [Postgres Database](https://www.postgresql.org/download/)

## Install
```
GO111MODULE=on go install -v github.com/FleexSecurity/dbns@latest
```

## Nuclei usage
```
dbns nuclei -u target.com
dbns nuclei -l all.txt
```
By default `dbns` ignores the results of `info`, but if you want to save them just add the `-i` flag.
```
dbns nuclei -l all.txt -i
```
`dbns nuclei -h` for more

## DB usage
DB is the command that returns the data from the Psql Database

Type of data that can be returned:
- `t: TemplateID`
- `h: Host`
- `s: Severity`
- `n: Name`
- `m: Matched`
- `g: Tags`

```
dbns db -p sm
```
`dbns db -h` for more

## Available commands
```
Available Commands:
  db          Retrieve data from db
  help        Help about any command
  nuclei      Nuclei Scanner command
```

# License
DBNS is distributed under Apache-2.0 License