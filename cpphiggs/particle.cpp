#include "particle.h"
#include <iostream>

Particle::Particle(std::string n) : name(std::move(n)) {}

void Particle::display() const {
    std::cout << "Particle: " << name << ", Mass: " << mass << " GeV" << std::endl;
}

ParticleInteraction::ParticleInteraction(std::string desc) : description(std::move(desc)) {}

void ParticleInteraction::display() const {
    std::cout << description << std::endl;
}
