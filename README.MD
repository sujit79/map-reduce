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

# Documentation to be followed

README at 
    https://github.com/BWbwchen/MapReduce/ <br/>
Intially how it was used 
```
    go run -race main.go -i 'input/files' -p 'wc.so' -r 1 -w 8 
```
Now supports -buildmode=plugin <br/>
How to Build <br/>
```
    go build -race -buildmode=plugin -o wc.so wc.go 
```
Master / Slave Examples <br/>
    # Word count
```
        go run -race cmd/master.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 8 &
        sleep 1
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 1 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 2 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 3 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 4 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 5 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 6 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 7 &
        go run -race cmd/worker.go -i 'txt/*' -p 'cmd/wc.so' -r 1 -w 8 
```


# Warranty

This code is an attempt at solving the Map/Reduce problem at a very generic level and has not be tested for the
experiment too. The input is just a string to the problem, not an file name as expected for an word count 
problem.
