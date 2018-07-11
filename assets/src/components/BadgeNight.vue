<template>
  <span class="badge badge-pill badge-primary">{{ value }}</span>
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
