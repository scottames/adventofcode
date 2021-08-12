#!/bin/env python3

from __future__ import annotations
from collections import Counter
from math import prod

DAY = "day10"
INPUT = "input/2020/" + DAY + ".txt"

def main():
    chain = parse_input(INPUT)
    print("Part 1: " + str(diffs(chain)))
    print("Part 2: " + str(arrangements(chain)))

def parse_input(input: str) -> List:
    with open(input) as txt:
        sorted_list = sorted(list(map(int,txt.read().rstrip().split("\n"))))
        return [0] + sorted_list + [sorted_list[-1]+3]

def diffs(list: List[int]) -> int:
    return prod(Counter([list[i]-list[i-1] for i in range(1,len(list))]).values())

def arrangements(list: List[int]) -> int:
    paths = [1]
    for i in range(1,len(list)):
        paths.append(
            sum(
                [
                    get_from_list(paths,i-n)
                        for n in range(1,4)
                        if i-n >= 0
                        and list[i]-list[i-n] <= 3
                ]
            )
        )
    return paths[-1]

def get_from_list (list: List, index: int) -> int:
    try:
        return list[index]
    except IndexError:
        return 0

if __name__ == "__main__":
    main()
