#!/bin/env python3

def halving(string, num):
    top = num
    bottom = 0
    i = 0
    for s in string:
        i += 1
        middle = (top + bottom) // 2
        if s == "B" or s == "R":
            bottom = middle + 1
            if i == len(string):
                return top
        elif s == "F" or s == "L":
            top = middle
            if i == len(string):
                return bottom
    return 0


def missing(lst):
    return [x for x in range(lst[0], lst[-1] + 1)
            if x not in lst]


def main():
    with open("input/2020/day5.txt") as txt:
        lines = txt.readlines()
        seats = []
        for line in lines:
            line = line.strip()
            if line == "":
                continue
            rows = line[:7]
            cols = line[7:]
            seats.append(
              halving(rows, 127) * 8 + halving(cols, 7)
            )
        seats.sort()
    print("--- Part 1 ---\n")
    print("Highest: " + str(seats[-1]))
    print()
    print("--- Part 2 ---\n")
    print("Missing: " + str(missing(seats)))


if __name__ == "__main__":
    main()

