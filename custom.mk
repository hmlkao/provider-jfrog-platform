### UP CLI
# up CLI was removed from build module, but crossplane CLI doesn't support 'xpkg append' subcommand
#   https://github.com/crossplane/build/commit/32266ec2df9b775cc23675131147009375fe9bb6
# So, it's added here as a custom workaround
UP_VERSION ?= v0.41.0
UP_CHANNEL ?= stable
UP := $(TOOLS_HOST_DIR)/up-$(UP_VERSION)

# up download and install
$(UP):
	@$(INFO) installing up $(UP_VERSION)
	@if [ ! -f $(UP) ]; then \
		curl -fsSLo $(UP) --create-dirs https://cli.upbound.io/$(UP_CHANNEL)/$(UP_VERSION)/bin/$(SAFEHOST_PLATFORM)/up?source=build || $(FAIL) && \
		chmod +x $(UP); \
	fi
	@$(OK) installing up $(UP_VERSION)

### Set package extensions
# 1: registry/org 2: repo
define xpkg.set-extensions
xpkg.set-extensions.$(1).$(2): $(UP)
	@$(INFO) Setting extensions for package $(1)/$(2):$(VERSION)
	@$(UP) version
	@$(UP) xpkg append \
		--extensions-root=./extensions \
		$(1)/$(2):$(VERSION)
	@$(OK) Extensions set for package $(1)/$(2):$(VERSION)
xpkg.set-extensions: xpkg.set-extensions.$(1).$(2)
endef
$(foreach r,$(XPKG_REG_ORGS), $(foreach x,$(XPKGS), $(eval $(call xpkg.set-extensions,$(r),$(x)))))

# Set extensions only for specific branches
set-extensions: ; @:
ifneq ($(filter $(RELEASE_BRANCH_FILTER),$(BRANCH_NAME)),)
set-extensions: $(foreach r,$(XPKG_REG_ORGS), $(foreach x,$(XPKGS),xpkg.set-extensions.$(r).$(x)))
endif
