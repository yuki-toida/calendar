<template>
  <div>
    <!-- <img v-bind:src="user.photo" class="rounded" width="auto" height="50">
    {{ user.name }}
    {{ user.id }} -->
    <div class="row my-2">
      <div class="col-sm-4">
        <div class="input-group">
          <input type="text" class="form-control" v-model="email" placeholder="example" aria-describedby="addon">
          <div class="input-group-append">
            <span class="input-group-text" id="addon">{{ $store.state.domain }}</span>
          </div>
        </div>
      </div>
      <div class="col-sm-2">
        <button class="btn btn-outline-secondary" type="button" v-on:click="search">検索</button>
      </div>
    </div>
    <table v-if="0 < events.length" class="table table-borderless table-hover">
      <tbody>
        <tr v-for="(event, index) in events" v-bind:key="index">
          <th scope="row">
            {{ format(event.date) }}
          </th>
          <td>
            <badge-day v-if="event.category == 'day'"/>
            <badge-night v-else/>
          </td>
          <td v-for="title in event.titles" v-bind:key="title">
            {{ title }}
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="text-secondary">
      参加履歴がありません<br/>
      {{ $route.params.id }} さんを誘ってみましょう
    </div>
  </div>
</template>

<script>
import http from '../http';
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'

export default {
  name: 'Search',
  components: {
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
  data: function() {
    return {
      user: null,
      events: [],
      email: this.$route.params.id.replace(this.$store.state.domain, ''),
    }
  },
  created: function() {
    this.fetch();
  },
  methods: {
    fetch() {
      http.get(`/search/${this.$route.params.id}`)
        .then((data) => {
          this.user = data.user;
          this.events = data.events;
        })
        .catch((error) => this.$toasted.show(error));
    },
    search() {
      if (this.email.includes('@')) {
        this.$toasted.show(this.$store.state.domain + '以降は不要です');
      } else {
        const id = this.email + this.$store.state.domain;
        this.$router.push({ name: 'Search', params: { id: id } });
        this.fetch();
      }
    },
    format(date) {
      if (typeof(date) == 'string') {
        date = new Date(date);
      }
      return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
    },
  }
}
</script>
