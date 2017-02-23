# go-marc

## Important

Too soon.

## Really important

Not all of MARC. Probably not ever. Just the `034` field so far.

## Tools

### 034-to-bbox

Currently this only supports `hdddmmss (hemisphere-degrees-minutes-seconds)` and `dddmmss (degrees-minutes-seconds)` notation:

```
./bin/034-to-bbox
2017/02/13 22:23:38 1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000 <-- input (MARC 034)
2017/02/13 22:23:38 -70.000000, -180.000000 84.000000, 180.000000 <-- output (decimal WSG84)
```

## See also

* https://www.loc.gov/marc/bibliographic/bd034.html

