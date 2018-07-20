<template>
  <div class="modal fade" id="modal" tabindex="-1" role="dialog" aria-labelledby="modalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">{{ title }}</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close" v-on:click="hide">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <div class="form-check form-check-inline">
            <label class="form-check-label">
              <input class="form-check-input" type="radio" name="inlineRadioOptions" value="day" v-model="category">
              <badge-day/>
            </label>
          </div>
          <div class="form-check form-check-inline">
            <label class="form-check-label">
              <input class="form-check-input" type="radio" name="inlineRadioOptions" value="night" v-model="category">
              <badge-night/>
            </label>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-outline-secondary" v-bind:class="{ disabled: category == null }" v-on:click="join">参加する</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import http from '../http';
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'

export default {
  name: 'Modal',
  props: ['date'],
  components: {
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
  data: function() {
    return {
      category: null,
    }
  },
  mounted: function() {
    this.$emit('show', false);
  },
  computed: {
    title: function() {
      return `${this.date.getFullYear()}年${this.date.getMonth() + 1}月${this.date.getDate()}日`;
    }
  },
  methods: {
    hide() {
      $('#modal').modal('hide');
      this.category = null;
      this.$emit('hide', false);
    },
    join() {
      if (this.category == null) return;
      http.post('/events', {category: this.category, date: this.date})
        .then((data) => {
          this.hide();
          this.$emit('join');
        })
        .catch((error) => {
          this.hide();
          this.$toasted.show(error);
        });
    },
  }
}
</script>
