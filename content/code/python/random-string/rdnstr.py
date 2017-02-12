#!/usr/bin/env python
# -*- coding:utf-8 -*-

import random

ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"

def random_string(strlen):
  result = ""
  for i in range(strlen):
    result += random.choice(ALPHABET)

  return result


if __name__ == '__main__':
  # for test purpose
  random.seed()
  print(random_string(5))
  print(random_string(6))
  print(random_string(7))
