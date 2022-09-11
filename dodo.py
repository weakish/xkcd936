import shutil
from typing import Optional

DOIT_CONFIG = {
        'action_string_formatting': 'both',
        'default_tasks': ['test', 'compile', 'install']
        }

def task_uninstall():
    path: Optional[str] = shutil.which('xkcd936')
    match path:
        case None:
            return None
        case _:
            return {'actions': [f'rm -f {path}']}

def task_test():
    return {'actions': ['go test ./xkcd936']}

def task_compile():
    return {
            'file_dep': ['main.go', 'xkcd936/xkcd936.go'],
            'actions': ['go build -o {targets}'],
            'targets': ['bin/xkcd936'],
            }

def task_install():
    return {
            'file_dep': ['bin/xkcd936'],
            'actions': ['cp {changed} {targets}'],
            'targets': ['/usr/local/bin/xkcd936'],
            }

def task_clear():
    return {'actions': ['rm -rf bin/xkcd936']} 


