#include <iostream>

// Higgs Field class
class HiggsField {
private:
    double vacuumExpectationValue;

public:
    HiggsField() : vacuumExpectationValue(246.0) {}  // GeV

    double getVacuumExpectationValue() const {
        return vacuumExpectationValue;
    }
};

// Particle class
class Particle {
private:
    std::string name;
    double mass;

public:
    Particle(std::string n) : name(n), mass(0.0) {}

    void interactWithHiggs(const HiggsField& higgs) {
        if (name == "W" || name == "Z") {
            mass = higgs.getVacuumExpectationValue();  // Simplified representation
        } else if (name == "photon") {
            mass = 0.0;
        }
    }

    void display() const {
        std::cout << "Particle: " << name << ", Mass: " << mass << " GeV" << std::endl;
    }
};

int main() {
    HiggsField higgs;

    Particle W("W");
    Particle Z("Z");
    Particle photon("photon");

    W.interactWithHiggs(higgs);
    Z.interactWithHiggs(higgs);
    photon.interactWithHiggs(higgs);

    W.display();
    Z.display();
    photon.display();

    return 0;
}
