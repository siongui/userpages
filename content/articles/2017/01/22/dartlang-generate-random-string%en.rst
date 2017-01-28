[Dart] Generate Random String From [a-z0-9]
###########################################

:date: 2017-01-22T00:18+08:00
:tags: Dart, String Manipulation, Random Number
:category: Dart
:summary: Generate a random string from [a-z0-9] in Dart_.
:adsu: yes


.. code-block:: dart

  import 'dart:core';
  import 'dart:math';

  const chars = "abcdefghijklmnopqrstuvwxyz0123456789";

  String RandomString(int strlen) {
    Random rnd = new Random(new DateTime.now().millisecondsSinceEpoch);
    String result = "";
    for (var i = 0; i < strlen; i++) {
      result += chars[rnd.nextInt(chars.length)];
    }
    return result;
  }

  void main() {
    print(RandomString(10));
    print(RandomString(20));
  }

.. adsu:: 2

Run `code <https://dartpad.dartlang.org/d6ed1551b254f04e41161abed5dae486>`_ on DartPad_:

.. raw:: html

  <iframe src='https://dartpad.dartlang.org/embed-dart.html?id=d6ed1551b254f04e41161abed5dae486' style='height:300px;width:100%;' frameborder='0'></iframe>

----

Tested on: DartPad_.

----

.. adsu:: 3

References:

.. [1] `dart random number - Google search <https://www.google.com/search?q=dart+random+number>`_

       `dart random number - DuckDuckGo search <https://duckduckgo.com/?q=dart+random+number>`_

       `dart random number - Bing search <https://www.bing.com/search?q=dart+random+number>`_

       `dart random number - Yahoo search <https://search.yahoo.com/search?p=dart+random+number>`_

       `dart random number - Baidu search <https://www.baidu.com/s?wd=dart+random+number>`_

       `dart random number - Yandex search <https://www.yandex.com/search/?text=dart+random+number>`_

.. [2] `dart string length - Google search <https://www.google.com/search?q=dart+string+length>`_

       `dart string length - DuckDuckGo search <https://duckduckgo.com/?q=dart+string+length>`_

       `dart string length - Bing search <https://www.bing.com/search?q=dart+string+length>`_

       `dart string length - Yahoo search <https://search.yahoo.com/search?p=dart+string+length>`_

       `dart string length - Baidu search <https://www.baidu.com/s?wd=dart+string+length>`_

       `dart string length - Yandex search <https://www.yandex.com/search/?text=dart+string+length>`_

.. [3] `Language Tour | Dart <https://www.dartlang.org/guides/language/language-tour>`_


.. _Dart: https://www.dartlang.org/
.. _DartPad: https://dartpad.dartlang.org/
