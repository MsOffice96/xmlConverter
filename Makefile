# CC=<컴파일러>
# CFLAGS=<컴파일 옵션>
# LDFLAGS=<링크 옵션>
# LDLIBS=<링크 라이브러리 목록>
# OBJS=<Object 파일 목록>
# TARGET=<빌드 대상 이름>
#  
# all: $(TARGET)
#  
# clean:
#     rm -f *.o
#     rm -f $(TARGET)
#  
# $(TARGET): $(OBJS)
# $(CC) -o $@ $(OBJS)


GOOS=linux # windows
GOARCH=amd64

BINARY_NAME = xmlConverter
MainFile = main.go

run: 
	go run main.go

windows: 
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}.exe ${MainFile}

linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} ${MainFile}

build:
	go build -o ${BINARY_NAME} main.go
	
clean:
	rm -rf ${BINARY_NAME}.*