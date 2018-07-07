<template>
  <div class="container">
		<calendar
      v-bind:events="$store.state.events"
      v-bind:startingDayOfWeek="1"
      v-bind:show-date="showDate"
			v-on:show-date-change="showDateChange"
      v-on:click-date="clickDate"
      v-on:click-event="clickEvent"
    />
  </div>
</template>

<script>
import http from '../http';
import Calendar from "vue-simple-calendar"

export default {
  components: {
    Calendar
  },
  data: function() {
    return {
      showDate: new Date()
    }
  },
  methods: {
    showDateChange(date) {
      this.showDate = date;
    },
    clickDate(date) {
      http.post('/', {
        userId: this.$store.state.user.id,
        date: date,
      })
      .then((data) => {
        this.$store.commit('addEvent', data.event);
      })
    },
    clickEvent(event) {
      http.put('/', {
        id: event.id,
        date: event.startDate,
      })
      .then((_) => {
        this.$store.commit('removeEvent', event);
      })
    },
  }
}
</script>
