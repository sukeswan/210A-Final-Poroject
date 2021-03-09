py:
	clear
	python3 Simon.py > Outputs/Python_Output.txt

go:
	clear
	go run Simon.go > Outputs/Go_Output.txt

js:
	clear
	node Simon.js > Outputs/JavaScript_Output.txt

all:
	clear
	node Simon.js > Outputs/JavaScript_Output.txt
	python3 Simon.py > Outputs/Python_Output.txt
	go run Simon.go > Outputs/Go_Output.txt

concat: 
	node Simon.js > Outputs/Outputs.txt
	python3 Simon.py >> Outputs/Outputs.txt
	go run Simon.go >> Outputs/Go_Output.txt
