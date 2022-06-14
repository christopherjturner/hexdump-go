# hexdump-go
Basically a clone of hexdump -C (supports no other output modes!) but in go optimized a bit to dump really quickly.

## How fast?
pretty fast i guess.

```
$ time cat 100m.test |  hexdump -C > /dev/null
real	0m22.205s
user	0m22.072s
sys	0m1.084s
```

```
$ time cat 100m.test |  ./hexdump-go  > /dev/null
real	0m3.328s
user	0m3.266s
sys	0m0.409s
```

