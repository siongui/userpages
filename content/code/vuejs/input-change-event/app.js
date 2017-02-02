'use strict';

var app = new Vue({
  el: '#app',
  data: {
    userinput: '',
    value: '',
    oldValue: ''
  },
  watch: {
    userinput: function(val, oldVal) {
      this.value = val;
      this.oldValue = oldVal;
    }
  }
})
