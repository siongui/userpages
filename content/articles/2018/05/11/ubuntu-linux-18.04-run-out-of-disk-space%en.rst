Ubuntu Linux 18.04 Run Out of Disk Space
########################################

:date: 2018-05-11T23:44+08:00
:tags: Bash, Commandline, Ubuntu Linux
:category: Ubuntu Linux
:summary: Upgrade from Ubuntu Linux 17.10 to 18.04 and disk space runs out
          quickly. And how I fix this issue.
:adsu: yes


Yesterday I upgraded my Ubuntu Linux from 17.10 to 18.04. I found a strange
phenomena. My free disk space (about 60GB) runs out quickly after several hours,
but I did not download anything that has such a big size.

I thought that maybe someone had the same problem as me, so I tried googling to
see if anyone had the same issue. But I found nothing. Maybe my search keywords
is not correct. So I tried searches again to find some ways to check where goes
wrong. I use ``sudo du -hs *`` starting from */*. And I found that a single log
file `/var/lib/gdm3/.local/share/xorg/Xorg.0.log` occupy all the free disk space
and are still growing until the system told me I have no disk space.

I used the keyword ``Xorg.0.log`` to search again to see if there is any
noticable issue with it. But I found nothing meaningful again. Then I tried to
use ``head -n 1000`` to print first 1000 lines of the log, and found that an
error message ``Touchpad: Read error 9`` keeps added to the log and finally used
up all my disk space.

I again use the error message to search again [1]_ and found that someone had
the same problem as me [2]_. Bingo! I got the root cause of the issue now. It
seems that Xorg replaces *input-synaptics* with *libinput*. The Ubuntu upgrader
did not remove *xserver-xorg-input-synaptics* package after upgrade and confuses
xorg package, so that the error message keeps added to the log.

The final solution is simple. Run

.. code-block:: bash

  $ sudo apt-get remove xserver-xorg-input-synaptics

After removing the package, reboot the system. And everything works fine now.

.. adsu:: 2

----

References:

.. [1] | `Elantech Touchpad: Read error 9 - Google search <https://www.google.com/search?q=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - DuckDuckGo search <https://duckduckgo.com/?q=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Ecosia search <https://www.ecosia.org/search?q=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Qwant search <https://www.qwant.com/?q=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Bing search <https://www.bing.com/search?q=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Yahoo search <https://search.yahoo.com/search?p=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Baidu search <https://www.baidu.com/s?wd=Elantech+Touchpad:+Read+error+9>`_
       | `Elantech Touchpad: Read error 9 - Yandex search <https://www.yandex.com/search/?text=Elantech+Touchpad:+Read+error+9>`_
.. [2] `[Resolved] Synaptics TouchPad: Read error 9 / Newbie Corner / Arch Linux Forums <https://bbs.archlinux.org/viewtopic.php?id=222031>`_
.. [3] `Ubuntu to remove xserver-xorg-input-synaptics on upgrade - Desktop - Ubuntu Community Hub <https://community.ubuntu.com/t/ubuntu-to-remove-xserver-xorg-input-synaptics-on-upgrade/3938>`_
