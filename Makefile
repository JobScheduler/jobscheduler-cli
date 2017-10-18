all:
	go build -o jobscheduler

clean:
	rm -rf jobscheduler

re: clean all
