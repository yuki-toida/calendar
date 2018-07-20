<template>
  <div class="badge badge-pill badge-success">{{ value }}</div>
</template>

<script>
export default {
  name: 'BadgeNight',
  props: ['text'],
  data: function() {
    return {
      value: this.text ? this.text : '夜Knowme',
    }
  }
}
</script>
