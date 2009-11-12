TARGET = goplot
SOURCES = *.go
OBJECTS = $(shell echo $(SOURCES) | sed -e 's,\.go,\.6,g')
CC = 6g
LD = 6l
all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(LD) -o $(TARGET) $(OBJECTS)

%.6 : %.go
	$(CC) $^

clean:
	rm -f $(TARGET) $(OBJECTS)
