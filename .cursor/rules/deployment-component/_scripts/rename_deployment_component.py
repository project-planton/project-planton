#!/usr/bin/env python3
"""
Deployment Component Rename Script

Renames a deployment component across the entire codebase:
- Copies component directory with new name
- Applies comprehensive find-replace patterns
- Updates cloud_resource_kind.proto registry
- Updates documentation
- Runs build pipeline (protos, build, test)

Usage:
  python3 .cursor/rules/deployment-component/_scripts/rename_deployment_component.py \
    --old-name KubernetesMicroservice \
    --new-name KubernetesDeployment \
    --new-id-prefix k8sdpl

  # Keep existing id_prefix
  python3 .cursor/rules/deployment-component/_scripts/rename_deployment_component.py \
    --old-name KubernetesMicroservice \
    --new-name KubernetesDeployment

Output: JSON object with success status, metrics, and build results
"""

import argparse
import json
import os
import re
import shutil
import subprocess
import sys
import time
from pathlib import Path
from typing import Dict, List, Tuple, Optional


def find_repo_root(start_dir: str) -> str:
    """Find repository root by looking for .git or go.mod"""
    current = os.path.abspath(start_dir)
    while True:
        if os.path.isdir(os.path.join(current, ".git")) or os.path.isfile(os.path.join(current, "go.mod")):
            return current
        parent = os.path.dirname(current)
        if parent == current:
            raise RuntimeError("Could not find repository root")
        current = parent


def to_lowercase(name: str) -> str:
    """Convert PascalCase to lowercase (KubernetesMicroservice -> kubernetesmicroservice)"""
    return name.lower()


def to_snake_case(name: str) -> str:
    """Convert PascalCase to snake_case (KubernetesMicroservice -> kubernetes_microservice)"""
    # Insert underscore before capital letters
    s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
    return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()


def to_kebab_case(name: str) -> str:
    """Convert PascalCase to kebab-case (KubernetesMicroservice -> kubernetes-microservice)"""
    return to_snake_case(name).replace('_', '-')


def to_space_separated(name: str) -> str:
    """Convert PascalCase to space separated (KubernetesMicroservice -> kubernetes microservice)"""
    return to_snake_case(name).replace('_', ' ')


def to_camel_case(name: str) -> str:
    """Convert PascalCase to camelCase (KubernetesMicroservice -> kubernetesMicroservice)"""
    if not name:
        return name
    return name[0].lower() + name[1:]


def to_upper_snake_case(name: str) -> str:
    """Convert PascalCase to UPPER_SNAKE_CASE (KubernetesMicroservice -> KUBERNETES_MICROSERVICE)"""
    return to_snake_case(name).upper()


def build_replacement_map(old_name: str, new_name: str) -> List[Tuple[str, str]]:
    """
    Build comprehensive replacement patterns for all naming conventions.
    Returns list of (old, new) tuples, ordered by specificity (most specific first).
    """
    patterns = [
        # PascalCase (most specific)
        (old_name, new_name),
        # camelCase
        (to_camel_case(old_name), to_camel_case(new_name)),
        # UPPER_SNAKE_CASE
        (to_upper_snake_case(old_name), to_upper_snake_case(new_name)),
        # snake_case
        (to_snake_case(old_name), to_snake_case(new_name)),
        # kebab-case
        (to_kebab_case(old_name), to_kebab_case(new_name)),
        # space separated (with quotes)
        (f'"{to_space_separated(old_name)}"', f'"{to_space_separated(new_name)}"'),
        # lowercase (least specific, should be last)
        (to_lowercase(old_name), to_lowercase(new_name)),
    ]
    
    # Remove duplicates while preserving order
    seen = set()
    unique_patterns = []
    for old, new in patterns:
        if old not in seen:
            seen.add(old)
            unique_patterns.append((old, new))
    
    return unique_patterns


def find_component_in_registry(repo_root: str, component_name: str) -> Optional[Dict]:
    """
    Find component in cloud_resource_kind.proto.
    Returns dict with: enum_value, provider, id_prefix, and the full enum entry text.
    """
    registry_path = os.path.join(
        repo_root,
        "apis/org/project_planton/shared/cloudresourcekind/cloud_resource_kind.proto"
    )
    
    if not os.path.isfile(registry_path):
        return None
    
    with open(registry_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Find enum entry pattern: ComponentName = NUMBER [(kind_meta) = { ... }];
    pattern = rf'^  {re.escape(component_name)}\s*=\s*(\d+)\s*\[.*?\];'
    match = re.search(pattern, content, re.MULTILINE | re.DOTALL)
    
    if not match:
        return None
    
    enum_value = int(match.group(1))
    
    # Extract metadata
    entry_text = match.group(0)
    
    # Extract provider
    provider_match = re.search(r'provider:\s*(\w+)', entry_text)
    provider = provider_match.group(1) if provider_match else None
    
    # Extract id_prefix
    id_prefix_match = re.search(r'id_prefix:\s*"([^"]+)"', entry_text)
    id_prefix = id_prefix_match.group(1) if id_prefix_match else None
    
    return {
        'enum_value': enum_value,
        'provider': provider,
        'id_prefix': id_prefix,
        'entry_text': entry_text,
        'registry_path': registry_path
    }


def get_component_directory(repo_root: str, provider: str, component_folder: str) -> str:
    """Get path to component directory"""
    return os.path.join(
        repo_root,
        "apis/org/project_planton/provider",
        provider,
        component_folder
    )


def delete_directory_if_exists(path: str) -> bool:
    """Delete directory if it exists. Returns True if deleted, False if didn't exist."""
    if os.path.exists(path):
        shutil.rmtree(path)
        return True
    return False


def copy_component_directory(src: str, dst: str) -> int:
    """Copy component directory. Returns number of files copied."""
    if not os.path.exists(src):
        raise RuntimeError(f"Source directory does not exist: {src}")
    
    shutil.copytree(src, dst)
    
    # Count files
    file_count = 0
    for root, dirs, files in os.walk(dst):
        file_count += len(files)
    
    return file_count


def apply_replacements_in_file(file_path: str, replacements: List[Tuple[str, str]]) -> int:
    """
    Apply all replacement patterns to a file.
    Returns number of replacements made.
    """
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
    except (UnicodeDecodeError, PermissionError):
        # Skip binary files or files we can't read
        return 0
    
    original_content = content
    replacement_count = 0
    
    for old, new in replacements:
        if old in content:
            content = content.replace(old, new)
            replacement_count += 1
    
    if content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
    
    return replacement_count


def apply_replacements_in_directory(directory: str, replacements: List[Tuple[str, str]]) -> int:
    """
    Apply replacements to all files in directory recursively.
    Returns total number of replacements made.
    """
    total_replacements = 0
    
    for root, dirs, files in os.walk(directory):
        for file in files:
            file_path = os.path.join(root, file)
            count = apply_replacements_in_file(file_path, replacements)
            total_replacements += count
    
    return total_replacements


def update_registry_entry(
    registry_path: str,
    old_name: str,
    new_name: str,
    new_id_prefix: Optional[str]
) -> None:
    """
    Update component entry in cloud_resource_kind.proto.
    Only updates the enum name and optionally the id_prefix.
    Preserves all other metadata.
    """
    with open(registry_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Find the enum entry
    pattern = rf'^  {re.escape(old_name)}\s*=\s*(\d+)\s*\[(.*?)\];'
    match = re.search(pattern, content, re.MULTILINE | re.DOTALL)
    
    if not match:
        raise RuntimeError(f"Could not find enum entry for {old_name}")
    
    enum_value = match.group(1)
    metadata = match.group(2)
    
    # Update id_prefix if provided
    if new_id_prefix:
        metadata = re.sub(
            r'id_prefix:\s*"[^"]+"',
            f'id_prefix: "{new_id_prefix}"',
            metadata
        )
    
    # Build new entry
    new_entry = f'  {new_name} = {enum_value} [{metadata}];'
    
    # Replace in content
    content = re.sub(
        rf'^  {re.escape(old_name)}\s*=\s*{enum_value}\s*\[.*?\];',
        new_entry,
        content,
        flags=re.MULTILINE | re.DOTALL
    )
    
    with open(registry_path, 'w', encoding='utf-8') as f:
        f.write(content)


def rename_icon_folder(repo_root: str, provider: str, old_folder: str, new_folder: str) -> Dict:
    """
    Rename icon folder if it exists.
    Returns dict with: exists, old_path, new_path, renamed
    """
    # Map provider to icon directory structure
    # For kubernetes/workload or kubernetes/addon, use just "kubernetes"
    icon_provider = provider.split('/')[0]
    
    old_icon_path = os.path.join(
        repo_root,
        "site/public/images/providers",
        icon_provider,
        old_folder
    )
    
    new_icon_path = os.path.join(
        repo_root,
        "site/public/images/providers", 
        icon_provider,
        new_folder
    )
    
    result = {
        'exists': os.path.exists(old_icon_path),
        'old_path': old_icon_path,
        'new_path': new_icon_path,
        'renamed': False
    }
    
    if result['exists']:
        # Delete target if exists
        if os.path.exists(new_icon_path):
            shutil.rmtree(new_icon_path)
        # Rename
        shutil.move(old_icon_path, new_icon_path)
        result['renamed'] = True
    
    return result


def run_command(cmd: List[str], cwd: str) -> Tuple[int, str, str]:
    """Run a command and return (exit_code, stdout, stderr)"""
    try:
        result = subprocess.run(
            cmd,
            cwd=cwd,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            check=False
        )
        return result.returncode, result.stdout, result.stderr
    except Exception as e:
        return 127, "", str(e)


def run_build_pipeline(repo_root: str) -> Dict:
    """
    Run the build pipeline: make protos, make build, make test.
    Returns dict with exit codes and output for each phase.
    """
    result = {
        'protos_exit_code': 0,
        'protos_stdout': '',
        'protos_stderr': '',
        'build_exit_code': 0,
        'build_stdout': '',
        'build_stderr': '',
        'test_exit_code': 0,
        'test_stdout': '',
        'test_stderr': '',
    }
    
    # Run make protos
    exit_code, stdout, stderr = run_command(['make', 'protos'], repo_root)
    result['protos_exit_code'] = exit_code
    result['protos_stdout'] = stdout
    result['protos_stderr'] = stderr
    
    if exit_code != 0:
        return result  # Stop on first failure
    
    # Run make build
    exit_code, stdout, stderr = run_command(['make', 'build'], repo_root)
    result['build_exit_code'] = exit_code
    result['build_stdout'] = stdout
    result['build_stderr'] = stderr
    
    if exit_code != 0:
        return result  # Stop on first failure
    
    # Run make test
    exit_code, stdout, stderr = run_command(['make', 'test'], repo_root)
    result['test_exit_code'] = exit_code
    result['test_stdout'] = stdout
    result['test_stderr'] = stderr
    
    return result


def main() -> int:
    parser = argparse.ArgumentParser(
        description='Rename a deployment component across the entire codebase'
    )
    parser.add_argument(
        '--old-name',
        required=True,
        help='Old component name in PascalCase (e.g., KubernetesMicroservice)'
    )
    parser.add_argument(
        '--new-name',
        required=True,
        help='New component name in PascalCase (e.g., KubernetesDeployment)'
    )
    parser.add_argument(
        '--new-id-prefix',
        help='New ID prefix (e.g., k8sdpl). If not provided, keeps existing.'
    )
    
    args = parser.parse_args()
    
    start_time = time.time()
    
    result = {
        'success': False,
        'old_component': args.old_name,
        'new_component': args.new_name,
        'old_folder': to_lowercase(args.old_name),
        'new_folder': to_lowercase(args.new_name),
        'old_id_prefix': None,
        'new_id_prefix': args.new_id_prefix,
        'enum_value': None,
        'provider': None,
        'files_modified': 0,
        'replacements_made': 0,
        'icon_folder_renamed': False,
        'protos_exit_code': 0,
        'build_exit_code': 0,
        'test_exit_code': 0,
        'duration_seconds': 0,
        'error': None
    }
    
    try:
        # Find repo root
        repo_root = os.environ.get('REPO_ROOT', find_repo_root(os.getcwd()))
        
        # 1. Validate: Find old component in registry
        component_info = find_component_in_registry(repo_root, args.old_name)
        if not component_info:
            result['error'] = f"Component {args.old_name} not found in cloud_resource_kind.proto"
            print(json.dumps(result), file=sys.stderr)
            return 1
        
        result['enum_value'] = component_info['enum_value']
        result['provider'] = component_info['provider']
        result['old_id_prefix'] = component_info['id_prefix']
        
        # Determine provider (handle kubernetes subdirectories)
        provider = component_info['provider']
        if provider == 'kubernetes':
            # Check if it's in workload or addon
            old_folder = to_lowercase(args.old_name)
            workload_path = os.path.join(repo_root, "apis/org/project_planton/provider/kubernetes/workload", old_folder)
            addon_path = os.path.join(repo_root, "apis/org/project_planton/provider/kubernetes/addon", old_folder)
            
            if os.path.exists(workload_path):
                provider = "kubernetes/workload"
            elif os.path.exists(addon_path):
                provider = "kubernetes/addon"
            # else keep as just "kubernetes"
        
        # 2. Get directory paths
        old_folder = to_lowercase(args.old_name)
        new_folder = to_lowercase(args.new_name)
        
        old_dir = get_component_directory(repo_root, provider, old_folder)
        new_dir = get_component_directory(repo_root, provider, new_folder)
        
        # Validate old directory exists
        if not os.path.exists(old_dir):
            result['error'] = f"Old component directory does not exist: {old_dir}"
            print(json.dumps(result), file=sys.stderr)
            return 1
        
        # Check if new component already exists in registry
        new_component_info = find_component_in_registry(repo_root, args.new_name)
        if new_component_info:
            result['error'] = f"Component {args.new_name} already exists in cloud_resource_kind.proto"
            print(json.dumps(result), file=sys.stderr)
            return 1
        
        # 3. Delete target directory if it exists
        if delete_directory_if_exists(new_dir):
            print(f"Deleted existing target directory: {new_dir}", file=sys.stderr)
        
        # 4. Copy component directory
        file_count = copy_component_directory(old_dir, new_dir)
        result['files_modified'] = file_count
        
        # 5. Build replacement map
        replacements = build_replacement_map(args.old_name, args.new_name)
        
        # 6. Apply replacements in new component directory
        replacements_in_component = apply_replacements_in_directory(new_dir, replacements)
        result['replacements_made'] += replacements_in_component
        
        # 7. Apply replacements in docs directory
        docs_dir = os.path.join(repo_root, "site/public/docs")
        if os.path.exists(docs_dir):
            replacements_in_docs = apply_replacements_in_directory(docs_dir, replacements)
            result['replacements_made'] += replacements_in_docs
        
        # 8. Update cloud_resource_kind.proto
        update_registry_entry(
            component_info['registry_path'],
            args.old_name,
            args.new_name,
            args.new_id_prefix
        )
        
        # 9. Rename icon folder (if exists)
        icon_result = rename_icon_folder(repo_root, provider, old_folder, new_folder)
        result['icon_folder_renamed'] = icon_result['renamed']
        
        if icon_result['exists']:
            print(f"Renamed icon folder: {icon_result['old_path']} -> {icon_result['new_path']}", file=sys.stderr)
        else:
            print(f"Icon folder not found (skipped): {icon_result['old_path']}", file=sys.stderr)
        
        # 10. Delete old component directory
        shutil.rmtree(old_dir)
        
        # 11. Run build pipeline
        build_results = run_build_pipeline(repo_root)
        result.update(build_results)
        
        # Check if all phases passed
        if (build_results['protos_exit_code'] == 0 and
            build_results['build_exit_code'] == 0 and
            build_results['test_exit_code'] == 0):
            result['success'] = True
        else:
            result['error'] = "Build pipeline failed"
        
    except Exception as e:
        result['error'] = str(e)
        print(json.dumps(result), file=sys.stderr)
        return 1
    finally:
        result['duration_seconds'] = round(time.time() - start_time, 2)
    
    print(json.dumps(result, indent=2))
    return 0 if result['success'] else 1


if __name__ == '__main__':
    sys.exit(main())

