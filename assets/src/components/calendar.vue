<template>
  <div>
    <p class="h5 my-3">
      <badge-day/>
      <badge-night/>
    </p>
    <simple-calendar
      v-bind:events="$store.state.events"
      v-bind:startingDayOfWeek="1"
      v-bind:show-date="showDate"
      v-on:show-date-change="showDateChange"
      v-on:click-date="clickDate"
      v-on:click-event="clickEvent"
    />
    <div class="modal fade" id="modal" tabindex="-1" role="dialog" aria-labelledby="modalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalLabel">{{ modalTitle }}</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-check form-check-inline">
              <label class="form-check-label">
                <input class="form-check-input" type="radio" name="inlineRadioOptions" value="day" v-model="modalCategory">
                <badge-day/>
              </label>
            </div>
            <div class="form-check form-check-inline">
              <label class="form-check-label">
                <input class="form-check-input" type="radio" name="inlineRadioOptions" value="night" v-model="modalCategory">
                <badge-night/>
              </label>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-outline-dark" v-bind:class="{ disabled: modalCategory == null }" v-on:click="leave">辞退する</button>
            <button type="button" class="btn btn-outline-dark" v-bind:class="{ disabled: modalCategory == null }" v-on:click="join">参加する</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import http from '../http';
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'
import SimpleCalendar from "vue-simple-calendar"

export default {
  name: 'Calendar',
  components: {
    'simple-calendar': SimpleCalendar,
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
  data: function() {
    return {
      showDate: new Date(),
      modalTitle: null,
      modalDate: null,
      modalCategory: null,
    }
  },
  methods: {
    showDateChange(date) {
      this.showDate = date;
    },
    showModal() {
      $('#modal').modal('show');
    },
    hideModal() {
      $('#modal').modal('hide');
      this.modalTitle = null;
      this.modalDate = null;
      this.modalCategory = null;
    },
    leave() {
      if (this.modalCategory == null) return;
      http.put('/events', {category: this.modalCategory, date: this.modalDate})
        .then((data) => {
          this.hideModal();
          this.$store.commit('removeEvent', data.event);
        })
        .catch((error) => {
          this.hideModal();
          this.$toasted.show(error);
        });
    },
    join() {
      if (this.modalCategory == null) return;
      http.post('/events', {category: this.modalCategory, date: this.modalDate})
        .then((data) => {
          this.hideModal();
          this.$store.commit('addEvent', data.event);
        })
        .catch((error) => {
          this.hideModal();
          this.$toasted.show(error);
        });
    },
    clickDate(date) {
      if (this.$store.getters.isSignIn) {
        this.modalTitle = date.getFullYear() + "年" + (date.getMonth() + 1) + "月"+ date.getDate() + "日";
        this.modalDate = date;
        this.showModal();
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
