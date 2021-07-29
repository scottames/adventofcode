#!/bin/env python3

from typing import List, Tuple, Dict

DAY = "day6"
INPUT = "input/2020/" + DAY + ".txt"


def main():
    groups = breakdown()
    print("--- Part 1 ---\n")
    print("Sum: " + str(sum_len(groups)))
    print()
    print("--- Part 2 ---\n")
    print("Sum: " + str(num_same_answers(groups)))


def sum_len(groups: List[Tuple[int, Dict]]) -> int:
    return sum([len(group[1]) for group in groups])


def num_same_answers(groups: List[Tuple[int, Dict]]) -> int:
    right = 0
    for group in groups:
        right += len([x for x in group[1].values() if x == group[0]])
    return right


def breakdown() -> List[Tuple[int, Dict]]:
    groups = []
    with open(INPUT) as txt:
        content = txt.read().rstrip()
        for raw_groups in content.split("\n\n"):
            answers = {}
            people = 0
            for raw_group in raw_groups.split("\n"):
                people += 1
                for a in raw_group:
                    if a in answers:
                        answers[a] = answers[a] + 1
                    else:
                        answers[a] = 1
            groups.append((people, answers))
    return groups


if __name__ == "__main__":
    main()
