Python v.s. Dart: Syntax Side by Side
#####################################

:date: 2013-11-27 16:03
:tags: Dart, Python
:category: Dart
:author: Siong-Ui Te
:summary: Comparison of Python and Dart syntax, side by side for easy reference


.. list-table:: Syntax Comparison Table
   :header-rows: 1
   :class: table-syntax-diff

   * - Python
     - Dart

   * - .. code-block:: python

         # import class/function from package
         from sys import path

     - .. code-block:: dart

         // import a class from library
         import 'dart:math' show Random;

   * - .. code-block:: python

         # create a list
         mylist = []

       `read official doc » <http://docs.python.org/2/tutorial/introduction.html#lists>`__

     - .. code-block:: dart

         // Use literals to create a list
         List mylist = [];

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#lists>`__

   * - .. code-block:: python

         # create a dictionary
         myHello = { 'hello': 'world' }

       `read official doc » <http://docs.python.org/2/tutorial/datastructures.html#dictionaries>`__

     - .. code-block:: dart

         // Map literal
         Map myHello = { 'hello': 'world' };

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#maps>`__

   * - .. code-block:: python

         # execution entry
         if __name__ == '__main__':
           print('program starts here')

       `read official doc » <http://docs.python.org/2/library/__main__.html>`__

     - .. code-block:: dart

         // execution entry
         main() {
           print('program starts here');
         }

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#ch02-main>`__

   * - .. code-block:: python

         # function
         def printNumber(aNumber):
           print('%d' % aNumber)

       `read official doc » <http://docs.python.org/2/tutorial/controlflow.html#defining-functions>`__

     - .. code-block:: dart

         // function
         void printNumber(int aNumber) {
           print('$aNumber');
         }

         // shorthand syntax of function
         printNumber(int aNumber) => print('$aNumber');

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#functions>`__

   * - .. code-block:: python

         # raise exception
         raise IOError('io error occurs')

       `read official doc » <http://docs.python.org/2/tutorial/errors.html#raising-exceptions>`__

     - .. code-block:: dart

         // throwing, or raising, an exception
         throw new IOException('io error occurs');

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#exceptions>`__

   * - .. code-block:: python

         # catch IOError exception
         try:
           doSomething()
         except IOError:
           # handle the error here 
           # if doSomething raise IOError

       `read official doc » <http://docs.python.org/2/tutorial/errors.html#handling-exceptions>`__

     - .. code-block:: dart

         // catch IOException
         try {
           doSomething();
         } on IOException {
           // handle the exception 
           // if doSomething throw IOException
         }

       `read official doc » <https://www.dartlang.org/docs/dart-up-and-running/contents/ch02.html#ch02-exceptions-catch>`__



See Also:

  `Dart Development Using Vim <{filename}./dart-development-using-vim#en.rst>`_

Side-by-side format inspired by:

  `CoffeeScript <http://coffeescript.org/>`_

  `Programming Languages - Hyperpolyglot <http://hyperpolyglot.org/>`_

