CC = gcc
CFLAGS = -Wall -Wextra -pedantic
TARGET = main
OBJS = main.o

all: $(TARGET)

$(TARGET): $(OBJS)
		$(CC) $(CFLAGS) -o $(TARGET) $(OBJS)

server.o: main.c
		$(CC) $(CFLAGS) -c main.c

clean:
		rm -f $(TARGET) $(OBJS)
