god:
	@read -p "Enter 0 for golangclosures, 1 for intsequencegenerator, 2 for sumgenerator: " number; \
	case $$number in \
	0) BINARY=golangclosures ;; \
	1) BINARY=intsequencegenerator ;; \
	2) BINARY=sumgenerator ;; \
	esac; \
	echo Selected Binary: $$BINARY; \
	go build -o $($BINARY)X $$BINARY.go; \
	./$($BINARY)X

clean_binaries:
	rm -f ./*X
