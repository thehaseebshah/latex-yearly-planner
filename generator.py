import itertools
from dataclasses import dataclass
from typing import Collection


@dataclass(frozen=True)
class Config:
    name: str | None
    files: list[str]

    @staticmethod
    def merge(confs: Collection['Config']) -> 'Config':
        return Config(
            name='.'.join(c.name for c in confs if c.name),
            files=sum([c.files for c in confs], [])
        )

    @property
    def command(self):
        cfg = ','.join(f'cfg/{f}.yaml' for f in self.files)
        name = self.name
        return f'PLANNER_YEAR=2026 PASSES=1 CFG="{cfg}" NAME="release/{name}" ./single.sh'

configs = [
    [Config(None, ['base'])],
    [Config('breadcrumb', ['template_breadcrumb', 'custom_layout'])],
    [Config('lh', [])],
    [Config('dotted', [])],
    [Config('nodailycal', [])],
    [Config('24h', [])]
]

for element in itertools.product(*configs):
    print(Config.merge(element).command)