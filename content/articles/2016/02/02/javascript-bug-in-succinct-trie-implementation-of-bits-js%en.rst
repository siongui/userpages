[JavaScript] Bug in Succinct Trie Implementation of Bits.js
###########################################################

:date: 2016-02-02T20:22+08:00
:tags: JavaScript, Data Structure
:category: JavaScript
:summary: Fix the bug in `succinct trie`_ implementation of `Bits.js`_ - Wrong
          trie_ insertion if words are not inserted in alphabetical order.
:adsu: yes


`Bits.js`_ ([1]_) implements a `succinct trie`_ for fast lookup of words in a
large dictionary and does not take much space at the same time. It works great
if the words are inserted in the trie_ in alphabetical order. When I ported the
code from JavaScript to Go_ ([2]_), I found that if the words are not sorted are
then inserted into the trie in alphabetical order, the resulting trie is wrong.

The following is my patch for `Bits.js`_:

.. code-block:: javascript

  diff --git a/reference/Bits.js b/reference/Bits.js
  index d5a3934..129bcd2 100644
  --- a/reference/Bits.js
  +++ b/reference/Bits.js
  @@ -453,6 +453,18 @@ Trie.prototype = {
           var node = this.cache[ this.cache.length - 1 ];

           for( i = commonPrefix; i < word.length; i++ ) {
  +            // fix the bug if words not inserted in alphabetical order
  +            var isLetterExist = false;
  +            for ( var j = 0; j < node.children.length; j++ ) {
  +                if (node.children[j].letter == word[i]) {
  +                    this.cache.push(node.children[j]);
  +                    node = node.children[j];
  +                    isLetterExist = true;
  +                    break;
  +                }
  +            }
  +            if (isLetterExist) { continue; }
  +
               var next = new TrieNode( word[i] );
               this.nodeCount++;
               node.children.push( next );

----

References:

.. [1] `Succinct Data Structures: Cramming 80,000 words into a Javascript file. <http://stevehanov.ca/blog/?id=120>`_
          (`source code <http://www.hanovsolutions.com/trie/Bits.js>`__)

.. [2] `siongui/go-succinct-data-structure-trie: Succinct Trie in Go - GitHub <https://github.com/siongui/go-succinct-data-structure-trie>`_

.. [3] `patch for Bits.js - fix the bug of wrong trie insertion if words not inserted in alphabetical order · siongui/go-succinct-data-structure-trie@22c54c0 · GitHub <https://github.com/siongui/go-succinct-data-structure-trie/commit/22c54c040b59408c45039a55dcc1b9e5daff93eb>`_

.. [4] `fix the bug of wrong trie insertion if words not inserted in alphabetical order · siongui/go-succinct-data-structure-trie@aff195b · GitHub <https://github.com/siongui/go-succinct-data-structure-trie/commit/aff195ba0f4bcf48428b2beeafaf501588728d31>`_

.. [5] `make patch file <https://www.google.com/search?q=make+patch+file>`_

.. [6] `How to create a patch - MoodleDocs <https://docs.moodle.org/dev/How_to_create_a_patch>`_


.. _succinct trie: https://www.google.com/search?q=succinct+trie
.. _Bits.js: http://www.hanovsolutions.com/trie/Bits.js
.. _trie: https://www.google.com/search?q=trie
.. _Go: https://golang.org/
