<template>
  <div class="badge badge-pill badge-danger">{{ value }}</div>
</template>

<script>
export default {
  name: 'BadgeDay',
  props: ['text'],
  data: function() {
    return {
      value: this.text ? this.text : '昼Knowme',
    }
  }
}
</script>
