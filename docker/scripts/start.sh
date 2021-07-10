#!/usr/bin/env bash
# Start the container

# shellcheck disable=SC1091
. /scripts/common.sh || exit 1

# Bootstrap on first run
if ! [ -d "$CONF_UNBOUND" ] || [ -z "$(find "$CONF_UNBOUND")" ] ; then
    cp -pr /default/unbound "$CONF_UNBOUND"
    cp -pr /default/blocklist "$CONF_BLOCKLIST"
    mkdir "$CONF_UNBOUND"/include
fi

update_root_files
update_blocklist
fix_ownership
start_unbound
