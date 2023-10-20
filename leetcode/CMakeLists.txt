cmake_minimum_required(VERSION 3.21.1)
project("leetcode")

set(CMAKE_CXX_STANDARD 17)

include_directories(include/)
add_executable(test 7_backtracking/day26/test.cpp)
add_executable(2_40 7_backtracking/day26/2_40.cpp)
add_executable(3_131 7_backtracking/day26/3_131.cpp)
add_executable(1_93 7_backtracking/day27/1_93.cpp)

add_executable(test27 7_backtracking/day27/test.cpp)
add_executable(2_78 7_backtracking/day27/2_78.cpp)
add_executable(3_90 7_backtracking/day27/3_90.cpp)

add_executable(1_491 7_backtracking/day28/1_491.cpp)
add_executable(2_46 7_backtracking/day28/2_46.cpp)
add_executable(3_47 7_backtracking/day28/3_47.cpp)
