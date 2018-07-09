<template>
  <SimpleCalendar
    v-bind:events="$store.state.events"
    v-bind:startingDayOfWeek="1"
    v-bind:show-date="showDate"
    v-on:show-date-change="showDateChange"
    v-on:click-date="clickDate"
    v-on:click-event="clickEvent"
  />
</template>

<script>
import http from '../http';
import SimpleCalendar from "vue-simple-calendar"

export default {
  components: {
    SimpleCalendar
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
      if (this.$store.getters.isSignIn) {
        http.post('/events', {date: date})
        .then((data) => {
          this.$store.commit('addEvent', data.event);
        })
      } else {
        console.log("サインインしてください");
      }
    },
    clickEvent(event) {
      if (this.$store.getters.isSignIn) {
        if (event.id.includes(this.$store.state.user.id)) {
          http.put('/events', {id: event.id, date: event.startDate})
          .then((data) => {
            this.$store.commit('removeEvent', event);
          })
        }
      } else {
        console.log("サインインしてください");
      }
    },
  }
}
</script>
