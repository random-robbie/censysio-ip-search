# censysio-ip-search
Provides a list of IP's for your search.

Important
----

```
Ensure you change MYKEY inside the censys-ip.go to your basic auth you can grab that from https://search.censys.io/api#/hosts/searchHosts when you are logged in
```

Install
----

```
git clone https://github.com/random-robbie/censysio-ip-search
cd censysio-ip-search
go build censys-ip.go
censys-ip -search apache
```
