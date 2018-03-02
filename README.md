# cidr-diff

CLI tool for excluding IP from one IP/subnet list from second list.

## Most useful when combined with aggregate

```
$ cat ip 
192.168.0.1
192.168.0.10/32
$ cat subnet
192.168.0.0/28
./cidr-diff -b subnet -w ip | aggregate -q
192.168.0.0/32
192.168.0.2/31
192.168.0.4/30
192.168.0.8/31
192.168.0.11/32
192.168.0.12/30
```

## Speed

Input data

``` shell
➜  cidr-diff git:(master) ✗ wc -l blacklist                   
64914 blacklist
➜  cidr-diff git:(master) ✗ grep  / blacklist | grep -c /32
32522
➜  cidr-diff git:(master) ✗ grep  / blacklist | grep -vc /32
408
```

Small whitelist

```
➜  cidr-diff git:(master) ✗ shuf -n 10 blacklist > whitelist  
➜  cidr-diff git:(master) ✗ time ./cidr-diff -b blacklist -w whitelist >/dev/null
./cidr-diff -b blacklist -w whitelist > /dev/null  0,13s user 0,04s system 100% cpu 0,174 total
```

Big whitelist

```
➜  cidr-diff git:(master) ✗ cat blacklist > whitelist                            
➜  cidr-diff git:(master) ✗ time ./cidr-diff -b blacklist -w whitelist >/dev/null
./cidr-diff -b blacklist -w whitelist > /dev/null  0,15s user 0,01s system 106% cpu 0,148 total
```
