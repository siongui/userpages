[JavaScript] Convert Text to Link in HTML
#########################################

:date: 2018-01-12T05:33+08:00
:tags: JavaScript, String Manipulation, Regular Expression
:category: JavaScript
:summary: Convert specific texts to clickable links (<a> tag) in HTML document
          via JavaScript.
:og_image: http://talkerscode.com/webtricks/images/convert_url.jpg
:adsu: yes


See demo first. (demo texts from the note in `SN.1.11`_)

| 　　「世尊；眾祐」，南傳作「世尊」
| 　　(bhagavā，音譯為「婆伽婆；婆伽梵；薄伽梵」，
| 　　義譯為「有幸者」，古譯為「尊祐」)，
| 　　菩提比丘長老英譯為「幸福者」(the Blessed One)。
| 　　請參看〈世尊譯詞的探討〉。

View page source of this web page, and you will find **世尊譯詞的探討** is plain
text in the page source. It is converted to a clickable link (<a> tag) via
JavaScript.

.. raw:: html

  <script>
  var textToUrlMapping = {
    "世尊譯詞的探討": "http://agama.buddhason.org/book/bb/bb21.htm",
  };

  function TextToLink(elm) {
    elm.innerHTML = elm.innerHTML.replace(/〈(.+)〉/g, function(text, str1) {
      if (textToUrlMapping.hasOwnProperty(str1)) {
        return '〈<a href="'+ textToUrlMapping[str1] +
               '" target="_blank">' + str1 +
               '</a>〉';
      }
      return str1;
    });
  }

  var t = document.querySelector(".line-block");
  TextToLink(t);
  </script>


The following is the source code of the demo.

**JavaScript**:

.. code-block:: javascript

  var textToUrlMapping = {
    "世尊譯詞的探討": "http://agama.buddhason.org/book/bb/bb21.htm",
  };

  function TextToLink(elm) {
    elm.innerHTML = elm.innerHTML.replace(/〈(.+)〉/g, function(text, str1) {
      if (textToUrlMapping.hasOwnProperty(str1)) {
        return '〈<a href="'+ textToUrlMapping[str1] +
               '" target="_blank">' + str1 +
               '</a>〉';
      }
      return str1;
    });
  }

  var t = document.querySelector(".line-block");
  TextToLink(t);

- Use querySelector_ to select the element which contains the texts.
- Call *TextToLink* function to convert the text in the element to link.
- In *TextToLink* function, extract the text in 〈〉 via the pattern
  **〈(.+)〉**.
- Get the URL from text-to-url mapping, and create the link by the text and url.

.. adsu:: 2

----

Tested on: ``Chromium 63.0.3239.84 on Ubuntu 17.10 (64-bit)``

----

**References**:

.. [1] | `javascript convert text to link - Google search <https://www.google.com/search?q=javascript+convert+text+to+link>`_
       | `javascript convert text to link - DuckDuckGo search <https://duckduckgo.com/?q=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Ecosia search <https://www.ecosia.org/search?q=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Qwant search <https://www.qwant.com/?q=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Bing search <https://www.bing.com/search?q=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Yahoo search <https://search.yahoo.com/search?p=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Baidu search <https://www.baidu.com/s?wd=javascript+convert+text+to+link>`_
       | `javascript convert text to link - Yandex search <https://www.yandex.com/search/?text=javascript+convert+text+to+link>`_

.. _querySelector: https://www.google.com/search?q=querySelector
.. _SN.1.11: http://agama.buddhason.org/SN/SN0011.htm
