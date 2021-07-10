#!/usr/bin/env bash

CONF_DIR=/conf
CONF_UNBOUND="$CONF_DIR/unbound"
CONF_BLOCKLIST="$CONF_DIR/blocklist"
ROOT_KEY="$CONF_UNBOUND/root.key"
ROOT_HINTS="$CONF_UNBOUND/root.hints"

fix_ownership() {
    chown -R unbound: "$CONF_UNBOUND"
}

update_blocklist() {
    if file_older_than "$CONF_UNBOUND"/include/blocklist.conf 180m; then
        build_blocklist \
            -sources "$CONF_BLOCKLIST"/sources.json \
            -output "$CONF_UNBOUND"/include/blocklist.conf
    fi
}

update_root_files() {
    if file_older_than "$ROOT_HINTS" 7d; then
        # Retrieve root hints file
        curl -so "$ROOT_HINTS" https://www.internic.net/domain/named.root &&
            # Setup root anchor
            unbound-anchor -r "$ROOT_HINTS" -a "$ROOT_KEY"
    fi
}

start_unbound() {
    unbound -d -c "$CONF_UNBOUND/unbound.conf"
}

reload_unbound() {
    unbound-control -qc "$CONF_UNBOUND/unbound.conf" reload
}

file_older_than() {
    local f=$1 a=$2
    if ! [ -e "$f" ]; then
        # File does not exist, needs updating
        return 0
    fi
    # File exists, check if it is old enough to be updated
    [[ "$a" =~ ^[0-9]+[md]$ ]] || return 1
    case "$a" in
        *m)
            [[ -n "$(find "$f" -mmin +"${a/m}")" ]] && return 0
            ;;
        *d)
            [[ -n "$(find "$f" -mtime +"${a/d}")" ]] && return 0
            ;;
    esac
    return 1
}
