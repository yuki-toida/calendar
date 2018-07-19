<template>
  <div>
    HOGE
  </div>
</template>

<script>
import http from '../http';
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'

export default {
  name: 'Images',
  components: {
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
  data: function() {
    return {
    }
  },
  created: function() {
    this.fetch();
  },
  methods: {
    fetch() {
      http.get('/images')
        .then((data) => {
          console.log(data);
        })
        .catch((error) => this.$toasted.show(error));
    },
  }
}
</script>
