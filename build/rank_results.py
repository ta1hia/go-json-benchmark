import os
from string import Template
from collections import defaultdict

BASE_PATH = os.path.dirname(os.path.abspath(__file__))
with open(os.path.join(BASE_PATH, 'results.out')) as f:
    d = f.read()

lines = d.split('\n')
lines = lines[3:-3]

marshal_avgs = defaultdict(int)
unmarshal_avgs = defaultdict(int)
for line in lines:
    parts = line.strip('Benchmark').split()
    lib = parts[0].split('-')[0].split('By')[-1].lower()
    if parts[0].startswith('Marshal'):
        marshal_avgs[lib] += int(parts[2])
    else:
        unmarshal_avgs[lib] += int(parts[2])

marshal = []
rank = 1
for lib, speed in sorted(marshal_avgs.items(), key=lambda kv:kv[1]):
    marshal.append('{} | {} | {} ns/op'.format(rank,lib,int(speed/3)))
    rank += 1

unmarshal = []
rank = 1
for lib, speed in sorted(unmarshal_avgs.items(), key=lambda kv:kv[1]):
    unmarshal.append('{} | {} | {} ns/op'.format(rank,lib,int(speed/3)))
    rank += 1

with open(os.path.join(BASE_PATH, 'README.md.template')) as f:
    readme = f.read()
    template = Template(readme)

with open(os.path.join(BASE_PATH, 'README.md'), 'w') as f:
    f.write(template.substitute(MarshalChart='\n'.join(marshal), UnmarshalChart='\n'.join(unmarshal)))

os.replace(os.path.join(BASE_PATH, 'README.md'), os.path.join(BASE_PATH, os.pardir, 'README.md'))
