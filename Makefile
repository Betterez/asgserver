default:
	@export GOPATH=$$GOPATH:$$(pwd) && go install asgserver
edit:
	@export GOPATH=$$GOPATH:$$(pwd) && atom .
edit2:
	@export GOPATH=$$GOPATH:$$(pwd) && code .
edit3:
	@export GOPATH=$$GOPATH:$$(pwd) && /home/tal/programs/idea/bin/idea.sh .
run: default
	@bin/asgserver
	@echo ""
clean:
	@rm -rf bin
test:
	@export GOPATH=$$GOPATH:$$(pwd) && go test ./...
test_ver:
	@export GOPATH=$$GOPATH:$$(pwd) && go test -v ./...
merge:
	git checkout master && git merge dev && git checkout dev && git push origin --all
