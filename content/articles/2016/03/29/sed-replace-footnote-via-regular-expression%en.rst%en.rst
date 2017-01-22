[sed] Replace Footnote by Regular Expression
############################################

:date: 2016-03-29T06:50+08:00
:tags: Bash, Commandline, sed, String Manipulation, Regular Expression,
       reStructuredText
:category: Bash
:summary: Replace footnote in reStructuredText_ format via `regular expression`_
          by sed_ stream editor.
:adsu: yes


Replace footnote in reStructuredText_ format via `regular expression`_ by sed_
stream editor.

Test file:

.. code-block:: txt

  南傳法句經 老品　JARAVAGGO 第　153　與　154　偈頌

  一五三、(註1-1)經多生輪迴，尋求造屋者(註1-2)，但未得見之，痛苦再再生。
  一五四、已見造屋者(註1-3)！不再造於屋(註1-4)。椽桷(註1-5)皆毀壞，棟梁(註1-6)亦摧折。我既證無為(註1-7)，一切愛盡滅。〔註一〕

  〔註1-1〕：以下二頌為釋迦牟尼佛在菩提樹下悟道的時候，心生歡喜， 自說此頌。後來又再阿難尊者的發問中而答以此頌。

  〔註1-2〕：指生死輪迴的原因。

  〔註1-3〕：喻情欲。

  〔註1-4〕：喻身體。

  〔註1-5〕：喻其他的一切煩惱欲。

  〔註1-6〕：喻無明。

  〔註1-7〕：即涅槃。

----

``(註1-1)`` => `` [1]\_ ``

.. code-block:: bash

  $ sed -r 's/\(註1-([0-9]{1})\)/ [\1]_ /g' test.txt

Output:

.. code-block:: txt

  南傳法句經 老品　JARAVAGGO 第　153　與　154　偈頌

  一五三、 [1]_ 經多生輪迴，尋求造屋者 [2]_ ，但未得見之，痛苦再再生。
  一五四、已見造屋者 [3]_ ！不再造於屋 [4]_ 。椽桷 [5]_ 皆毀壞，棟梁 [6]_ 亦摧折。我既證無為 [7]_ ，一切愛盡滅。〔註一〕

  〔註1-1〕：以下二頌為釋迦牟尼佛在菩提樹下悟道的時候，心生歡喜， 自說此頌。後來又再阿難尊者的發問中而答以此頌。

  〔註1-2〕：指生死輪迴的原因。

  〔註1-3〕：喻情欲。

  〔註1-4〕：喻身體。

  〔註1-5〕：喻其他的一切煩惱欲。

  〔註1-6〕：喻無明。

  〔註1-7〕：即涅槃。

----

``〔註1-1〕：`` => `` .. [1] ``

.. code-block:: bash

  $ sed -r 's/^〔註1-([0-9]{1})〕：/.. [\1] /g' test.txt

Output:

.. code-block:: txt

  南傳法句經 老品　JARAVAGGO 第　153　與　154　偈頌

  一五三、(註1-1)經多生輪迴，尋求造屋者(註1-2)，但未得見之，痛苦再再生。
  一五四、已見造屋者(註1-3)！不再造於屋(註1-4)。椽桷(註1-5)皆毀壞，棟梁(註1-6)亦摧折。我既證無為(註1-7)，一切愛盡滅。〔註一〕

  .. [1] 以下二頌為釋迦牟尼佛在菩提樹下悟道的時候，心生歡喜， 自說此頌。後來又再阿難尊者的發問中而答以此頌。

  .. [2] 指生死輪迴的原因。

  .. [3] 喻情欲。

  .. [4] 喻身體。

  .. [5] 喻其他的一切煩惱欲。

  .. [6] 喻無明。

  .. [7] 即涅槃。


----

Tested on ``Ubuntu Linux 15.10``, ``sed 4.2.2-6.1``.

----

References:

.. [1] `[Vim] Replace Footnote by Regular Expression <{filename}../28/vim-replace-footnote-via-regular-expression%en.rst>`_

.. [2] `[Python] Convert Footnote to reStructuredText Format <{filename}../../02/15/python-re-convert-to-rst-footnote%en.rst>`_

.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _regular expression: https://www.google.com.tw/search?q=regular+expression
.. _reStructuredText: https://www.google.com.tw/search?q=reStructuredText
