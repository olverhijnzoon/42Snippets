install:
	pip3 install requests

listener:
	printf "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nConnection: close\r\n\r\nhello 42Snippets" | nc -l 25000 &

run: listener
	python3 main.py
