TRAINER      = ba t
TRAINER_OPTS = --no-browser --debug
AGENT_BIN    = agent
DOCKER_IMAGE = go-agent
DOCKER_OPTS  = --no-cache

build:
	go build -o $(AGENT_BIN)
	docker build . -t $(DOCKER_IMAGE) $(DOCKER_OPTS)

run-test: build
	$(TRAINER) --agent $(DOCKER_IMAGE) --map training-dojo $(TRAINER_OPTS)

run: build
	$(TRAINER) --agent $(DOCKER_IMAGE) --map island $(TRAINER_OPTS)
