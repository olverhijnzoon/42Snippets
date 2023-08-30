#include "fermion.h"

void Fermion::interactWithHiggs(const HiggsField& higgs) {
    if (name == "electron") {
        mass = 0.511;
    } else if (name == "neutrino") {
        mass = 0.00001;
    } else if (name == "up quark" || name == "down quark") {
        mass = 2.3;
    }
}

std::unique_ptr<ParticleInteraction> Fermion::interact() {
    if (name == "up quark") {
        return std::make_unique<ParticleInteraction>("u -> d + W+");
    }
    return nullptr;  // No specific interaction for other fermions in this model
}
