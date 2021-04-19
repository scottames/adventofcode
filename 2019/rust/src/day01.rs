#[aoc_generator(day1)]
pub fn fuel_calculation(input: &str) -> Vec<i32> {
    input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<_>>()
}

#[aoc(day1, part1)]
pub fn day1_part1(modules: &Vec<i32>) -> i32 {
    let mut fuel: i32 = 0;

    for module in modules {
        fuel += module / 3 - 2;
    }

    return fuel;
}

#[aoc(day1, part2, map_fold)]
pub fn day1_part2_map_fold(modules: &Vec<i32>) -> i32 {
    modules.iter().map(calculate_fuel).fold(0, |acc, x| acc + x)
}

#[aoc(day1, part2, for_loop)]
pub fn day1_part2_for_loop(modules: &Vec<i32>) -> i32 {
    let mut fuel: i32 = 0;

    for module in modules {
        fuel += calculate_fuel(module)
    }
    return fuel;
}

fn calculate_fuel(mass: &i32) -> i32 {
    let fuel = mass / 3 - 2;
    if fuel <= 0 {
        return 0;
    }
    return fuel + calculate_fuel(&fuel);
}
