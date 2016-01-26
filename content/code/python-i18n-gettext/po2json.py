#!/usr/bin/env python
# -*- coding:utf-8 -*-

import re
import json

def getPOPath(locale, domain, localeDir):
  return localeDir + "/" + locale + "/LC_MESSAGES/" + domain + ".po"

def extractFromPOFile(poPath):
  with open(poPath, 'r') as f:
    tuples = re.findall(r'msgid "(.+)"\nmsgstr "(.+)"', f.read())
  return tuples

def PO2JSON(locales, domain, localeDir):
  # create PO-like json data for i18n
  obj = {}
  for locale in locales:
    # English is default language
    if locale == "en_US": continue

    obj[locale] = {}
    tuples = extractFromPOFile( getPOPath(locale, domain, localeDir) )
    for tuple in tuples:
      obj[locale][tuple[0].decode('utf-8')] = tuple[1].decode('utf-8')
      #obj[locale][tuple[0]] = tuple[1]

  return json.dumps(obj)

if __name__ == '__main__':
  locales = ["zh_TW", "vi_VN"]
  domain = "messages"
  localeDir = "locale"
  print(PO2JSON(locales, domain, localeDir))
