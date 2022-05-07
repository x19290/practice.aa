from pathlib import Path
from subprocess import list2cmdline


def main():
    data = Path(__file__).with_name(r"data.txt").read_text()
    data = data.split("\n----\n")[:-1]
    g = data.__iter__()
    for y in g:
        feed, expected = g.__next__(), g.__next__()
        feed = feed.splitlines(keepends=False)
        expected = expected.splitlines(keepends=False)
        assert feed.__len__() == expected.__len__()
        for f, e in zip(feed, expected):
            y = list2cmdline((f,))
            assert y == e


if __name__ == r"__main__":
    main()
