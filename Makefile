py:
	clear
	python3 Simon.py > Outputs/Python_Output.txt

# c:

# go:

js:
	clear
	node Simon.js > Outputs/JavaScript_Output.txt

all:
	clear
	node Simon.js > Outputs/JavaScript_Output.txt
	python3 Simon.py > Outputs/Python_Output.txt

concat: 
	node Simon.js > Outputs/Outputs.txt
	python3 Simon.py >> Outputs/Outputs.txt
