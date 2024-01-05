make:
	g++ -o main ./main.cpp ./src/*.cpp ./lib/libraylib.a
	./main
