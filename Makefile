PY?=python
PELICAN?=pelican
PELICANOPTS=

BASEDIR=$(CURDIR)
THEMEDIR=$(BASEDIR)/theme
CSSDIR=$(THEMEDIR)/static/css
SCSSDIR=$(THEMEDIR)/styling
INPUTDIR=$(BASEDIR)/content
OUTPUTDIR=$(BASEDIR)/output
PLUGINSDIR=$(BASEDIR)/plugins
CONFFILE=$(BASEDIR)/pelicanconf.py
PUBLISHCONF=$(BASEDIR)/publishconf.py

# pelican plugins
I18N_SUBSITES_DIR=$(PLUGINSDIR)/i18n_subsites

# pip3 (python3-pip) on Ubuntu 20.04 gives following warning while installing requirements.txt:
# WARNING: The scripts pelican, pelican-import, pelican-quickstart and pelican-themes are installed in '${HOME}/.local/bin' which is not on PATH.
export PATH := ${PATH}:${HOME}/.local/bin


DEBUG ?= 0
ifeq ($(DEBUG), 1)
	PELICANOPTS += -D
endif

RELATIVE ?= 0
ifeq ($(RELATIVE), 1)
	PELICANOPTS += --relative-urls
endif

default: devserver

help:
	@echo 'Makefile for a pelican Web site                                           '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make html                           (re)generate the web site          '
	@echo '   make clean                          remove the generated files         '
	@echo '   make regenerate                     regenerate files upon modification '
	@echo '   make publish                        generate using production settings '
	@echo '   make serve [PORT=8000]              serve site at http://localhost:8000'
	@echo '   make serve-global [SERVER=0.0.0.0]  serve (as root) to $(SERVER):80    '
	@echo '   make devserver [PORT=8000]          serve and regenerate together      '
	@echo '                                                                          '
	@echo 'Set the DEBUG variable to 1 to enable debugging, e.g. make DEBUG=1 html   '
	@echo 'Set the RELATIVE variable to 1 to enable relative urls                    '
	@echo '                                                                          '

html: scss
	$(PELICAN) $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS)

clean:
	[ ! -d $(OUTPUTDIR) ] || rm -rf $(OUTPUTDIR)

regenerate: scss
	$(PELICAN) -r $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS)

serve:
ifdef PORT
	$(PELICAN) -l $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS) -p $(PORT)
else
	$(PELICAN) -l $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS)
endif

serve-global:
ifdef SERVER
	$(PELICAN) -l $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS) -p $(PORT) -b $(SERVER)
else
	$(PELICAN) -l $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS) -p $(PORT) -b 0.0.0.0
endif


devserver: scss
ifdef PORT
	$(PELICAN) -lr $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS) -p $(PORT)
else
	$(PELICAN) -lr $(INPUTDIR) -o $(OUTPUTDIR) -s $(CONFFILE) $(PELICANOPTS)
endif

publish: scss
	$(PELICAN) $(INPUTDIR) -o $(OUTPUTDIR) -s $(PUBLISHCONF) $(PELICANOPTS)

scss:
	[ -d $(CSSDIR) ] || mkdir -p $(CSSDIR)
	$(PY) -mscss < $(SCSSDIR)/style.scss -I $(SCSSDIR) -o $(CSSDIR)/style.css

download:
	# download pelican i18n_subsites plugin
	[ ! -d $(I18N_SUBSITES_DIR) ] || rm -rf $(I18N_SUBSITES_DIR)
	mkdir -p $(I18N_SUBSITES_DIR)
	wget -P $(I18N_SUBSITES_DIR) https://raw.githubusercontent.com/getpelican/pelican-plugins/master/i18n_subsites/__init__.py
	wget -P $(I18N_SUBSITES_DIR) https://raw.githubusercontent.com/getpelican/pelican-plugins/master/i18n_subsites/i18n_subsites.py
	# download normalize.css
	wget -O $(SCSSDIR)/_normalize801.scss https://necolas.github.io/normalize.css/8.0.1/normalize.css


.PHONY: html help clean regenerate serve serve-global devserver publish download scss
