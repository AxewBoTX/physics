#ifndef SRC_H
#define SRC_H

// Vector Class
class Vector {
public:
  double x;
  double y;

  Vector(float _x, float _y);

  float mag();
  Vector get();
  Vector &add(Vector &v);
  Vector &sub(Vector &v);
  Vector &multi(double n);
};

#endif
