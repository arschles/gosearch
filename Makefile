GIT_SHA=$(shell git rev-parse --short HEAD)

.PHONY: gitsha
gitsha:
	@echo ${GIT_SHA}

.PHONY: dockerimage
dockerimage:
	@echo "arschles/gosearch:$(GIT_SHA)"

.PHONY: dockerbuild
dockerbuild:
	docker build -t arschles/gosearch:$(GIT_SHA) .

.PHONY: dockerrun
dockerrun:
	docker run --rm -p 8080:8080 -e BING_SEARCH_KEY=${BING_SEARCH_KEY} arschles/gosearch:$(GIT_SHA)

.PHONY: dockerpush
dockerpush: dockerbuild
	docker push arschles/gosearch:$(GIT_SHA)

.PHONY: acideploy
acideploy: dockerpush
	az container create -g twitch -n gosearch --image arschles/gosearch:$(GIT_SHA) --dns-name-label=gosearch --environment-variables BING_SEARCH_KEY=${BING_SEARCH_KEY} PROD="prod" --ports 8080
	az container restart -n gosearch -g twitch

.PHONY: acilogs
acilogs:
	az container logs -n gosearch -g twitch --follow
