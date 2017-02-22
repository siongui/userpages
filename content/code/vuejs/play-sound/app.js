'use strict';

var app = new Vue({
  el: '#app',
  methods: {
    play: function(event) {
      this.$refs.audioElm.play();
    }
  }
})
