#!/usr/bin/env python3

import shutil
from typing import Optional
import doit

DOIT_CONFIG = {
        'action_string_formatting': 'both',
        'backend': 'sqlite3',
        'default_tasks': ['test', 'compile', 'install']
        }

def task_uninstall():
    path: Optional[str] = shutil.which('xkcd936')
    match path:
        case None:
            return {'actions': []}
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
            'actions': ['cp bin/xkcd936 {prefix}/xkcd936'],
            'params':[{'name':'prefix',
                       'default': '/usr/local/bin',
                       'short': 'p',
                       'long':'prefix',
                       'help': 'install prefix'}],
            'verbosity': 2
            }

def task_clear():
    return {'actions': ['rm -rf bin/xkcd936']} 

if __name__ == '__main__':
    doit.run(globals())
