! 42Snippets
! Fortran Distance

! This Fortran program calculates the Euclidean, Manhattan and Chebyshev distances between two points in a 3D space. 
! It defines a derived type named Point that has three coordinates (x, y, z) 
! and functions for distance computation in a module named Geometry. 
! These functions compute the Euclidean and Manhattan distances between two points. 
! The main program creates two points with random coordinates, calculates the distances between them, 
! and prints the coordinates of the points and the distances.

! Euclidean Distance: The most common use of distance. In simple terms, it is the shortest distance between two points, and it can be used for all types of data.
! Manhattan Distance (L1 norm): Also known as city block distance. It is the distance between two points in a grid-based system (like a chess board or city blocks).
! Chebyshev Distance (L∞ norm): It is a distance measure between two vectors or points in N-dimensional vector space. It is also known as maximum value distance.
! Minkowski Distance: It is a generalized distance metric that includes the Euclidean, Manhattan, and Chebyshev distances as special cases for different values of a parameter called p. The formula for the Minkowski distance between two points p1(x1, y1, z1) and p2(x2, y2, z2) in a 3D space is (|x1-x2|^p + |y1-y2|^p + |z1-z2|^p)^(1/p). For p=2, this gives the Euclidean distance; for p=1, the Manhattan distance; and for p->∞, the Chebyshev distance.

! Define a module named Geometry
module Geometry
  implicit none
  ! Define a derived data type named Point
  type :: Point
     real :: x
     real :: y
     real :: z
  end type Point
contains
  ! Function to compute Euclidean distance between two points
  real function euclidean_distance(point1, point2)
    type(Point), intent(in) :: point1
    type(Point), intent(in) :: point2
    euclidean_distance = sqrt((point1%x - point2%x)**2 + (point1%y - point2%y)**2 + (point1%z - point2%z)**2)
  end function euclidean_distance

  ! Function to compute Manhattan distance between two points
  real function manhattan_distance(point1, point2)
    type(Point), intent(in) :: point1
    type(Point), intent(in) :: point2
    manhattan_distance = abs(point1%x - point2%x) + abs(point1%y - point2%y) + abs(point1%z - point2%z)
  end function manhattan_distance
  
  ! Function to compute Chebyshev distance between two points
  real function chebyshev_distance(point1, point2)
    type(Point), intent(in) :: point1
    type(Point), intent(in) :: point2
    chebyshev_distance = max(abs(point1%x - point2%x), abs(point1%y - point2%y), abs(point1%z - point2%z))
  end function chebyshev_distance
end module Geometry

program Main
  use Geometry
  implicit none
  type(Point) :: p1, p2
  real :: euclidean_dist, manhattan_dist, chebyshev_dist
  
  ! Initialize the random number generator
  call random_seed()

  ! Initialize points with random numbers
  call random_number(p1%x)
  call random_number(p1%y)
  call random_number(p1%z)

  call random_number(p2%x)
  call random_number(p2%y)
  call random_number(p2%z)
  
  euclidean_dist = euclidean_distance(p1, p2)
  manhattan_dist = manhattan_distance(p1, p2)
  chebyshev_dist = chebyshev_distance(p1, p2)
  
  print*, 'The coordinates of the first point are: ', p1%x, p1%y, p1%z
  print*, 'The coordinates of the second point are: ', p2%x, p2%y, p2%z
  print*, 'The Euclidean distance between points is: ', euclidean_dist
  print*, 'The Manhattan distance between points is: ', manhattan_dist
  print*, 'The Chebyshev distance between points is: ', chebyshev_dist
end program Main
