#!/bin/env python3

from typing import Dict, List

DAY = "day7"
INPUT = "input/2020/" + DAY + ".txt"
SHINY_GOLD = "shiny gold"

def main():
    bags = parse_input(INPUT)
    # print(bags)
    print("--- Part 1 ---\n")
    print("Num bags holds " + SHINY_GOLD + ": " + str(bags.num_holds(SHINY_GOLD)))
    print()
    print("--- Part 2 ---\n")
    print("Num bags " + SHINY_GOLD + " can contain: " + str(bags.holds_num(SHINY_GOLD)))

def parse_input(input: str) -> Dict:
    with open(input) as txt:
        content = txt.read().rstrip()
        return Bags(content)

class Bag:
    name = ""
    holds = {}

    def holds_bags(self) -> Dict[str, int]:
        return self.holds

    def __init__(self, bagStr: str) -> None:
        self.__parse(bagStr)

    def __parse(self, bagStr: str) -> None:
        split = [x.strip() for x in bagStr.split("bags contain")]
        self.name = split[0]
        self.holds = self.__parse_holds(split[1])

    def __parse_holds(self, holdsStr: str) -> Dict[str, int]:
        holds = {}
        if "no other bags" in holdsStr:
            return holds

        for contain in holdsStr.split(", "):
            tmp = contain.split(" ")
            num = int(tmp[0].strip())
            bag = tmp[1] + " " + tmp[2]
            holds[bag] = num

        return holds

class Bags:
    bags = {}

    def __init__(self, bagStr: str) -> None:
        self.bags = {}
        for line in bagStr.split("\n"):
            bag = Bag(line)
            self.bags[bag.name] = bag

    def holds_num(self, bag: str) -> int:
        if bag not in self.bags.keys():
            return 0

        count = 0
        for b, n in self.bags[bag].holds_bags().items():
            count = count + n + n * self.holds_num(b)

        return count

    def num_holds(self, bag: str) -> int:
        bags = []
        bs = [bag]
        while len(bs) > 0:
            bs = self.__num_holds(bs)
            if len(bs) > 0:
                for b in bs:
                    bags.append(b)

        return len(set(bags))

    def __num_holds(self, list: List[str]) -> List[str]:
        bags = []
        for self_bag in self.bags.values():
            for bag in list:
                if bag in self_bag.holds_bags().keys():
                    bags.append(self_bag.name)
        return bags

if __name__ == "__main__":
    main()
