#!/bin/env python3

from typing import Dict, List, Tuple, Union
from utils.print import print_part

DAY = "day3"
INPUT = "input/2021/" + DAY + ".txt"
EXAMPLE = """00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010"""


def parse(input: str) -> List[str]:
    return input.splitlines()


def reverse_bits(bin_string: str) -> str:
    return "".join("1" if x == "0" else "0" for x in bin_string)


def base2_from_str(string: str) -> int:
    return int(string, base=2)


def col_counts(lines: List[str]) -> dict[int, int]:
    col_counts: Dict[int, int] = {}

    for line in lines:
        for col in range(0, len(line)):
            i = int(line[col])
            col_counts[col] = col_counts.get(col, 0) + i

    return col_counts


def calc_gamma(col_counts: dict[int, int], total: int) -> str:
    gamma = ""

    for col, count in col_counts.items():
        result = ""
        if count > total / 2:
            result = "1"
        else:
            result = "0"
        gamma = gamma + result

    return gamma


def o2_rating(lines: List[str]) -> str:
    result = lines.copy()
    for col in range(len(result[0])):
        if len(result) == 1:
            break
        bits_at_col = [bits[col] for bits in result]
        common_bit = "1" if bits_at_col.count("1") >= len(result) / 2 else "0"
        result = [bits for bits in result if bits[col] == common_bit]

    return result[0]


def co2_rating(lines: List[str]) -> str:
    result = lines.copy()
    for col in range(len(result[0])):
        if len(result) == 1:
            break

        bits_at_col = [bits[col] for bits in result]
        common_bit = "0" if bits_at_col.count("1") >= len(result) / 2 else "1"
        result = [bits for bits in result if bits[col] == common_bit]

    return result[0]


def main():
    # input = parse(EXAMPLE)
    with open(INPUT, "r") as txt:
        input = parse(txt.read().rstrip())
        col_count = col_counts(input)

        gamma = calc_gamma(col_count, len(input))
        epsilon = reverse_bits(gamma)

        print_part(1)
        print(base2_from_str(gamma) * base2_from_str(epsilon))

        print()

        print_part(2)
        o2 = base2_from_str(o2_rating(input))
        co2 = base2_from_str(co2_rating(input))

        print(o2 * co2)


if __name__ == "__main__":
    main()
