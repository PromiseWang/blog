cmake_minimum_required(VERSION 3.21.1)
project(leetcode)

set(CMAKE_CXX_STANDARD 17)

include_directories(Header/)
add_executable(test 7_backtracking/day26/test.cpp)
add_executable(2_40 7_backtracking/day26/2_40.cpp)
add_executable(3_131 7_backtracking/day26/3_131.cpp)
