# Makefile for compiling the project

# Compiler and flags
CC := cc
CFLAGS := -Wall -W -pedantic -std=c99
CXX := g++
CXXFLAGS :=

# Directories
SRC_DIR := src
OBJ_DIR := dist

# C source files
C_SOURCES := $(wildcard $(SRC_DIR)/*.c)
C_OBJECTS := $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(C_SOURCES))

# Go source files
GO_SOURCES := $(wildcard $(SRC_DIR)/**/*.go)
GO_BINARIES := $(patsubst $(SRC_DIR)/%.go, $(OBJ_DIR)/%, $(GO_SOURCES))

# Targets
all: $(OBJ_DIR) pegasus pegasusedit traceroute getIP mac sniffer $(GO_BINARIES)

$(OBJ_DIR):
	mkdir -p $(OBJ_DIR)

pegasus: $(C_OBJECTS)
	$(CC) $(CFLAGS) $^ -o $(OBJ_DIR)/$@

pegasusedit: $(SRC_DIR)/pegasus-edit/editor.c
	$(CC) $(CFLAGS) $^ -o $(OBJ_DIR)/$@

traceroute: $(SRC_DIR)/tools/traceroute/route.c
	$(CC) $(CFLAGS) $^ -o $(OBJ_DIR)/$@

getIP: $(SRC_DIR)/tools/getIP/getIP.cc
	$(CXX) $(CXXFLAGS) $^ -o $(OBJ_DIR)/$@

mac: $(SRC_DIR)/tools/mac-spoof/mac.cc
	$(CXX) $(CXXFLAGS) $^ -o $(OBJ_DIR)/$@

sniffer: $(SRC_DIR)/tools/packet-sniffer/sniffer.c
	$(CC) $^ -lpcap -o $(OBJ_DIR)/$@

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@

$(OBJ_DIR)/%: $(SRC_DIR)/%.go
	go build -o $@ $<

clean:
	rm -rf $(OBJ_DIR)

.PHONY: all clean
