#!/usr/bin/env python3
from subprocess import list2cmdline
from sys import argv, stdout
b = stdout.buffer
for y in argv[1:]:
	b.write(br"<%s>\n" % y)
	b.write(br"<%s>\n" % list2cmdline([y]))
