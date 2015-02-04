Evaluation of Multivariate Gaussian with NumPy
##############################################

:tags: Hidden_Markov_model, multivariate gaussian, multivariate normal, NumPy, Python
:category: Math
:summary: Eva


To implement a continuous `HMM <http://en.wikipedia.org/wiki/Hidden_Markov_model>`_, it involves the evaluation of `multivariate Gaussian <http://en.wikipedia.org/wiki/Multivariate_normal_distribution>`_ (multivariate normal distribution). This post gives description of how to evaluate multivariate Gaussian with `Numpy <http://www.numpy.org/>`_.

The formula for multivariate Gaussian used for continuous HMM is:

.. role:: raw-html(raw)
   :format: html

.. rubric:: :raw-html:`<a href="http://www.codecogs.com/eqnedit.php?latex=N(o;\mu,\Sigma)&space;=\frac{1}{\sqrt{(2\pi)^{n}\left&space;|&space;\Sigma&space;\right&space;|}}e^{-\frac{1}{2}(o-\mu)'\Sigma^{-1}(o-\mu)}" target="_blank"><img src="http://latex.codecogs.com/gif.latex?N(o;\mu,\Sigma)&space;=\frac{1}{\sqrt{(2\pi)^{n}\left&space;|&space;\Sigma&space;\right&space;|}}e^{-\frac{1}{2}(o-\mu)'\Sigma^{-1}(o-\mu)}" title="N(o;\mu,\Sigma) =\frac{1}{\sqrt{(2\pi)^{n}\left | \Sigma \right |}}e^{-\frac{1}{2}(o-\mu)'\Sigma^{-1}(o-\mu)}" /></a>`
   :class: align-center

where :math:`o` is vector extracted from observation, :math:`\mu` is mean vector, and :math:`\Sigma` is covariance matrix.

For the computation and implementation to be easily done, covariance matrix is assume to be diagonal without loss of much performance. The following is the code snippet:

.. code-block:: python

  dimension = mean_vector.size
  detDiagCovMatrix = numpy.sqrt(numpy.prod(numpy.diag(covariance_matrix)))
  frac = (2*numpy.pi)**(-dimension/2.0) * (1/detDiagCovMatrix)
  fprime = feature_vector - mean_vector
  fprime **= 2
  if log:
    logValue = -0.5*numpy.dot(fprime, 1/numpy.diag(covariance_matrix))
    logValue += numpy.log(frac)
    return logValue
  else:
    return frac * numpy.exp(-0.5*numpy.dot(fprime, 1/numpy.diag(covariance_matrix)))

Both mean vector and covariance matrix are of type numpy.ndarray with proper dimension. The code is self-explanatory. If more details are needed, please follow the link in [1]_.

----

References:

.. [1] `Source Code Reference <http://projects.scipy.org/scikits/browser/trunk/learn/scikits/learn/machine/em/densities.py?rev=447>`_.

       PS: Original link is invalid now. See `archive 1 <http://scikit-learn.sourcearchive.com/documentation/0.3-2/densities_8py-source.html>`_ or `archive 2 <http://www.sourcecodebrowser.com/python-scipy/0.6.0/densities_8py_source.html>`_ for the source code.

.. [2] `David Cournapeau <http://www.ar.media.kyoto-u.ac.jp/members/david/softwares/em/index.html>`_ (author of source code in [1]_)

.. [3] `Numpy Example List <http://wiki.scipy.org/Numpy_Example_List>`_
