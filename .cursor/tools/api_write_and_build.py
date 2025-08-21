#!/usr/bin/env python3
import argparse
import hashlib
import json
import os
import subprocess
import sys
from typing import List, Tuple


def find_repo_root(start_dir: str) -> str:
    current = os.path.abspath(start_dir)
    while True:
        if os.path.isdir(os.path.join(current, ".git")) or os.path.isfile(os.path.join(current, "go.mod")):
            return current
        parent = os.path.dirname(current)
        if parent == current:
            return start_dir
        current = parent


def api_path(repo_root: str, provider: str, kind_folder: str) -> Tuple[str, str]:
    rel = os.path.join("apis", "project", "planton", "provider", provider, kind_folder, "v1", "api.proto")
    return os.path.join(repo_root, rel), rel


def ensure_parent(path: str) -> List[str]:
    created = []
    parent = os.path.dirname(path)
    if not os.path.isdir(parent):
        os.makedirs(parent, exist_ok=True)
        created.append(parent)
    return created


def norm(seg: str) -> str:
    s = seg.strip().lower().replace("_", "")
    if ".." in s or s.startswith("/") or s.startswith("~"):
        raise ValueError("invalid segment")
    return s


def run_build(repo_root: str) -> Tuple[int, str, str]:
    apis_dir = os.path.join(repo_root, "apis")
    try:
        p = subprocess.run(["make", "-C", apis_dir, "build"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, check=False)
        return p.returncode, p.stdout, p.stderr
    except Exception as exc:
        return 127, "", str(exc)


def main() -> int:
    parser = argparse.ArgumentParser(description="Write api.proto and run apis build")
    parser.add_argument("--provider", required=True)
    parser.add_argument("--kindfolder", required=True)
    g = parser.add_mutually_exclusive_group(required=True)
    g.add_argument("--stdin", action="store_true")
    g.add_argument("--content-file")
    args = parser.parse_args()

    try:
        provider = norm(args.provider)
        kind = norm(args.kindfolder)
    except Exception as exc:
        print(json.dumps({"wrote": False, "build_succeeded": False, "error": f"invalid inputs: {exc}"}))
        return 2

    try:
        if args.stdin:
            content = sys.stdin.read()
        else:
            with open(os.path.abspath(args.content_file), "r", encoding="utf-8") as f:
                content = f.read()
    except Exception as exc:
        print(json.dumps({"wrote": False, "build_succeeded": False, "error": f"failed to read content: {exc}"}))
        return 3

    repo_root = os.environ.get("REPO_ROOT", find_repo_root(os.getcwd()))
    abs_path, rel_path = api_path(repo_root, provider, kind)

    result = {
        "wrote": False,
        "path": abs_path,
        "relative_path": rel_path,
        "bytes_written": 0,
        "created_dirs": [],
        "sha256": "",
        "build_succeeded": False,
        "build_exit_code": 0,
        "build_stdout": "",
        "build_stderr": "",
    }

    try:
        result["created_dirs"] = ensure_parent(abs_path)
        with open(abs_path, "w", encoding="utf-8", newline="\n") as f:
            f.write(content)
        result["bytes_written"] = len(content.encode("utf-8"))
        result["wrote"] = True
        result["sha256"] = hashlib.sha256(content.encode("utf-8")).hexdigest()

        code, out, err = run_build(repo_root)
        result["build_exit_code"] = code
        result["build_stdout"] = out
        result["build_stderr"] = err
        result["build_succeeded"] = code == 0
        print(json.dumps(result))
        return 0 if code == 0 else 4
    except Exception as exc:
        result["error"] = str(exc)
        print(json.dumps(result))
        return 1


if __name__ == "__main__":
    sys.exit(main())


