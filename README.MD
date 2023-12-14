# Introduction

Helps map reduce to analyse large amount of contextual data, used to find context between a key and value.
Keys are Strings which have a large chuck of data that it could be matching with values.

# Experimental Code

Code is untested, It is basically a way to solve the classical Map/Reduce word count program to provide
lexical analysis to find potential word mappings which are there

# Experiment

An input of Names as "[Tom, Carl, Carly, Smith, Carly, Smith]" should return a mapping as "[1, 2, 3, 4, 3, 4]".

# Warranty

This code is an attempt at solving the Map/Reduce problem at a very generic level and has not be tested for the
experiment too. The input is just a string to the problem, not an file name as expected for an word count 
problem.
