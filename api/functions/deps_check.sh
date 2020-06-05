deps_check() {
  if ! builtin command -v "$1" >/dev/null 2>&1; then
    printf '%s\n' "$2" >&2
    exit 1
  fi
}
