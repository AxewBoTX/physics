#include "src.h"
#include <random>

Perceptron::Perceptron(int n) {
  std::random_device rd;
  std::mt19937 gen(rd());
  std::uniform_real_distribution<double> distribution(-1.0, 1.0);
  double randomDouble = distribution(gen);
  for (int i = 0; i < n; i++) {
    this->weights[i] = randomDouble;
  }
};
int Perceptron::feed_forward(std::vector<int> &input) {
  double sum = 0;
  for (int i = 0; i < this->weights.size(); i++) {
    sum += input[i] * this->weights[i];
  }
  return activate_neuron(sum);
}
