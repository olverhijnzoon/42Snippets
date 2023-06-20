BINARY=geometry

compile_fortran:
	gfortran -c $(BINARY).f90 -o $(BINARY).o

build_fortran:
	gfortran -o $(BINARY)X $(BINARY).o

run_fortran:
	./$(BINARY)X

fort: compile_fortran build_fortran run_fortran

clean_fortran:
	rm -f $(BINARY).o

clean_binaries:
	rm -f ./*X
