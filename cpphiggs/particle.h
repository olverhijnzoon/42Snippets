#pragma once
#include <string>
#include <memory>
#include "higgs_field.h"

// Forward declaration
class ParticleInteraction;

// Abstract base class representing a generic particle
class Particle {
protected:
    std::string name;  // Name of the particle
    double mass{0.0};  // Mass of the particle in GeV

public:
    explicit Particle(std::string n);
    virtual ~Particle() = default;

    virtual void interactWithHiggs(const HiggsField& higgs) = 0;
    virtual std::unique_ptr<ParticleInteraction> interact() = 0;  // New interaction method
    virtual void display() const;
};

// Class representing a particle interaction
class ParticleInteraction {
private:
    std::string description;

public:
    explicit ParticleInteraction(std::string desc);
    void display() const;
};
