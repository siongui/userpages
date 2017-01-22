HTTP POST (multipart/form-data) on Android
##########################################

:tags: Android, HTTP POST, Java, multipart/form-data
:category: Android
:summary: Upload a file to the server in your Android app.
:adsu: yes


Since the coming of cloud computing, sometimes we would like to let the `Android <http://en.wikipedia.org/wiki/POST_%28HTTP%29>`_ application send `HTTP POST <http://en.wikipedia.org/wiki/POST_%28HTTP%29>`_ request to a server and get result back from the server. Especially, we would like to do HTTP POST which includes `multipart form data <http://stackoverflow.com/questions/4526273/what-does-enctype-multipart-form-data-mean>`_ (file and variables). This post is going to show you how to do this.

Assume the data to be posted contains a file whose name is **uploadedFile**, a variable whose name is **to**, and a variable whose name is **from**. The following is the code snippet for posting the data to the server:

.. show_github_file:: siongui userpages content/articles/2012/05/24/http-post-multipart-form-data-on-android.java

Hope this example would be helpful for those who are interested!
