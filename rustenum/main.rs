// Define the factions as enums, each with some characters as variants
enum Human {
    Neo,
    Trinity,
    Morpheus,
}

enum Machine {
    AgentSmith,
    TheArchitect,
}

// Define a function that prints a message based on a character
fn print_message(character: Character) {
    match character {
        Character::Human(human) => match human {
            Human::Neo => println!("Neo: The One."),
            Human::Trinity => println!("Trinity: Neo's love."),
            Human::Morpheus => println!("Morpheus: The captain of the Nebuchadnezzar."),
        },
        Character::Machine(machine) => match machine {
            Machine::AgentSmith => println!("Agent Smith: A sentient Agent."),
            Machine::TheArchitect => println!("The Architect: The creator of the Matrix."),
        },
    }
}

// Define an enum for Character that can be either a Human or a Machine
enum Character {
    Human(Human),
    Machine(Machine),
}

fn main() {
    // Create some characters
    let neo = Character::Human(Human::Neo);
    let agent_smith = Character::Machine(Machine::AgentSmith);

    // Print messages for the characters
    print_message(neo);
    print_message(agent_smith);
}