<template>
  <div>
    <div v-if="0 < events.length" class="row mb-3">
      <div class="col-md-3">
        <select v-model="selected" class="form-control">
          <option v-for="event in events" v-bind:key="event.date" v-bind:value="event">
            {{ formatDay(event.date) }}
          </option>
        </select>
      </div>
      <div class="col-md-9">
        <div class="input-group">
          <div class="custom-file">
            <input v-on:change="changeFile" type="file" class="custom-file-input" id="customFile" lang="ja" accept="image/*" required="">
            <label class="custom-file-label" for="customFile">{{ uploadFileName }}</label>
          </div>
          <div class="input-group-append">
            <button v-on:click="upload" type="button" class="btn btn-outline-secondary">アップロード</button>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <dl v-for="key in Object.keys(images)" v-bind:key="key" class="mb-3">
          <dt>{{ format(key) }}</dt>
          <dd v-for="image in images[key]" v-bind:key="image" class="d-inline">
            <img v-bind:src="image" height="240" width="auto">
          </dd>
        </dl>
      </div>
    </div>
  </div>
</template>

<script>
import http from '../http';

export default {
  name: 'Images',
  data: function() {
    return {
      images: {},      
      events: [],
      selected: null,
      uploadFileName: 'ファイル選択...',
      uploadFile: null,
    }
  },
  created: function() {
    this.fetch();
  },
  methods: {
    fetch() {
      http.get('/images')
        .then((data) => {
          this.images = data.images;
          this.events = data.events;
        })
        .catch((error) => this.$toasted.show(error));
    },
    changeFile(e) {
      e.preventDefault();
      const file = e.target.files[0];
      this.uploadFileName = file.name;
      this.uploadFile = file;
    },
    upload(e) {
      if (this.selected) {
        const date = new Date(this.selected.date);
        let formData = new FormData();
        formData.append('year', date.getFullYear());
        formData.append('month', date.getMonth() + 1);
        formData.append('day', date.getDate());
        formData.append('category', this.selected.category);
        formData.append('file', this.uploadFile);
        http.post('/upload', formData)
          .then(() => this.fetch())
          .catch((error) => this.$toasted.show(error));
      } else {
        this.$toasted.show('日付が選択されていません');
      }
    },
    format(date) {
      if (typeof(date) == 'string') {
        date = new Date(date);
      }
      return `${date.getFullYear()}年${date.getMonth() + 1}月`;
    },
    formatDay(date) {
      if (typeof(date) == 'string') {
        date = new Date(date);
      }
      return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
    },
  }
}
</script>
