#!/usr/bin/env python
# -*- coding: utf-8 -*-
from __future__ import division, print_function, absolute_import


from struct import unpack
from zlib import decompress
import re
filename = '/Library/Dictionaries/英辞郎.dictionary/Contents/Body.data'
filename = '/System/Library/Assets/com_apple_MobileAsset_DictionaryServices_dictionaryOSX/1a184f41ee27bbb203ba596a25c8ea104c13d98f.asset/AssetData/Sanseido The WISDOM English-Japanese Japanese-English Dictionary.dictionary/Contents/Resources/Body.data'
f = open(filename, 'rb')

def gen_entry():
    f.seek(0x40)
    limit = 0x40 + unpack('i', f.read(4))[0]
    f.seek(0x60)
    while f.tell()<limit:
        h = f.read(4)
        sz, = unpack('i', h)
        print(h)
        print(sz)
        buf = decompress(f.read(sz)[8:])
        print(buf)
        return
        for m in re.finditer(b'<d:entry[^\n]+', buf):
            entry = m.group().decode()
            title = re.search('d:title="(.*?)"', entry).group(1)
            yield title, entry

for word, definition in gen_entry():
    pass
#    print(word, definition)
