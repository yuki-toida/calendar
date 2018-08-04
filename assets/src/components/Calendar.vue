<template>
  <div>
    <img src="/static/img/header.jpg" alt="" class="img-thumbnail mb-3">
    <div v-if="0 < pictures.length" class="row mb-3">
      <div v-for="img in pictures" v-bind:key="img" class="col">
        <img v-bind:src="img" width="auto" height="100">
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div v-if="event" class="border border-bottom-0 py-3">
          <ul class="list-inline ml-3 mb-0">
            <li class="list-inline-item">
              <badge-day v-if="event.category == 'day'"/>
              <badge-night v-else/>
              {{ eventDate }}
            </li>
            <li class="list-inline-item">
              <ul class="list-inline mb-0">
                <li v-for="title in event.titles" v-bind:key="title" class="list-inline-item">
                  {{ title }}
                </li>
              </ul>
            </li>
            <li class="list-inline-item">
              <button type="button" class="btn btn-outline-dark btn-sm" v-on:click="leave">参加を取り消す</button>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="border border-bottom-0 py-3">
          <div class="ml-3 d-flex">
            <div>
              <badge-day/><span class="ml-1">残り{{ dayRestCouples }}組</span>
            </div>
            <div class="ml-4">
              <badge-night/><span class="ml-1">残り{{ nightRestCouples }}組</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <simple-calendar
      v-bind:events="events"
      v-bind:startingDayOfWeek="1"
      v-bind:show-date="showDate"
      v-bind:monthNameFormat="narrow"
      v-on:show-date-change="showDateChange"
      v-on:click-date="clickDate"
      v-on:click-event="clickEvent"
    />
    <modal v-if="modal" v-bind:date="targetDate" v-on:join="fetch" v-on:show="showModal" v-on:hide="modal = false"/>
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
      events: [],
      event: null,
      dayRestCouples: null,
      nightRestCouples: null,
      modal: false,
      targetDate: null,
      pictures: [],
    }
  },
  created: function() {
    this.fetch();
  },
  computed: {
    eventDate: function() {
      const date = new Date(this.event.date);
      return `${date.getMonth() + 1}月${date.getDate()}日参加予定`;
    }
  },
  methods: {
    fetch() {
      http.get('/events')
        .then((data) => {
          this.events = data.events;
          this.event = data.event;
          this.dayRestCouples = data.dayRestCouples;
          this.nightRestCouples = data.nightRestCouples;
          this.pictures = data.pictures;
        })
        .catch((error) => this.$toasted.show(error));
    },
    showDateChange(date) {
      this.showDate = date;
    },
    leave() {
      http.put('/events', {date: this.event.date, category: this.event.category})
        .then((data) => {
          this.fetch();
        })
        .catch((error) => {
          this.$toasted.show(error);
        });
    },
    clickDate(date) {
      if (this.$store.getters.isSignIn) {
        this.targetDate = date;
        this.modal = true;
        this.showModal();
      } else {
        this.$toasted.show("サインインしてください");
      }
    },
    clickEvent(event) {
      const id = event.id.split(':').pop();
      this.$router.push({ name: 'Search', params: { id: id } });
    },
    showModal() {
      $('#modal').modal('show');
    }
  }
}
</script>
