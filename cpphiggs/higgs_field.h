#ifndef HIGGS_FIELD_H
#define HIGGS_FIELD_H

/**
 * @class HiggsField
 * 
 * @brief Represents the Higgs Field in the Standard Model of particle physics.
 * 
 * The Higgs Field is a scalar field with a non-zero vacuum expectation value (VEV),
 * leading to the spontaneous breaking of the electroweak symmetry. This results in
 * the W and Z bosons acquiring mass. The excitation of this field is the Higgs boson.
 * 
 * This class models the Higgs potential and its VEV.
 */
class HiggsField {
private:
    // Parameters for the Higgs potential V(Φ) = μ^2 |Φ|^2 + λ |Φ|^4
    double mu2;
    double lambda;

    // Vacuum Expectation Value (VEV)
    double vacuumExpectationValue;

public:
    HiggsField(double mu2, double lambda);

    // Calculate the VEV based on the potential
    void calculateVEV();

    // Getter for the VEV
    double getVacuumExpectationValue() const;
};

#endif // HIGGS_FIELD_H
