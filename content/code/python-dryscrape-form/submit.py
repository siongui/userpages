#!/usr/bin/env python
# -*- coding:utf-8 -*-

import dryscrape

# make sure you have xvfb installed
dryscrape.start_xvfb()

root_url = 'YOUR_BASE_URL'

if __name__ == '__main__':
  # set up a web scraping session
  sess = dryscrape.Session(base_url = root_url)

  # we don't need images
  sess.set_attribute('auto_load_images', False)

  # visit webpage
  sess.visit('YOUR_RELATIVE_PATH_TO_BASE_URL')
  # search for iframe with id="mainframe"
  frame = sess.at_xpath('//*[@id="mainframe"]')

  # get the URL of iframe
  frameURL = root_url + frame['src']
  # visit the URL of iframe
  sess2 = dryscrape.Session()
  sess2.visit(frameURL)

  # fill in the form in iframe
  name = sess2.at_xpath('//*[@id="username"]')
  name.set("John")
  pid = sess2.at_xpath('//*[@id="person_id"]')
  pid.set("Q123446589")
  year = sess2.at_xpath('//*[@id="bornyear"]')
  year.set("2000")
  mobile = sess2.at_xpath('//*[@id="mobile"]')
  mobile.set("5631365976")

  # submit form
  name.form().submit()

  # save a screenshot of the web page
  sess2.render("test.png")
  print("Session rendered")
