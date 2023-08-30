#include <iostream>
#include <memory>
#include <string>
#include <vector>

// Higgs Field class
class HiggsField {
private:
    constexpr static double vacuumExpectationValue = 246.0;  // GeV

public:
    double getVacuumExpectationValue() const {
        return vacuumExpectationValue;
    }
};

// Base Particle class
class Particle {
protected:
    std::string name;
    double mass{0.0};

public:
    explicit Particle(std::string n) : name(std::move(n)) {}
    virtual void interactWithHiggs(const HiggsField& higgs) = 0;
    virtual void display() const {
        std::cout << "Particle: " << name << ", Mass: " << mass << " GeV" << std::endl;
    }
};

// Bosons
class Boson : public Particle {
public:
    using Particle::Particle;
    void interactWithHiggs(const HiggsField& higgs) override {
        if (name == "W" || name == "Z") {
            mass = higgs.getVacuumExpectationValue();  // Simplified representation
        }
        // photon remains massless by default
    }
};

// Fermions
class Fermion : public Particle {
public:
    using Particle::Particle;
    void interactWithHiggs(const HiggsField& higgs) override {
        if (name == "electron") {
            mass = 0.511;  // MeV, simplified representation
        } else if (name == "neutrino") {
            mass = 0.00001;  // MeV, very tiny mass
        } else if (name == "up quark" || name == "down quark") {
            mass = 2.3;  // MeV for up quark, simplified representation
        }
    }
};

int main() {
    auto higgs = std::make_unique<HiggsField>();

    std::vector<std::unique_ptr<Particle>> particles;
    particles.push_back(std::make_unique<Boson>("W"));
    particles.push_back(std::make_unique<Boson>("Z"));
    particles.push_back(std::make_unique<Boson>("photon"));
    particles.push_back(std::make_unique<Fermion>("electron"));
    particles.push_back(std::make_unique<Fermion>("neutrino"));
    particles.push_back(std::make_unique<Fermion>("up quark"));
    particles.push_back(std::make_unique<Fermion>("down quark"));

    for (const auto& particle : particles) {
        particle->interactWithHiggs(*higgs);
        particle->display();
    }

    return 0;
}
