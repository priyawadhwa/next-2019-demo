import Vue from 'vue'
import App from './App.vue'
import VueFirestore from 'vue-firestore';

Vue.use(VueFirestore);

new Vue({
  el: '#app',
  render: h => h(App)
})
