#pragma once
#include "particle.h"

// Class representing bosons derived from Particle
class Boson : public Particle {
public:
    using Particle::Particle;  // Inherit the constructor of the base class
    virtual ~Boson() override = default;

    void interactWithHiggs(const HiggsField& higgs) override;
    std::unique_ptr<ParticleInteraction> interact() override;
};
