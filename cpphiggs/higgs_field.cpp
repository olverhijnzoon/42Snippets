#include "higgs_field.h"
#include <cmath>

// Constructor for the HiggsField class
HiggsField::HiggsField(double mu2, double lambda) : mu2(mu2), lambda(lambda), vacuumExpectationValue(0.0) {
    calculateVEV();
}

// Method to calculate the Vacuum Expectation Value (VEV) based on the Higgs potential
void HiggsField::calculateVEV() {
    // VEV v = sqrt(-μ^2 / (2λ))
    if (lambda != 0.0) {
        vacuumExpectationValue = std::sqrt(-mu2 / (2 * lambda));
    }
}

// Getter method to retrieve the calculated VEV
double HiggsField::getVacuumExpectationValue() const {
    return vacuumExpectationValue;
}
