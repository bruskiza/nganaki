	.PHONY: all help build test clean

	# Define a variable to hold the name of the project
	PROJECT_NAME := nganaki

	# A special variable for targets we want to hide from the help output
	HIDE_HELP := .PHONY all

	# Default target is 'help'
	all: help

	help: ## Show this help message.
		@echo ""
		@echo "** $(PROJECT_NAME) Makefile Usage **"
		@echo ""
		@echo "Available commands (make <command>):"
		@echo ""
		@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk ' \
			BEGIN {FS = ":.*?## "}; \
			{printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}; \
		'
		@echo ""


	build: ## Compile the source code and create the main executable.
		@echo "--- Building $(PROJECT_NAME) ---"
		# Example: Compile a simple C program (replace with your actual build command)
		go build -o $(PROJECT_NAME) main.go
		@echo "Build complete. Executable: $(PROJECT_NAME)"

	test: ## Run all unit and integration tests for the project.
		@echo "--- Running Tests ---"
		go test -coverprofile=coverage.out ./...
		
		go run gotest.tools/gotestsum@latest
		@echo "All tests passed successfully."

	coverage: ## Generate code coverage report.
		@echo "--- Generating Code Coverage Report ---"
		go tool cover -html=coverage.out -o coverage.html
		@echo "Coverage report generated: coverage.html"
		open coverage.html

	clean: ## Remove all generated files (executables, object files, etc.).
		@echo "--- Cleaning up build artifacts ---"
		# Example: Remove the compiled executable (replace with your actual clean commands)
		# rm -f $(PROJECT_NAME) *.o
		@echo "Cleanup finished."


