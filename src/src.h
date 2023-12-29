#include <vector>

#ifndef SRC_H
#define SRC_H

// Vector Class
class Vector {
public:
  double x;
  double y;

  Vector(double _x = 0, double _y = 0);

  double mag();
  Vector get();
  Vector &add(Vector &v);
  Vector &sub(Vector &v);
  Vector &multi(double n);
  Vector &div(double n);
  Vector &normalize();
  Vector &limit(double max);
  static Vector AddVectors(Vector &v1, Vector &v2);
  static Vector SubVectors(Vector &v1, Vector &v2);
  static Vector Vector_multi(Vector &v1, double n);
  static Vector Vector_div(Vector &v1, double n);
};

// Physical Body Class
class Body {
public:
  double mass;
  Vector location = Vector(0, 0);
  Vector vel = Vector(0, 0);
  Vector acc = Vector(0, 0);
  double angle;

  Body(double _x, double _y, double _mass, double _angle);
  void update();
  void handle_force_stacking();
  void apply_force(Vector &force);
  void vel_direction();
  void handle_input();
};

// Physical World Class
class World {
public:
  std::vector<Body> bodies;
  void add_body(Body &body);
  void generate_friction();
};

// Important functions and raylib addons
void debug_body(Body &body);
void DrawRotationalLine(Body &body);

#endif
