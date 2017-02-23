'use strict';

var app = new Vue({
  el: '#app',
  methods: {
    play: function(event) {
      var a = this.$refs.audioElm;
      if (a.paused) {
        a.play();
      } else {
        a.pause();
      }
    }
  }
})
