BINARY_DIR=bin/

SERVER_NAME=server
TERM_NAME=term
TEST_NAME=test
SERVER=cmd/server.go
TERM=cmd/term.go
TEST=cmd/test_main.go

all: compile

compile: | $(BINARY_DIR)
	@go build -o $(BINARY_DIR)$(SERVER_NAME) $(SERVER)
	@go build -o $(BINARY_DIR)$(TERM_NAME) $(TERM)
	@go build -o $(BINARY_DIR)$(TEST_NAME) $(TEST)
	@printf "\033[33mbuild\033[0m\n"

$(BINARY_DIR):
	@mkdir -p bin
	@printf "\033[36mcreate dir binary dir\033[0m\n"

run_server:
	@printf "\033[33mRUN SERVER\033[0m\n"
	@$(BINARY_DIR)$(SERVER_NAME)

run_term: compile
	@printf "\033[33mRUN TERM\033[0m\n"
	@$(BINARY_DIR)$(TERM_NAME)

run_test:
	@printf "\033[33mRUN TERM\033[0m\n"
	@$(BINARY_DIR)$(TEST_NAME)

clean:
	@rm -rf $(INFO_DIR) $(BINARY_DIR)
	@printf "\033[31mdeleted $(INFO_DIR)\033[0m\n\033[31mdeleted $(BINARY_DIR)\033[0m\n\033[33mClean ok\033[0m\n"