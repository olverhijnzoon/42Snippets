#include "boson.h"

void Boson::interactWithHiggs(const HiggsField& higgs) {
    if (name == "W" || name == "Z") {
        mass = higgs.getVacuumExpectationValue();
    }
}

std::unique_ptr<ParticleInteraction> Boson::interact() {
    return nullptr;  // No specific interaction for bosons in this model
}
