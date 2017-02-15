[AWS] Upgrade from Ubuntu 12.04 to 14.04
########################################

:date: 2016-05-01T12:25+08:00
:tags: Amazon Web Services (AWS), Bash, Ubuntu Linux, Commandline
:category: Amazon Web Services (AWS)
:summary: Upgrade from AWS_ EC2_ t1.micro_ running `Ubuntu 12.04 LTS`_ with
          Apache-2.2 to t2.nano_ running `Ubuntu 14.04 LTS`_ with Apache-2.4.
:og_image: https://upload.wikimedia.org/wikipedia/commons/1/1d/AmazonWebservices_Logo.svg
:adsu: yes


Upgrade from AWS_ EC2_ t1.micro_ running `Ubuntu 12.04 LTS`_ with Apache-2.2 to
t2.nano_ running `Ubuntu 14.04 LTS`_ with Apache-2.4. The following is the
problems I have.

access / on this server
+++++++++++++++++++++++

Upgrade to Apache-2.4 from 2.2. Got the following message:

.. code-block:: txt

  You don't have permission to access / on this server

`Solution <http://stackoverflow.com/a/14671738>`__:

Add

.. code-block:: conf

  	<Directory />
  		Require all granted
  	</Directory>

to `/etc/apache2/sites-available/000-default.conf`

.. code-block:: conf

  <VirtualHost *:80>

  	...

  	<Directory />
  		Require all granted
  	</Directory>
  </VirtualHost>

.. adsu:: 2

enable apache file and directory indexing
+++++++++++++++++++++++++++++++++++++++++

Solution [6]_: Edit `/etc/apache2/apache2.conf`. Find:

.. code-block:: conf

  <Directory />
  	Options FollowSymLinks
  	AllowOverride None
  	Require all denied
  </Directory>

Change to:

.. code-block:: conf

  <Directory />
  	Options Indexes FollowSymLinks
  	AllowOverride None
  	Require all denied
  </Directory>

Then restart apache2:

.. code-block:: sh

  $ sudo service apache2 restart
  # or
  $ sudo /etc/init.d/apache2 restart

.. adsu:: 3

unable to resolve host
++++++++++++++++++++++

Upgrade from Ubuntu 12.04 to 14.04. Got the following message:

.. code-block:: sh

  unable to resolve host ip-172-30-0-19

Sulution [4]_: add the following line to `/etc/hosts`

.. code-block:: conf

  127.0.1.1 ip-172-30-0-19


----

References:

.. [1] `Migrate from t1.micro to t2.micro Amazon AWS - Stack Overflow <http://stackoverflow.com/questions/26676933/migrate-from-t1-micro-to-t2-micro-amazon-aws>`_

.. [2] `[AWS] Create/Migrate Linux Users on Amazon EC2 <{filename}../../04/30/aws-create-or-migrate-linux-users-on-ec2%en.rst>`_

.. [3] `apache 2.2 - NameVirtualHost has no effect - Server Fault <http://serverfault.com/questions/576939/namevirtualhost-has-no-effect>`_
.. adsu:: 4
.. [4] `AWS Developer Forums: sudo: unable to resolve host ... <https://forums.aws.amazon.com/message.jspa?messageID=495274>`_

.. [5] `apache - Error message “Forbidden You don't have permission to access / on this server” <http://stackoverflow.com/questions/10873295/error-message-forbidden-you-dont-have-permission-to-access-on-this-server>`_

.. [6] | `apache list files in directory - Google search <https://www.google.com/search?q=apache+list+files+in+directory>`_
       | `How do I enable apache file and directory indexing under Linux or UNIX? <http://www.cyberciti.biz/faq/enabling-apache-file-directory-indexing/>`_


.. _SSH: https://www.google.com/search?q=SSH
.. _AWS: https://aws.amazon.com/
.. _EC2: https://aws.amazon.com/ec2/
.. _t1.micro: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts_micro_instances.html
.. _t2.nano: https://aws.amazon.com/blogs/aws/ec2-update-t2-nano-instances-now-available/
.. _Ubuntu 14.04 LTS: https://aws.amazon.com/marketplace/pp/B00JV9TBA6/
.. _Ubuntu 12.04 LTS: https://aws.amazon.com/marketplace/pp/B007Z5YWX4/
.. _mount: http://linux.die.net/man/8/mount
