#!/usr/bin/env bash
# Update DNS configuration

# shellcheck disable=SC1091
. /scripts/common.sh || exit 1

update_root_files
update_blocklist
fix_ownership
reload_unbound
