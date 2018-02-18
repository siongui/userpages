Trick to Get Instagram Live Video Replay
########################################

:date: 2018-02-18T22:11+08:00
:tags: Go, Golang, Web Scrape, Instagram
:category: Go
:summary: Trick to get Instagram post live video items.
:og_image: https://simplymeasured.com/wp-content/uploads/2016/12/IMG_0556-576x1024.png
:adsu: yes


When I use ChromeIGStory_, I found it can download the live video replay that
is shared to user stories. I know that the main endpoint of Instagram private
API which ChromeIGStory_ uses is the following *reels_tray* feed:

::

  https://i.instagram.com/api/v1/feed/reels_tray/

So I think maybe there is another endpoint that can get information of live
video replay, I looked at the source code and found there is a *LIVE_API* and
it seems to be the answer I need, so I tried but failed.

After several days, I looked at the source code of ChromeIGStory_ more closely,
but still cannot found any endpoint which returns information of live video
replay. Today I made some Google searches, and found the following thread:

- `get last shared live video · Issue #1370 · mgp25/Instagram-API · GitHub <https://github.com/mgp25/Instagram-API/issues/1370>`_

In the thread, I found the information is returned in the *reels_tray* feed,
which I already know how to access. But I cannot see any information of live
video replay in the returned JSON. So I tried again to see what other developers
send HTTP request, and I tried to set *cache-control* and *x-ig-capabilities* in
the HTTP request header, but it still did not work.

Then I searched agagin, and found maybe I need to add *ig_sig_key*, but it is
too difficult to get the *ig_sig_key*, so I did not try.

Finally, I made more Google searches with the keyword
*post live item instagram github* and found the replay_ section in
`instagram_private_api_extensions`_, and I found the *User-Agent* used in the
Python code is different from mine. My *User-Agent* is

::

  Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+

But they use:


::

  Instagram 10.26.0 (iPhone8,1; iOS 10_2; en_US; en-US; scale=2.00; gamut=normal; 750x1334) AppleWebKit/420+

I changed my *User-Agent* in my Go code and *BINGO*! It works! Now I see a new
field called *post-live* in the returned JSON in *reels_tray* feed.

It took me a lot of time to know the solution, so I write this post to take note
of my new finding!

.. adsu:: 2

----

References:

.. [1] `tricks to get post live · siongui/instago@3cee0a0 · GitHub <https://github.com/siongui/instago/commit/3cee0a066a3f3798f5988efc99469ca3761210dd>`_

.. _Instagram: https://www.instagram.com/
.. _ChromeIGStory: https://github.com/CaliAlec/ChromeIGStory
.. _replay: https://github.com/ping/instagram_private_api_extensions#replay
.. _instagram_private_api_extensions: https://github.com/ping/instagram_private_api_extensions
