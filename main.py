# coding=utf-8

import os
import argparse
import ipaddress


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-b', '--blacklist', help='Black list of IP hosts/networks')
    parser.add_argument('-w', '--whitelist', help='White list of IP hosts/networks')
    return parser.parse_args()


def file2ipset(filename):
    output = set()
    if not os.path.isfile(filename):
        return output
    with open(filename) as __file:
        for line in __file.readlines():
            line = line.strip()
            if '/' in line:
                for ip in ipaddress.ip_network(line).hosts():
                    output.add(str(ip))
            else:
                output.add(line)
    return output


def main():
    args = parse_args()
    for ip in (file2ipset(args.blacklist) - file2ipset(args.whitelist)):
        print(ip)


if __name__ == '__main__':
    main()
