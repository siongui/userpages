#!/bin/bash
# wget webpages from http://www.aia.or.th/prayerXX.htm
# Google search: bash print number with leading zeros
# run this script with the following command on Ubuntu 15.10
# $ bash wget-webpages.sh
# the following command is not working:
# $ sh wget-webpages.sh

for i in {01..99}
do
  wget "http://www.aia.or.th/prayer$i.htm" -P prayer
  sleep .3
done
