[AWS] Create/Migrate Linux Users on Amazon EC2
##############################################

:date: 2016-04-30T22:50+08:00
:tags: Amazon Web Services (AWS), Bash, Ubuntu Linux, Commandline, cat command
:category: Amazon Web Services (AWS)
:summary: Create or migrate user account with SSH_ login only (no password
          login) on AWS_ EC2_ t2.nano_ with `Ubuntu 14.04 LTS`_.
:og_image: https://upload.wikimedia.org/wikipedia/commons/1/1d/AmazonWebservices_Logo.svg


Create or migrate user account with SSH_ login only (no password login) on AWS_
EC2_ t2.nano_ with `Ubuntu 14.04 LTS`_.


Create User Account
+++++++++++++++++++

.. code-block:: sh

  # login to EC2 running *Ubuntu 14.04*
  $ ssh -i my_key.pem ubuntu@my_EC2_IP
  # assume the name of the user account is *foo*
  $ sudo adduser foo --disabled-password
  # su to *foo*
  $ sudo su - foo
  # Create a *.ssh* directory for the *authorized_keys* file
  $ mkdir .ssh
  $ chmod 700 .ssh
  # Create the public and private keys for ssh login
  # Don't use a paraphrase -- just hit enter
  $ ssh-keygen -b 1024 -f fookey -t dsa
  # After keygen finished,
  # two files named *fookey* and *fookey.pub* will be generated.
  # *fookey* is private key and *fookey.pub* is public key
  $ cat fookey.pub > .ssh/authorized_keys
  $ chmod 600 .ssh/authorized_keys

Download *fookey* to your local machine. Now you can SSH login as *foo*:

.. code-block:: sh

  $ ssh -i fookey foo@my_EC2_IP

For more details, see [2]_, [4]_, and [5]_.


Migrate User Account
++++++++++++++++++++

My case is simple. Every user has his own block device as home directory. So
create user account as described above, mount_ the user's block device back to
his home directory. Then

.. code-block:: sh

  $ sudo chown foo:foo /home/foo -R

For more details or more sophisticated case, see [6]_ and [7]_.

----

References:

.. [1] `Migrate from t1.micro to t2.micro Amazon AWS - Stack Overflow <http://stackoverflow.com/questions/26676933/migrate-from-t1-micro-to-t2-micro-amazon-aws>`_

.. [2] `Managing User Accounts on Your Linux Instance - Amazon Elastic Compute Cloud <http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/managing-users.html>`_

.. [3] `Run Dictionary and Tipitaka Websites on Amazon Web Services <https://github.com/siongui/pali/blob/master/docs/AWS.md>`_

.. [4] `Manage multiple Linux Users on 1 Amazon EC2 Instance | the lost logbook <http://utkarshsengar.com/2011/01/manage-multiple-accounts-on-1-amazon-ec2-instance/>`_

.. [5] `How to add new users to EC2 and give SSH Key access <http://www.ampedupdesigns.com/blog/show?bid=44>`_

.. [6] `chown recursively changed permissions - Ask Ubuntu <http://askubuntu.com/questions/502110/chown-recursively-changed-permissions>`_

.. [7] `Linux: Changing UIDs and GIDs for a user by Stuart Colville <https://muffinresearch.co.uk/linux-changing-uids-and-gids-for-user/>`_


.. _SSH: https://www.google.com/search?q=SSH
.. _AWS: https://aws.amazon.com/
.. _EC2: https://aws.amazon.com/ec2/
.. _t2.nano: https://aws.amazon.com/blogs/aws/ec2-update-t2-nano-instances-now-available/
.. _Ubuntu 14.04 LTS: https://aws.amazon.com/marketplace/pp/B00JV9TBA6/
.. _mount: http://linux.die.net/man/8/mount
