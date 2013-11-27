Python v.s. Dart: Syntax Side by Side
##########################################################################

:date: 2013-11-27 12:23
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

     - .. code-block:: dart

         // Use literals to create a list
         List mylist = [];

   * - .. code-block:: python

         # create a dictionary
         myHello = { 'hello': 'world' }

     - .. code-block:: dart

         // Map literal
         Map myHello = { 'hello': 'world' };

   * - .. code-block:: python

         # execution entry
         if __name__ == '__main__':
           print('program starts here')

     - .. code-block:: dart

         // execution entry
         main() {
           print('program starts here');
         }




Side-by-side format inspired by:

  `CoffeeScript <http://coffeescript.org/>`_

  `Programming Languages - Hyperpolyglot <http://hyperpolyglot.org/>`_

