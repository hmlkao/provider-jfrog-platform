
# 1: registry/org 2: repo
define xpkg.set-extensions
xpkg.set-extensions.$(1).$(2): $(UP)
	@$(INFO) Setting extensions for package $(1)/$(2):$(VERSION)
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
