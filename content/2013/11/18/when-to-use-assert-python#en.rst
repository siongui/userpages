[Python] When to use assert
###########################

:modified: 2013-11-29 12:18
:tags: Python
:category: Python
:author: Siong-Ui Te
:summary: Steven D'Aprano explains when to use assert at Python mailing list.

Link:

  `When to use assert <https://mail.python.org/pipermail/python-list/2013-November/660401.html>`_
  (Simplified Chinese Translation `1 <http://www.oschina.net/translate/when-to-use-assert>`_,
  `2 <http://www.linuxeden.com/html/news/20131207/146183.html>`_,
  `3 <http://www.pythoner.cn/home/blog/when-to-use-assert/>`_)

The following is his whole post:

::

  The question of when to use the assert statement comes up occasionally, 
  usually in response to somebody misusing it, so I thought I'd write a 
  post describing when and why to use assertions, and when not to.

  For those who aren't aware of it, Python's "assert" checks a condition, 
  if it is true it does nothing, and if it is false it raises an 
  AssertionError with an optional error message. For example:

  py> x = 23
  py> assert x > 0, "x is not zero or negative"
  py> assert x%2 == 0, "x is not an even number"
  Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
  AssertionError: x is not an even number


  Many people use asserts as a quick and easy way to raise an exception if 
  an argument is given the wrong value. But this is wrong, dangerously 
  wrong, for two reasons. The first is that AssertionError is usually the 
  wrong error to give when testing function arguments. You wouldn't write 
  code like this:

  if not isinstance(x, int):
      raise AssertionError("not an int")

  you'd raise TypeError instead. "assert" raises the wrong sort of 
  exception.

  But, and more dangerously, there's a twist with assert: it can be 
  compiled away and never executed, if you run Python with the -O or -OO 
  optimization flags, and consequently there is no guarantee that assert 
  statements will actually be run. When using assert properly, this is a 
  feature, but when assert is used inappropriately, it leads to code that 
  is completely broken when running with the -O flag.

  When should use assert? In no particular order, assertions should be used 
  for:

  * defensive programming;
  * runtime checks on program logic;
  * checking contracts (e.g. pre-conditions and post-conditions);
  * program invariants; and 
  * checked documentation.

  (It's also acceptable to use assert when testing code, as a sort of quick-
  and-dirty poor man's unit testing, so long as you accept that the tests 
  simply won't do anything if you run with the -O flag. And I sometimes use 
  "assert False" in code to mark code branches that haven't been written 
  yet, and I want them to fail. Although "raise NotImplementedError" is 
  probably better for that, if a little more verbose.)

  Opinions on assertions vary, because they can be a statement of 
  confidence about the correctness of the code. If you're certain that the 
  code is correct, then assertions are pointless, since they will never 
  fail and you can safely remove them. If you're certain the checks can 
  fail (e.g. when testing input data provided by the user), then you dare 
  not use assert since it may be compiled away and then your checks will be 
  skipped.

  It's the situations in between those two that are interesting, times when 
  you're certain the code is correct but not *quite* absolutely certain. 
  Perhaps you've missed some odd corner case (we're all only human). In 
  this case an extra runtime check helps reassure you that any errors will 
  be caught as early as possible rather than in distant parts of the code.

  (This is why assert can be divisive. Since we vary in our confidence 
  about the correctness of code, one person's useful assert is another 
  person's useless runtime test.)

  Another good use for asserts is checking program invariants. An invariant 
  is some condition which you can rely on to be true unless a bug causes it 
  to become false. If there's a bug, better to find out as early as 
  possible, so we make a test for it, but we don't want to slow the code 
  down with such tests. Hence assert, which can be turned on in development 
  and off in production.

  An example of an invariant might be, if your function expects a database 
  connection to be open when it starts, and promises that it will still be 
  open when it returns, that's an invariant of the function:

  def some_function(arg):
      assert not DB.closed()
      ... # code goes here
      assert not DB.closed()
      return result


  Assertions also make good checked comments. Instead of writing a comment:

  # when we reach here, we know that n > 2

  you can ensure it is checked at runtime by turning it into an assert:

  assert n > 2

  Assertions are also a form of defensive programming. You're not 
  protecting against errors in the code as it is now, but protecting 
  against changes which introduce errors later. Ideally, unit tests will 
  pick those up, but let's face it, even when tests exist at all, they're 
  often incomplete. Build-bots can be down and nobody notices for weeks, or 
  people forget to run tests before committing code. Having an internal 
  check is another line of defence against errors sneaking in, especially 
  those which don't noisily fail but cause the code to malfunction and 
  return incorrect results.

  Suppose you have a series of if...elif blocks, where you know ahead of 
  time what values some variable is expected to have:

  # target is expected to be one of x, y, or z, and nothing else.
  if target == x:
      run_x_code()
  elif target == y:
      run_y_code()
  else:
      run_z_code()


  Assume that this code is completely correct now. But will it stay 
  correct? Requirements change. Code changes. What happens if the 
  requirements change to allow target = w, with associated action 
  run_w_code? If we change the code that sets target, but neglect to change 
  this block of code, it will wrongly call run_z_code() and Bad Things will 
  occur. It would be good to write this block of code defensively, so that 
  it will either be correct, or fail immediately, even in the face of 
  future changes.

  The comment at the start of the block is a good first step, but people 
  are notorious for failing to read and update comments. Chances are it 
  will soon be obsolete. But with an assertion, we can both document the 
  assumptions of this block, and cause a clean, immediate failure if they 
  are violated:

  assert target in (x, y, z)
  if target == x:
      run_x_code()
  elif target == y:
      run_y_code()
  else:
      assert target == z
      run_z_code()


  Here, the assertions are both defensive programming and checked 
  documentation. I consider this to be a far superior solution than this:

  if target == x:
      run_x_code()
  elif target == y:
      run_y_code()
  elif target == z:
      run_z_code()
  else:
      # This can never happen. But just in case it does...
      raise RuntimeError("an unexpected error occurred")


  This tempts some helpful developer to "clean it up" by removing the 
  "unnecessary" test for value==c and removing the "dead code" of the 
  RuntimeError. Besides, "unexpected error" messages are embarrassing when 
  they occur, and they will.

  Design by contract is another good use of assertions. In design by 
  contract, we consider that functions make "contracts" with their callers. 
  E.g. something like this:

  "If you pass me an non-empty string, I guarantee to return the first 
  character of that string converted to uppercase."

  If the contract is broken by either the function or the code calling it, 
  the code is buggy. We say that functions have pre-conditions (the 
  constraints that arguments are expected to have) and post-conditions (the 
  constraints on the return result). So this function might be coded as:

  def first_upper(astring):
      assert isinstance(astring, str) and len(astring) > 0
      result = astring[0].upper()
      assert isinstance(result, str) and len(result) == 1
      assert result == result.upper()
      return result


  The aim of Design By Contract is that in a correct program, the pre-
  conditions and post-conditions will always hold. Assertions are typically 
  used, since (so the idea goes) by the time we release the bug-free 
  program and put it into production, the program will be correct and we 
  can safely remove the checks.

  Here's my advice when *not* to use assertions:

  * Never use them for testing user-supplied data, or for anything 
    where the check must take place under all circumstances.

  * Don't use assert for checking anything that you expect might fail
    in the ordinary use of your program. Assertions are for extraordinary
    failure conditions. Your users should never see an AssertionError;
    if they do, it's a bug to be fixed.

  * In particular, don't use assert just because it's shorter than an
    explicit test followed by a raise. Assert is not a shortcut for
    lazy coders.

  * Don't use them for checking input arguments to public library 
    functions (private ones are okay) since you don't control the 
    caller and can't guarantee that it will never break the 
    function's contract.

  * Don't use assert for any error which you expect to recover from.
    In other words, you've got no reason to catch an AssertionError
    exception in production code.

  * Don't use so many assertions that they obscure the code.



  -- 
  Steven


