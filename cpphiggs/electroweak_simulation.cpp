#include <iostream>
#include <memory>
#include <string>

// Higgs Field class
class HiggsField {
private:
    constexpr static double vacuumExpectationValue = 246.0;  // GeV

public:
    double getVacuumExpectationValue() const {
        return vacuumExpectationValue;
    }
};

// Particle class
class Particle {
private:
    std::string name;
    double mass{0.0};

public:
    explicit Particle(std::string n) : name(std::move(n)) {}

    void interactWithHiggs(const HiggsField& higgs) {
        if (name == "W" || name == "Z") {
            mass = higgs.getVacuumExpectationValue();  // Simplified representation
        }
        // photon remains massless by default
    }

    void display() const {
        std::cout << "Particle: " << name << ", Mass: " << mass << " GeV" << std::endl;
    }
};

int main() {
    auto higgs = std::make_unique<HiggsField>();

    auto W = std::make_unique<Particle>("W");
    auto Z = std::make_unique<Particle>("Z");
    auto photon = std::make_unique<Particle>("photon");

    W->interactWithHiggs(*higgs);
    Z->interactWithHiggs(*higgs);
    photon->interactWithHiggs(*higgs);

    W->display();
    Z->display();
    photon->display();

    return 0;
}
