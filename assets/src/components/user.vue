<template>
  <div>
    Hoge
  </div>
</template>

<script>
// import http from '../http';
// import Calendar from "vue-simple-calendar"

// export default {
//   components: {
//     Calendar
//   },
//   data: function() {
//     return {
//       showDate: new Date()
//     }
//   },
//   methods: {
//     showDateChange(date) {
//       this.showDate = date;
//     },
//     clickDate(date) {
//       if (this.$store.getters.isSignIn) {
//         http.post('/events', {date: date})
//         .then((data) => {
//           this.$store.commit('addEvent', data.event);
//         })
//       } else {
//         console.log("サインインしてください");
//       }
//     },
//     clickEvent(event) {
//       if (this.$store.getters.isSignIn) {
//         if (event.id.includes(this.$store.state.user.id)) {
//           http.put('/events', {id: event.id, date: event.startDate})
//           .then((data) => {
//             this.$store.commit('removeEvent', event);
//           })
//         }
//       } else {
//         console.log("サインインしてください");
//       }
//     },
//   }
// }
</script>
