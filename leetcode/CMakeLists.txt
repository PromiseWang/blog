cmake_minimum_required(VERSION 3.21.1)
project(leetcode)

set(CMAKE_CXX_STANDARD 17)

include_directories(Header/)
add_executable(test 7_backtracking/day26/test.cpp)
add_executable(2_40 7_backtracking/day26/2_40.cpp)
add_executable(3_131 7_backtracking/day26/3_131.cpp)
add_executable(1_93 7_backtracking/day27/1_93.cpp)

add_executable(test27 7_backtracking/day27/test.cpp)
add_executable(2_78 7_backtracking/day27/2_78.cpp)
add_executable(3_90 7_backtracking/day27/3_90.cpp)
