Google App Engine Datastore Example : Dictionary Application
############################################################

:date: 2012-06-07 12:57
:modified: 2015-04-07 22:00
:tags: Python, Web application, web, Google App Engine, XML
:category: Google App Engine
:summary: Simple usage example of Google App Engine Python Datastore.
:adsu: yes


I have a dictionary website which needs to store over 50,000+ xml files of words
in the datastore. The following is the snapshot of part of the xml fils.

.. code-block:: bash

  $ ls
  ebhi.xml             ekaṃsena.xml          ekunasattati.xml  esaṃ.xml
  edhati.xml           ekamsikata.xml        ekūnasattati.xml  esanam.xml
  edha.xml             ekaṁsikatā.xml        ekunasatthi.xml   esānaṃ.xml
  edhi.xml             ekamsika.xml          ekūnasaṭṭhi.xml   esana.xml
  edh.xml              ekaṃsika.xml          ekunasiti.xml     esāna.xml
  edisaka.xml          ekanika.xml           ekūnāsīti.xml     esanā.xml
  edisa.xml            ekānika.xml           ekunatimsati.xml  esani.xml
  ehi-bhikkhu.xml      ekantam.xml           ekūnatiṃsati.xml  esanī.xml
  ehibhikkhu.xml       ekantaṃ.xml           ekunavisati.xml   esanta.xml
  ...

I use `Google App Engine`_ to host my website, and it takes me a while to figure
out how to put 50,000+ xml on GAE_ datastore and retrieve them efficiently. The
following is how I do it.

Before we start, we first need to define some terminology, which comes from GAE.
Objects in the GAE datastore are known as *entities*. Each entity in the
Datastore has a *key* that uniquely identifies it. The key consists of the
following components:

  - The kind_ of the entity, which categorizes it for the purpose of Datastore
    queries

  - An identifier_ for the individual entity

So now we define four terms: *entity*, *key*, *kind*, *id (identifier)*. Refer
to `Entities, Properties, and Keys - Python — Google Cloud Platform`_ for more
details.

In my application, I defile a *kind* called *PaliWord* for storing xml files in
the datastore.

.. code-block:: python

  from google.appengine.ext import ndb

  ...

  class PaliWord(ndb.Model):
    xmlfilename = ndb.StringProperty()
    xmlfiledata = ndb.TextProperty()

To store one xml file in the datastore, I use file name of the xml file without
`.xml` extension as *id* for manipulating the *entity*. The name and data of the
file is store in the *entity*.

.. code-block:: python

  def storeToNDB(filename, filedata):
    # id = filename without .xml extension
    paliword = PaliWord(id = filename[0:-4],
                        xmlfilename = filename,
                        xmlfiledata = filedata)
    paliword.put()
    return '%s : ok' % filename

To retrieve specific *entity*, we simply supply *id*, i.e., the file name
without `.xml` extension, to the `get_by_id()`_ of *PaliWord* *kind*. The
*entity* will be returned by this call if exists. Then we can further process
the *entity* the way we want.

.. code-block:: python

  def lookup(word):
    paliword = PaliWord.get_by_id(word)
    if (paliword):
      return decodeXML(paliword.xmlfilename, paliword.xmlfiledata.encode('utf8'))
    else:
      return u'查無此字(No Such Word)'

In my dictionary application, the use of App Engine Datastore is very simple and
straight forward because of the characteristics of my data. If you have a
similar application and want to know how to use GAE Datastore, I hope this would
be helpful.



.. _Google App Engine: https://cloud.google.com/appengine/

.. _GAE: https://cloud.google.com/appengine/

.. _kind: https://cloud.google.com/appengine/docs/python/datastore/entities#Python_Kinds_and_identifiers

.. _identifier: https://cloud.google.com/appengine/docs/python/datastore/entities#Python_Kinds_and_identifiers

.. _Entities, Properties, and Keys - Python — Google Cloud Platform: https://cloud.google.com/appengine/docs/python/datastore/entities

.. _get_by_id(): https://cloud.google.com/appengine/docs/python/datastore/modelclass#Model_get_by_id
