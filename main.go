package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"flag"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func Hosts(CIDR string, output map[string]bool) {
	ip, mask, err := net.ParseCIDR(CIDR)
	if err == nil {
		for ip := ip.Mask(mask.Mask); mask.Contains(ip); inc(ip) {
			output[ip.String()] = true
		}
	}
}

func Host(CIDR string, output map[string]bool) {
	ip := net.ParseIP(CIDR)
	output[ip.String()] = true
}

func file2IPMap(filename string) map[string]bool {
	var CIDR string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	output := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		CIDR = scanner.Text()
		if strings.LastIndexByte(CIDR, '/') == -1 {
			Host(CIDR, output)
		} else {
			Hosts(CIDR, output)
		}
	}
	return output
}

func main() {
	blacklist_file := flag.String("b", "tests/data/blacklists/ip_https_plus", "Blacklist filename")
	whitelist_file := flag.String("w", "tests/data/workarounds/workaround.ip_https_plus_whitelist", "Blacklist filename")
	flag.Parse()
	blacklist := file2IPMap(*blacklist_file)
	whitelist := file2IPMap(*whitelist_file)
	for key := range (blacklist) {
		if whitelist[key] != true {
			fmt.Printf("%s/32\n", key)
		}
	}
}
