#!/usr/bin/env python
# -*- coding: utf-8 -*-

# pylint: disable=missing-docstring

import sys

from csv import reader

from jinja2 import Template


TEMPLATE = '''
  - id: {{ id }}
    rank: {{ rank }}
    suit: {{ suit }}
    title: {{ title }}
    meaning:
      general:
        keywords: >-
          {{ general_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ general_description | wordwrap(width=71, wrapstring='\n          ') }}
      longterm:
        keywords: >-
          {{ longterm_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ longterm_description | wordwrap(width=71, wrapstring='\n          ')}}
      mercury:
        keywords: >-
          {{ mercury_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ mercury_description | wordwrap(width=71, wrapstring='\n          ') }}
      venus:
        keywords: >-
          {{ venus_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ venus_description | wordwrap(width=71, wrapstring='\n          ') }}
      mars:
        keywords: >-
          {{ mars_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ mars_description | wordwrap(width=71, wrapstring='\n          ') }}
      jupiter:
        keywords: >-
          {{ jupiter_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ jupiter_description | wordwrap(width=71, wrapstring='\n          ') }}
      saturn:
        keywords: >-
          {{ saturn_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ saturn_description | wordwrap(width=71, wrapstring='\n          ') }}
      uranus:
        keywords: >-
          {{ uranus_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ uranus_description | wordwrap(width=71, wrapstring='\n          ') }}
      neptune:
        keywords: >-
          {{ neptune_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ neptune_description | wordwrap(width=71, wrapstring='\n          ') }}
      pluto:
        keywords: >-
          {{ pluto_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ pluto_description | wordwrap(width=71, wrapstring='\n          ') }}
      result:
        keywords: >-
          {{ result_keywords | wordwrap(width=71, wrapstring='\n          ') }}
        description: >-
          {{ result_description | wordwrap(width=71, wrapstring='\n          ') }}

'''

#  [ 0] id
#  [ 1] karta
#  [ 2] vkratce
#  [ 3] gruppe
# [ 4] kugaenko
# [ 5] inglish
#  [ 6] nazvan
#  [ 7] dolgo1
#  [ 8] dolgo2
# [ 9] pluton1
# [10] pluton2
# [11] rezul1
# [12] rezul2
# [13] merkur1
# [14] merkur2
# [15] venera1
# [16] venera2
# [17] mars1
# [18] mars2
# [19] upiter1
# [20] upiter2
# [21] saturn1
# [22] saturn2
# [23] uran1
# [24] uran2
# [25] neptun1
# [26] neptun2


def main():
    filepath = sys.argv[1]

    with open(filepath, 'r') as csv_file:
        csv_reader = reader(csv_file.readlines(), delimiter='|')

    tmpl = Template(TEMPLATE)

    idx = 1
    print('---')
    print('deck:')
    for row in csv_reader:
        card = row[1].decode(encoding='utf-8')
        values = dict(
            id=idx,
            rank=card[:-1],
            suit=card[-1],
            title=row[6]
                .decode(encoding='utf-8'),
            general_keywords=row[2]
                .decode(encoding='utf-8'),
            general_description=row[3]
                .decode(encoding='utf-8'),
            longterm_keywords=row[7]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            longterm_description=row[8]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            mercury_keywords=row[13]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            mercury_description=row[14]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            venus_keywords=row[15]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            venus_description=row[16]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            mars_keywords=row[17]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            mars_description=row[18]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            jupiter_keywords=row[19]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            jupiter_description=row[20]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            saturn_keywords=row[21]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            saturn_description=row[22]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            uranus_keywords=row[23]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            uranus_description=row[24]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            neptune_keywords=row[25]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            neptune_description=row[26]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            pluto_keywords=row[9]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            pluto_description=row[10]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
            result_keywords=row[11]
                .decode(encoding='utf-8')
                .replace(u'КЛЮЧЕВЫЕ СЛОВА: ', ''),
            result_description=row[12]
                .decode(encoding='utf-8')
                .replace(u'ЗНАЧЕНИЕ: ', ''),
        )
        print(tmpl.render(**values).encode(encoding='utf-8'))
        idx = idx + 1


if __name__ == "__main__":
    sys.exit(main())
