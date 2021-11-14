Copy Files to Android While Preserving Timestamp
################################################

:date: 2021-11-15T00:53+08:00
:tags: Bash, adb, Commandline, Android
:category: Bash
:summary: Copy files to Android smart phone without losing timestamps.
:og_image: https://i0.wp.com/www.infonautics.ch/blog/wp-content/uploads/2020/06/Copy-Files-With-Dates_EN.png?resize=750%2C340&ssl=1
:adsu: yes


Recently I bought a `Google Pixel XL`_ and tried to copy files to it. I found
that the timestamp of the copied files is the time of the creation of the copy,
not last modified date in original filesystem. After some googling [1]_, I found
that the only solution to the earlier version of Android is rooting the Android
devices.

Luckily enough, The Google Pixel XL can run Android 10 and it is possible to use
adb_ push to keep timestamp because the filesystem is changed from FUSE to
SDCardFS.

On Ubuntu system, it is easy to install adb: [3]_

.. code-block:: bash

  $ sudo apt-get install android-tools-adb android-tools-fastboot

Enable USB Debugging on Google Pixel XL [4]_. Open a terminal and use USB cable
to connect your phone and Ubuntu machine. Allow USB debugging on your phone,
in the Ubuntu terminal, we can see the Android device by:

.. code-block:: bash

  $ adb devices

You can see your phone in the list. Push the files to Android phone by:

.. code-block:: bash

  $ adb push myfolder /sdcard/Download

After pushing files finished, we can use adb shell to check the files on Android
phone:

.. code-block:: bash

  # Login to Android shell
  $ adb shell
  # Go to the Download folder
  $ cd /sdcard/Download
  # See the timestamp of files
  $ ls -al

If you are running Android 10 on Google Pixel XL, you can see the timestamp of
files is preserved after the push!

Hope this helps those who have the same issue as me.

Tested on: ``Ubuntu Linux 20.04``.

----

References:

.. [1] | `copy file to android preserve timestamp - Google search <https://www.google.com/search?q=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - DuckDuckGo search <https://duckduckgo.com/?q=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Ecosia search <https://www.ecosia.org/search?q=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Qwant search <https://www.qwant.com/?q=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Bing search <https://www.bing.com/search?q=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Yahoo search <https://search.yahoo.com/search?p=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Baidu search <https://www.baidu.com/s?wd=copy+file+to+android+preserve+timestamp>`_
       | `copy file to android preserve timestamp - Yandex search <https://www.yandex.com/search/?text=copy+file+to+android+preserve+timestamp>`_

.. [2] | `adb push preserve timestamp - Google search <https://www.google.com/search?q=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - DuckDuckGo search <https://duckduckgo.com/?q=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Ecosia search <https://www.ecosia.org/search?q=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Qwant search <https://www.qwant.com/?q=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Bing search <https://www.bing.com/search?q=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Yahoo search <https://search.yahoo.com/search?p=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Baidu search <https://www.baidu.com/s?wd=adb+push+preserve+timestamp>`_
       | `adb push preserve timestamp - Yandex search <https://www.yandex.com/search/?text=adb+push+preserve+timestamp>`_

.. [3] `How to Install ADB on Windows, macOS, and Linux <https://www.xda-developers.com/install-adb-windows-macos-linux/>`_
.. [4] `How to Enable USB Debugging on Google Pixel or Pixel XL (Android 7.0) <https://www.syncios.com/android/how-to-debug-google-pixel.html>`_
.. [5] `failed to authenticate to <ip> | adb wifi - Stack Overflow <https://stackoverflow.com/questions/19485467/failed-to-authenticate-to-ip-adb-wifi>`_

.. [6] | `root google pixel google photos - Google search <https://www.google.com/search?q=root+google+pixel+google+photos>`_
       | `root google pixel google photos - DuckDuckGo search <https://duckduckgo.com/?q=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Ecosia search <https://www.ecosia.org/search?q=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Qwant search <https://www.qwant.com/?q=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Bing search <https://www.bing.com/search?q=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Yahoo search <https://search.yahoo.com/search?p=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Baidu search <https://www.baidu.com/s?wd=root+google+pixel+google+photos>`_
       | `root google pixel google photos - Yandex search <https://www.yandex.com/search/?text=root+google+pixel+google+photos>`_
.. [7] | `[MOD][ROOT] Unlimited Original Quality Google Photo Storage | XDA Forums <https://forum.xda-developers.com/t/mod-root-unlimited-original-quality-google-photo-storage.3969433/>`_
       | `Is Google Photos unlimited after rooting? | XDA Forums <https://forum.xda-developers.com/t/is-google-photos-unlimited-after-rooting.3722725/>`_

.. _Google Pixel XL: https://support.google.com/pixelphone/answer/7158570
.. _adb: https://developer.android.com/studio/command-line/adb
