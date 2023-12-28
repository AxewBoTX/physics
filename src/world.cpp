#include "math.h"
#include "src.h"

// Appending new objects to the world
void World::add_body(Body &body) { this->bodies.push_back(body); }

// Generating friction
void World::generate_friction() {
  for (int i = 0; i < this->bodies.size(); i++) {
    this->bodies[i].apply_force(this->bodies[i].vel.get().multi(-1).multi(0.5));
    if (this->bodies[i].vel.x < 0.3 && this->bodies[i].vel.x > -0.3) {
      this->bodies[i].vel.x = 0;
    }
    if (this->bodies[i].vel.y < 0.3 && this->bodies[i].vel.y > -0.3) {
      this->bodies[i].vel.y = 0;
    }
  }
}
