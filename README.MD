# Introduction

Helps map reduce to analyse large amount of contextual data, used to find context between a key and value.
Keys are Strings which have a large chuck of data that it could be matching with values.

# Experimental Code

Code is untested, It is basically a way to solve the classical Map/Reduce word count program to provide
lexical analysis to find potential word mappings which are there

# Experiment

An input of Names as "[Tom, Carl, Carly, Smith, Carly, Smith, Jason, Tom, Carl, Matthew]" should return a mapping as <br/> 

Map Job output "[(Tom, 0) (Carl, 0) (Carly 0) (Smith, 0) (Carly, 3) (Smith, 4) (Jason, 5) (Tom, 1) (Carl, 2) (Mathew, 0)]" <br/>
Reduce Job output "[(Tom, 1) (Carl, 2) (Carly, 3) (Smith, 4) (Jason, 5) (Mathew, 6)]" <br/>

# Warranty

This code is an attempt at solving the Map/Reduce problem at a very generic level and has not be tested for the
experiment too. The input is just a string to the problem, not an file name as expected for an word count 
problem.
