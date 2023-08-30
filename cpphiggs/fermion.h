#ifndef FERMION_H
#define FERMION_H

#include "particle.h"

/**
 * @class Fermion
 * 
 * @brief Represents fermionic particles in the Standard Model of particle physics.
 * 
 * Fermions are one of the two fundamental classes of subatomic particles, the other being bosons.
 * They are typically associated with matter and include particles such as quarks and leptons.
 * In the context of the electroweak theory, fermions acquire mass through their interaction
 * with the Higgs field via the mechanism of spontaneous electroweak symmetry breaking.
 * 
 * This class provides specific implementations for interactions with the Higgs field and
 * models the behavior of fermions like electrons, neutrinos, and quarks.
 */
class Fermion : public Particle {
public:
    using Particle::Particle;  // Inherit the constructor of the base class

    virtual ~Fermion() override = default;

    /**
     * @brief Interact with the Higgs field to determine the mass of the fermion.
     * 
     * This method updates the mass of the fermion based on its interaction with the Higgs field.
     * Different fermions have different masses, and this interaction models how they acquire
     * their respective masses in the presence of the Higgs field.
     * 
     * @param higgs The Higgs field with which the fermion interacts.
     */
    void interactWithHiggs(const HiggsField& higgs) override;

    /**
     * @brief Simulate the interaction of the fermion with other particles.
     * 
     * This method returns a description of the fermion's interaction with other particles.
     * For instance, an up quark might change to a down quark while emitting a W+ boson.
     * 
     * @return A unique pointer to a ParticleInteraction object describing the interaction.
     */
    std::unique_ptr<ParticleInteraction> interact() override;
};

#endif // FERMION_H
