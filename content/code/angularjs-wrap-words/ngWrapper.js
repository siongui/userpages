'use strict';


angular.module('wordWrapper', ['tooltip']).
  directive('wrapEveryWord', ['wrapStringOfWords', function(wrapStringOfWords) {
    /**
     * wrap every word in ELEMENT and add tooltip to every word.
     */
    return {
      restrict: 'A',
      link: function(scope, elm, attrs) {
        wrapStringOfWords.traverse(elm[0])
      }
    };
  }]).

  factory('wrapStringOfWords', ['processStringOfWords', function(processStringOfWords) {
    /**
     * find all words in the element
     * @param {DOM element}
     */
    function traverse(elm) {
      // 1: element node
      if (elm.nodeType == 1) {
        for (var i=0; i<elm.childNodes.length; i++)
          // recursively call self to process
          traverse(elm.childNodes[i]);
        return;
      }

      // 3: text node
      if (elm.nodeType == 3) {
        var string = elm.nodeValue;
        if (string.replace(/\s*/, '') !== '')
          // string is not whitespace
          elm.parentNode.replaceChild(processStringOfWords.toDom(string), elm);
        return;
      }
    }

    var serviceInstance = { traverse: traverse };
    return serviceInstance;
  }]).

  factory('processStringOfWords', ['AddTooltipToElement', function(AddTooltipToElement) {

    function toDom(string) {
      // wrap all words in span
      var spanContainer = htmlStr2Dom(markInSpan(string));

      // add tooltip to every word
      for (var i=0; i<spanContainer.childNodes.length; i++) {
        var node = spanContainer.childNodes[i];
        var tagName = node.tagName;
        if (tagName && tagName.toLowerCase() === 'span') {
          AddTooltipToElement.add(node)
        }
      }
      return spanContainer;
    }

    function htmlStr2Dom(string) {
      // @see http://stackoverflow.com/questions/3103962/converting-html-string-into-dom-elements
      // @see http://stackoverflow.com/questions/888875/how-to-parse-html-from-javascript-in-firefox
      var spanContainer = document.createElement('span');
      spanContainer.innerHTML = string;
      return spanContainer;
    }

    function markInSpan(string) {
      return string.replace(/[AaBbCcDdEeGgHhIiJjKkLlMmNnOoPpRrSsTtUuVvYyĀāĪīŪūṀṁṂṃŊŋṆṇṄṅÑñṬṭḌḍḶḷ]+/g, '<span>$&</span>');
    }

    var serviceInstance = { toDom: toDom };
    return serviceInstance;
  }]);
