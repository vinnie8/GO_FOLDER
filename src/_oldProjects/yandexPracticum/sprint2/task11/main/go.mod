module main

go 1.18
//write this -vvv- to help copiler search module in local directory
replace somemodule => ../somemodule

// use command-> go get somemodule 
require somemodule v0.0.0-00010101000000-000000000000
