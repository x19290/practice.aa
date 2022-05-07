#!/usr/bin/env python3


def main():
    from pathlib import Path
    data = Path(__file__).with_suffix(r".txt").read_text()
    data = data.split("\n----\n")
    feed0, feed1 = data[1], data[4]
    _expected(feed0)
    _expected(feed1)


def _expected(feed):
    from subprocess import list2cmdline
    for y in feed.splitlines(keepends=False):
        print(list2cmdline((y,)))


if __name__ == r'__main__':
    main()
