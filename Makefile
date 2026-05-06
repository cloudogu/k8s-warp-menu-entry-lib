# Set these to the desired values
PROJECT_NAME=k8s-warp-menu-entry-lib
ARTIFACT_ID=k8s-warp-menu-entry-crd
APPEND_CRD_SUFFIX=false
VERSION=0.0.1

MAKEFILES_VERSION=10.9.0

include build/make/variables.mk

include build/make/self-update.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-integration.mk
include build/make/test-unit.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/release.mk
include build/make/mocks.mk

include build/make/digital-signature.mk
include build/make/k8s-controller.mk

