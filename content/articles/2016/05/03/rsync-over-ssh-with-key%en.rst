rsync over SSH with Key
#######################

:date: 2016-05-03T21:44+08:00
:tags: Amazon Web Services (AWS), Bash, Ubuntu Linux, Commandline, rsync
:category: Bash
:summary: Example of copy local files to remote machine which allows only SSH_
          key login via rsync_ command.
:og_image: https://upload.wikimedia.org/wikipedia/en/thumb/1/11/Newrsynclogo.png/160px-Newrsynclogo.png
:adsu: yes


Assume you have a remote machine which allows SSH_ login with key generated in
[1]_. The user account is `foo` and the key is stored in the file `~/fookey`.
The IP of the remote machine is `1.2.3.4`.

You want to copy the content of local `~/local_dir/` to remote `~/remote_dir` of
user `foo` via the following rsync_ command:

.. code-block:: sh

  $ rsync -avz ~/local_dir/ foo@1.2.3.4:~/remote_dir/ -e 'ssh -i ~/fookey'

----

References:

.. [1] `[AWS] Create/Migrate Linux Users on Amazon EC2 <{filename}../../04/30/aws-create-or-migrate-linux-users-on-ec2%en.rst>`_

.. [2] | `rsync over ssh - Google search <https://www.google.com/search?q=rsync+over+ssh>`_
       | `Rsync Examples, SSH, Backup @ Calomel.org <https://calomel.org/rsync_tips.html>`_
.. adsu:: 2
.. [3] | `rsync over ssh with key - Google search <https://www.google.com/search?q=rsync+over+ssh+with+key>`_
       | `ssh - Specify identity file (id_rsa) with rsync - Unix & Linux Stack Exchange <http://unix.stackexchange.com/questions/127352/specify-identity-file-id-rsa-with-rsync>`_

.. [4] `rsync over ftp - Google search <https://www.google.com/search?q=rsync+over+ftp>`_


.. _SSH: https://www.google.com/search?q=SSH
.. _rsync: https://www.google.com/search?q=rsync
