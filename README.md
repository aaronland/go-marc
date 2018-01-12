# go-marc

## Important

Too soon.

## Really important

Not all of MARC. Probably not ever. Just the `034` field so far.

## Tools

### marc-034

Currently this only supports `hdddmmss (hemisphere-degrees-minutes-seconds)` and `dddmmss (degrees-minutes-seconds)` notation:

```
./bin/marc-034
2017/02/13 22:23:38 1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000 <-- input (MARC 034)
2017/02/13 22:23:38 -70.000000, -180.000000 84.000000, 180.000000 <-- output (decimal WSG84)
```

### marc-034d

```
curl 'http://localhost:8080/bbox?034=1%23%24aa$b22000000%24dW1800000%24eE1800000%24fN0840000%24gS0700000`
```

## Docker

[Yes](Docker}, for `marc-034d` at least.

## See also

* https://www.loc.gov/marc/bibliographic/bd034.html

