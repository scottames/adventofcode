#!/bin/env python3

from collections import Counter
from math import prod
from typing import List

DAY = "day10"
INPUT = "input/2020/" + DAY + ".txt"


def main():
    chain = parse_input(INPUT)
    print("Part 1: " + str(diffs(chain)))
    print("Part 2: " + str(arrangements(chain)))


def parse_input(day_input: str) -> List:
    with open(day_input) as txt:
        sorted_chain = sorted(list(map(int, txt.read().rstrip().split("\n"))))
        return [0] + sorted_chain + [sorted_chain[-1] + 3]


def diffs(chain: List[int]) -> int:
    return prod(Counter([chain[i] - chain[i - 1] for i in range(1, len(chain))]).values())


def arrangements(chain: List[int]) -> int:
    paths = [1]
    for i in range(1, len(chain)):
        paths.append(
            sum(
                [
                    get_from_chain(paths, i - n)
                    for n in range(1, 4)
                    if i - n >= 0 and chain[i] - chain[i - n] <= 3
                ]
            )
        )
    return paths[-1]


def get_from_chain(chain: List, index: int) -> int:
    try:
        return chain[index]
    except IndexError:
        return 0


if __name__ == "__main__":
    main()
