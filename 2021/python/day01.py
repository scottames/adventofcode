#!/bin/env python3

from typing import List
from utils.print import print_part

DAY = "day1"
INPUT = "input/2021/" + DAY + ".txt"


def main():
    list = parse()

    print_part(1)
    print("Number of increases: ", str(increases_in_list(list)))
    print()

    print_part(2)
    print(
        "Number of three sequential increases: ",
        str(triple_sequential_increases_in_list(list)),
    )
    print()


def parse() -> List[int]:
    ints = []
    with open(INPUT, "r") as txt:
        ints = [int(i) for i in txt.read().rstrip().split("\n")]

    return ints


def increases_in_list(ints: List[int]) -> int:
    result = [n for i, n in enumerate(ints) if i != 0 and n > ints[i - 1]]

    return len(result)


def triple_sequential_increases_in_list(ints: List[int]) -> int:
    return increases_in_list([a + b + c for a, b, c in zip(ints, ints[1:], ints[2:])])


if __name__ == "__main__":
    main()
