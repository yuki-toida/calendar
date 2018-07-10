<template>
  <div>
    <p v-if="user" class="h5 my-3">
      <!-- <img v-bind:src="user.photo" class="rounded" width="auto" height="50"> -->
      {{ user.name }}
      {{ user.id }}
    </p>
    <p v-else class="h5 my-3">
      {{ $route.params.id }}
    </p>
    <table class="table table-hover">
      <tbody>
        <tr v-for="(event, index) in events" v-bind:key="index">
          <th scope="row">
            {{ event.date }}
            <badge-day v-if="event.category == 'day'"/>
            <badge-night v-else/>
          </th>
          <td v-for="name in event.names" v-bind:key="name">
            {{ name }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import http from '../http';
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'

export default {
  name: 'User',
  components: {
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
  data: function() {
    return {
      user: null,
      events: [],
    }
  },
  created: function() {
    http.get(`/users/${this.$route.params.id}`)
      .then((data) => {
        console.log(data);
        this.user = data.user;
        this.events = data.events;
      })
      .catch((error) => {
        this.$toasted.show(error);
      })
  },
  methods: {
  }
}
</script>
