Pip Install from Setup.py
#########################

:tags: Python
:category: Python
:author: Siong-Ui Te
:summary: Pip Install from Setup.py

Question: Assume that I git clone a repository from Github.
I make some modifications for this repository and would like
to install the modified repository locally using pip. How do I do?

Answer:

.. code-block:: bash

  $ cd path_to_the_modified_repo
  $ pip install -e .


Reference:

 1. `pip install a local git repository`_

 2. `pip to install using a local setup.py`_

 3. `stackoverflow question`_


.. _`pip install a local git repository`: http://stackoverflow.com/questions/14159482/pip-install-a-local-git-repository
.. _`pip to install using a local setup.py`: https://groups.google.com/d/topic/python-virtualenv/Z11FHJDYKEk
.. _`stackoverflow question`: http://stackoverflow.com/questions/2087148/can-i-use-pip-instead-of-easy-install-for-python-setup-py-install-dependen 
