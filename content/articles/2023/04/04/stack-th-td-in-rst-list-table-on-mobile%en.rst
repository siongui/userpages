[CSS] Stack td and th in reStructuredText list table on Mobile Screen
#####################################################################

:date: 2023-04-04T22:53+07:00
:tags: reStructuredText, CSS, Responsive Web Design, SCSS
:category: CSS
:summary: Stack *td* and *th* elements in reStructuredText list table on small
          screen.
:og_image: https://www.maxpixel.net/static/photo/1x/Table-List-Personal-Data-Data-Names-Column-428312.jpg
:adsu: yes


Stack *td* and *th* elements in `reStructuredText list table`_ on mobile devices
with small screen. [1]_ [2]_ [3]_

**Demo**: `Bhikkhu Pāṭimokkha <https://siongui.github.io/pali-chanting/bhikkhu-patimokkha/>`_

**reStructuredText code**:

.. code-block:: rst

   .. list-table:: Pubba-kicca
      :header-rows: 1
      :class: stack-th-td-on-mobile
      :widths: auto

      * - Pāli
        - English
        - Chinese

      * - Okāsaṃ me bhante thero detu, pāṭimokkhaṃ uddesituṃ
        - May the senior monk give me the opportunity to recite the Pāṭimokkha
        - 願長老給我機會背誦波羅提木叉

**SCSS code**:

.. code-block:: scss

   $screen-xs-max: 767;
   $smartphone: "(max-width: #{$screen-xs-max}px)";

   // Google search: stack td on mobile
   // Stack two <td> over each other
   // https://stackoverflow.com/a/15571109
   .stack-th-td-on-mobile {
     @media #{$smartphone} {
       th,td {
         display: block;
         clear: both;
       }
     }
   }

----

**Bonus** [4]_ [5]_ [6]_

Add background color for th and td elements.

.. code-block:: scss

   $screen-xs-max: 767;
   $smartphone: "(max-width: #{$screen-xs-max}px)";

   // Google search: stack td on mobile
   // Stack two <td> over each other
   // https://stackoverflow.com/a/15571109
   .stack-th-td-on-mobile {
     @media #{$smartphone} {
       th,td {
         display: block;
         clear: both;
       }
     }

     // https://stackoverflow.com/a/28186676
     // https://www.w3schools.com/cssref/sel_nth-child.php
     tr > td:first-child,th:first-child {
       background-color: None;
     }
     tr > td:nth-child(2),th:nth-child(2) {
       background-color: Azure;
     }
     tr > td:last-child,th:last-child {
       background-color: Cornsilk;
     }
   }

----

Tested on: Python ``3.10.6``, ``Pelican 4.8.0``.

----

References:

.. [1] | `stack td on mobile - Google search <https://www.google.com/search?q=stack+td+on+mobile>`_
       | `stack td on mobile - DuckDuckGo search <https://duckduckgo.com/?q=stack+td+on+mobile>`_
       | `stack td on mobile - Ecosia search <https://www.ecosia.org/search?q=stack+td+on+mobile>`_
       | `stack td on mobile - Qwant search <https://www.qwant.com/?q=stack+td+on+mobile>`_
       | `stack td on mobile - Bing search <https://www.bing.com/search?q=stack+td+on+mobile>`_
       | `stack td on mobile - Yahoo search <https://search.yahoo.com/search?p=stack+td+on+mobile>`_
       | `stack td on mobile - Baidu search <https://www.baidu.com/s?wd=stack+td+on+mobile>`_
       | `stack td on mobile - Yandex search <https://www.yandex.com/search/?text=stack+td+on+mobile>`_

.. [2] `html - Stack two <td> over each other - Stack Overflow <https://stackoverflow.com/a/15571109>`_

.. [3] `stack th td on mobile · siongui/pali-chanting@6b9c4ca <https://github.com/siongui/pali-chanting/commit/6b9c4ca4ddefa316a9ebee3cc00ade1f4994d6c6>`_

.. [4] `add td background color · siongui/pali-chanting@97a981a <https://github.com/siongui/pali-chanting/commit/97a981a60488b73420a200d8ecea31559a3e0b3f>`_

.. [5] `html - How to set background color of td in CSS - Stack Overflow <https://stackoverflow.com/a/28186676>`_

.. [6] `CSS :nth-child() Selector <https://www.w3schools.com/cssref/sel_nth-child.php>`_

.. _reStructuredText list table: https://docutils.sourceforge.io/docs/ref/rst/directives.html#list-table
