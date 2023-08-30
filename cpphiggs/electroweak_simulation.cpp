#include <iostream>
#include <memory>
#include <string>
#include <vector>
#include "higgs_field.h"

// Forward declaration
class ParticleInteraction;

// Base Particle class
class Particle {
protected:
    std::string name;
    double mass{0.0};

public:
    explicit Particle(std::string n) : name(std::move(n)) {}
    virtual ~Particle() = default;

    virtual void interactWithHiggs(const HiggsField& higgs) = 0;
    virtual std::unique_ptr<ParticleInteraction> interact() = 0;  // New interaction method

    virtual void display() const {
        std::cout << "Particle: " << name << ", Mass: " << mass << " GeV" << std::endl;
    }
};

// Particle Interaction class
class ParticleInteraction {
private:
    std::string description;

public:
    explicit ParticleInteraction(std::string desc) : description(std::move(desc)) {}

    void display() const {
        std::cout << description << std::endl;
    }
};

// Bosons
class Boson : public Particle {
public:
    using Particle::Particle;
    virtual ~Boson() override = default;

    void interactWithHiggs(const HiggsField& higgs) override {
        if (name == "W" || name == "Z") {
            mass = higgs.getVacuumExpectationValue();
        }
    }

    std::unique_ptr<ParticleInteraction> interact() override {
        return nullptr;  // No specific interaction for bosons in this model
    }
};

// Fermions
class Fermion : public Particle {
public:
    using Particle::Particle;
    virtual ~Fermion() override = default;

    void interactWithHiggs(const HiggsField& higgs) override {
        if (name == "electron") {
            mass = 0.511;
        } else if (name == "neutrino") {
            mass = 0.00001;
        } else if (name == "up quark" || name == "down quark") {
            mass = 2.3;
        }
    }

    std::unique_ptr<ParticleInteraction> interact() override {
        if (name == "up quark") {
            return std::make_unique<ParticleInteraction>("u -> d + W+");
        }
        return nullptr;  // No specific interaction for other fermions in this model
    }
};

int main() {
    auto higgs = std::make_unique<HiggsField>();

    std::vector<std::unique_ptr<Particle>> particles;
    particles.push_back(std::make_unique<Boson>("W"));
    particles.push_back(std::make_unique<Boson>("Z"));
    particles.push_back(std::make_unique<Boson>("photon"));
    particles.push_back(std::make_unique<Fermion>("electron"));
    particles.push_back(std::make_unique<Fermion>("neutrino"));
    particles.push_back(std::make_unique<Fermion>("up quark"));
    particles.push_back(std::make_unique<Fermion>("down quark"));

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
