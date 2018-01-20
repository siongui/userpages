Avoid Vim overwriting indention settings
########################################

:date: 2017-11-03T01:22:00+08:00
:author: Ruey-Shi Rau (timrau)
:tags: python, tech, vim
:category: timrau's sandbox
:summary: I have tons of files in Python and other languages written as indented by 2 spaces, so I have the following settings in my ~/.vimrc:   Howev...
:og_image: https://scontent.fkhh1-2.fna.fbcdn.net/v/t1.0-9/16864052_10154594252128318_1036650617995059131_n.jpg?oh=8e888e99ca174d1daaa5cda10f426779&oe=5AFCB2D4

.. raw:: html

  I have tons of files in Python and other languages written as indented by 2 spaces, so I have the following settings in my ~/.vimrc:<br/>
  <script src="https://gist.github.com/timrau/f118690004942a7a9d5b1e50291708b9.js"></script> However, after I installed Vim 8.0 and python-mode for myself, I keep getting shiftwidth rewritten as 4 when I open Python files. I tried uninstalling python-mode, but it doesn&#39;t help.<br/>
  The reason is <a href="https://stackoverflow.com/a/21074056/718379" target="_blank">here</a>, and there is an <a href="https://github.com/vim/vim/issues/989" target="_blank">issue</a> filed against this behavior.<br/>
  <br/>
  To avoid modifying tons of files, I added the following line to my ~/.vimrc to turn this feature off.
  <script src="https://gist.github.com/timrau/4d31601804f539cde108d8de644e2c6e.js"></script>

----

`post <https://timrau.blogspot.com/2017/11/avoid-vim-overwriting-indention-settings.html>`_
by
`timrau <{filename}/pages/en/timrau.rst>`_
