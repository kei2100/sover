sover
==

Just sort SemVer strings

```bash
$ sover v0.10.0 v0.2.0 v0.1.0 v0.1.0-rc.0
v0.1.0-rc.0
v0.1.0
v0.2.0
v0.10.0

$ git tag -l
2.0.0
2.0.1
2.0.10
2.0.2
v1.0.0
v1.0.0-alpha
v1.0.0-rc.0
v1.0.0-rc.1
v1.0.0-rc.1+tag
v1.10.0
v1.2.0

$ git tag -l | sover
v1.0.0-alpha
v1.0.0-rc.0
v1.0.0-rc.1
v1.0.0-rc.1+tag
v1.0.0
v1.2.0
v1.10.0
2.0.0
2.0.1
2.0.2
2.0.10
```

## install

`$ go get github.com/kei2100/sover`

or

download binary
