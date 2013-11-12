GNU/Linux Running on Browser with Network and Graphics Support
##############################################################

:tags: JavaScript, Linux
:category: JavaScript
:author: Siong-Ui Te
:summary: GNU/Linux Running on Browser with Network and Graphics Support

In my `previous post <{filename}../02/linux-run-on-browser#en.rst>`_, we see
somebody running GNU/Linux on the browser. `Slashdot <http://slashdot.org/>`_
posted a story_ that shows GNU/Linux running on browser with network and
graphics support. Very interesting!

  `live demo <http://jor1k.widgetry.org/>`_ (or `here <http://s-macke.github.io/jor1k/>`_)

  `source code <https://github.com/s-macke/jor1k>`_

  `more details about the project <https://github.com/s-macke/jor1k/wiki>`_

  `network support by Ben Burns <http://www.benjamincburns.com/2013/11/10/jor1k-ethmac-support.html>`_

The author (`Sebastian Macke <https://github.com/s-macke>`_) also left a
`comment <http://linux.slashdot.org/comments.pl?sid=4437617&cid=45403173>`_ on Slashdot:

::

  A year ago, when I started the project it was simply interest in learning Javascript. I was fascinated by the emulator from Fabrice Bellard http://www.bellard.org/jslinux/ [bellard.org] 
  I am a programmer focused on simulations/emulations and performance. I was also interested in learning the internals of how a computer nowadays works. The x86 CPU is way to compilcated. You lose the clear sight for stupid details like the A20 gate. The OpenRISC project is perfect. It is a CPU with a very easy and clear CPU. Nothing historic. It has even some similarities with byte code, which makes it very fast if you emulate it properly. I optimized especially for running Linux violating the specification a bit.
  The whole CPU with MMU fits in around 1000 lines of code. During that day I never expected to get that far. Now with all three cores and devices it needs around 7000 lines of code.

  I have a list of useful things you can do with it:
  1. Use it as an education system of the Linux system or other tools. For example you could write a git tutorial with live examples.
  2. This emulator provides an alternative way to port old software to run on modern systems. In direct comparison to the project Emscripten it is slow, but the porting could be much easier. For terminal applications probably no porting is neccessary at all (e. g. Frotz).
  3. The emulated OpenRISC CPU is very easy and contains around 1000 lines of code. So it is the perfect example to learn how emulation works.
  4. With network support it allows you to access other computers within the Web Browser providing ready to use tools. (Even an encrypted chat is possible if you run the sshd daemon)
  5. Use it as a speedtest for Javascript engines.
  6. It is an advertisement for the OpenRISC project.

  You can also read the motivation of Ben Burns in his Blog: http://www.benjamincburns.com/2013/11/10/jor1k-ethmac-support.html [benjamincburns.com] 
  And I have to admit that I did the wayland support last time only to get some news. :)


.. _story: http://linux.slashdot.org/story/13/11/12/161227/linux-kernel-running-in-javascript-emulator-with-graphics-and-network-support

