<template>
  <div>
    <img src="/static/img/header.jpg" alt="" class="img-thumbnail">
    <div v-if="0 < images.length" class="row my-2">
      <div v-for="img in images" v-bind:key="img" class="col">
        <img v-bind:src="img" width="auto" height="100">
      </div>
    </div>
    <div class="row my-2">
      <div class="col">
        <div class="my-2">
          <badge-day/><span class="ml-3">{{ thisMonth }} 残席 {{ dayRestCount }}</span>
        </div>
        <div class="my-2">
          <badge-night/><span class="ml-3">{{ thisMonth }} 残席 {{ nightRestCount }}</span>
        </div>
      </div>
    </div>
    <div v-if="uploadText" class="row mt-2 mb-4">
      <div class="col">
        <div class="font-weight-bold mb-2">{{ uploadText }}</div>
        <div class="input-group">
          <div class="custom-file">
            <input v-on:change="changeFile" type="file" class="custom-file-input" id="customFile" lang="ja" accept="image/*" required="">
            <label class="custom-file-label" for="customFile">ファイル選択...</label>
          </div>
          <div class="input-group-append">
            <button v-on:click="upload" type="button" class="btn btn-outline-secondary">アップロード</button>
          </div>          
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div v-if="event" class="card border-bottom-0 rounded-0">
          <div class="card-body d-flex justify-content-between align-items-center">
            <p class="card-text mb-0">
              {{ eventDate }}
              <badge-day v-if="event.category == 'day'"/>
              <badge-night v-else/>
            </p>
            <div class="d-flex align-items-center">
              <ul class="list-inline mb-0 d-flex">
                <li v-for="title in event.titles" v-bind:key="title" class="list-inline-item mb-0 mr-2">
                  {{ title }}
                </li>
              </ul>
              <button type="button" class="btn btn-outline-dark btn-sm" v-on:click="leave">参加を取り消す</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <simple-calendar
      v-bind:events="events"
      v-bind:startingDayOfWeek="1"
      v-bind:show-date="showDate"
      v-on:show-date-change="showDateChange"
      v-on:click-date="clickDate"
      v-on:click-event="clickEvent"
    />
    <modal v-if="showModal" v-bind:date="targetDate" v-on:join="fetch" v-on:hide="showModal = false"/>
  </div>
</template>

<script>
import SimpleCalendar from "vue-simple-calendar"
import Modal from './Modal.vue'
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'
import http from '../http';
import { functions } from 'firebase';

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
      dayRestCount: null,
      nightRestCount: null,
      showModal: false,
      targetDate: null,
      images: [],
      uploadText: null,
      uploadFile: null,
    }
  },
  created: function() {
    this.fetch();
  },
  computed: {
    thisMonth: function() {
      return `${this.now.getFullYear()}年${this.now.getMonth() + 1}月`;
    },
    eventDate: function() {
      const date = new Date(this.event.date);
      return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
    }
  },
  methods: {
    fetch() {
      http.get('/events')
        .then((data) => {
          this.events = data.events;
          this.event = data.event;
          this.dayRestCount = data.dayRestCount;
          this.nightRestCount = data.nightRestCount;
          this.images = data.images;
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
        this.showModal = true;
      } else {
        this.$toasted.show("サインインしてください");
      }
    },
    clickEvent(event) {
      const id = event.id.split(':').pop();
      this.$router.push({ name: 'Search', params: { id: id } });
    },
    changeFile(e) {
      e.preventDefault();
      this.uploadFile = e.target.files[0];
    },
    upload(e) {
      const date = new Date(this.event.date);
      let formData = new FormData();
      formData.append('year', date.getFullYear());
      formData.append('month', date.getMonth() + 1);
      formData.append('day', date.getDate());
      formData.append('category', this.event.category);
      formData.append('file', this.uploadFile);
      http.post('/upload', formData)
        .then((data) => {
          console.log(data);
        })
        .catch((error) => this.$toasted.show(error));
    },
  }
}
</script>

<style>
.custom-file-input:lang(ja) ~ .custom-file-label::after {
  content: "選択";
}
.custom-file {
  max-width: 20rem;
  overflow: hidden;
}
.custom-file-label {
  white-space: nowrap;
}
</style>
