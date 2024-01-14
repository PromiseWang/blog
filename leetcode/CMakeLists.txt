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

add_executable(3_37 7_backtracking/day29/3_37.cpp)
add_executable(day29 7_backtracking/day29/test.cpp)
add_executable(1_332 7_backtracking/day29/1_332.cpp)

add_executable(1_455 8_greed/day31/1_455.cpp)
add_executable(2_376 8_greed/day31/2_376.cpp)
add_executable(3_53 8_greed/day31/3_53.cpp)

add_executable(1_122 8_greed/day32/1_122.cpp)
add_executable(2_55 8_greed/day32/2_55.cpp)
add_executable(3_45 8_greed/day32/3_45.cpp)

add_executable(1_1005 8_greed/day33/1_1005.cpp)
add_executable(2_134 8_greed/day33/2_134.cpp)
add_executable(day33_test 8_greed/day33/test.cpp)
add_executable(3_135 8_greed/day33/3_135.cpp)

add_executable(1_860 8_greed/day34/1_860.cpp)
add_executable(day34_test 8_greed/day34/test.cpp)
add_executable(2_406 8_greed/day34/2_406.cpp)
add_executable(3_452 8_greed/day34/3_452.cpp)
