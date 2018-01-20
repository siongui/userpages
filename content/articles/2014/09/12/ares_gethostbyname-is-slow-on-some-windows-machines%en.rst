ares_gethostbyname() is slow on some Windows machines
#####################################################

:date: 2014-09-12T23:08:00+08:00
:author: Ruey-Shi Rau (timrau)
:tags: tech
:category: timrau's sandbox
:summary: The function ares_gethostbyname()  in  c-ares  library, as of version 1.10.0, is slow on some Windows machines. ...
:og_image: https://scontent.fkhh1-2.fna.fbcdn.net/v/t1.0-9/16864052_10154594252128318_1036650617995059131_n.jpg?oh=8e888e99ca174d1daaa5cda10f426779&oe=5AFCB2D4

.. raw:: html

  The function <a href="http://linux.die.net/man/3/ares_gethostbyname" target="_blank">ares_gethostbyname()</a> in <a href="http://c-ares.haxx.se/" target="_blank">c-ares</a> library, as of version 1.10.0, is slow on some Windows machines.<br/>
  The more unconnected network interface the machine has, the slower the DNS lookup may be.<br/>
  This <a href="https://github.com/bagder/c-ares/pull/12" target="_blank">pull request</a> intended to solve this issue, although it is not merged into code base yet.<br/>
  <br/>
  Per Microsoft&#39;s document on <a href="http://msdn.microsoft.com/en-us/library/windows/desktop/aa365915(v=vs.85).aspx" target="_blank">GetAdaptersAddresses function</a>, GetAdaptersAddresses(), called by get_DNS_AdaptersAddresses() in c-ares library,<br/>
  <blockquote>
  is implemented only as a synchronous function. The GetAdaptersAddresses function requires a significant amount of network resources and time to complete since all of the low-level network interface tables must be traversed.
  </blockquote>
  Thus it makes sense to postpone calls to get_DNS_AdaptersAddresses() later. The only concern is that I&#39;m not sure whether the functionality is still correct or not.
  <div style="clear: both;"></div>

----

`post <https://timrau.blogspot.com/2014/09/c-ares-is-slow-on-windows.html>`_
by
`timrau <{filename}/pages/en/timrau.rst>`_
