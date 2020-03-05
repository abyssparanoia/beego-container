# note: call scripts from /scripts

format:
	$(call format)

build:
	go build -o default-server ./cmd/default

build-helper:
	go build -o helper ./cmd/helper

test:
	go test `go list ./... | grep -v pkg/dbmodels`

mockgen:
	$(call mockgen_app ,default)
	$(call format)

define format
	go fmt ./... && goimports -w ./ && go mod tidy

endef

define mockgen_app
	$(eval USECASE_LIST := $(call get_usecase_list,$1))
	$(foreach file, $(USECASE_LIST), $(call mockgen_usecase,$1,$(shell basename $(file))))
	$(eval SERVICE_LIST := $(call get_service_list,$1))
	$(foreach file, $(SERVICE_LIST), $(call mockgen_service,$1,$(shell basename $(file))))
	$(eval REPOSITORY_LIST := $(call get_repository_list,$1))
	$(foreach file, $(REPOSITORY_LIST), $(call mockgen_repository,$1,$(shell basename $(file))))

endef


define get_usecase_list
	$(shell	find ./internal/$1/usecase -mindepth 1 -maxdepth 1 -type f ! -name "*_impl*.go")

endef

define mockgen_usecase
	$(shell mockgen -source ./internal/$1/usecase/$2 -destination ./internal/$1/usecase/mock/$2)

endef

define get_service_list
	$(shell	find ./internal/$1/domain/service -mindepth 1 -maxdepth 1 -type f )

endef

define mockgen_service
	$(shell mockgen -source ./internal/$1/domain/service/$2 -destination ./internal/$1/domain/service/mock/$2)

endef

define get_repository_list
	$(shell	find ./internal/$1/domain/repository -mindepth 1 -maxdepth 1 -type f )
	
endef

define mockgen_repository
	$(shell mockgen -source ./internal/$1/domain/repository/$2 -destination ./internal/$1/domain/repository/mock/$2)

endef