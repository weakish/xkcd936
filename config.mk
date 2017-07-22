# paths
PREFIX = /usr/local

# cc options
CFLAGS += -std=c11 -Wall -Wextra -Wtype-limits -Wconversion -Ofast
CC = clang

CSTDFLAGS = -Wc90-c99-compat -Wc99-c11-compat -Wc++-compat
