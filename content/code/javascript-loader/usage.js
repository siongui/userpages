/**
 * The file names of JavaScript files to be loaded
 */
var jsNames = ['a.js',
               'b.js',
               'c.js',
               'd.js'];

/**
 * A JavaScript file A is dependent on a JavaScript file B if B must be loaded
 * before A is loaded. To describe the dependencies, use:
 *
 * 'file name of A' : 'file name of B'
 *
 * If a JavaScript file A is not dependent on other JavaScript, use:
 *
 * 'file name of A' : null
 *
 * If a JavaScript file is dependent on multiple JavaScript files, the
 * multiple files is separated by comma.
 */
var jsDependencies = {
  'a.js': null,
  'b.js': 'a.js',
  'c.js': 'a.js',
  'd.js': 'b.js, c.js'
};

/**
 * The locations of JavaScript files, usually are URLs
 */
var jsLocations = {
  'a.js': 'http://example.com/static/js/a.js',
  'b.js': 'http://example.com/static/js/b.js',
  'c.js': 'http://example.com/static/js/c.js',
  'd.js': 'http://example.com/static/js/d.js'
};

var myloader = new MySimpleJSLoader(jsNames, jsDependencies, jsLocations);
