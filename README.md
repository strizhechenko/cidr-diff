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
