make:
	g++ -o main \
	./main.cpp ./src/*.cpp \
	-I ./lib/raylib \
	-L ./lib/raylib \
	-l raylib
	./main
