#[aoc_generator(day2)]
pub fn intcode_parser(input: &str) -> Vec<i32> {
  input
    .split(",")
    .map(|s| s.parse::<i32>().unwrap())
    .collect::<Vec<i32>>()
}

#[aoc(day2, part1)]
pub fn day2_part1(mem: &Vec<i32>) -> i32 {
  computer(mem, Some(12), Some(2))
}

#[aoc(day2, part2)]
pub fn day2_part2(mem: &Vec<i32>) -> i32 {
  let (n, v) = complete_maneuver(mem, 19690720);
  return 100 * n + v;
}

fn computer(mem_input: &Vec<i32>, noun_fix: Option<i32>, verb_fix: Option<i32>) -> i32 {
  let mut mem = mem_input.clone();

  let start_position = 0;

  if let Some(n) = noun_fix {
    mem[1] = n;
  }

  if let Some(v) = verb_fix {
    mem[2] = v;
  }

  computer_inter(&mut mem, start_position);

  return mem[0];
}

fn computer_inter(mem: &mut Vec<i32>, pos: usize) {
  let step = 4;
  let noun_position = 1;
  let verb_position = 2;

  let mut position = pos;

  loop {
    let addr = mem[&position + 3];
    let opcode = mem[position];
    let new_noun_pos = &position + &noun_position;
    let new_verb_pos = &position + &verb_position;

    match opcode {
      1 => {
        mem[addr as usize] =
          mem[mem[new_noun_pos as usize] as usize] + mem[mem[new_verb_pos as usize] as usize];
        position += &step;
      }
      2 => {
        mem[addr as usize] =
          mem[mem[new_noun_pos as usize] as usize] * mem[mem[new_verb_pos as usize] as usize];
        position += &step;
      }
      _ => break,
    }
  }
}

fn complete_maneuver(mem_input: &Vec<i32>, desired_output: i32) -> (i32, i32) {
  for i in 0..100 {
    for j in 0..100 {
      let result: i32 = computer(mem_input, Some(i), Some(j));
      if result == desired_output {
        return (i, j);
      }
    }
  }
  return (0, 0);
}
