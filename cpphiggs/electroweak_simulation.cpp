#include <iostream>
#include <memory>
#include <string>
#include <vector>
#include "higgs_field.h"
#include "particle.h"
#include "boson.h"
#include "fermion.h"

std::vector<std::unique_ptr<Particle>> createParticles() {
    std::vector<std::unique_ptr<Particle>> particles;
    particles.push_back(std::make_unique<Boson>("W"));
    particles.push_back(std::make_unique<Boson>("Z"));
    particles.push_back(std::make_unique<Boson>("photon"));
    particles.push_back(std::make_unique<Fermion>("electron"));
    particles.push_back(std::make_unique<Fermion>("neutrino"));
    particles.push_back(std::make_unique<Fermion>("up quark"));
    particles.push_back(std::make_unique<Fermion>("down quark"));
    return particles;
}

int main() {
    std::cout << "# 42Snippets" << std::endl;
    std::cout << "## Higgs" << std::endl;

    // Initialize Higgs Field
    double mu2_value = -2.0;
    double lambda_value = 0.5;
    auto higgs = std::make_unique<HiggsField>(mu2_value, lambda_value);
    std::cout << "Higgs Field VEV: " << higgs->getVacuumExpectationValue() << " GeV" << std::endl;

    // Create and interact particles
    auto particles = createParticles();
    for (const auto& particle : particles) {
        particle->interactWithHiggs(*higgs);
        particle->display();
        auto interaction = particle->interact();
        if (interaction) {
            interaction->display();
        }
    }

    return 0;
}