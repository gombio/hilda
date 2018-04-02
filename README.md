# Config

Optionally create `server.yaml` file alongside `main.go` file with example data:

```
request: false
http: true
services: false

```
where 'true' means disable report component

If you want to override setting from file use flags e.g.

```
hilda test http://example.com -r
hilda test http://example.com -r -w=false
hilda test http://example.com -s=true
```
