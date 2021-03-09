# 210A-Final-Project
## Surya Keswani and Donnie Stewart

This project tests the performance of JavaScript, Golang, and Python3 by measuring time and space metrics. Simon, a lightweight block cipher was implemented in each of the 3 languages and performance metrics were run on each implementation across 10 runs. 

The `/Outputs` folder has the results for each language as well as all the results concated in the `Outputs.txt` file. 

To run this project: 

1. Clone the repo or download the source code
2. `cd` into the local repo you just cloned/downloaded 
3. Run the corresponsding make command listed below 

`make js`: runs the javascript implementation of Simon and updates the results in `/Outputs/JavaScript_Output.txt`
`make py`: runs the python3 implementation of Simon and updates the results in `/Outputs/Python_Output.txt`
`make go`: runs the golang implementation of Simon and updates the results in `/Outputs/Go_Output.txt`
`make all`: runs the all 3 implementations in and updates the correpsonding text files 
`make concat`: runs the all 3 implementations and updates the results in te single text file `/Outputs/Output.txt`