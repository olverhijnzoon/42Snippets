#ifndef PARTICLE_H
#define PARTICLE_H

#include <string>
#include <memory>
#include "higgs_field.h"

/**
 * @class Particle
 * 
 * @brief Abstract base class representing a generic particle in the Standard Model of particle physics.
 * 
 * The Particle class provides a foundation for modeling different types of subatomic particles,
 * such as bosons and fermions. It defines common attributes like name and mass and provides
 * virtual methods for interactions with the Higgs field and other particles.
 * 
 * Derived classes should provide specific implementations for these interactions.
 */
class Particle {
protected:
    std::string name;  // Name of the particle
    double mass{0.0};  // Mass of the particle in GeV

public:
    explicit Particle(std::string n);
    virtual ~Particle() = default;

    /**
     * @brief Interact with the Higgs field.
     * 
     * This virtual method should be overridden by derived classes to provide specific
     * implementations for how different particles interact with the Higgs field.
     * 
     * @param higgs The Higgs field with which the particle interacts.
     */
    virtual void interactWithHiggs(const HiggsField& higgs) = 0;

    /**
     * @brief Simulate the interaction of the particle with other particles.
     * 
     * This virtual method should be overridden by derived classes to provide specific
     * descriptions of particle interactions.
     * 
     * @return A unique pointer to a ParticleInteraction object describing the interaction.
     */
    virtual std::unique_ptr<ParticleInteraction> interact() = 0;

    /**
     * @brief Display the particle's name and mass.
     */
    virtual void display() const;
};

/**
 * @class ParticleInteraction
 * 
 * @brief Represents the interaction of a particle with other particles.
 * 
 * The ParticleInteraction class provides a way to model and describe the interactions
 * between different particles. For instance, an up quark might change to a down quark
 * while emitting a W+ boson. This class encapsulates such interactions.
 */
class ParticleInteraction {
private:
    std::string description;  // Description of the interaction

public:
    explicit ParticleInteraction(std::string desc);

    /**
     * @brief Display the description of the particle interaction.
     */
    void display() const;
};

#endif // PARTICLE_H
