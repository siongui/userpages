/**
 * @fileoverview Simple JavaScript Loader (with dependency handling)
 */


/**
 * Class to load JavaScript files according to dependency.
 *
 * @param {Array} jsNames The file names of JavaScript files to be loaded.
 * @param {object} jsDependencies The dependencies of JavaScript files to be
 *                                loaded.
 * @param {object} jsLocations The locations of JavaScript files to be loaded.
 * @constructor
 */
MySimpleJSLoader = function(jsNames, jsDependencies, jsLocations) {
  // check data sanity
  if (Object.prototype.toString.apply(jsNames) != '[object Array]')
    throw "In MySimpleJSLoader constructor: jsNames is not Array";

  if (typeof jsDependencies != 'object')
    throw "In MySimpleJSLoader constructor: jsDependencies is not object";

  if (typeof jsDependencies != 'object')
    throw "In MySimpleJSLoader constructor: jsLocations is not object";

  for (var i=0; i < jsNames.length; i++) {
    if (!jsDependencies.hasOwnProperty(jsNames[i])) {
      throw "In MySimpleJSLoader constructor: no dependency info of " +
            jsNames[i];
    }
    if (!jsLocations.hasOwnProperty(jsNames[i])) {
      throw "In MySimpleJSLoader constructor: no location info of " +
            jsNames[i];
    }
  }

  /**
   * The dependencies of JavaScript files to be loaded.
   * @const
   * @type {object}
   * @private
   */
  this.jsDependencies_ = jsDependencies;

  /**
   * Indicate if the JavaScript file is loaded.
   * @type {object}
   * @private
   */
  this.isLoaded_ = {};
  /**
   * Contains the content of the JavaScript file
   * @type {object}
   * @private
   */
  this.jsContent_ = {};
  /**
   * Indicate if the content of JavaScript file is loaded.
   * @type {object}
   * @private
   */
  this.isJSContentDownloaded_ = {};

  // initialize internal private variable
  for (var i=0; i < jsNames.length; i++) {
    this.isLoaded_[ jsNames[i] ] = false;
    this.jsContent_[ jsNames[i] ] = null;
    this.isJSContentDownloaded_[ jsNames[i] ] = false;
  }

  // Load all JavaScript files
  for (var i=0; i < jsNames.length; i++) {
    this.loadJSFile(jsNames[i], jsLocations[ jsNames[i] ]);
  }
};


/**
 * Load one JavaScript file by XMLHttpRequest.
 * @param {string} jsName The name of the JavaScript file to be loaded.
 * @param {string} jsLocation The location of the JavaScript file to be loaded.
 * @private
 */
MySimpleJSLoader.prototype.loadJSFile = function(jsName, jsLocation) {
  /**
   * XMLHttpRequest variable.
   * @type {object}
   * @private
   */
  var xmlhttp;

  if (window.XMLHttpRequest) { xmlhttp = new XMLHttpRequest(); }
  else { xmlhttp = new ActiveXObject("Microsoft.XMLHTTP"); }

  xmlhttp.onreadystatechange = function(jsName, jsLocation) {
    if (xmlhttp.readyState == 4) {
      if (xmlhttp.status == 200) {
        // The content of the JavaScript file is downloaded.
        this.jsContent_[jsName] = xmlhttp.responseText;
        this.isJSContentDownloaded_[jsName] = true;

        if ( this.isDependencySatisfied(jsName) ) {
          this.insertJS(jsName);
        }
      } else {
        // fail to get the content of the JavaScript file
        throw 'cannot load file ' + jsName + ' at ' + jsLocation;
      }
    }
  }.bind(this, jsName, jsLocation);

  xmlhttp.open("GET", jsLocation, true);
  xmlhttp.send();
};


/**
 * Check if the dependency of the JavaScript file is satisfied.
 * @param {string} jsName The name of the JavaScript file to be checked.
 * @private
 */
MySimpleJSLoader.prototype.isDependencySatisfied = function(jsName) {
  if (this.jsDependencies_[jsName] == null) {
    console.log(jsName + ' has no dependency.');
    console.log(jsName + ' dependency satisfied? true');
    console.log('---');
    return true;
  }

  /**
   * The JavaScript files that the JavaScript file with jsName is dependent on.
   * @type {Array}
   * @private
   */
  var dependentJSFiles = this.jsDependencies_[jsName].split(',');

  for (var i=0; i < dependentJSFiles.length; i++) {
    // Remove whitespace in the beginning and end of the string
    dependentJSFiles[i] = dependentJSFiles[i].replace(/(^\s+)|(\s+$)/g, "");
  }

  console.log(jsName + ' depends on :');
  for (var i=0; i < dependentJSFiles.length; i++) {
    console.log(dependentJSFiles[i]);
  }

  /**
   * Indicate whether all dependencies is satisfied.
   * @type {boolean}
   * @private
   */
  var isAllDependenciesSatisfied = true;
  for (var i=0; i < dependentJSFiles.length; i++) {
    if (!this.isLoaded_[dependentJSFiles[i]]) {
      isAllDependenciesSatisfied = false;
    }
  }

  if (isAllDependenciesSatisfied) {
    console.log(jsName + ' dependency satisfied? true');
  } else {
    console.log(jsName + ' dependency satisfied? false');
  }
  console.log('---');

  return isAllDependenciesSatisfied;
};


/**
 * Insert the JavaScript file to document.
 * @param {string} jsName The name of the JavaScript file to be inserted.
 * @private
 */
MySimpleJSLoader.prototype.insertJS = function(jsName) {
  /**
   * DOM element of HTML script tag
   * @type {DOM Element}
   * @private
   */
  var script = document.createElement('script');
  script.setAttribute("type", "text/javascript");
  /**
   * The content of the JavaScript file
   * @type {DOM Element}
   * @private
   */
  var textNode = document.createTextNode( this.jsContent_[jsName] );
  script.appendChild(textNode);
  document.getElementsByTagName("head")[0].appendChild(script);

  this.isLoaded_[jsName] = true;

  console.log(jsName + ' loaded');
  console.log('---');

  // Load other JavaScript files dependent on this inserted JavaScript file.
  this.checkAfterInsertion();
};


/**
 * When a JavaScript is inserted to the docuement, dependencies of other
 * JavaScript files may be satisfied after the insertion. As a result, we need
 * to check dependencies of all JavaScript files after insertion.
 * @private
 */
MySimpleJSLoader.prototype.checkAfterInsertion = function() {
  for (var jsName in this.isLoaded_) {
    if (!this.isLoaded_[jsName]
        && this.isJSContentDownloaded_[jsName]
        && this.isDependencySatisfied(jsName)) {
      this.insertJS(jsName);
    }
  }
};
