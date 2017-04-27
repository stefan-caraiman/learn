#!/bin/ruby

types_of_people = 10
x = "There are #{types_of_people} types of people."
y = "Hahah."
z = 'Types: #{types_of_people}' # will not replace the string, leaves it as it is
puts z
puts x + y
puts "bye "*2

a_multi_line = """I'm a very
long string, so good
bye!"""
puts a_multi_line