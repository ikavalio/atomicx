#!/usr/bin/env python3

import re
import os.path
import sys
import subprocess

pattern_fdef = re.compile('__asm__\("github_com_ikavalio_atomicx." "(\w+)"\)')

tmap = {
    'Int32': 'int32',
    'Int64': 'int64',
    'Uint32': 'uint32',
    'Uint64': 'uint64',
    'Pointer': 'unsafe.Pointer',
    'Uintptr': 'uintptr',
}

cfile = sys.argv[1]
ofile = sys.argv[2]

if os.path.isfile(cfile):
    out = ['package atomicx', '', 'import "unsafe"', '']
    p2 = subprocess.run(['cpp', cfile], stdout=subprocess.PIPE)
    for match in pattern_fdef.finditer(p2.stdout.decode('ascii')):
        fname = match.group(1)
        fparts = re.findall('[A-Z][^A-Z]*', fname)
        cmd = fparts[0]
        if cmd == 'Clear':
            out.append("//%s does atomic clear operation on the bool *addr. Should be used in conjunction with TestAndSet."  % fname)
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)
            out.append("func %s(addr *bool)" % fname)
        elif cmd == 'Test':
            out.append("//%s does atomic test-and-set operation on the bool *addr (atomically *addr = *addr ? *addr : true)." % fname)
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)
            out.append("func %s(addr *bool) bool" % fname)
        elif cmd == 'Nand' or cmd == 'Xor' or cmd == 'Or' or cmd == 'And' or cmd == 'Add' or cmd == 'Swap':
            if cmd == 'Swap':
                out.append("//%s atomically stores new into *addr and returns the previous *addr value." % fname)
            else:
                out.append("//%s applies %s to delta and *addr and returns the new value." % (fname, cmd.lower()))
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)

            argtype = tmap[fparts[1]]
            out.append("func %s(addr *%s, value %s) %s" % (fname, argtype, argtype, argtype))
        elif cmd == 'Store':
            out.append("//%s atomically stores val into *addr." % fname)
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)

            argtype = tmap[fparts[1]]
            out.append("func %s(addr *%s, val %s)" % (fname, argtype, argtype))
        elif cmd == 'Load':
            out.append("//%s atomically loads *addr." % fname)
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)

            argtype = tmap[fparts[1]]
            out.append("func %s(addr *%s) %s" % (fname, argtype, argtype))
        elif cmd == 'Compare':
            argtype = tmap[fparts[4]]
            out.append("//%s executes the compare-and-swap operation for an %s value." % (fname, argtype))
            out.append("//extern github_com_ikavalio_atomicx.%s" % fname)
            out.append("func %s(addr *%s, old, new %s) bool" % (fname, argtype, argtype))

        out.append('\n')

    with open(ofile, 'w') as f:
        f.write('\n'.join(out))
else:
    print('Not such file: ', cfile)
