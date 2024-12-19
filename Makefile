test-all:
	@find . -name "go.mod" ! -path "./go.mod" -execdir go test -cover ./... \;
