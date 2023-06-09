use rand::Rng;
use rand::seq::SliceRandom;

#[derive(Debug)]
pub enum Human {
    Neo,
    Trinity,
    Morpheus,
}

#[derive(Debug)]
pub enum Machine {
    AgentSmith,
    TheArchitect,
}

#[derive(Debug)]
pub enum Character {
    Human(Human),
    Machine(Machine),
}

pub struct MatrixWorld {
    characters: Vec<Character>,
}

impl MatrixWorld {
    pub fn new() -> MatrixWorld {
        MatrixWorld {
            characters: vec![],
        }
    }

    pub fn add_character(&mut self, character: Character) {
        self.characters.push(character);
    }

    pub fn get_character(&self, name: &str) -> Option<&Character> {
        for character in &self.characters {
            match character {
                Character::Human(human) => match human {
                    Human::Neo if name == "Neo" => return Some(character),
                    Human::Trinity if name == "Trinity" => return Some(character),
                    Human::Morpheus if name == "Morpheus" => return Some(character),
                    _ => {},
                },
                Character::Machine(machine) => match machine {
                    Machine::AgentSmith if name == "Agent Smith" => return Some(character),
                    Machine::TheArchitect if name == "The Architect" => return Some(character),
                    _ => {},
                },
            }
        }

        None
    }

    pub fn populate(&mut self) {
        let mut rng = rand::thread_rng();
        let character_count = rng.gen_range(1..=5);

        let mut characters: Vec<Character> = vec![
            Character::Human(Human::Neo),
            Character::Human(Human::Trinity),
            Character::Human(Human::Morpheus),
            Character::Machine(Machine::AgentSmith),
            Character::Machine(Machine::TheArchitect),
        ];

        characters.shuffle(&mut rng);

        for character in characters.into_iter().take(character_count) {
            self.add_character(character);
        }
    }
}

fn main() {

    println!("# 42Snippets");
    println!("## Enums in Rust");

    /*
        This Rust program simulates a simple model of the world from the Matrix movie. It uses Rust's powerful enum and match features to represent various characters and their factions (Humans and Machines), and provides functionality to interact with them.

        The main building blocks of the program are three enums: Human, Machine, and Character. Human and Machine represent the two factions and list some members of each, while Character is a more general type that can hold a Human or Machine variant. This makes use of Rust's ability to nest enums, reflecting the hierarchical relationship between characters and factions.

        The MatrixWorld struct represents the world of the Matrix. It stores a list of characters and provides methods to manipulate and query this list. These methods include:

          - new: This static method creates a new, empty Matrix world.
          - add_character: This method adds a character to the world.
          - get_character: This method returns a reference to a character, given the character's name as a string. It uses Rust's match feature to iterate through the list of characters and check each one's name.
          - populate: This method populates the Matrix world with a random number of characters. It ensures that each character only appears once in the world by creating a list of all possible characters, shuffling it, and then taking a certain number from the front. The randomness is achieved using methods from the rand crate.
    
        The main function creates a new Matrix world, populates it, checks for the presence of a specific character, and prints all characters. It serves as a simple demonstration of how the program's features can be used. 
    */

    let mut matrix = MatrixWorld::new();
    matrix.populate();

    // Uncomment this to add characters manually
    // matrix.add_character(Character::Human(Human::Neo));
    // matrix.add_character(Character::Machine(Machine::AgentSmith));

    let character = matrix.get_character("Neo");
    match character {
        Some(character) => println!("Character found: {:?}", character),
        None => println!("Character not found"),
    }

    // Print who is in the Matrix
    println!("### MATRIX meta ###");
    for character in &matrix.characters {
        println!("{:?}", character);
    }
}
