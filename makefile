make:
	g++ -o main.a ./main.cpp ./src/*.cpp ./lib/libraylib.a
	./main.a
