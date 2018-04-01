Uploaded Name of Photos/Videos of gphotosuploader
#################################################

:date: 2018-04-01T23:18+08:00
:tags: Go, Golang
:category: Go
:summary: The uploaded name of photos/video of gphotosuploader includes full
          file path. Remove the directory path and leave only file name.
:og_image: https://www.gstatic.com/images/branding/product/2x/photos_96dp.png
:adsu: yes


I tried gphotosuploader_ today and found it works great, except one small
problem. When I uploaded files, for example,

::

  /home/myusername/test.jpg

When I check the photo in the online Google Photos, I found the name of the
photo is not *test.jpg*. It is */home/myusername/test.jpg* instead. This is not
what I expected, because the uploaded filenames via online web interface does
not include full path. So I traced the code and found that it is easy to fix
this problem (for me).

In *api/upload.go*

.. code-block:: patch

   	"fmt"
   	"io"
   	"os"
  +	"path"
   	"regexp"
   	"time"
 
  @@ -47,7 +48,7 @@ func NewUploadOptionsFromFile(file *os.File) (*UploadOptions, error) {
   		Stream:   file,
   		FileSize: info.Size(),
 
  -		Name:      file.Name(),
  +		Name:      path.Base(file.Name()),
  		Timestamp: info.ModTime().Unix() * 1000,
   	}, nil
   }

The modification is made in the fork of original repo, i.e.,
*github.com/siongui/gphotosuploader*. I commit the change and use the following
command to build the package to see if my commit works:

.. code-block:: bash

  $ go get -u github.com/siongui/gphotosuploader

After build finished, I run the tool and found nothing changed. The file name
still included full path. I looked at the code again and found that some of the
source files include the packages of sub-directories of original repo:

::

  github.com/simonedegiacomi/gphotosuploader/api
  github.com/simonedegiacomi/gphotosuploader/auth

So I modified the name to my forked repo and re-build my forked package:

.. code-block:: bash

  $ go build -a github.com/siongui/gphotosuploader

Finally everything worked as expected.

.. adsu:: 2

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

----

References:

.. [1] `GitHub - simonedegiacomi/gphotosuploader: Unofficial Google Photos uploader and Go library <https://github.com/simonedegiacomi/gphotosuploader>`_

.. _gphotosuploader: https://github.com/simonedegiacomi/gphotosuploader
