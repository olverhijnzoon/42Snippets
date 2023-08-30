#ifndef BOSON_H
#define BOSON_H

#include "particle.h"

/**
 * @class Boson
 * 
 * @brief Represents bosonic particles in the electroweak theory of the Standard Model.
 * 
 * Bosons are one of the two fundamental classes of subatomic particles, the other being fermions.
 * In the context of the electroweak theory, bosons are vector particles that mediate the weak
 * and electromagnetic forces. The most notable bosons in this theory are the W and Z bosons,
 * which mediate the weak force, and the photon, which mediates the electromagnetic force.
 * 
 * This class provides specific implementations for interactions with the Higgs field, 
 * which gives mass to the W and Z bosons through the mechanism of spontaneous electroweak symmetry breaking.
 */
class Boson : public Particle {
public:
    using Particle::Particle;  // Inherit the constructor of the base class

    virtual ~Boson() override = default;

    /**
     * @brief Interact with the Higgs field to determine the mass of the boson.
     * 
     * This method updates the mass of the boson based on its interaction with the Higgs field.
     * In the electroweak theory, the W and Z bosons acquire mass through their interaction with
     * the Higgs field, while the photon remains massless.
     * 
     * @param higgs The Higgs field with which the boson interacts.
     */
    void interactWithHiggs(const HiggsField& higgs) override;

    /**
     * @brief Simulate the interaction of the boson with other particles.
     * 
     * This method returns a description of the boson's interaction with other particles.
     * For the purpose of this simulation, specific interactions for bosons are not modeled.
     * 
     * @return A unique pointer to a ParticleInteraction object describing the interaction.
     */
    std::unique_ptr<ParticleInteraction> interact() override;
};

#endif // BOSON_H
