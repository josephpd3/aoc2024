use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

#[derive(Debug)]
struct Input {
    left_list: Vec<i32>,
    right_list: Vec<i32>,
    right_counts: HashMap<i32, i32>,
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn get_input() -> Input {
    let mut left_list: Vec<i32> = Vec::new();
    let mut right_list: Vec<i32> = Vec::new();
    let mut right_counts: HashMap<i32, i32> = HashMap::new();

    if let Ok(lines) = read_lines("../data/input.txt") {
        for line in lines {
            if let Ok(ip) = line {
                let parts: Vec<&str> = ip.split_whitespace().collect();
                let left: i32 = parts[0].parse::<i32>().unwrap();
                let right: i32 = parts[1].parse::<i32>().unwrap();
                left_list.push(left);
                right_list.push(right);
                let count = right_counts.entry(right).or_insert(0);
                *count += 1;
            }
        }
    }

    left_list.sort();
    right_list.sort();

    Input {
        left_list,
        right_list,
        right_counts,
    }
}

fn abs_difference(a: i32, b: i32) -> i32 {
    if a > b {
        a - b
    } else {
        b - a
    }
}

fn get_distance(input : &Input) -> i32 {
    let mut distance = 0;
    for i in 0..input.left_list.len() {
        let left = input.left_list[i];
        let right = input.right_list[i];
        distance += abs_difference(left, right);
    }
    distance
}

fn get_similarity(input : &Input) -> i32 {
    let mut similarity = 0;
    for i in 0..input.left_list.len() {
        let left = input.left_list[i];
        let right_count = input.right_counts.get(&left).unwrap_or(&0);
        similarity += left * *right_count;
    }
    similarity
}

fn main() {
    let input = get_input();
    let distance = get_distance(&input);
    let similarity = get_similarity(&input);
    println!("Distance: {:?}", distance);
    println!("Similarity: {:?}", similarity);
}
