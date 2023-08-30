#ifndef HIGGS_FIELD_H
#define HIGGS_FIELD_H

/**
 * @class HiggsField
 * 
 * @brief Represents the Higgs Field in the Standard Model of particle physics.
 * 
 * The Higgs Field is a scalar field that permeates all of space. When particles
 * interact with this field, they acquire mass. The strength of their interaction
 * determines how much mass they obtain. The excitation of this field is what we
 * observe as the Higgs boson.
 * 
 * This class provides a simplified representation of the Higgs Field, primarily
 * its vacuum expectation value, which is crucial for particles to acquire mass.
 */
class HiggsField {
private:
    constexpr static double vacuumExpectationValue = 246.0;  // GeV

public:
    double getVacuumExpectationValue() const;
};

#endif // HIGGS_FIELD_H
