<template>
  <div>
    <img src="/static/img/header.jpg" alt="" class="img-thumbnail">
    <div class="row my-3">
      <div class="col-sm-6">
        <p>
          <badge-day/><span class="ml-3">{{ thisMonth }} 残席 {{ $store.state.dayEventRest }}</span>
        </p>
        <p>
          <badge-night/><span class="ml-3">{{ thisMonth }} 残席 {{ $store.state.nightEventRest }}</span>
        </p>
      </div>
      <div class="col-sm-6">
        <div v-if="$store.state.myEvent" class="card">
          <div class="card-body">
            <p class="card-text">
              {{ myEventDate }}
              <badge-day v-if="$store.state.myEvent.category == 'day'" v-bind:text="'昼'"/>
              <badge-night v-else v-bind:text="'夜'"/>
            </p>
            <ul class="list-inline">
              <li v-for="title in $store.state.myEvent.titles" v-bind:key="title" class="list-inline-item">
                {{ title }}
              </li>
            </ul>
            <button type="button" class="btn btn-outline-dark btn-sm" v-on:click="leave">参加を取り消す</button>
          </div>
        </div>
      </div>
    </div>
    <simple-calendar
      v-bind:events="$store.state.events"
      v-bind:startingDayOfWeek="1"
      v-bind:show-date="showDate"
      v-on:show-date-change="showDateChange"
      v-on:click-date="clickDate"
      v-on:click-event="clickEvent"
    />
    <modal v-if="showModal" v-bind:date="targetDate" v-on:hide="showModal = false"/>
  </div>
</template>

<script>
import SimpleCalendar from "vue-simple-calendar"
import Modal from './Modal.vue'
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'
import http from '../http';

export default {
  name: 'Calendar',
  components: {
    'simple-calendar': SimpleCalendar,
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
    'modal': Modal,
  },
  data: function() {
    return {
      now: new Date(),
      showDate: new Date(),
      showModal: false,
      targetDate: null,
    }
  },
  computed: {
    thisMonth: function() {
      return `${this.now.getFullYear()}年${this.now.getMonth() + 1}月`;
    },
    myEventDate: function() {
      const date = new Date(this.$store.state.myEvent.date);
      return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
    }
  },
  methods: {
    showDateChange(date) {
      this.showDate = date;
    },
    leave() {
      http.put('/events', {date: this.$store.state.myEvent.date, category: this.$store.state.myEvent.category})
        .then((data) => {
          window.location.reload();
        })
        .catch((error) => {
          this.$toasted.show(error);
        });
    },
    clickDate(date) {
      if (this.$store.getters.isSignIn) {
        this.targetDate = date;
        this.showModal = true;
      } else {
        this.$toasted.show("サインインしてください");
      }
    },
    clickEvent(event) {
      const id = event.id.split(':').pop();
      this.$router.push({ name: 'User', params: { id: id } });
    },
  }
}
</script>
