#!/usr/bin/env bash
set -euo pipefail

repo_root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root_dir"

bad_root="apis/project-planton"

if [[ -e "$bad_root" ]]; then
  echo "ERROR: Found incorrect path: $bad_root" >&2
  echo "Use 'apis/project/planton' (slash) instead of 'apis/project-planton' (hyphen)." >&2
  echo "Please move or delete the contents of: $bad_root" >&2
  exit 1
fi

bad_dirs=( $(find apis -path "apis/internal/generated" -prune -o -type d -name "project-planton" -print 2>/dev/null || true) )
if [[ ${#bad_dirs[@]} -gt 0 ]]; then
  echo "ERROR: Found directories named 'project-planton' under 'apis/'. Should be 'project/planton'." >&2
  printf '%s\n' "${bad_dirs[@]}" >&2
  exit 1
fi

echo "Path guard passed: no 'apis/project-planton' misuse detected."

