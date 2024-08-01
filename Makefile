NAME := go-windows-junction
DESCRIPTION := A tool to manipulate Windows junctions and Unix hard links.
COPYRIGHT := 2024 © Andrea Funtò
LICENSE := MIT
LICENSE_URL := https://opensource.org/license/mit/
VERSION_MAJOR := 0
VERSION_MINOR := 0
VERSION_PATCH := 1
VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)
MAINTAINER=dihedron.dev@gmail.com
VENDOR=dihedron.dev@gmail.com
PRODUCER_URL=https://github.com/dihedron/
DOWNLOAD_URL=$(PRODUCER_URL)${NAME}
METADATA_PACKAGE=$$(grep "module .*" go.mod | sed 's/module //gi')/commands

_RULES_MK_OMIT_GO_GENERATE=1
_RULES_MK_MINIMUM_VERSION=202408011410
include rules.mk
