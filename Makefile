NAME = jobscheduler
DIST_DIR = dist

$(NAME):
	go build -o $(NAME)

clean:
	rm -rf $(NAME) $(DIST_DIR)

re: clean $(NAME)

dir:
	mkdir -p $(DIST_DIR)

all: $(NAME) darwin-386 darwin-amd64 linux-arm linux-arm64 linux-386 linux-amd64 windows-386 windows-amd64 freebsd-386 freebsd-amd64

darwin-386:
	GOOS=darwin GOARCH=386 go build -o $(DIST_DIR)/$(NAME)-darwin-386

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(NAME)-darwin-amd64

linux-arm:
	GOOS=linux GOARCH=arm go build -o $(DIST_DIR)/$(NAME)-linux-arm

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(NAME)-linux-arm64

linux-386:
	GOOS=linux GOARCH=386 go build -o $(DIST_DIR)/$(NAME)-linux-386

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(NAME)-linux-amd64

windows-386:
	GOOS=windows GOARCH=386 go build -o $(DIST_DIR)/$(NAME)-windows-386

windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(NAME)-windows-amd64

freebsd-386:
	GOOS=freebsd GOARCH=386 go build -o $(DIST_DIR)/$(NAME)-freebsd-386

freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 go build -o $(DIST_DIR)/$(NAME)-freebsd-amd64
