adb push error: remote couldn't create file: Operation not permitted
####################################################################

:date: 2022-01-02T23:20+08:00
:tags: Commandline, Android, adb
:category: Bash
:summary: adb push error caused by illegal character in file name on Android.
:og_image: https://c.pxhere.com/photos/85/1a/android_android_phone_blur_close_up_data_device_display_electronics-1515453.jpg!d
:adsu: yes


When I tried to adb push file to Google Pixel 3 XL running Android 12. The
following error occurred:

.. code-block:: bash

  $ adb push story-2022-01-01T19:02:47+08:00-1641034967.mp4 /sdcard/DCIM/Camera/
  adb: error: failed to copy 'story-2022-01-01T19:02:47+08:00-1641034967.mp4' to '/sdcard/DCIM/Camera/story-2022-01-01T19:02:47+08:00-1641034967.mp4': remote couldn't create file: Operation not permitted

I tried to Google search the error message but the posts I found told me to
re-mount or root the phone, which did not work for me.

Before this I have successfully adb pushed files to the phone so the probably
the problem is not about mount or root. Then I tried to search
*illegal character android file name* and found that the problem was the ``:``
character is not allowed in the file name on Android 12. After replacing the
``:`` with allowed character, I successfully pushed the file to the phone.

----

References:

.. [1] `filenames - What characters allowed in file names on Android? - Stack Overflow <https://stackoverflow.com/questions/2679699/what-characters-allowed-in-file-names-on-android>`_
