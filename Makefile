CDKTF_GENERATOR=.bin/cdktf-provider-gen
CDKTF_PACKAGES=$(patsubst %.cdktf.yaml,%,$(wildcard *.cdktf.yaml))

.PHONY : all
all: install $(CDKTF_PACKAGES)

.PHONY: install
install:
	go build -o $(CDKTF_GENERATOR) github.com/sourcegraph/cdktf-provider-gen/cmd/cdktf-provider-gen

.PHONY: targets
targets:
	@echo $(CDKTF_PACKAGES)

%: ./%.cdktf.yaml
	$(CDKTF_GENERATOR) --config $@