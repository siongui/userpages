[Makefile] Convert Traditional Chinese PO file to Simplified Chinese via OpenCC
###############################################################################

:date: 2016-03-12T21:24+08:00
:tags: Bash, Makefile, Commandline, sed, String Manipulation, i18n, Locale,
       gettext, Conversion of Traditional and Simplified Chinese, OpenCC
:category: Bash
:summary: Convert Traditional Chinese PO file to Simplified Chinese via OpenCC_
          and sed_ in Makefile_.


Convert Traditional Chinese PO file to Simplified Chinese via OpenCC_ and sed_
in Makefile_.

.. code-block:: bash

  twpo2cn:
  	@echo "\033[92mCheck if OpenCC exists ...\033[0m"
  	@[ -x $(shell command -v opencc 2> /dev/null) ] || sudo apt-get install opencc
  	@echo "\033[92mCreating zh_CN PO from zh_TW PO ...\033[0m"
  	@[ -d $(LOCALE_DIR)/zh_CN/LC_MESSAGES/ ] || mkdir -p $(LOCALE_DIR)/zh_CN/LC_MESSAGES/
  	@opencc -c zht2zhs.ini -i $(LOCALE_DIR)/zh_TW/LC_MESSAGES/messages.po -o $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po
  	@sed 's/zh_TW/zh_CN/' -i $(LOCALE_DIR)/zh_CN/LC_MESSAGES/messages.po

----

Tested on ``Ubuntu Linux 15.10``, ``opencc 0.4.3-2build1``, `GNU make 4.0-8.2`_.

----

References:

.. [1] `makefile test command exists <https://www.google.com/search?q=makefile+test+command+exists>`_

       `make - Check if a program exists from a Makefile - Stack Overflow <http://stackoverflow.com/questions/5618615/check-if-a-program-exists-from-a-makefile>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _OpenCC: http://opencc.byvoid.com/
.. _GNU make 4.0-8.2: http://packages.ubuntu.com/wily/make
