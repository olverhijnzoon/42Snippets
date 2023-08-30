#pragma once
#include "particle.h"

// Class representing fermions derived from Particle
class Fermion : public Particle {
public:
    using Particle::Particle;  // Inherit the constructor of the base class
    virtual ~Fermion() override = default;

    void interactWithHiggs(const HiggsField& higgs) override;
    std::unique_ptr<ParticleInteraction> interact() override;
};
