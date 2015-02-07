Multiple Files Upload with Google App Engine Python
###################################################

:tags: Google App Engine, html, JavaScript, multipart/form-data, multiple files upload, Python
:category: Web Development
:summary: Let website users upload multiple files to Google App Engine (Python) servers.


On client side (**HTML** code for browsers):

.. code-block:: html

  <strong>Upload Files:</strong>
  <form method="post" action="/postmultifiles" enctype="multipart/form-data">
    <input type="file" name="filesToUpload[]" id="filesToUpload" multiple="" onChange="makeFileList();" />
    <input type="submit" value="Send">
  </form>
  <p><strong>Files You Selected:</strong></p>
  <ul id="fileList"><li>No Files Selected</li></ul>

On client side (**JavaScript** code for browsers):

.. code-block:: javascript

  function makeFileList() {
    var input = document.getElementById("filesToUpload");
    var ul = document.getElementById("fileList");
    while (ul.hasChildNodes()) {
      ul.removeChild(ul.firstChild);
    }
    for (var i = 0; i < input.files.length; i++) {
      var li = document.createElement("li");
      li.innerHTML = input.files[i].name;
      ul.appendChild(li);
    }
    if(!ul.hasChildNodes()) {
      var li = document.createElement("li");
      li.innerHTML = 'No Files Selected';
      ul.appendChild(li);
    }
  }

The above *HTML/JavaScript* code comes from [1]_. Put the code on your webpage.
When users choose multiple files, the files will be listed on the webpage.

| ----

On server side (**Python** code for Google App Engine):

.. code-block:: python

  class MultipleFilePage(webapp2.RequestHandler):
    def post(self):
      for file_data in self.request.POST.getall('filesToUpload[]'):
        file_name = file_data.filename
        file_content = file_data.value
        # do something here...

The code for receiving multiple files is easy. Please check the above code, which comes from [2]_.

----

References:

.. [1] `Multiple File Upload Input <http://davidwalsh.name/multiple-file-upload>`_

.. [2] `Receive multi file post with google app engine <http://stackoverflow.com/questions/1503526/receive-multi-file-post-with-google-app-engine>`_

.. [3] `Reading local files in JavaScript - HTML5 Rocks <http://www.html5rocks.com/en/tutorials/file/dndfiles/>`_
