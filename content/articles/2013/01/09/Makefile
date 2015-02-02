CC=gcc

main: func test
	$(CC) -o main remove-string-trailing-newline.o test-remove-string-trailing-newline.o

func: remove-string-trailing-newline.c
	$(CC) -c remove-string-trailing-newline.c

test: test-remove-string-trailing-newline.c
	$(CC) -c test-remove-string-trailing-newline.c

clean:
	rm *.o main
