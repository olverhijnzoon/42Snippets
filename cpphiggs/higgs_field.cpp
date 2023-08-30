#include "higgs_field.h"
#include <cmath>

HiggsField::HiggsField(double mu2, double lambda) : mu2(mu2), lambda(lambda), vacuumExpectationValue(0.0) {
    calculateVEV();
}

void HiggsField::calculateVEV() {
    // VEV v = sqrt(-μ^2 / (2λ))
    if (lambda != 0.0) {
        vacuumExpectationValue = std::sqrt(-mu2 / (2 * lambda));
    }
}

double HiggsField::getVacuumExpectationValue() const {
    return vacuumExpectationValue;
}
