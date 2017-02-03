'use strict';

var app = new Vue({
  el: '#suggestApp',
  data: {
    userinput: '',
    suggestedWords: []
  },
  computed: {
    isShowSuggestMenu: function() {
      return (this.suggestedWords.length > 0);
    },
    styleObject: function() {
      var inputElm = document.querySelector('input[autocomplete=off]');
      return {
        left: inputElm.offsetLeft + 'px',
        minWidth: inputElm.offsetWidth + 'px'
      };
    }
  },
  watch: {
    /**
     * Watch *userinput* change.
     * If change, generate 10 random strings as suggested words.
     * Implement your own suggest function here.
     */
    userinput: function(val, oldVal) {

      // min(inclusive), max(inclusive)
      function getRandomInt(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
      }

      function randomString(strlen) {
        const chars = "abcdefghijklmnopqrstuvwxyz0123456789";
        var result = "";
        for (var i=0; i<strlen; i++) {
          result += chars[getRandomInt(0,35)];
        }
        return result;
      }

      // empty *suggestedWords*
      while(this.suggestedWords.length > 0) {
        this.suggestedWords.pop();
      }

      if (val.length > 0) {
        // append 10 random strings to *suggestedWords*
        for (var i=0; i<10; i++) {
          this.suggestedWords.push(randomString(val.length));
        }
      }
    }
  }
})
