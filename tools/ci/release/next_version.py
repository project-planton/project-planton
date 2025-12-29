#!/usr/bin/env python3
"""Calculate next semantic version based on git tags."""

import re
import subprocess
import sys

# Strict semver pattern: vX.Y.Z where X, Y, Z are digits only
SEMVER_PATTERN = re.compile(r"^v(\d+)\.(\d+)\.(\d+)$")


def get_latest_tag() -> str:
    """Get the latest version tag matching strict vX.Y.Z pattern."""
    result = subprocess.run(
        ["git", "tag", "--list", "v*", "--sort=-v:refname"],
        capture_output=True,
        text=True,
    )

    if result.returncode != 0:
        return "v0.0.0"

    # Find the first tag that matches strict semver pattern
    for tag in result.stdout.strip().split("\n"):
        tag = tag.strip()
        if tag and SEMVER_PATTERN.match(tag):
            return tag

    return "v0.0.0"


def bump_version(current: str, bump_type: str) -> str:
    """Bump the version according to semver rules."""
    match = SEMVER_PATTERN.match(current)
    if not match:
        raise ValueError(f"Invalid version format: {current}")

    major, minor, patch = int(match.group(1)), int(match.group(2)), int(match.group(3))

    if bump_type == "major":
        return f"v{major + 1}.0.0"
    elif bump_type == "minor":
        return f"v{major}.{minor + 1}.0"
    elif bump_type == "patch":
        return f"v{major}.{minor}.{patch + 1}"
    else:
        raise ValueError(f"Invalid bump type: {bump_type}. Use major, minor, or patch")


def main():
    bump_type = sys.argv[1] if len(sys.argv) > 1 else "patch"

    if bump_type not in ("major", "minor", "patch"):
        print(f"Error: Invalid bump type '{bump_type}'. Use major, minor, or patch", file=sys.stderr)
        sys.exit(1)

    current = get_latest_tag()
    next_version = bump_version(current, bump_type)
    print(next_version)


if __name__ == "__main__":
    main()

